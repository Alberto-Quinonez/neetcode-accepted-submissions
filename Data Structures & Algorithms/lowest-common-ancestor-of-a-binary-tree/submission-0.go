/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// found either
	if root == p || root == q {
		return root
	}
	// explore left
	left := lowestCommonAncestor(root.Left,p,q)
	// explore right
	right := lowestCommonAncestor(root.Right,p,q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}



// func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if p.Val < root.Val && q.Val < root.Val {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	} else if p.Val > root.Val && q.Val > root.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	return root
// }