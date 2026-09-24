/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root *TreeNode) int {
	if root ==nil{
		return 0
	}
	l:=helper(root.Left)
	if l==-1{return -1}
	r:=helper(root.Right)
	if r==-1{return -1}

	if math.Abs(float64(l-r))>1{
		return -1
	}
	return 1+max(l,r)
}
func isBalanced(root *TreeNode) bool {
	output := helper(root)
	if output == -1 {
		return false
	}else {
		return true
	}
    
}
