/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidNode(root, math.Inf(1), math.Inf(-1))
}

func isValidNode(node *TreeNode, lessThan float64, largerThan float64) bool {
	if node == nil {
		return true
	}
	val := float64(node.Val)
	if val >= lessThan || val <= largerThan {
		return false
	}
	least := lessThan
	if val < lessThan {
		least = val
	}
	largest := largerThan
	if val > largerThan {
		largest = val
	}
	return isValidNode(node.Left, least, largerThan) &&         isValidNode(node.Right, lessThan, largest)
}