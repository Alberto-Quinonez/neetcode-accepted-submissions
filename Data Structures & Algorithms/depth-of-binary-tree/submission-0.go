/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    return trav(root)
}

func trav(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// we know the current node exists, so contributes at least 1 depth
	return 1 + max(trav(node.Left), trav(node.Right))
}
