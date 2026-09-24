// import "slices"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//OPtimized solution!!
func help(root *TreeNode,maxsofar int,count *int){
	if root==nil{
		return 
	}
	
	if root.Val>=maxsofar{
		*count++
		maxsofar=root.Val

	}
	help(root.Left,maxsofar,count)
	help(root.Right,maxsofar,count)

}
func goodNodes(root *TreeNode) int {
	count:=0
	maxsofar:=math.MinInt64
	help(root,maxsofar,&count)
	return count
}


// import "slices"
// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */

// func pathfinder(root * TreeNode,target *TreeNode, path *[]int)*TreeNode{
// 	if root==nil {return nil}
// 	if root == target {
// 		*path=append(*path,root.Val)//path is only addressof slice,  *path is the slice
// 		return root
// 		}
	
// 	left:= pathfinder(root.Left,target,path)
// 	if left!=nil{
// 		*path=append(*path,root.Val)
// 		return root
// 	}
// 	right:=pathfinder(root.Right,target,path)
// 	if right!=nil{
// 		*path=append(*path,root.Val)
// 		return root
// 	}
// 	return nil 
// } 
// func traversal( originalroot *TreeNode,currentroot *TreeNode,count *int){
// 	if currentroot ==nil{return }

// 	path:=[]int{}
// 	pathfinder(originalroot,currentroot,&path)
	
// 	if len(path)>0 && slices.Max(path)==path[0]{
// 		*count++
// 	}

// 	traversal(originalroot,currentroot.Left,count) 
// 	traversal(originalroot,currentroot.Right,count) 
// }

// func goodNodes(root *TreeNode) int {
// 	count:=0
// 	traversal(root,root,&count)

//     return count
// }
