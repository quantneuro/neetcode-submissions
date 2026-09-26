/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    


	var lca func( root *TreeNode,pe *bool, qe *bool) bool

	lca=func(root *TreeNode, pe *bool, qe *bool) bool{
		
		
		if root==p{
			*pe=true
		}
		if root==q{
			*qe=true
		}
		if *pe && *qe{
			return true
		}
		if root ==nil{
			return false
		}

		l:=lca(root.Left,pe,qe)
		if l{return true}
		r:=lca(root.Right,pe,qe)
		if r{return true}

		return l||r
	}

	var traverse func(root *TreeNode, p *TreeNode, q *TreeNode)*TreeNode

	traverse=func(root *TreeNode, p *TreeNode, q *TreeNode)*TreeNode{
		if root==nil{
			return nil
		}
		

		l:=traverse(root.Left,p,q)
		if l!=nil{return l}
		r:=traverse(root.Right,p,q)
		if r!=nil{return r}

		pe:=false
		qe:=false
		if lca(root,&pe,&qe){
			return root
		}
		
		return nil

	}

	
	return traverse(root,p,q)
	
}
