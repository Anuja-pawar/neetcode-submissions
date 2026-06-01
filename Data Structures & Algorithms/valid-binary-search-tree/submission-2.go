/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return isValidate(root, nil, nil)
}
func isValidate(node *TreeNode, left *int, right *int) bool{
    if node == nil{
        return true
    }
    val := node.Val
    if left != nil && node.Val <= *left {
        return false
    }
    
    if right != nil && node.Val >= *right {
        return false
    }
    return isValidate(node.Left, left, &val) && isValidate(node.Right, &val, right)
}

