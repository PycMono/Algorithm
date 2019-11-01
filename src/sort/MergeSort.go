package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func Test() {
	array := make([]int, 0, 6)
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)

	for i := 10; i > 0; i-- {
		array = append(array, r1.Intn(100))
	}

	tempArray := sort(array)
	for _, v := range tempArray {
		fmt.Print(fmt.Sprintf("%d,", v))
	}
}

// 归并排序
func sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	left := sort(array[mid:])
	right := sort(array[:mid])
	fmt.Println(fmt.Sprintf("left_len=%d,right_len=%d,mid=%d", len(left), len(right), mid))
	return merge(left, right)
}

// 排序合并
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
