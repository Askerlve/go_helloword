package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question",flag.ExitOnError)

func init()  {
	/**
	flag.ExitOnError的含义是，当命令后跟--help或者参数设置的不正确的时候，在打印命令参数使用说明后以状态码2结束当前程序.
	状态码2代表用户错误地使用了命令,而flag.PanicOnError与之的区别是在最后抛出“运行时恐慌（panic）"
	 */
	flag.CommandLine = flag.NewFlagSet("",flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
		flag.PrintDefaults()
	}
	//flag.StringVar(&name,"name","everyone","The greeting object.")
	cmdLine.StringVar(&name,"name","everyone","The greeting object.")
}

func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
