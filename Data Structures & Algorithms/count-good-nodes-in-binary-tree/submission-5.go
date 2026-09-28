/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	good:=0
    
	var dfs func(root *TreeNode,maxi int)

	dfs=func(root *TreeNode,maxi int){

		if root==nil{
			return 
		}
		if root.Val>=maxi{
			maxi=root.Val
			good++
		}

		dfs(root.Left,maxi)
		dfs(root.Right,maxi)
	}
	dfs(root,root.Val)

	return good
}
