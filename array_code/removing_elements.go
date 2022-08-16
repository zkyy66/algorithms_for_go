package array_code

import "fmt"

// 数组移除元素
func removingElements(nums []int, val int) int {
	elementCount := len(nums)
	if elementCount <= 0 {
		return 0
	}
	slowPoint := 0
	for fastPoint := 0; fastPoint <= elementCount; fastPoint++ {
		if nums[fastPoint] != val {
			nums[slowPoint] = nums[fastPoint]
			slowPoint++
		}
	}
	fmt.Println(nums)
	return slowPoint
}
