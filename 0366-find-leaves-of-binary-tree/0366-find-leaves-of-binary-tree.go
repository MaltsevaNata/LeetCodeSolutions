/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    var res [][]int
    for root.Left != nil || root.Right != nil {
        res = append(res, removeLeaves(root, nil))
    }
    res = append(res, []int{root.Val})
    return res
}

func removeLeaves(root *TreeNode, parent *TreeNode) []int {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        if parent != nil {
            if parent.Left == root {
                parent.Left = nil
            } else {
                parent.Right = nil
            }
        } 
        return []int{root.Val}
    }
    var res []int
    res = append(res, removeLeaves(root.Left, root)...)
    res = append(res, removeLeaves(root.Right, root)...)
    return res
}