package classpath

import (
	"archive/zip"
	"errors"
	"goJvm/main/util"
	"io/ioutil"
	"log"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	//log.Println("newZipEntry:" , path)
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Panic(err)
	}
	return &ZipEntry{absDir}
}

func (this *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(this.absDir)
	if util.PanicError(err) {
		return nil, nil, err
	}
	defer r.Close()
	// #NOTE 2019/7/14 range切片会返回index和element
	for _, f := range r.File {
		if f.Name == className {
			classFile, err := f.Open()
			if util.PanicError(err) {
				return nil, nil, err
			}
			defer classFile.Close()
			data, err := ioutil.ReadAll(classFile)
			if util.PanicError(err) {
				return nil, nil, err
			}
			return data, this, nil
		}
	}
	return nil, nil, errors.New("class not Found: " + className)
}

func (this *ZipEntry) String() string {
	return this.absDir
}
