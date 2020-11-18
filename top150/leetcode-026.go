package top150

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastIndex := 0
	res := 1

	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur == nums[lastIndex] {
			continue
		} else {
			lastIndex++
			nums[lastIndex] = cur
			res++
		}
	}
	return res
}
