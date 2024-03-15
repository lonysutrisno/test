package hrank

import (
	"fmt"
	"sort"
)

func ThreeSum() [][]int {

	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	var res [][]int
	arrmap := prepare2sum(nums)

	for _, m := range arrmap {
		//loop map key and value
		for _, v := range m.othermember {
			//loop map value
			if m.resultMember[0]+m.resultMember[1]+v == 0 {
				temp := append(m.resultMember, v)
				sort.Ints(temp)
				res = append(res, temp)
			}

		}
	}
	if len(res) > 0 {
		res = distinctSlice(res)
	}
	return res
}

type Maps struct {
	resultMember []int
	othermember  []int
}

func prepare2sum(nums []int) []Maps {
	var arrmap []Maps
	var mymap Maps
	for idxI, i := range nums {
		for j := idxI + 1; j < len(nums); j++ {
			mymap.othermember = removeElementsByIndices(nums, []int{idxI, j})
			mymap.resultMember = []int{i, nums[j]}

			arrmap = append(arrmap, mymap)
		}

	}
	return arrmap
}

func removeElementsByIndices(slice []int, indices []int) []int {

	// Create a map to store indices to be removed
	indicesToRemove := make(map[int]bool)
	for _, idx := range indices {
		indicesToRemove[idx] = true
	}

	// Create a new slice without the elements at the specified indices
	result := make([]int, 0, len(slice)-len(indicesToRemove))
	for idx, value := range slice {
		if !indicesToRemove[idx] {
			result = append(result, value)
		}
	}
	return result
}
func distinctSlice(sliceOfSlice [][]int) [][]int {
	// Create a map to track the occurrence of each slice
	seen := make(map[string]bool)
	var result [][]int

	// Iterate over each slice in the slice of slices
	for _, slice := range sliceOfSlice {
		// Convert the slice to a string representation
		key := fmt.Sprintf("%v", slice)

		// If the slice has not been seen before, add it to the result slice
		if _, ok := seen[key]; !ok {
			result = append(result, slice)
			seen[key] = true // Mark the slice as seen
		}
	}

	return result
}
