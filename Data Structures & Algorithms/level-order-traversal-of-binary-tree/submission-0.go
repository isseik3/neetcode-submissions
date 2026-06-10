/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
	if root == nil {
		return ans
	}
	// BFS
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		elements := []int{}
		for i := 0; i < size; i++ {
			node := queue[0]
			elements = append(elements, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, elements)
	}
	return ans

}
