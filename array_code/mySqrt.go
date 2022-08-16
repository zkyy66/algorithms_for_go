package array_code

// 69. x 的平方根
// 二分查找
func mySqrt(x int) int {
	left := 0
	right := x
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
			res = mid
		}
	}
	return res
}
