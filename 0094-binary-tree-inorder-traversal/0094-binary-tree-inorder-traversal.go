/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) (res []int) {
    return nodeInorderTraversal(root, res)
}

func nodeInorderTraversal(node *TreeNode, res []int) []int {
    if node == nil {
        return res
    }
    res = nodeInorderTraversal(node.Left, res)
    res = append(res, node.Val)
    res = nodeInorderTraversal(node.Right, res)
    return res
}