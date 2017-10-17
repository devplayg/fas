package preprocessor

import (
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/h2non/filetype.v1"
	//	"strconv"
	//"path/filepath"
	"crypto/md5"
	"io/ioutil"
	"time"

	//	"github.com/devplayg/golibs/checksum"
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

func (this *Preprocessor) enqueue(name string, depth int) error { // uuid, p_uuid, g_uuid
	defer func() {
		<-this.c
	}()

	// Sleep
	time.Sleep(100 * time.Millisecond)

	// Read file
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Info("###x")
		this.errChan <- err
		count := 1
		for err != nil && count <= 15 {
			time.Sleep(1 * time.Second)
			data, err = ioutil.ReadFile(name)
			if err != nil {
				this.errChan <- err
			}
			count++
		}

		if err != nil {
			this.errChan <- errors.New("Failed to read file: " + name)
			return err
		}
	}

	// Get MD5 checksum
	b := md5.Sum(data)
	md5 := hex.EncodeToString(b[:16])
	log.Debugf("Checksum=%s", md5)

	// Filetype
	filetype, err2 := filetype.Match(data)
	if err2 != nil {
		this.errChan <- err
		return nil
	}
	log.Infof("File type: %s. MIME: %s\n", filetype.Extension, filetype.MIME.Value)

	// Move file
	//	err = os.Rename(name, filepath.Join(this.homedir, "/storage", md5[0:2] ,md5+".bin"))
	err = os.Rename(name, filepath.Join(this.homedir, "/storage", md5+".bin"))

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
