package main

import (
	"fmt"
	"reflect"
)

type insertionSort struct {

}

func (insertionSort insertionSort) sort(arr []int) []int  {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		if i != preIndex + 1 {
			arr[preIndex+1] = current
		}
	}
	return arr
}

func main() {
	var testArr []int = []int{21,32,21,45,5,4,87,9,1,36}
	var sort *insertionSort = new(insertionSort)
	fmt.Println(reflect.TypeOf(sort))
	fmt.Println(sort.sort(testArr))
	fmt.Println(testArr)
}
