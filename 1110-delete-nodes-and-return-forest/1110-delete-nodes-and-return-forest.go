/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    var del = make(map[int]struct{}) // {3, 5}
    
    for _, val := range to_delete {
        del[val] = struct{}{}
    }
    
    return dfs(root, del)
}

func dfs(root *TreeNode, del map[int]struct{}) []*TreeNode { //{1}
    var res []*TreeNode
    if root == nil { // 1
        return res
    }
    if _, ok := del[root.Val]; ok {
        if root.Left != nil {
            res = append(res, dfs(root.Left, del)...) 
        }
        if root.Right != nil {
            res = append(res, dfs(root.Right, del)...)
        }
    } else {
        res = append(res, root)
        res = add(res, root, root.Left, del) // {1}, 2
        res = add(res, root, root.Right, del) // {1}, 3
    }
    return res
}

func add(res []*TreeNode, parent *TreeNode, root *TreeNode, del map[int]struct{}) []*TreeNode { // {1}, 3
    if root == nil {
        return res
    }
    if _, ok := del[root.Val]; ok {
        if root.Left != nil {
            res = append(res, dfs(root.Left, del)...) 
        }
        if root.Right != nil {
            res = append(res, dfs(root.Right, del)...)
        }
        if parent.Right == root {
            parent.Right = nil
        } else {
            parent.Left = nil
        }
    } else {
        res = add(res, root, root.Left, del)
        res = add(res, root, root.Right, del)
    }
    return res
}