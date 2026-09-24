/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode, subRoot *TreeNode)bool{
	if root==nil && subRoot ==nil {
		return true
	}
	if root ==nil || subRoot ==nil{
		return false
	}
	
	

	if !(root.Val==subRoot.Val){
		return false
	}
	left:=helper(root.Left,subRoot.Left)
	right:=helper(root.Right,subRoot.Right)

	return left && right



}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root ==nil{return false}

	if root.Val==subRoot.Val{
		if helper(root,subRoot){
			return true
		}
	}
	left:=isSubtree(root.Left,subRoot)
	right:=isSubtree(root.Right,subRoot)

	return left || right 
	
    
}
