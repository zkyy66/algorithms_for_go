package array_code

//367有效的完全平方数

func isPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		mid := (left + right) / 2
		res := mid * mid
		if res < num {
			left = mid + 1
		} else if res > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
