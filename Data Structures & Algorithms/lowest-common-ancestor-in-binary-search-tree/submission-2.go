/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func help(root*TreeNode,p*TreeNode,q*TreeNode,pe* bool,qe* bool )bool{

	if root==nil {return *pe && *qe}
	if root==p{*pe = true}
	if root==q{*qe = true}

	if *pe && *qe{ return true}

	left:=help(root.Left,p,q,pe,qe)
	right:=help(root.Right,p,q,pe,qe)

	return left || right


 }

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if root==nil{return nil}
	pe:=false
	qe:=false
	left:=lowestCommonAncestor(root.Left,p,q)
	if left!=nil{return left}
	right:=lowestCommonAncestor(root.Right,p,q)
	if right!=nil{return right}
	
	if help(root,p,q,&pe,&qe){
		return root
	} 
	return nil
	

	
}
