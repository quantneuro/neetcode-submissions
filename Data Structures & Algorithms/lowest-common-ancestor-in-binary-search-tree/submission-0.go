/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func exists (root *TreeNode,p *TreeNode,q *TreeNode, pe *bool,qe *bool)bool{

	if *pe && *qe {return true}
	if root == nil {return *pe && *qe}
	if p==root {*pe = true}
	if q==root {*qe = true}

	left :=exists(root.Left,p,q,pe,qe)
	right:=exists(root.Right,p,q,pe,qe)

	return left ||right


}
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {


	if root==nil{return root}
	pe:=false
	qe:=false
	

	left:=lowestCommonAncestor(root.Left,p,q)
	if left != nil {
		return left
	}
	right:=lowestCommonAncestor(root.Right,p,q)
	if right != nil {
		return left
	}
	

	
	if exists(root,p,q,&pe,&qe) {
		return root
	}
	return nil
	
    
}
