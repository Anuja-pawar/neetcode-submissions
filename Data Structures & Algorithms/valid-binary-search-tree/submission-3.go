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
func isValidate(node *TreeNode, min *int, max *int) bool{
    if node == nil{
        return true
    }
    val := node.Val
    if min != nil && node.Val <= *min {
        return false
    }
    
    if max != nil && node.Val >= *max {
        return false
    }
    return isValidate(node.Left, min, &val) && isValidate(node.Right, &val, max)
}

