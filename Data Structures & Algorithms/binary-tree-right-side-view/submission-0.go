/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}
	ans := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		last := 0
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			last = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, last)
	}
	return ans
}
