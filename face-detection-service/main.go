package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/ubuntu/face-detection-demo/appstate"
	"github.com/ubuntu/face-detection-demo/comm"
	"github.com/ubuntu/face-detection-demo/datastore"
	"github.com/ubuntu/face-detection-demo/detection"
	"github.com/ubuntu/face-detection-demo/messages"
)

var (
	wgwebcam         *sync.WaitGroup
	wgservices       *sync.WaitGroup
	shutdownwebcam   chan interface{}
	shutdownservices chan interface{}
	rootdir          string
)

//go:generate protoc --go_out=../messages/ --proto_path ../messages/ ../messages/communication.proto
func main() {
	var err error

	// Set main set of directories
	rootdir = os.Getenv("SNAP")
	if rootdir == "" {
		if rootdir, err = filepath.Abs(path.Join(filepath.Dir(os.Args[0]), "..")); err != nil {
			log.Fatal(err)
		}
	}
	datadir := os.Getenv("SNAP_DATA")
	if datadir == "" {
		datadir = rootdir
	}

	// check if we are in broken mode and remove database if it's the case
	appstate.CheckIfBroken(rootdir)
	if appstate.BrokenMode {
		datastore.WipeDB(datadir)
		detection.WipeScreenshots(datadir)
	}

	// Load logos and set arctefacts destination directory
	detection.InitLogos(path.Join(rootdir, "images"), datadir)

	// channels synchronization
	wgwebcam = new(sync.WaitGroup)
	wgservices = new(sync.WaitGroup)
	shutdownwebcam = make(chan interface{})
	shutdownservices = make(chan interface{})

	// handle user generated stop requests
	userstop := make(chan os.Signal)
	signal.Notify(userstop, syscall.SIGINT, syscall.SIGTERM)

	actions := make(chan *messages.Action, 2)

	// detect available cameras at startup
	detection.DetectCameras()

	// prepare settings and data
	datastore.LoadSettings(datadir)
	datastore.StartDB(datadir, shutdownservices, wgservices)

	fmt.Println(datastore.DB.Stats)

	// starts external communications channel
	comm.SetSocketDir(datadir)
	comm.StartSocketListener(actions, shutdownservices, wgservices)
	comm.StartServer(rootdir, datadir, actions)

	// starts camera if it was already started last time
	if datastore.FaceDetection() {
		detection.StartCameraDetect(rootdir, shutdownwebcam, wgwebcam)
	}

mainloop:
	for {
		fmt.Println("main loop")
		select {
		case action := <-actions:
			fmt.Println("new action received")
			if processaction(action) {
				break mainloop
			}
		case <-userstop:
			quit()
			break mainloop
		}
	}

	// Ensure webcam and services stopped
	wgwebcam.Wait()
	wgservices.Wait()
}

// process action and return true if we need to quit (exit mainloop)
// TODO: use quit channel (renamed userstop to quit) and send data there. Remove the bool True/False
func processaction(action *messages.Action) bool {
	if action.FaceDetection == messages.Action_FACEDETECTION_ENABLE {
		detection.StartCameraDetect(rootdir, shutdownwebcam, wgwebcam)
		fmt.Println("Received camera on")
	} else if action.FaceDetection == messages.Action_FACEDETECTION_DISABLE {
		detection.EndCameraDetect()
		fmt.Println("Received camera off")
	}
	if action.RenderingMode == messages.Action_RENDERINGMODE_FUN {
		datastore.SetRenderingMode(datastore.FUNRENDERING)
		comm.WSserv.SendAllClients(&messages.WSMessage{
			Type:          "renderingmode",
			RenderingMode: datastore.RenderingMode()})
	} else if action.RenderingMode == messages.Action_RENDERINGMODE_NORMAL {
		datastore.SetRenderingMode(datastore.NORMALRENDERING)
		comm.WSserv.SendAllClients(&messages.WSMessage{
			Type:          "renderingmode",
			RenderingMode: datastore.RenderingMode()})
	}
	// camera is offsetted by 1 for the client (0, protobuf default means no change)
	cameranum := int(action.Camera) - 1
	if cameranum > -1 && cameranum != datastore.Camera() {
		datastore.SetCamera(cameranum)
		comm.WSserv.SendAllClients(&messages.WSMessage{
			Type:   "newcameraactivated",
			Camera: cameranum + 1})
		if datastore.FaceDetection() {
			fmt.Println("Change active camera")
			go detection.RestartCamera(rootdir, shutdownwebcam, wgwebcam)
		}
	}
	if action.QuitServer {
		quit()
		return true
	}
	return false
}

func quit() {
	fmt.Println("quit server")
	// wait for webcam to shutdown, then ask services to shutdown
	close(shutdownwebcam)
	wgwebcam.Wait()
	close(shutdownservices)
}
