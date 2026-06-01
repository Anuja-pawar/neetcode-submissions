# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        def dfs(node, maxValSoFar):
            if not node:
                return 0
            res = 0
            if node.val >= maxValSoFar:
                res = 1
                maxValSoFar = node.val
            
            res += dfs(node.left, maxValSoFar)
            res += dfs(node.right, maxValSoFar)

            return res
    
        return dfs(root, root.val)
        