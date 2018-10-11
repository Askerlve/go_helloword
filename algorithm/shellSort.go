package main

import (
	"fmt"
	"reflect"
)

type shellSort struct {

}

func (shellSort shellSort) sort(arr []int) []int  {
	length := len(arr)
	gap := 1
	for gap < length {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			if j != i-gap {
				arr[j+gap] = temp
			}
		}
		gap = gap / 3
	}
	return arr
}

func main() {
	var testArr []int = []int{21,32,21,45,5,4,87,9,1,36}
	var sort *shellSort = new(shellSort)
	fmt.Println(reflect.TypeOf(sort))
	fmt.Println(sort.sort(testArr))
	fmt.Println(testArr)
}
