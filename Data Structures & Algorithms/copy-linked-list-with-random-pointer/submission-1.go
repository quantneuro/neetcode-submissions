/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	dummy:=&Node{}
    cur:=dummy
	t:=head
	store:=make(map[*Node]*Node)
	for t!=nil{
		temp:=&Node{
			Val : t.Val,
		}
		store[t]=temp

		cur.Next=temp
		cur=temp
		t=t.Next
	}
	t=head
	cur=dummy.Next
	for t!=nil{
		currentrandom:=t.Random

		val,_:=store[currentrandom]
		cur.Random=val
		cur=cur.Next
		t=t.Next
	}
	return dummy.Next

	
}
