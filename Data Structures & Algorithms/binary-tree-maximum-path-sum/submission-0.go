/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxPath := root.Val
    _ = dfs(root, &maxPath)

    return maxPath

}
func dfs(node *TreeNode, maxPath *int) int{
    if node == nil{
        return 0
    }
    leftMax := dfs(node.Left, maxPath)
    rightMax := dfs(node.Right, maxPath)
    leftMax = max(leftMax, 0)
    rightMax = max(rightMax, 0)

    *maxPath = max(*maxPath, node.Val + leftMax + rightMax)

    return node.Val + max(leftMax, rightMax)
}
