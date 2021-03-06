package main

import "fmt"

/*
	go 变量查找过程：
				1.代码引用变量的时候总是最优先查找当前代码块中的那个变量。
				2.若当前代码块没有申明以此为名的变量，name程序会沿着代码块的嵌套关系一层一层网上寻找。
				3.一般情况下，程序会一直找到当前代码包代表的那层代码块。如果任未找到，GO语言编译器将报错
 */

//main包级代码块变量
var block = "package"

func main (){
	//main函数代码块变量
	block := "function"
	{
		//main函数中{}代码块变量
		block := "inner"
		fmt.Printf("The block is %s.\n",block)
	}

	fmt.Printf("The block is %s.\n",block)
}

/*
	重名变量与变量重申明的区别:
	1.变量重声明中的变量一定是在某一个代码块中。
	2.变量申明是对同一个变量的多次申明，这里的变量只有一个。而重名变量中涉及的变量肯定是多个的。
	3.不论对变量重申明多少次，器类型必须一致，具体遵循第一次申明的类型。而重名变量之间则不存在这个限制，它的类型可以使任意的。
	4.如果重名变量说实在的代码块之间存在直接或间接的嵌套关系，那么他们之间一定会存在“屏蔽”的现象。但这种现象绝对不会在变量重申明的场景下出现。
 */
