package engine

//func Enqueue() error {

//}

//import (
//	log "github.com/sirupsen/logrus"
//)

type Dispatcher struct {
	storageDir string
}

func NewDispatcher(storageDir string) *Dispatcher {
	// Insert into DB
	return &Dispatcher{
		storageDir: storageDir,
	}

	return nil
}

func (this *Dispatcher) Enqueue(fp string) error {
	return nil
}

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
