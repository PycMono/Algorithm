package main

import (
	"fmt"
	"moqikaka.com/Algorithm/src/search"
	"moqikaka.com/Algorithm/src/sum"
	"moqikaka.com/Algorithm/src/tree"

	"math"
	"moqikaka.com/algorithm/src/sort"
)

func main() {
	fmt.Println("------------------------插入排序结果------------------------")
	sort.InsrtSort1()

	fmt.Println("------------------------冒泡排序结果------------------------")
	sort.BubbleSort1()

	fmt.Println("------------------------两数之和------------------------")
	sum.TwoSum()

	//NJVIhHWPQWHIMIMUJqqLFNzULWGC
	// "NJVIhHWPQWHIMIMUJqqLFNzULWGC"
	str := "GCNJVIhHWPQWHIMIMUJqqLFNzULW"
	left := 3063
	right := 9697
	offest := right - left
	offest = offest % len(str)
	//fmt.Println(offest % len(str))
	//fmt.Println(len(str))
	if offest > 0 {
		str1 := str[len(str)-offest:]
		str2 := str[:len(str)-offest]
		str = str1 + str2
	} else if offest < 0 {
		offest = int(math.Abs(float64(offest)))
		str1 := str[:offest]
		str2 := str[offest:]
		str = str2 + str1
	}

	fmt.Println("------------------------二分查找------------------------")
	search.Test()

	fmt.Println("------------------------归并排序------------------------")
	sort.Test()

	fmt.Println("------------------------二叉树------------------------")
	tree.Test()
}
