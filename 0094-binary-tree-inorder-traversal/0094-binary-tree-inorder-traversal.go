/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res = make([]int, 0, 100)
    nodeInorderTraversal(root, &res)
    return res
}


func nodeInorderTraversal(node *TreeNode, res *[]int)  {
    if node == nil {
        return
    }
    nodeInorderTraversal(node.Left, res)
    *res = append(*res, node.Val)
    nodeInorderTraversal(node.Right, res)
    return 
}