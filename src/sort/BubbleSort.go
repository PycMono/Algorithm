package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡算法
func BubbleSort() {
	var length = 100000
	var array []int

	for i := 0; i < length; i++ {
		array = append(array, int(rand.Intn(100000)))
	}

	fmt.Println(array)

	beforeTime := time.Now().Unix()

	// 冒泡排序
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j > 0; j-- {
			if array[j] >= array[j-1] {
				continue
			}

			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	endTime := time.Now().Unix()
	fmt.Println(endTime - beforeTime)

	fmt.Println(array)
}
