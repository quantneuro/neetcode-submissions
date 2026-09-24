/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func help(root*TreeNode,maxVal int, minVal int)bool{
	if root ==nil{return true}
	if !(root.Val>minVal &&root.Val<maxVal){return false}

	left  := help(root.Left,root.Val,math.MinInt)
	if !left{return false}

	right := help(root.Right,math.MaxInt,root.Val)
	if !right{return false}

	return left && right
	

}

func isValidBST(root *TreeNode) bool {
	if root == nil { return true }
	if !help(root,math.MaxInt,math.MinInt){
		return false
	}

	left:=isValidBST(root.Left)
	if !left{return false}
	right:=isValidBST(root.Right)
	if !right{return false}

	return left && right
    
}
