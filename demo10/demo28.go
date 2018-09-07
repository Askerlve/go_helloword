package main

import (
	"errors"
	"fmt"
)

type Operate func(x,y int) int

//方案一
func calculate(x int,y int,op Operate)(int,error){
	if(op == nil){
		return 0,errors.New("param error")
	}
	return op(x,y),nil
}

//方案二
type CalculateFunc func(x int,y int)(int,error)

//闭包
func genCalculator(op Operate) CalculateFunc  {
	return func(x int, y int) (int, error) {
		if(op == nil) {
			return 0,errors.New("param error")
		}
		return op(x,y),nil
	}
}

func main(){
	//方案一
	x,y := 12,23
	op := func(x,y int) int {
		return x + y
	}
	result,err := calculate(x,y,op)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	//方案二
	x,y = 1,2
	add := genCalculator(op)
	result,err = add(x,y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}
