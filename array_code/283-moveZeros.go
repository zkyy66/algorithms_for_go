/**
 * @Date 2022/6/27
 * @Name 283-moveZeros
 * @VariableName
https://leetcode.cn/problems/move-zeroes/solution/pei-yang-suan-fa-si-wei-cong-bao-li-ceng-c24w/
https://leetcode.cn/problems/move-zeroes/solution/wo-de-fang-fa-bi-guan-fang-jie-fa-geng-j-kdac/
**/
package array_code

func MoveZeros(nums []int) {
	n := len(nums)
	if n <= 0 {
		return
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
		fast++
	}
}
func MoveZerosTwo(nums []int) []int {
	n := len(nums)
	if n <= 0 {
		return nil
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
