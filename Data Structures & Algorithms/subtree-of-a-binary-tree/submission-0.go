/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func same(r *TreeNode,sr * TreeNode)bool{
	if r==nil && sr==nil{
		return true
	}else if r==nil|| sr==nil {
		return false
	}else{
		if r.Val!=sr.Val{
		return false
		}
	}
	lv:=same(r.Left,sr.Left)
	rv:=same(r.Right,sr.Right)

	return lv && rv

}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    
	if root==nil {
		return false
	}
	
	if same(root,subRoot){
		return true
	}
	

	l:=isSubtree(root.Left,subRoot)
	r:=isSubtree(root.Right,subRoot)

	return l || r


}
