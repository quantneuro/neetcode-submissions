// import(
// 	"slices"
// )
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//from original root to current root
func total(root *TreeNode, maxsofar int, count *int) {
	if root==nil{return }

	if maxsofar<=root.Val{
		*count++
		maxsofar=root.Val
	}
	total(root.Left,maxsofar,count)
	total(root.Right,maxsofar,count)
	return

}
func goodNodes(root *TreeNode) int {
	count:=0
	total(root,root.Val,&count)
	return count
}

// func goodNodes(root *TreeNode) int {
// 	count:=0
// 	traversal(root,root,&count)
// 	return count
    
// }

// func path(root *TreeNode,target *TreeNode,pathh*[]int)*TreeNode{
// 	if root==nil{
// 		return nil
// 	}
// 	if root==target{
// 		*pathh=append(*pathh,root.Val)//first value that goes in is the root that is being checked
// 		return root
// 	}
// 	left:=path(root.Left,target,pathh)
// 	if left!=nil{
// 		*pathh=append(*pathh,root.Val)
// 		return left
// 	}
// 	right:=path(root.Right,target,pathh)
// 	if right !=nil{
// 		*pathh=append(*pathh,root.Val)
// 		return right
// 	}

// 	return nil


// }
// func traversal(originalroot*TreeNode,root *TreeNode,count *int){
// 	if root==nil{return }
// 	pathh:=[]int{}
// 	check:=path(originalroot,root,&pathh)
// 	if check !=nil && slices.Max(pathh)==pathh[0]{
// 		*count++
// 	}
	
// 	traversal(originalroot, root.Left, count)
// 	traversal(originalroot, root.Right, count)
	
	
// 	return 

// }


