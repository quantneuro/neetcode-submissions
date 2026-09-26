/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {

	var SameTree func(p *TreeNode, q *TreeNode)bool


	SameTree=func(p *TreeNode, q *TreeNode)bool{
		if p==nil && q==nil{
			return true
		}else if p==nil ||q==nil{
			return false
		}
		if p.Val!=q.Val{
			return false
		}

		l:=SameTree(p.Left,q.Left)
		if !l{
			return false
		}
		r:=SameTree(p.Right,q.Right)
		if !r{
			return false
		}

		return l && r

	}

	return SameTree(p,q)
	
}
