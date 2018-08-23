package main

import (
	"flag"
	"fmt"
)

/**
	定义变量的几种方式
 */
func main(){
	//var name string
	//flag.StringVar(&name,"name","everyone","the greeting object!")

	// 方式1。
	// 类型*string代表的是字符串的指针类型，而不是字符串类型
	// flag.String函数返回的结果值的类型是*string而不是string,若不加*则得到的将是一个字符串指针，通过一个“*”把这个字符串的指针值指向的字符串值取出来
	//var name = *flag.String("name", "everyone", "The greeting object.")

	// 方式2。
	name := *flag.String("name", "everyone", "The greeting object.")
	value := "haha"

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
	fmt.Printf("Hello, %v!\n", value)
}
