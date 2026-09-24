/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    copy:=&Node{}//remember copy is a dummy node and the copied list actually starts at copy.Next
	copyparse:=copy
	check :=make(map[*Node]*Node)
	
	parse:=head
	for parse!=nil {
		temp :=&Node{}
		temp.Val = parse.Val
		copyparse.Next =temp
		copyparse=temp
		
		check[parse]=temp
		parse=parse.Next
	
	}

	copyparse=copy.Next
	parse=head
	for parse!=nil {
		currentRandom:=parse.Random

		corresponding :=check[currentRandom]

		copyparse.Random=corresponding

		parse=parse.Next
		copyparse=copyparse.Next

	}

	return copy.Next
	
}
