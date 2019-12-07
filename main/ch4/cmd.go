package main

import (
	"flag"
	"fmt"
	"goJvm/main/ch4/jvm/rtda"
	"os"
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

func TestLocalVar() {
	local := rtda.NewLocalVar(100)
	local.SetInt(0, 1)
	local.SetLong(1, int64(1)<<62-1)
	local.SetLong(4, -(int64(1)<<62 - 1))
	local.SetFloat(6, 0.123313423)
	local.SetFloat(7, -0.324223)
	local.SetDouble(8, 0.3245345435)
	local.SetDouble(10, -0.3242453)
	local.SetRef(11, nil)
	fmt.Println(local.GetInt(0))
	fmt.Println(local.GetLong(1))
	fmt.Println(local.GetLong(4))
	fmt.Println(local.GetFloat(6))
	fmt.Println(local.GetFloat(7))
	fmt.Println(local.GetDouble(8))
	fmt.Println(local.GetDouble(10))
	fmt.Println(local.GetRef(11))
}

func TestOperandStack() {
	stack := rtda.NewOperandStack(100)
	stack.PushInt(1)
	stack.PushLong(int64(1)<<62 - 1)
	stack.PushLong(-(int64(1)<<62 - 1))
	stack.PushFloat(0.123313423)
	stack.PushFloat(-0.324223)
	stack.PushDouble(0.3245345435)
	stack.PushDouble(-0.3242453)
	stack.PushRef(nil)
	fmt.Println(stack.PopRef())
	fmt.Println(stack.PopDouble())
	fmt.Println(stack.PopDouble())
	fmt.Println(stack.PopFloat())
	fmt.Println(stack.PopFloat())
	fmt.Println(stack.PopLong())
	fmt.Println(stack.PopLong())
	fmt.Println(stack.PopInt())
}

func startJvm(cmd *Cmd) {
	fmt.Println("Test Local Var")
	TestLocalVar()
	fmt.Println("Test Operand Stack")
	TestOperandStack()
}
