/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {

    var invert func(root *TreeNode)*TreeNode

    invert= func(root * TreeNode)*TreeNode{
        if root==nil{
            return root
        }
        l:=invert(root.Left)
        r:=invert(root.Right)
        root.Left=r
        root.Right=l
        return root
    }
    return invert(root)
    
}
