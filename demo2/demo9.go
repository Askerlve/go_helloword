package main

import (
	"io"
	"os"
	"fmt"
)

func main()  {
	var err error
	//对err进行了从新申明
	n,err := io.WriteString(os.Stdout,"hello,everyone!\n")
	if err != nil {
		fmt.Printf("error:%v\n",err)
	}
	fmt.Printf("%d byte(s) were written.\n",n)
}
