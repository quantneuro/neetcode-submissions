/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root==nil{
		return []int{}
	}
	
	rightview:=[]int{}

	q:=[]*TreeNode{}
	q=append(q,root)

	for len(q)!=0{
		curlevel:=len(q)
		last:=0
		for i:=0;i<curlevel;i++{
			current:=q[0]
			q=q[1:]
			last=current.Val
			
			if current.Left!=nil{
				q=append(q,current.Left)
			}
			if current.Right!=nil{
				q=append(q,current.Right)
			}
		}
		rightview=append(rightview,last) 
	}
	return rightview


}
