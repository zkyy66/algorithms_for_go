/**
 * @Date 2022/7/2
 * @Name 209-minimumSizeSubarraySum
 * @VariableName
**/
//https://programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html#%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3
//窗口就是 满足其和 ≥ s 的长度最小的 连续 子数组
//窗口的起始位置如何移动：如果当前窗口的值大于s了，窗口就要向前移动了（也就是该缩小了）。
//窗口的结束位置如何移动：窗口的结束位置就是遍历数组的指针，窗口的起始位置设置为数组的起始位置就可以了。
//https://programmercarl.com/%E6%95%B0%E7%BB%84%E6%80%BB%E7%BB%93%E7%AF%87.html#%E5%8F%8C%E6%8C%87%E9%92%88%E6%B3%95
package array_code

func minSubArrayLen(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	index := 0
	sum := 0
	result := n + 1
	for j := 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - index + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[index]
			index++
		}
	}
	if result == n+1 {
		return 0
	} else {
		return result
	}
}
