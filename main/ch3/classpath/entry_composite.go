package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	//log.Println("newCompositeEntry:" , pathList)
	pathSep := getClassPathSeparator()
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathSep) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// #IMP 2019/7/15 前面的this就是指向自己的指针
// 和Java一样，需要一个指向自己的指针，只不过这里显示地写了出来而已
// func (this *CompositeEntry) readClass(className string)([]byte,Entry,error){
func (this CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range this {
		data, result, _ := entry.readClass(className)
		if result != nil {
			return data, result, nil
		}
	}
	return nil, nil, errors.New("Class not found error: " + className)
}

func (this CompositeEntry) String() string {
	if this == nil {
		return ""
	}
	pathSep := getClassPathSeparator()
	result := this[0].String()
	for i := 1; i < len(this); i++ {
		result += pathSep + this[i].String()
	}
	return result
}
