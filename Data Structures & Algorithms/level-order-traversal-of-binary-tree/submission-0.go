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
	queue:=[]*TreeNode{}
	
	queue=append(queue,root)
	currentroot:=queue[0]
	for len(queue)!=0 && currentroot!=nil{
		currentlevelsize:=len(queue)
		valuesincurrentlevel:=[]int{}

		for ;currentlevelsize>0;currentlevelsize--{
			currentroot=queue[0]
			queue=queue[1:]
			valuesincurrentlevel=append(valuesincurrentlevel, currentroot.Val)
			queue = append(queue, currentroot.Left)
			queue = append(queue, currentroot.Right)
		}
		currentroot=queue[0]
		lot=append(lot,valuesincurrentlevel)
	}
	return lot
}
