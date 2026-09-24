/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorder(root *TreeNode, inord *[]int,k int){
	if root ==nil {return}

	inorder(root.Left,inord,k)
	*inord=append(*inord,root.Val)
	inorder(root.Right,inord,k)
	return
}
func kthSmallest(root *TreeNode, k int) int {
	inord:=[]int{}
	
	inorder(root,&inord,k)
	return inord[k-1]
	

}
