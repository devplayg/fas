package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/howeyc/fsnotify"
	log "github.com/sirupsen/logrus"
)

const (
	DefaultWatchDir = "/fas"
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
		watchDir = fs.String("d", os.TempDir()+DefaultWatchDir, "Watching directory")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])
	log.Infof("Starting server. Target dir: %s", *watchDir)

	// Start watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Error("event:", ev)

			case err := <-watcher.Error:
				log.Error(err.Error())
			}
		}
	}()

	//	err = watcher.Watch()

	watcher.WatchFlags(*watchDir, fsnotify.FSN_CREATE)
	if err != nil {
		log.Fatal(err)
	}

	// Stop
	waitForSignals()
	watcher.Close()
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
