package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			for j := mid; j < len(nums) && nums[j] <= target; j++ {
				right = j
			}
			for j := mid; j >= 0 && nums[j] == target; j-- {
				left = j
			}
			break
		}
	}
	return right - left + 1
}
