package top150

/**
* @author fan_mei
* @date 2020/11/19 18:47
**/

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	aim := binarySearch(nums, 0, len(nums), target)
	if aim == -1 {
		return res
	}
	left := aim
	right := aim
	for left > 0 && nums[left-1] == target {
		left = binarySearch(nums, 0, left-1, target)
	}

	for right < len(nums)-1 && nums[right+1] == target {
		right = binarySearch(nums, right+1, len(nums), target)
	}
	res[0] = left
	res[1] = right
	return res
}

func binarySearch(nums []int, left, right, target int) int {
	if len(nums) == 0 {
		return -1
	}
	mid := 0
	for left < right {
		mid = (right-left)>>1 + left

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if len(nums) > left && nums[left] == target {
		return left
	}
	return -1
}
