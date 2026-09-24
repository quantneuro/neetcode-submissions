/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(root*TreeNode,traversal *[]int){
	if root==nil{return}
	dfs(root.Left,traversal)
	*traversal=append(*traversal,root.Val)
	dfs(root.Right,traversal)
	return
}
func kthSmallest(root *TreeNode, k int) int {
	traversal:=[]int {}
	dfs(root,&traversal)

	return traversal[k-1]

    
}
