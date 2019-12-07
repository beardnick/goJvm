package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) *Classpath  {
	cp := Classpath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return &cp
}

func (this *Classpath) ReadClass(className string)([]byte, Entry,error)  {
	className = className + ".class"
	if data, entry, err := this.bootClassPath.readClass(className); err == nil{
		return data, entry, err
	}
	if data, entry, err := this.extClassPath.readClass(className); err == nil{
		return data, entry, err
	}
	return this.userClassPath.readClass(className)
}


func (this *Classpath) String() string {
	return this.userClassPath.String()
}



func (this *Classpath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
//	jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib", "*")
	this.bootClassPath = newEntry(jreLibPath)
//	jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir,"lib","ext","*")
	this.extClassPath = newEntry(jreExtPath)
}


func (this *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	this.userClassPath = newEntry(cpOption)
}


func getJreDir(jreOption string) string {
	if jreOption != "" && exits(jreOption){
		return jreOption
	}
	if exits("./jre"){
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != ""{
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder !")
}

func exits(path string)bool {
	if _, err := os.Stat(path); err != nil{
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}


