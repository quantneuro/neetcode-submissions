/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var dbt func(root *TreeNode)int
    counter:=0
    dbt=func(root *TreeNode)int{
        if root==nil{return 0}
        l:=dbt(root.Left)
        r:=dbt(root.Right)
        counter=max(l+r,counter)

        return 1+max(l,r)

    }
    dbt(root)
    return counter
    
}
