/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchIndex(inorder [] int,currentroot int ) int{
	//can't do binary search since the question is only for bt , only bst inorder are increasing 
	// l:=0
	// r:=len(inorder)-1
	// mid:=0
	// for l<=r{
	// 	mid:= l+(r-l)/2
	// 	if inorder[mid]<currentroot{
	// 		l=mid+1
	// 	}else if inorder[mid]>currentroot{
	// 		r=mid-1
	// 	}else{
	// 		return mid
	// 	}

	// }

	foundindex:=0
	for i,v:=range inorder{
		if v==currentroot {
			foundindex=i
			break
		}
	}
	return foundindex



}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
		return nil
	}
	currentroot:=preorder[0]
	// rootindexinorder
	rii:=searchIndex(inorder,currentroot)

	left:=buildTree(preorder[1:rii+1],inorder[:rii])

	right:=buildTree(preorder[rii+1:],inorder[rii+1:])

	return &TreeNode{
		Val:currentroot,
		Left:left,
		Right:right,
	}


}
