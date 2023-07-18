/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    var result int
    if root.Val >= low && root.Val<=high {
        result += root.Val
    }
    result += rangeSumBST(root.Left, low, high)
    result += rangeSumBST(root.Right, low, high)
    return result
}