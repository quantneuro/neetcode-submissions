/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {

	var balance func(root *TreeNode)int 
	balance = func(root *TreeNode)int{
		if root==nil{return 0}
		l:=balance(root.Left)
		if l==-1{return -1}
		r:=balance(root.Right)
		if r==-1{return -1}
		if math.Abs(float64(l-r))>1{return -1}

		return 1+max(l,r)
	}
	f:=balance(root)
	if f==-1{
		return false
	}else{
		return true
	}
	

    
}
