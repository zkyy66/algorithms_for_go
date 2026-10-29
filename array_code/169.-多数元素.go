package array_code

//摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。

// 算法可以分为两个阶段：

// 对抗阶段：分属两个候选人的票数进行两两对抗抵消
// 计数阶段：计算对抗结果中最后留下的候选人票数是否有效

// 作者：江不知
// 链接：https://leetcode.cn/problems/majority-element/solutions/154941/tu-jie-mo-er-tou-piao-fa-python-go-by-jalan/

func majorityElement(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	major := 0
	count := 0
	for _, val := range nums {
		if count == 0 {
			major = val
		}
		if major == val {
			count += 1
		} else {
			count -= 1
		}
	}
	return major
}

func majorityElementV(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	if n < 2 {
		return nums[0]
	}
	hash := make(map[int]int)
	for i := 0; i < n; i++ {
		hash[nums[i]]++
		count, _ := hash[nums[i]]
		if count > n/2 {
			return nums[i]
		}
	}
	return 0
}
