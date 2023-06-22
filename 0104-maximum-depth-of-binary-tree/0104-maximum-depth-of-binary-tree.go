/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } 
    var leftDepth, rightDepth int
    max := 1
    if root.Left != nil {
        leftDepth = maxDepth(root.Left)
    }
    if root.Right != nil {
        rightDepth = maxDepth(root.Right)
    }
    if leftDepth > rightDepth {
        max += leftDepth
    } else {
        max += rightDepth
    }
    return max
}
