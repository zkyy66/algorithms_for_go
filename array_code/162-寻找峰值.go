package array_code

import (
	"math"
	"math/rand"
)

//给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//由于题目保证了 nums[i] !=nums[i+1]，那么数组 nums 中最大值两侧的元素一定严格小于最大值本身。因此，最大值所在的位置就是一个可行的峰值位置。
//我们对数组 nums 进行一次遍历，找到最大值对应的位置即可。
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/find-peak-element/solutions/998152/xun-zhao-feng-zhi-by-leetcode-solution-96sj/
//时间复杂度：O(n)，其中 n 是数组 nums 的长度。
//
//空间复杂度：O(1)。
func findPeakElement(nums []int) int {
	idx := 0
	for key, val := range nums {
		if val > nums[idx] {
			idx = key
		}
	}
	return idx
}

//我们首先在 [0,n) 的范围内随机一个初始位置 i，随后根据 nums[i−1],nums[i],nums[i+1] 三者的关系决定向哪个方向走：
//
//如果 nums[i−1]<nums[i]>nums[i+1]，那么位置 i 就是峰值位置，我们可以直接返回 i 作为答案；
//
//如果 nums[i−1]<nums[i]<nums[i+1]，那么位置 i 处于上坡，我们需要往右走，即 i←i+1；
//
//如果 nums[i−1]>nums[i]>nums[i+1]，那么位置 i 处于下坡，我们需要往左走，即 i←i−1；
//
//如果 nums[i−1]>nums[i]<nums[i+1]，那么位置 i 位于山谷，两侧都是上坡，我们可以朝任意方向走。
//
//如果我们规定对于最后一种情况往右走，那么当位置 i 不是峰值位置时：
//
//如果 nums[i]<nums[i+1]，那么我们往右走；
//
//如果 nums[i]>nums[i+1]，那么我们往左走。
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/find-peak-element/solutions/998152/xun-zhao-feng-zhi-by-leetcode-solution-96sj/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func findPeakElementTwo(nums []int) int {
	n := len(nums)
	idx := rand.Intn(n)
	
	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	
	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}
	
	return idx
}

//如果 nums[i]<nums[i+1]，并且我们从位置 i 向右走到了位置 i+1，那么位置 i 左侧的所有位置是不可能在后续的迭代中走到的。
//
//这是因为我们每次向左或向右移动一个位置，要想「折返」到位置 i 以及其左侧的位置，我们首先需要在位置 i+1 向左走到位置 i，但这是不可能的。
//
//并且根据方法二，我们知道位置 i+1 以及其右侧的位置中一定有一个峰值，因此我们可以设计出如下的一个算法：
//
//对于当前可行的下标范围 [l,r]，我们随机一个下标 i；
//
//如果下标 i 是峰值，我们返回 i 作为答案；
//
//如果 nums[i]<nums[i+1]，那么我们抛弃 [l,i] 的范围，在剩余 [i+1,r] 的范围内继续随机选取下标；
//
//如果 nums[i]>nums[i+1]，那么我们抛弃 [i,r] 的范围，在剩余 [l,i−1] 的范围内继续随机选取下标。
//
//在上述算法中，如果我们固定选取 i 为 [l,r] 的中点，那么每次可行的下标范围会减少一半，成为一个类似二分查找的方法，时间复杂度为 O(logn)。
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/find-peak-element/solutions/998152/xun-zhao-feng-zhi-by-leetcode-solution-96sj/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func findPeakElementThree(nums []int) int {
	n := len(nums)
	
	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	
	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

//https://leetcode.cn/problems/find-peak-element/solutions/998152/xun-zhao-feng-zhi-by-leetcode-solution-96sj/
