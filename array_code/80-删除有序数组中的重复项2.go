package array_code
/**
通用解法
为了让解法更具有一般性，我们将原问题的「保留 2 位」修改为「保留 k 位」。

对于此类问题，我们应该进行如下考虑：

由于是保留 k 个相同数字，对于前 k 个数字，我们可以直接保留
对于后面的任意数字，能够保留的前提是：与当前写入的位置前面的第 k 个元素进行比较，不相同则保留

举个🌰，我们令 k=2，假设有如下样例

[1,1,1,1,1,1,2,2,2,2,2,2,3]

首先我们先让前 2 位直接保留，得到 1,1
对后面的每一位进行继续遍历，能够保留的前提是与当前位置的前面 k 个元素不同（答案中的第一个 1），因此我们会跳过剩余的 1，将第一个 2 追加，得到 1,1,2
继续这个过程，这时候是和答案中的第 2 个 1 进行对比，因此可以得到 1,1,2,2
这时候和答案中的第 1 个 2 比较，只有与其不同的元素能追加到答案，因此剩余的 2 被跳过，3 被追加到答案：1,1,2,2,3

作者：宫水三叶
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/solutions/702970/gong-shui-san-xie-guan-yu-shan-chu-you-x-glnq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
**/
func removeDuplicatesForTwo(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	var res func(k int) int
	res = func(k int) int {
		u := 0
		for _ v := range nums {
			if u < k || nums[u-k] != v {
				nums[u] = v
				u++
			}
		}
		return u
	}
	return res(2)
}
//双指针
func removeDuplicates(nums []int) int {
	//初始
	stackIndex := 2
	n := len(nums)
	for i:=2;i<n;i++{
		if nums[i] != nums[stackIndex-2] {
			nums[stackIndex] = nums[i]
			stackIndex++
		}
	}
	//return stackIndex
	return min(stackIndex,len(nums))
}
func removeDuplicates(nums []int) int {
	k := 0
	for _,v := range nums {
		if k < 2 ||nums[k-2] != v {
			nums[k] = v
			k++
		}
	}
	return k
}