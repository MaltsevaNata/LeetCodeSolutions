/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    var parents = make(map[int]*TreeNode) // {1: nil}
	parents[root.Val] = nil

	findParents(root, parents) // root
    var visited = make(map[int]struct{})
	return findNodes(target, k, parents, visited)

}

func findNodes(root *TreeNode, k int, parents map[int]*TreeNode, visited map[int]struct{}) []int { // 3, 1, {1: nil, 2: 1, 3: 1}, {2, 1, 3}
    if root == nil {
        return []int{}
    }
    if _, ok := visited[root.Val]; ok {
        return []int{}
    }
    if k == 0 {
        return []int{root.Val}
    }
    if root.Left == nil && root.Right == nil && parents[root.Val] == nil {
        return []int{}
    }

    var res []int
    visited[root.Val] = struct{}{}
    res = append(res, findNodes(root.Left, k - 1, parents, visited)...) // []
    res = append(res, findNodes(root.Right, k - 1, parents, visited)...) // []
    res = append(res, findNodes(parents[root.Val], k - 1, parents, visited)...)// []
    return res
}


func findParents(root *TreeNode, parents map[int]*TreeNode) { // 1, {1: nil}, 1
	if root == nil {
        return 
    }
    if root.Left != nil {
        parents[root.Left.Val] = root
        findParents(root.Left, parents)
    }
    if root.Right != nil {
        parents[root.Right.Val] = root
        findParents(root.Right, parents)
    }

    return 
}
