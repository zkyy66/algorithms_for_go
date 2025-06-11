package binary_tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		stack = append(stack, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return stack
}

func preorderTraversalForStack(root *TreeNode) []int {
	return nil
}
