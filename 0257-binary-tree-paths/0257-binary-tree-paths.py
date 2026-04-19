class Solution:
    def __init__(self):
        self.paths = []

    def traverse(self, root: Optional[TreeNode], result: str) -> None:
        if not root:
            return
        result += str(root.val)
        if not root.left and not root.right:
            self.paths.append(result)
        result += "->"
        self.traverse(root.left, result)
        self.traverse(root.right, result)


    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        self.traverse(root, "")
        return self.paths