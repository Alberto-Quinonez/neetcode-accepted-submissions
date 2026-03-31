/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    m := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		// hit leaf, there is no change in max
		if root == nil{
			return 0
		}
		hLeft := dfs(root.Left)
		hRight := dfs(root.Right)
		// update global max
		m = max(m, hLeft + hRight)
		return 1 + max(hLeft,hRight)
	}
	dfs(root)
	return m
}
