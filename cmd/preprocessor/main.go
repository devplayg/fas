package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//	"time"

	"github.com/devplayg/fas/preprocessor"
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

	// Check directories

	// Open logging channel
	errChan := make(chan error, 1)
	go logDrain(errChan)

	// Start preprocessor
	preprocessor := preprocessor.NewPreprocessor(*homeDir)
	if err := preprocessor.Start(errChan); err != nil {
		log.Fatal(err)
	}

	// Stop
	waitForSignals()
	preprocessor.Stop()
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

func logDrain(errChan <-chan error) {
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Error(err.Error())
			}
		}
	}
}
