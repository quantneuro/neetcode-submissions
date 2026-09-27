/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
    queue:=[]*TreeNode{}

	queue=append(queue,root)
	res:=[][]int{}
	for len(queue)!=0{

		length:=len(queue)
		cur:=[]int{}
		for i:=0;i<length;i++{
			current :=queue[0]
			queue=queue[1:]
			
			cur=append(cur,current.Val)
			if current.Left!=nil{
				queue=append(queue,current.Left)
			}
			if current.Right!=nil{
				queue=append(queue,current.Right)
			}
		}
		res=append(res,cur)
	}

	return res
}
