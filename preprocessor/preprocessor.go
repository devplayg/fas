package preprocessor

import (
	"os"
	//	"strconv"
	//"path/filepath"
	"time"

	"github.com/devplayg/golibs/checksum"
	//	"github.com/devplayg/golibs/strings"
	"github.com/fsnotify/fsnotify"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Preprocessor struct {
	homedir string
	w       *fsnotify.Watcher
	c       chan bool
	errChan chan<- error
}

func NewPreprocessor(homedir string) *Preprocessor {
	return &Preprocessor{
		homedir: homedir,
		c:       make(chan bool, 3),
	}
}

func (this *Preprocessor) Start(errChan chan<- error) error {

	//	this.errChan = errChan
	//	this.init()
	//	var err error
	//	this.w, err = fsnotify.NewWatcher()
	//	if err != nil {
	//		this.errChan <- err
	//		return err
	//	}
	//	go func() {
	//		for {
	//			select {
	//			case ev := <-this.w.Event:
	//				if ev.IsCreate() {
	//					this.c <- true
	//                    ev.
	//					//					log.Infof("Event: %s", ev)
	//					//					go Enqueue(ev.Name, c)
	//					//					log.Debug(ev.Name)
	//					go this.enqueue(ev.Name, 0)
	//				}
	//			case err := <-this.w.Error:
	//				this.errChan <- err
	//			}
	//		}
	//	}()
	//	this.w.WatchFlags(this.homedir+"/watch", fsnotify.FSN_CREATE)
	//	this.w.WatchFlags(this.homedir+"/user", fsnotify.FSN_CREATE)

	this.errChan = errChan
	this.init()
	var err error

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		this.errChan <- err
		return err
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:

				//log.Println("event:", event)
				//this.enqueue(event.Name, 0)
				if event.Op == fsnotify.Create {
					this.c <- true
					this.enqueue(event.Name, 0)
				}

				//				if event.Op&fsnotify.Write == fsnotify.Write {
				//					log.Println("modified file:", event.Name)
				//				}
			case err := <-watcher.Errors:
				//				log.Println("event:###")
				this.errChan <- err
			}
		}
	}()
	err = watcher.Add(this.homedir + "/watch")
	err = watcher.Add(this.homedir + "/user")

	if err != nil {
		this.errChan <- err
		return err
	}

	return nil
}

func (this *Preprocessor) Stop() error {
	this.w.Close()
	return nil
}

func (this *Preprocessor) enqueue(filepath string, depth int) error { // uuid, p_uuid, g_uuid
	defer func() {
		<-this.c
	}()
	var err error

	time.Sleep(100 * time.Millisecond)

	// check lock
	file, err := os.Open(filepath)
	file.Close()
	if err != nil {
		log.Infof("Waiting until file is unlocked: %s", filepath)
		time.Sleep(15000 * time.Millisecond)
	}

	// Get checksum
	t1 := time.Now()
	log.Debugf("Calculating checksum: %s", filepath)
	md5, err := checksum.GetMd5File(filepath)
	if err != nil {
		this.errChan <- err
		return err
	}
	log.Infof("Checksum=%s, exectime=%s", md5, time.Since(t1))
	//filepath.
	//filepath.
	// Move file
	//err = os.Rename(filepath, this.homedir+"/storage/"+md5[0:2]+"/"+md5+".bin")
	//filepath.
	err = os.Rename(filepath, this.homedir+"/storage/"+md5+".bin")
	if err != nil {
		this.errChan <- err
		return err
	}

	// Issue UUID
	uuid := uuid.NewV4()
	log.Debug(uuid)

	// Get type

	// Get size

	// Is archive?

	//

	//

	// Write log
	//	time.Sleep(2 * time.Second)
	//	log.Debug("Enqueued: " + filepath)
	return nil
}

func (this *Preprocessor) init() {
	//	for i := 0x00; i <= 0xff; i++ {
	//		str := strconv.FormatInt(int64(i), 16)
	//		str = strings.StrPadLeft(str, 2, "0")
	//		dir := this.homedir + "/storage/" + str
	//		//		log.Info(dir)
	//		if _, err := os.Stat(dir); os.IsNotExist(err) {
	//			e := os.Mkdir(dir, 0755)
	//			if e != nil {
	//				this.errChan <- err
	//			}
	//		}

	//	}
}
