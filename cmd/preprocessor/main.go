package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//	"time"

	"github.com/devplayg/fas"
	//	"github.com/howeyc/fsnotify"
	log "github.com/sirupsen/logrus"
)

const (
	DefaultHomeDIr = "/home/fas"
)

var (
	fs *flag.FlagSet
)

func init() {

	// Set logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: true,
	})
	log.SetLevel(log.InfoLevel)
	//	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//	if err == nil {
	//		log.SetOutput(file)
	//	} else {
	//		log.Error("Failed to log to file, using default stderr")
	//		log.SetOutput(os.Stdout)
	//	}
}

func main() {
	// Set flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		homeDir = fs.String("homedir", DefaultHomeDIr, "Home directory")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])
	log.Infof("Starting server. Homedir: %s", *homeDir)

	// Start watcher
	//	watcher, err := fsnotify.NewWatcher()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	c := make(chan bool, 3)
	go engine.Enqueue("/home/fas/user/1.png", c)
	go engine.Enqueue("/home/fas/user/2.png", c)
	go engine.Enqueue("/home/fas/user/3.png", c)
	go engine.Enqueue("/home/fas/user/4.png", c)

	//	go func() {
	//		for {
	//			select {
	//			case ev := <-watcher.Event:
	//				if ev.IsCreate() {
	//					c <- true
	//					log.Info("Request enqueuing: ", ev.Name)
	//					go engine.Enqueue(ev.Name, c)
	//				}
	//			case err := <-watcher.Error:
	//				log.Error(err.Error())
	//			}
	//		}
	//	}()

	//	// Start watching
	//	watcher.WatchFlags(*homeDir+"/watch", fsnotify.FSN_CREATE)
	//	watcher.WatchFlags(*homeDir+"/user", fsnotify.FSN_CREATE)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	// Stop
	waitForSignals()
	//	watcher.Close()
}

func printHelp() {
	fmt.Println("fap [options]")
	fs.PrintDefaults()
}

func waitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		log.Println("Signal received, shutting down...")
	}
}
