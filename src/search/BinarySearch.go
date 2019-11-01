package search

import (
	"fmt"
)

func Test() {
	//array := make([]int, 0, 1000000000)
	//for i := 0; i < 1000000000; i++ {
	//	array = append(array, i)
	//}
	//
	//beforeTime := time.Now().Unix()
	//fmt.Println(binarySearch(array, 8))
	//fmt.Println(binarySearch(array, 1000000000-1))
	//endTime := time.Now().Unix()
	//fmt.Println(fmt.Sprintf("二分查找所用的时间%d", endTime-beforeTime))
}

// 二分查找
// 参数：
// array：源数据
// value：目标数据
func binarySearch(array []int, value int) int {
	low := 0
	high := len(array) - 1
	mid := 0
	num := 0
	for {
		// 如果数据未找到退出循环
		if low > high {
			fmt.Println(fmt.Sprintf("二分查找所用的次数为%d", num))
			return -1
		}

		mid = low + (high-low)/2
		if array[mid] == value {
			fmt.Println(fmt.Sprintf("二分查找所用的次数为%d", num))
			return mid
		} else if array[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}

		num++
	}
}
