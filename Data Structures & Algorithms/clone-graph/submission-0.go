/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    
	seen:=make(map[*Node]bool)
	clone:=make(map[*Node]*Node)//key is orginal nodes//value is newly made nodes

	var dfsbuild func(node * Node)*Node

	dfsbuild=func(node * Node)*Node{
		if node==nil{
			return nil
		}
		seen[node]=true
		currentneighborslist:=node.Neighbors
		newNode:=&Node{
			Val:node.Val,
		}
		clone[node]=newNode
		for i:=0;i<len(currentneighborslist);i++{
			currentneighbor:=currentneighborslist[i]
			if !seen[currentneighbor]{
				newNode.Neighbors=append(newNode.Neighbors,dfsbuild(currentneighbor))
			}else{
				newcurrentneighbor:=clone[currentneighbor]
				newNode.Neighbors=append(newNode.Neighbors,newcurrentneighbor)
			}
			
		}
		return newNode
		
	}

	return dfsbuild(node)

}
