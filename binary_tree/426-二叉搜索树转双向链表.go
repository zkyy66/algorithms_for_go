package binary_tree

/**
 *
 * @param pRootOfTree TreeNode类
 * @return TreeNode类
 */
func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	var f func(root *TreeNode)
	var pre *TreeNode = nil
	var head *TreeNode = nil
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		if pre == nil {
			head = root
		} else {
			root.Left, pre.Right = pre, root
		}
		pre = root
		f(root.Right)
	}
	f(pRootOfTree)
	return head
}
