package main

import (
	"flag"
	"fmt"
	"goJvm/main/ch3/classfile"
	"goJvm/main/ch3/classpath"
	"os"
	"strings"
)

type Cmd struct {
	help      bool
	version   bool
	classPath string
	Xjre      string
	class     string
	args      []string
}

func printUsage() {
	// #TODO 2019/7/14 这个help也太简单了吧，基本没啥帮助，写一个更好的help
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func parseCmd() *Cmd {
	cmd := Cmd{}
	flag.Usage = printUsage
	// #NOTE 2019/7/14 将要注入值的域， 对应的命令行参数，默认值，参数说明
	flag.BoolVar(&cmd.help, "help", false, "print help message")
	flag.BoolVar(&cmd.version, "version", false, "print program version")
	flag.StringVar(&cmd.classPath, "cp", "", "classpath")
	flag.StringVar(&cmd.Xjre, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	// #NOTE 2019/7/14 逸出时自动在堆上分配空间，所以不会导致返回的指针失效
	return &cmd
}

func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.Xjre, cmd.classPath)
	fmt.Printf("classpath: %v class: %v args: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
	}
	cf := &classfile.ClassFile{}
	cf, err = cf.Parse(classData)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	cf.PrintClassInfo()
}
