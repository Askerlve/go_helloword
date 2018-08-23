package main

import (
	"flag"
	"fmt"
)

func main(){
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("hello,%v!\n",*name)
}


func getTheFlag() *string {
	return flag.String("name","everyone","oh,shit!")
}