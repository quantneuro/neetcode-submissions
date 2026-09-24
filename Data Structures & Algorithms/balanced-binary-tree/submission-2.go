/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  import (
//     "math"
// )
func helper(root *TreeNode) int {
	if root==nil {return 0}

	ld:=helper(root.Left)
	rd:=helper(root.Right)

	if ld==-1{return -1}
	if rd==-1{return -1}

	diff:= rd-ld
	if diff<0{
		diff=-diff
	}
	if diff<=1{
		return 1+max(ld,rd)
	}else{
		return -1
	}

}

func isBalanced(root *TreeNode) bool {
	return helper(root)!=-1
    
}
