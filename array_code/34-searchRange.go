package array_code

//34
import "sort"

//https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	leftBoundary := leftBoundary(nums, target)
	rightBoundary := rightBoundary(nums, target)
	if leftBoundary == -2 || rightBoundary == -2 {
		return []int{-1, -1}
	}
	if rightBoundary-leftBoundary > 1 {
		return []int{leftBoundary + 1, rightBoundary - 1}
	}
	return []int{-1, -1}
}

//确定左边界
func leftBoundary(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	elementOffset := -2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
			elementOffset = right
		} else {
			left = mid + 1
		}
	}
	return elementOffset
}

//确定右边界
func rightBoundary(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	elementOffset := -2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
			elementOffset = left
		}
	}
	return elementOffset
}

func searchRangeS(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
