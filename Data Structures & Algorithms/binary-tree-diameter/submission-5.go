/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root*TreeNode,counter *int)int{
    if root==nil{return 0}
    l:=helper(root.Left,counter)
    r:=helper(root.Right,counter)
    *counter = max(l+r,*counter)

    return 1+max(l,r)
}
func diameterOfBinaryTree(root *TreeNode) int {
    counter :=0
    helper(root,&counter)
    return counter
    
}
