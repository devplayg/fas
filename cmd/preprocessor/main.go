package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultWatchDir = "/tmp/fas"
)

var (
	//	stats = expvar.NewMap("server")
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
		watchDir = fs.String("d", DefaultWatchDir, "Watching directory")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	log.Info("Starting server..")
	log.Infof("Target: %s", *watchDir)

}

func printHelp() {
	fmt.Println("fap [options]")
	fs.PrintDefaults()
}
