/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	copied:=make(map[*Node]bool)
	newval:=make(map[*Node]*Node)

	var dfscopy func(node *Node)*Node
	dfscopy = func(node *Node)*Node{
		if node==nil{
			return nil
		}
		if copied[node]{
			return newval[node]
		}

		newNode:=&Node{
			Val:node.Val,
		    Neighbors:[]*Node{},
			}
		newval[node]=newNode
		copied[node]=true

		if len(node.Neighbors)==0{
			// newNode.Neighbors=[]*Node{}
			return newNode
		}

		for i:=0;i<len(node.Neighbors);i++{
			nn:=dfscopy(node.Neighbors[i])
			
			
			newNode.Neighbors=append(newNode.Neighbors,nn)
			
		}



		return newNode

	}

	return dfscopy(node)
    
}
