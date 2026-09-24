/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	lot:=[][]int {}
	if root ==nil{
		return lot
	}
	
	queue:=[]*TreeNode{}
	
	queue=append(queue,root)
	
	for len(queue)!=0 {
		currentlevelsize:=len(queue)
		valuesincurrentlevel:=[]int{}

		for ;currentlevelsize>0;currentlevelsize--{
			currentroot:=queue[0]
			queue=queue[1:]
			valuesincurrentlevel=append(valuesincurrentlevel, currentroot.Val)
			if currentroot.Left!=nil{
			queue = append(queue, currentroot.Left)
			}
			if currentroot.Right!=nil{
			queue = append(queue, currentroot.Right)
			}
		}

		lot=append(lot,valuesincurrentlevel)
	}
	return lot
}
