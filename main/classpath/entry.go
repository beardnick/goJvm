package classpath

import (
	"runtime"
	"strings"
)

type Entry interface {
	readClass(className string)([]byte,Entry,error)
	String() string
}

// #TODO 2019/7/15 write newEntry code
func newEntry(path string)Entry {
	// find a needle in a haystack
	if strings.Contains(path,getClassPathSeparator()){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".jar") ||
	strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".jar"){
		return newZipEntry(path)
	}
	return newDirEntry(path)

}


func getClassPathSeparator()  string {
	switch runtime.GOOS{
	case "windows":
		return ";"
		break
	case "linux":
	case "darwin":
		return ":"
		break
	}
	// #NOTE 2019/7/15 这个地方设计得不错，使用零值代替nil
	return ""
}
