package hrank

import "fmt"

func Arraysumbiggest() (result int) {
	array := []int{-2, 1, -3, 4, -1, 2, 1, -5, 6}
	// var temp int
	maxSum := array[0]
	currSum := array[0]
	for i := 0; i < len(array); i++ {
		// if i < len(array)-1 {
		currSum = max(array[i], currSum+array[i])
		maxSum = max(maxSum, currSum)

		// }
	}
	// for i := 0; i < len(array); i++ {
	// 	temp = array[i]
	// 	for j := i; j < len(array); j++ {
	// 		if j < len(array)-1 {
	// 			temp = temp + array[j+1]
	// 		}
	// 	}

	// 	if result < temp {
	// 		result = temp
	// 	}
	fmt.Println(maxSum)

	// }

	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
