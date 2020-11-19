package top150

/**
* @author fan_mei
* @date 2020/11/19 17:33

需要先确定mid 左/右 侧是否被截断过
如果没有 则可以继续根据单调性进行二分
如果有 那么可以判断相反方向的条件 结果就在另外半段上

**/

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	left, right := 0, length-1

	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		}

		//left ~ mid 是否被截断
		if nums[left] <= nums[mid] {
			if nums[mid] >= target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
