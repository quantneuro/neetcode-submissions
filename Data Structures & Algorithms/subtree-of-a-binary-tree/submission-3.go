/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

    var subtree func(root *TreeNode, subRoot *TreeNode) bool
	
	subtree=func(root *TreeNode, subRoot *TreeNode)bool {
		if root==nil && subRoot==nil{
			return true
		}else if root==nil || subRoot==nil{
			return false
		}

		l:=subtree(root.Left,subRoot.Left)
		if !l{return false}
		r:=subtree(root.Right,subRoot.Right)
		if !r{return false}

		return l&&r
	}

	var traverse func(root *TreeNode) bool

	traverse=func(root *TreeNode)bool  {
		if root==nil{
			return false
		}
		if root.Val==subRoot.Val{
			check:=subtree(root,subRoot)
			if check{
				return true
			}
		}
		l:=traverse(root.Left)
		if l{return true}
		r:=traverse(root.Right)
		if r{return true}

		return l||r
	}

	return traverse(root)
	
}
