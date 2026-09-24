/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 //dfs inroder way

// func dfs(root*TreeNode,traversal *[]int){
// 	if root==nil{return}
// 	dfs(root.Left,traversal)
// 	*traversal=append(*traversal,root.Val)
// 	dfs(root.Right,traversal)
// 	return
// }
// func kthSmallest(root *TreeNode, k int) int {
// 	traversal:=[]int {}
// 	dfs(root,&traversal)

// 	return traversal[k-1]

    
// }
func help(root *TreeNode,limit int,counter *int,ans *int){
	if root==nil{return }
	
	help(root.Left,limit,counter,ans)
	*counter++
	if *counter ==limit {
		*ans=root.Val
		return
	}
	help(root.Right,limit,counter,ans)
	return

}
func kthSmallest(root *TreeNode, k int) int {
	counter:=0
	answer:=0
	help(root,k,&counter,&answer)
	return answer

}