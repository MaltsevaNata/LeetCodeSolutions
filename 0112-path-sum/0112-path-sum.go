
func countSum(root *TreeNode, sum int, leaves []int) []int {
	if root == nil {
		return leaves
	}
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		leaves = append(leaves, sum)
		return leaves
	}

	leaves = countSum(root.Left, sum, leaves)

	leaves = countSum(root.Right, sum, leaves)
	return leaves
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	leaves := countSum(root, 0, []int{})
	for _, leaf := range leaves {
		if leaf == targetSum {
			return true
		}
	}
	return false
}