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
		return 0
	}

	hr := helper(root.Right) 
	if hr == -1 {return -1}
	hl := helper(root.Left) 
	if hl==-1 {return -1}

	diff:=math.Abs(float64(hl-hr))
	if diff>1{
		return -1
	}

	return max(hl,hr)+1

}

func isBalanced(root *TreeNode) bool {
	return helper(root)!=-1
    
}
