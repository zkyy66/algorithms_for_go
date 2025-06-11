package binary_tree

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		result = append(result, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return result
}
