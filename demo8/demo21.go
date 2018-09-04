package main

import "fmt"

/**
对于一个通道来说：
1.发送操作之间是互斥的，接收操作之间也是互斥的。
2.发送操作和接收操作对元素的操作都是不可切分的。
3.发送操作在完全完成之前会被阻塞，接收操作也是如此。
 */
func main()  {
	ch1 := make(chan int,3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <- ch1
	elem2 := <- ch1

	fmt.Printf("the frist element received from channel ch1:%v\n",elem1)
	fmt.Printf("the second element received from channel ch1:%v\n",elem2)
	fmt.Printf("the third element received from channel ch1:%v\n",<-ch1)
}
