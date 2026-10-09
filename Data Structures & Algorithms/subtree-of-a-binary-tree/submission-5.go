/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var traverse func(root *TreeNode, subRoot *TreeNode) bool
	var check func(root *TreeNode, subRoot *TreeNode) bool
	check = func(root *TreeNode, subRoot *TreeNode) bool{
		if subRoot==nil && root==nil{return true}
		if subRoot==nil || root==nil{return false}

		
		if root.Val !=subRoot.Val{return false}
		l:=check(root.Left,subRoot.Left)
      	if !l{return false}
       	r:=check(root.Right,subRoot.Right)
       	if !r{return false}


      	 return true
 
	}

	traverse = func(root *TreeNode, subRoot *TreeNode) bool{
		if root==nil || subRoot==nil{return false}


		if root.Val==subRoot.Val {
			val := check(root,subRoot)
			if val{return val}
		}
		l:=traverse(root.Left,subRoot)
		if l{return true}
		r:=traverse(root.Right,subRoot)
		if r{return true}
		return false
	}

	
	return traverse(root,subRoot)
    
}
