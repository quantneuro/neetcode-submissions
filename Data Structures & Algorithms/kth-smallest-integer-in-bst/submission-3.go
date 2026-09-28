/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {

	var inorder func(root *TreeNode)
	inord:=[]int{}

	inorder= func(root *TreeNode){
		if root==nil{return}
		inorder(root.Left)
		inord=append(inord,root.Val)
		inorder(root.Right)
	}
	inorder(root)
	
	return inord[k-1]


    
}
