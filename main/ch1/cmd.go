package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	help      bool
	version   bool
	classPath string
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
	flag.StringVar(&cmd.classPath, "cp", "", "classpah")
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
}
