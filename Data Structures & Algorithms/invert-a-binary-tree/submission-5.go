/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    stack := []*TreeNode{root}

    for len(stack) > 0{
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        curr.Left, curr.Right = curr.Right, curr.Left

        if curr.Left != nil{
            stack = append(stack, curr.Left)
        }
        if curr.Right != nil{
            stack = append(stack, curr.Right)
        }

    }

    return root
}
