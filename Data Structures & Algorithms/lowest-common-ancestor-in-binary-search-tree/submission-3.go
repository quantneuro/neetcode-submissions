/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root *TreeNode, p *TreeNode, q *TreeNode, pe *bool,qe *bool)bool{
	if root == p {*pe=true}
	if root == q {*qe=true}
	if *pe && *qe {return true}
	if root ==nil{return false}

	left:=helper(root.Left,p,q,pe,qe)
	right:=helper(root.Right,p,q,pe,qe)

	return left ||right

}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root==nil{return nil}

	left:=lowestCommonAncestor(root.Left,p,q)
	if left!=nil{return left}
	right:=lowestCommonAncestor(root.Right,p,q)
	if right !=nil {return right}
	pe,qe:=false,false
    if helper(root,p,q,&pe,&qe){
		return root
	}
	return nil
	
}
