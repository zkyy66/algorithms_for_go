/**
 * @Date 2022/8/16
 * @Name 209-minimum-size-subarray-sum
 * @VariableName
**/
package array_code

func minSubArrayLen(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	i := 0
	sum := 0
	result := lens + 1
	for j := 0; j < lens; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == lens+1 {
		return 0
	} else {
		return result
	}
}
