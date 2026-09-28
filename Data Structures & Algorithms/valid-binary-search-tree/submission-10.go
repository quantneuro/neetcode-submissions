/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

	var dfs func(root * TreeNode, mini int, maxi int)bool

	dfs=func(root *TreeNode, mini int, maxi int)bool{
		if root==nil{
			return true
			}
		if !(mini>=root.Val) || !(maxi<=root.Val){
			return false
		}

		l:=dfs(root.Left,mini,root.Val)
		if !l{return false}
		r:=dfs(root.Right,root.Val,maxi)
		if !r{return false}

		return l && r

	}

	return dfs(root,math.MinInt32,math.MaxInt32)
    
}
