package classpath

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//log.Println("newDirEntry:" , path)
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Panic(err)
	}
	return &DirEntry{absDir}
}

func (this *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(this.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, this, err
}

func (this *DirEntry) String() string {
	return this.absDir
}
