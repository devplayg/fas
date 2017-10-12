package engine

import (
	"os"
	"time"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func Enqueue(filepath string, c <-chan bool) error {
	//	defer return nil
	// Get checksum

	// Move file
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Error(err)
		<-c
		return err
	}

	// Issue UUID
	id := issueId()
	log.Info(id)

	// Write log

	time.Sleep(2 * time.Second)
	log.Info("Enqueued: " + filepath)
	//	<-c

	return nil

}

func issueId() uuid.UUID {
	return uuid.NewV4()
}

//func Enqueue() error {

//}

//type Dispatcher struct {
//	storageDir string
//}

//func NewDispatcher(storageDir string) *Dispatcher {
//	// Insert into DB
//	return &Dispatcher{
//		storageDir: storageDir,
//	}

//	return nil
//}

//func (this *Dispatcher) Enqueue(c chan<- bool, fp string) error {
//	return nil
//}

//func NewWatcher(datadir string) *Watcher {
//	return &Watcher{
//		datadir: datadir,
//	}
//}

//func (this *Watcher) Start(errChan chan<- error) error {
//	log.Info("watcher started")

//	//	watcher, err := fsnotify.NewWatcher()
//	//	if err != nil {
//	//		return err
//	//	}

//	//	//	done := make(chan bool)

//	//	// Process events
//	//	go func() {
//	//		for {
//	//			select {
//	//			case ev := <-watcher.Event:
//	//				log.Println("event:", ev)
//	//			case err := <-watcher.Error:
//	//				log.Println("error:", err)
//	//			}
//	//		}
//	//	}()

//	//	err = watcher.Watch("testDir")
//	//	if err != nil {
//	//		log.Fatal(err)
//	//	}

//	//	// Hang so program doesn't exit
//	//	<-done

//	//	/* ... do stuff ... */
//	//	watcher.Close()

//	return nil
//}
