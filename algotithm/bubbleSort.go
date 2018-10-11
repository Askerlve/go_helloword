package main

import (
	"fmt"
	"reflect"
)

type bubbleSort struct {

}

func (bubbleSort bubbleSort) sort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var testArr []int = []int{21,32,21,45,5,4,87,9,1,36}
	var sort *bubbleSort = new(bubbleSort)
	fmt.Println(reflect.TypeOf(sort))
	fmt.Println(sort.sort(testArr))
	fmt.Println(testArr)
}
