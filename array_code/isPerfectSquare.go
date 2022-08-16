package array_code

import "math"

//有效的完全平方数

func isPerfectSquarevers(num int) bool {
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

// 通过暴力循环法
func versionTwo(num int) bool {
	for x := 1; x*x <= num; x++ {
		if x*x == num {
			return true
		}
	}
	return false
}

// 内置函数库解决
func versionOne(num int) bool {
	x := int(math.Sqrt(float64(num)))
	return x*x == num

}
