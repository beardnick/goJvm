package classpath

import (
	"log"
	"runtime"
	"strings"
)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// #TODO 2019/7/15 write newEntry code
func newEntry(path string) Entry {
	log.Println("newEntry :", path)
	// find a needle in a haystack
	// path1:path2
	if strings.Contains(path, getClassPathSeparator()) {
		return newCompositeEntry(path)
	}
	// path/*
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// class.zip
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".jar") {
		return newZipEntry(path)
	}
	// path
	return newDirEntry(path)

}

func getClassPathSeparator() string {
	// #IMP 2019-08-01 注意go的case语句不用break会自动跳转回去
	//fmt.Println("pathsep :", runtime.GOOS)
	switch runtime.GOOS {
	case "windows":
		return ";"
	case "linux":
		return ":"
	case "darwin":
		return ":"
	}
	// #NOTE 2019/7/15 这个地方设计得不错，使用零值代替nil
	return ""
}
