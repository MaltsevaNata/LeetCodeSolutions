/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func nodeTraversal(result []int, root *Node) []int {
	if root != nil {
		result = append(result, root.Val)
		for _, node := range root.Children {
			if len(node.Children) == 0 {
				result = append(result, node.Val)
			} else {
				result = nodeTraversal(result, node)
			}
		}
	}
	return result
}

func preorder(root *Node) (result []int) {
	return nodeTraversal(result, root)
}