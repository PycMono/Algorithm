package sort

import (
	"fmt"
	"math/rand"
	"time"
)

//原理
//直接插入排序，也是一种非常简单的排序算法。
//
//第一轮先从第二个元素开始，和第一个比较，如果较小就交换位置，本轮结束。第二轮从第三个元素开始，先与第二个比较，如果较小就与第二个交换，交换后再于第一个比较。如此循环直至最后一个元素完成比较逻辑。
//
//
//
//复杂度
//最好的情况下，直接插入排序只需进行n-1次比较，0次的交换。平均下来时间复杂度为 O(n^2)。
//
//由于是每个元素逐个与有序的队列进行比较，所以不会出现相同数值的元素在排序完成后交换位置。所以直接插入排序是种稳定的排序算法。

// 插入排序
func InsrtSort() {
	var length = 100000
	var array []int

	for i := 0; i < length; i++ {
		array = append(array, int(rand.Intn(100000)))
	}

	fmt.Println(array)

	beforeTime := time.Now().Unix()

	// 插入排序
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			// 加快排序
			if array[j] >= array[j-1] {
				break
			}

			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	endTime := time.Now().Unix()
	fmt.Println(endTime - beforeTime)

	fmt.Println(array)
}

// 插入排序
func InsrtSort1() {
	var length = 10
	var array []int

	for i := 0; i < length; i++ {
		array = append(array, int(rand.Intn(500)))
	}

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if array[j-1] <= array[j] {
				break
			}

			array[j-1], array[j] = array[j], array[j-1]
		}
	}

	fmt.Println(array)
}
