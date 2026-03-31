/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var queue []*TreeNode
	queue = append(queue,root)
	// continue until queue is empty
	for len(queue) > 0 {
		// process all all values in queue
		var innerResult []int
		levelSize := len(queue)
		for i:=0;i<levelSize;i++ {
			// pop
			node := queue[0]
			queue = queue[1:]
			// only add leaf nodes to result
			if node != nil {
				innerResult = append(innerResult, node.Val)
				// push children to back of queue
				if node.Left != nil {
					queue = append(queue,node.Left)
				}
				if node.Right != nil {
					queue = append(queue,node.Right)
				}
			}
		}
		// append new innerResult to existing result once level is done processing
		if len(innerResult) > 0 {
			result = append(result, innerResult)
		}
	}
	return result
}