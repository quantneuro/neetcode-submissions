/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var count int = 0
func height(root * TreeNode) int  {
	if root==nil{return -1}
	    
        ld:=1+height(root.Left)
	    rd:=1+height(root.Right)
        count=max(ld+rd,count);

	return max(ld,rd);
}
func diameterOfBinaryTree(root *TreeNode) int {
    count=0
    if root==nil {return 0}
    height(root)
	return count;
}
