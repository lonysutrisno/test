package hrank

func RemoveDuplicates() int {
	nums := []int{1, 2, 2, 2, 3, 3, 4, 4, 5}
	var res int
	for idx, v := range nums {
		if idx+1 < len(nums) {
			if v == nums[idx+1] {
				newSlice := append(nums[:idx], nums[idx+1:]...)
				nums = newSlice

			} else {
				res++
			}
		}
	}
	res++
	return res
}
