package watcher

import (
	//	"errors"

	//	"github.com/howeyc/fsnotify"
	log "github.com/sirupsen/logrus"
)

type Watcher struct {
	datadir string
}

func NewWatcher(datadir string) *Watcher {
	return &Watcher{
		datadir: datadir,
	}
}

func (this *Watcher) Start(errChan chan<- error) error {
	log.Info("watcher started")

	//	watcher, err := fsnotify.NewWatcher()
	//	if err != nil {
	//		return err
	//	}

	//	//	done := make(chan bool)

	//	// Process events
	//	go func() {
	//		for {
	//			select {
	//			case ev := <-watcher.Event:
	//				log.Println("event:", ev)
	//			case err := <-watcher.Error:
	//				log.Println("error:", err)
	//			}
	//		}
	//	}()

	//	err = watcher.Watch("testDir")
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	// Hang so program doesn't exit
	//	<-done

	//	/* ... do stuff ... */
	//	watcher.Close()

	return nil
}
