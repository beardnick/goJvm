package classpath

import (
	"goJvm/main/util"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type WildcardEntry []Entry

func newWildcardEntry(path string) WildcardEntry {
	wildcardEntry := []Entry{}
	if strings.Index(path, "*") != -1{
		path = strings.TrimSuffix(path,  "*")
	}
	path, _ = filepath.Abs(path)
	dir, err := ioutil.ReadDir(path)
	if util.PanicError(err){
		return nil
	}
	for _, fileInfo := range dir{
		validSuffixes := [...]string{".class", ".zip", ".jar", ".JAR",".ZIP"}
		ok := false
		for _, suffix := range validSuffixes {
			if strings.HasSuffix(fileInfo.Name(),suffix){
				ok = true
				break
			}
		}
		if ok {
			wildcardEntry = append(wildcardEntry, newEntry(path + fileInfo.Name()))
		}
	}
	return wildcardEntry
}

func (this WildcardEntry) readClass(className string)([]byte,Entry,error){
	return CompositeEntry{&this}.readClass(className)
}

func (this WildcardEntry) String() string{
	return CompositeEntry{&this}.String()
}
