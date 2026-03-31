/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// -1 means the tree is unbalanced
func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := dfs(root.Left)
	rHeight := dfs(root.Right)
	// check for sentinel
	if lHeight == -1 || rHeight == -1 {
		return -1
	}
	diff := lHeight - rHeight
	if diff > 1 || diff < -1 {
		return -1
	}
	return 1 + max(lHeight, rHeight)
}