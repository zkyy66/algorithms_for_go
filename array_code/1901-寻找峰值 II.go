package array_code

//一个 2D 网格中的 峰值 是指那些 严格大于 其相邻格子(上、下、左、右)的元素。
//
//给你一个 从 0 开始编号 的 m x n 矩阵 mat ，其中任意两个相邻格子的值都 不相同 。找出 任意一个 峰值 mat[i][j] 并 返回其位置 [i,j] 。
//
//你可以假设整个矩阵周边环绕着一圈值为 -1 的格子。
//
//要求必须写出时间复杂度为 O(m log(n)) 或 O(n log(m)) 的算法
func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	low, high := 0, m-1
	for low < high {
		i := (low + high) / 2
		j := maxElement(mat[i])
		if i-1 >= 0 && mat[i][j] < mat[i-1][j] {
			high = i - 1
			continue
		}
		if i+1 < m && mat[i][j] > mat[i+1][j] {
			low = i + 1
			continue
		}
		return []int{i, j}
	}
	return nil
}
func maxElement(nums []int) int {
	i := 0
	for j := range nums {
		if nums[i] > nums[j] {
			i = j
		}
	}
	return i
}
