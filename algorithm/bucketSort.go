package main

import (
	"fmt"
)

func bucketSort(arr []int, bucketSize int) []int {

	length := len(arr)
	if length == 0 {
		return arr
	}
	max := arr[0]
	min := arr[0]

	for i := 0; i < length; i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}

	bucketCount := (max - min) / bucketSize + 1
	buckets := make([][]int,bucketCount)

	// 利用映射函数将数据分配到各个桶中
	for i := 0; i < length; i++ {
		index := (arr[i] - min) / bucketSize
		buckets[index] = append(buckets[index],arr[i])
	}

	arrIndex := 0
	for _,bucket := range buckets {

		if len(bucket) <= 0 {
			continue
		}
		bucket = insertSort(bucket)

		for _,v := range bucket {
			arr[arrIndex] = v
			arrIndex++
		}

	}

	return arr
}

func insertSort(arr []int) []int  {
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
	fmt.Print("排序前")
	fmt.Println(testArr)
	fmt.Print("排序后")
	fmt.Println(bucketSort(testArr,5))
}
