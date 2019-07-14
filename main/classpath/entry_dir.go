package classpath

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string)*DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string)([]byte,Entry,error){
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string{
	return self.absDir
}


