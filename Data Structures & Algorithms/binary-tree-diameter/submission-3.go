/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxlen int = 0
func maxlength(root *TreeNode)int{
    if root ==nil{return -1}
    
    left:=maxlength(root.Left)
    
    right:=maxlength(root.Right)
    

  maxlen=max(maxlen,left+right+2)//because =2, is neede read insight in dheets if you cant' recall
    
    return 1+max(right,left)

}

func diameterOfBinaryTree(root *TreeNode) int {
    maxlength(root)
    return maxlen
}
