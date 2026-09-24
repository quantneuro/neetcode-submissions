// import
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func check(root*TreeNode)int{
	if root==nil{return 0}
	
	left:=check(root.Left)
	if left == -1{return -1}
	right:=check(root.Right)
	if right == -1{return -1}

	if math.Abs(float64(left-right))>1{
		return -1
	}
	return 1+max(left,right)
}
func isBalanced(root *TreeNode) bool {
	return check(root)!=-1
    
}
