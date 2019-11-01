package sum

import (
	"fmt"
	"testing"
)

func TwoSum() {
	array := make([]int, 10)
	tempMap := make(map[int]int, 10)

	array[0] = 2
	array[1] = 7
	array[2] = 5
	array[3] = 4
	array[4] = 10
	for _, item := range array {
		tempMap[item] = 0
	}

	for index, item := range array {
		result := 9 - item
		if _, exists := tempMap[result]; exists {
			fmt.Println(index)
		}
	}
}

func TestName(t *testing.T) {
	TwoSum()
}
