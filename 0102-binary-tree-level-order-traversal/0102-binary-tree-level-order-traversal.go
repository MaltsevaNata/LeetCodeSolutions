/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func nodeLevelTraversal(result [][]int, level int, root *TreeNode) [][]int {
	if root != nil {
		if len(result) < (level + 1) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		result = nodeLevelTraversal(result, level+1, root.Left)
		result = nodeLevelTraversal(result, level+1, root.Right)
	}
	return result
}

func levelOrder(root *TreeNode) (result [][]int) {
	result = nodeLevelTraversal(result, 0, root)
	return result
}