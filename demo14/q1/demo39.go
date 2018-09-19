package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//以协程的方式运行（异步)
		go func() {
			fmt.Println(i)
		}()
	}
}