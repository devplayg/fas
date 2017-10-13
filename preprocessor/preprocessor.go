package preprocessor

import (
	"os"
	"strconv"
	"time"

	"github.com/devplayg/golibs/checksum"
	"github.com/devplayg/golibs/strings"
	"github.com/howeyc/fsnotify"
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

	this.errChan = errChan
	this.init()
	var err error
	this.w, err = fsnotify.NewWatcher()
	if err != nil {
		this.errChan <- err
		return err
	}

	go func() {
		for {
			select {
			case ev := <-this.w.Event:
				if ev.IsCreate() {
					this.c <- true
					log.Debug("Request enqueuing: ", ev.Name)
					//					go Enqueue(ev.Name, c)
					go this.enqueue(ev.Name, 0)
				}
			case err := <-this.w.Error:
				this.errChan <- err
			}
		}
	}()
	this.w.WatchFlags(this.homedir+"/watch", fsnotify.FSN_CREATE)
	this.w.WatchFlags(this.homedir+"/user", fsnotify.FSN_CREATE)

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

	// Get checksum
	time.Sleep(10 * time.Millisecond)
	md5, err := checksum.GetMd5File(filepath)
	if err != nil {
		log.Error(err.Error())
		this.errChan <- err
		return err
	}
	log.Info("MD5: " + md5)

	// Move file
	err = os.Rename(filepath, this.homedir+"/storage/"+md5[0:2]+"/"+md5+".bin")
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
	for i := 0x00; i <= 0xff; i++ {
		str := strconv.FormatInt(int64(i), 16)
		str = strings.StrPadLeft(str, 2, "0")
		dir := this.homedir + "/storage/" + str
		//		log.Info(dir)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			e := os.Mkdir(dir, 0755)
			if e != nil {
				this.errChan <- err
			}
		}

	}
}
