package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mholt/archiver"
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

	archiver.Zip.Open("C:/dev/src/github.com/devplayg/fas/test/root.zip", "C:/dev/src/github.com/devplayg/fas/extract/")
	//	archiver.Tar.Open("C:/dev/src/github.com/devplayg/fas/test/root.tar", "C:/dev/src/github.com/devplayg/fas/extract/")
	//	archiver.TarGz.Open("C:/dev/src/github.com/devplayg/fas/test/root.tar.gz", "C:/dev/src/github.com/devplayg/fas/extract/")
	//	archiver.Zip.Open("C:/dev/src/github.com/devplayg/fas/test/root.7z", "C:/dev/src/github.com/devplayg/fas/extract/")

}

func printHelp() {
	fmt.Println("fap [options]")
	fs.PrintDefaults()
}

func CheckErr(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}
