package util

import (
	"fmt"
	"log"
)

// #TODO 2019/7/14 使用反射提供打印的类的信息

func PanicError(err error)bool  {
	if err != nil {
		log.Panic(err)
		return true
	}
	return false
}


func FatalError(err error) bool  {
	if err != nil {
		log.Fatal(err)
		return true
	}
	return false
}

func PrintError(err error)bool{
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
