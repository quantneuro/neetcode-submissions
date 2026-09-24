/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root *TreeNode) int{
	if root == nil {
		return -1
	}

	hr := helper(root.Right) 
	
	hl := helper(root.Left) 

	diff:=math.Abs(float64(hl-hr))
	if diff>1{
		return -2
	}

	return max(hl,hr)+1

}

func isBalanced(root *TreeNode) bool {
	return helper(root)!=-2
    
}
