/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    
	// seen:=make(map[*Node]bool)
	newcopy:=make(map[*Node]*Node)
	var dfscopy func(node *Node)*Node

	dfscopy = func(node * Node)*Node{
		if node==nil {
			return nil
		}
		if v,ok:=newcopy[node];ok{
			return v
		}
		
		newNode:=&Node{
			Val:node.Val,
		}
		// seen[node]=true
		newcopy[node]=newNode
		for i:=0;i<len(node.Neighbors);i++{
			a:=dfscopy(node.Neighbors[i])
			if a!=nil{

			newNode.Neighbors =append(newNode.Neighbors,a)
		}

		
		}
		return newNode

	}

	return dfscopy(node)
}
