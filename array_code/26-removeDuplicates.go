package array_code

/**
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
*/
//如果数组 nums 的长度为 0，则数组不包含任何元素，因此返回 0。
//当数组 nums 的长度大于 0 时，数组中至少包含一个元素，在删除重复元素之后也至少剩下一个元素，因此 nums[0] 保持原状即可，从下标 1开始删除重复元素。
//定义两个指针 fast 和 slow 分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 1
//假设数组 nums 的长度为 n。将快指针 fast 依次遍历从 1 到 n-1 的每个位置，对于每个位置，如果nums[fast]不等于nums
//[fast−1],说明 nums[fast] 和之前的元素都不同，因此将 nums[fast] 的值复制到 nums[slow]，然后将 slow 的值加 1，即指向下一个位置。
//遍历结束之后，从 nums[0] 到 nums[slow−1] 的每个元素都不相同且包含原数组中的每个不同的元素，因此新的长度即为 slow，返回 slow 即可。

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solution/shan-chu-pai-xu-shu-zu-zhong-de-zhong-fu-tudo/
func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

//https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solution/gui-zi-mo-tu-jie-shuang-zhi-zhen-duo-yu-9f1hj/
func RemoveDuTwo(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	}
	index := 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

func RemoveDuThree(nums []int) (int, []int) {
	numARR := []int{1, 1, 2, 2, 3, 3, 4, 4}
	n := len(numARR)
	if n <= 0 {
		return 0, nil
	}
	seen := make(map[int]bool)
	result := make([]int, 0)
	
	for _, num := range numARR {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return len(result), result
}

func Kuaimanzhizhen(nums []int) (int, []int) {
	n := len(nums)
	if n == 0 {
		return 0, nil
	}
	numArr := []int{1, 2, 2, 3, 3, 4, 4, 5}
	length := len(numArr)
	//如果slow默认值为0，那么for循环里面的slow++需要在numArr[slow] = numArr[fast]之前，并且返回值修改为slow+1, numArr[:slow+1]
	slow := 1
	for fast := 1; fast < length; fast++ {
		if numArr[fast] != numArr[fast-1] {
			numArr[slow] = numArr[fast]
			slow++
		}
	}
	return slow, numArr[:slow]
}
