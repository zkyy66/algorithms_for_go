package array_code

// 双指针法
// 双指针法（快慢指针法）： 通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。
//
// 定义快慢指针
//
// 快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
// 慢指针：指向更新 新数组下标的位置
// 很多同学这道题目做的很懵，就是不理解 快慢指针究竟都是什么含义，所以一定要明确含义，后面的思路就更容易理解了。
func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}
	res := 0
	for i := 0; i < numsLen; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

func pointsFun(nums []int, val int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	res := 0
	for i := 0; i < lens; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
