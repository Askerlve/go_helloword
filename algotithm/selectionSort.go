package main

import (
	"fmt"
	"reflect"
)

type selectionSort struct {

}

func (selectionSort selectionSort) sort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}

	}
	return arr
}

func main() {
	var testArr []int = []int{21,32,21,45,5,4,87,9,1,36}
	var sort *selectionSort = new(selectionSort)
	fmt.Println(reflect.TypeOf(sort))
	fmt.Println(sort.sort(testArr))
	fmt.Println(testArr)
}
