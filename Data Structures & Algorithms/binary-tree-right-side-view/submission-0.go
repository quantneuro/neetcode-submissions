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
	queue:=[]*TreeNode{root}
	// bfs:=[][]int{}
	onlyright:=[]int{}
	// queue=append(queue,root)

	for len(queue)!=0{

		ll:=len(queue)//ll=level length
		vl:=[]int{}//values of this level

			for ;ll>0;ll--{
				currentroot:=queue[0]
				queue=queue[1:]
				
				vl=append(vl,currentroot.Val)
				if currentroot.Left!=nil{
				queue = append(queue,currentroot.Left)
				}
				if currentroot.Right!=nil{
				queue = append(queue,currentroot.Right)
				}
			}
			vln:=len(vl)-1//rightmost value
			onlyright=append(onlyright,vl[vln])
	}
	return onlyright
}
