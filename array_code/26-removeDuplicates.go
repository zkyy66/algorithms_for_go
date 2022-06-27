package array_code

//如果数组 \textit{nums}nums 的长度为 00，则数组不包含任何元素，因此返回 00。
//
//当数组 \textit{nums}nums 的长度大于 00 时，数组中至少包含一个元素，在删除重复元素之后也至少剩下一个元素，因此 \textit{nums}[0]nums[0] 保持原状即可，从下标 11 开始删除重复元素。
//
//定义两个指针 \textit{fast}fast 和 \textit{slow}slow 分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 11。
//
//假设数组 \textit{nums}nums 的长度为 nn。将快指针 \textit{fast}fast 依次遍历从 11 到 n-1n−1 的每个位置，对于每个位置，如果 \textit{nums}[\textit{fast}] \ne \textit{nums}[\textit{fast}-1]nums[fast]
//
//​
// =nums[fast−1]，说明 \textit{nums}[\textit{fast}]nums[fast] 和之前的元素都不同，因此将 \textit{nums}[\textit{fast}]nums[fast] 的值复制到 \textit{nums}[\textit{slow}]nums[slow]，然后将 \textit{slow}slow 的值加 11，即指向下一个位置。
//
//遍历结束之后，从 \textit{nums}[0]nums[0] 到 \textit{nums}[\textit{slow}-1]nums[slow−1] 的每个元素都不相同且包含原数组中的每个不同的元素，因此新的长度即为 \textit{slow}slow，返回 \textit{slow}slow 即可。
//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solution/shan-chu-pai-xu-shu-zu-zhong-de-zhong-fu-tudo/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
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
			nums[index+1] = nums[i]
			index++
		}
	}
	return index + 1
}
