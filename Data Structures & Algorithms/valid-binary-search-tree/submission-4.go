
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bst(root *TreeNode, minlimit int, maxlimit int)bool{
	if root==nil{ return true }

	if !(root.Val>minlimit && root.Val<maxlimit){
		return false
	}
	//updating minlimit and maxlimit 
	left:=bst(root.Left,minlimit,root.Val)
	if left ==false{return false}
	right:=bst(root.Right,root.Val,maxlimit)
	if right ==false{return false}

	return true ||false

}

func isValidBST(root *TreeNode) bool {
	return bst(root,math.MinInt64,math.MaxInt64)
    
}
