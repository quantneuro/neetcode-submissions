/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    
	var build func(preorder []int, inorder []int) *TreeNode

	build=func(preorder []int, inorder []int) *TreeNode{
		if len(preorder)==0{
			return nil
		}
		value:=preorder[0]
		
		i:=0
		for ;i<len(inorder);i++{
			if value ==inorder[i]{
				break
			}
		}
		left:=inorder[:i]
		right:=inorder[i+1:]

		newRoot:=TreeNode{
			Val:value,
			Left:build(preorder[1:i+1],left),
			Right:build(preorder[i+1:],right),
		}

		return &newRoot
		
	}
	return build(preorder,inorder)
}
