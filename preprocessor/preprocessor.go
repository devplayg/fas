package preprocessor

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	//	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/mholt/archiver"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/filetype.v1"
)

type File struct {
	UUID      string
	P_UUID    string
	G_UUID    string
	Name      string
	Path      string
	MD5       string
	Size      uint64
	Priority  uint8
	Type      string
	Extension string
	Mime      string
	State     uint16
	Depth     int
	Children  []File
}

func newFile(path, p_uuid, g_uuid string, depth int) (*File, error) {
	file := File{
		Path:     path,
		UUID:     uuid.NewV4().String(),
		P_UUID:   p_uuid,
		G_UUID:   g_uuid,
		Name:     filepath.Base(path),
		Depth:    depth,
		Children: make([]File, 0, 1),
	}

	// Sleep
	time.Sleep(100 * time.Millisecond)

	// Read file
	data, err := ioutil.ReadFile(file.Path)
	if err != nil {
		log.Error(err)
		count := 1
		for err != nil && count <= 15 {
			time.Sleep(1 * time.Second)
			data, err = ioutil.ReadFile(file.Path)
			if err != nil {
				log.Error(err)
			}
			count++
		}

		if err != nil {
			//			this.errChan <- errors.New("Failed to read file: " + file.Path)
			return &file, err
		}
	}

	// Get MD5 checksum
	b := md5.Sum(data)
	file.MD5 = hex.EncodeToString(b[:16])
	file.Size = uint64(len(data))

	// Filetype
	filetype, _ := filetype.Match(data)
	file.Type = filetype.Extension
	file.Mime = filetype.MIME.Value
	//	log.Infof("File type: %s. MIME: %s\n", filetype.Extension, filetype.MIME.Value)
	//	spew.Dump(file)
	// Move file
	//	//	err = os.Rename(name, filepath.Join(this.homedir, "/storage", md5[0:2] ,md5+".bin"))
	//	err = os.Rename(file.Path, filepath.Join(this.homedir, "/storage", file.MD5+".bin"))
	//	if err != nil {
	//		this.errChan <- err
	//		return &file, err
	//	}

	return &file, nil
}

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
					file, err := newFile(event.Name, "", "", 0)
					if err != nil {
						errChan <- err
						<-this.c
						//						continue
					} else {
						this.enqueue(file)
					}

					//					// Issue UUID
					//					uuid := uuid.NewV4().String()
					//					file := File{
					//						Path:   event.Name,
					//						Depth:  0,
					//						UUID:   uuid,
					//						G_UUID: uuid,
					//					}
					//					uuid :=
					//					file := newFile(event.Name, p_uuid, g_uuid)

				}
				//				if event.Op&fsnotify.Write == fsnotify.Write {
				//					log.Println("modified file:", event.Name)
				//				}
			case err := <-watcher.Errors:
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

func (this *Preprocessor) enqueue(file *File) error { // uuid, p_uuid, g_uuid, -- priority, filesize, type,
	log.Infof("Mime: %s, Type: %s\n", file.Mime, file.Type)
	defer func() {
		<-this.c

	}()

	var err error
	// Extract
	if match, _ := regexp.MatchString("^(zip|7z)$", file.Type); match {
		tempDir := this.tempMkdir()

		switch file.Type {
		case "zip":
			err = archiver.Zip.Open(file.Path, tempDir)
		case "7z":
			err = archiver.TarGz.Open(file.Path, tempDir)
			//			log.Info("#")
			//			fmt.Println("Linux.")
		default:
			log.Info("#")
			// freebsd, openbsd,
			// plan9, windows...
			//			fmt.Printf("%s.", os)
		}

		if err != nil {
			this.errChan <- err
		}
		//		if file.Type == "zip" {
		//			err := archiver.Zip.Open(file.Path, tempDir)
		//		} else
	}

	//	log.Info(tempDir)

	//	// Extract
	//	tempDir := "tempDir"
	//	ReadDir(tempDir)
	//	for ReadFile(tempDir) {
	//		f := newFile(path, file.UUID, file.G_UUID, file.Depth+1)

	//		this.enqueue(f)
	//	}

	// Move file
	//	err = os.Rename(name, filepath.Join(this.homedir, "/storage", md5[0:2] ,md5+".bin"))
	err = os.Rename(file.Path, filepath.Join(this.homedir, "/storage", file.MD5+".bin"))
	if err != nil {
		this.errChan <- err
		return err
	}
	log.Infof("Rename complete: %s ", file.Path)

	// Issue UUID
	//	file.UUID = uuid.NewV4().String()

	// Get type

	// Get size

	// Is archive?

	//

	//

	// Write log
	//	time.Sleep(2 * time.Second)
	//	log.Debug("Enqueued: " + filepath)
	//	spew.Dump(file)

	return nil
}

func (this *Preprocessor) tempMkdir() string {
	dir, err := ioutil.TempDir("/home/fas/temp", "fas")
	if err != nil {
		this.errChan <- err
	}
	return dir
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
