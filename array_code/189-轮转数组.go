package array_code

import "slices"

func rotate(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
