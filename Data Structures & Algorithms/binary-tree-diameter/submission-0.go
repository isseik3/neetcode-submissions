/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		right := height(node.Right)
		if left + right > best {
			best = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	height(root)
	return best
}
