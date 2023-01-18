/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func nodeTraversal(result *[]int, root *Node) []int {
	if root != nil {
		res := append(*result, root.Val)
		result = &res
		for _, node := range root.Children {
			if len(node.Children) == 0 {
				res := append(*result, node.Val)
				result = &res
			} else {
				trav := nodeTraversal(result, node)
				result = &trav
			}
		}
	}
	return *result
}

func preorder(root *Node) (result []int) {
	result = nodeTraversal(&result, root)
	return result
}