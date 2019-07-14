package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"goJvm/main/util"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry  {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Panic(err)
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string)([]byte,Entry,error){
	r, err := zip.OpenReader(self.absDir)
	if util.PanicError(err) {
		return nil, nil, err
	}
	defer r.Close()
	// #NOTE 2019/7/14 range切片会返回index和element
	for _, f := range r.File{
		if f.Name == className{
			classFile, err := f.Open()
			if util.PanicError(err){
				return nil,nil,err
			}
			defer classFile.Close()
			data , err := ioutil.ReadAll(classFile)
			if util.PanicError(err){
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil,nil,errors.New("class not Found: " + className)
}

func (self *ZipEntry) String() string{
	return self.absDir
}
