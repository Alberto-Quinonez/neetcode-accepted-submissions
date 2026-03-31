/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min,max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	// left side needs to be less than current root
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}