package binary_tree

func postorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		result = append(result, root.Val)
	}
	traversal(root)
	return result
}
