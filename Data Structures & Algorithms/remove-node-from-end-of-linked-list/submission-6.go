/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func length(head *ListNode)int{
	
	temp:=head
	x:=0
	for temp!=nil{
		temp=temp.Next
		x++
	}
	return x

}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	x:=length(head)
	println(x)
	if n==x{
		head=head.Next
		return head
	}
	posdelete:= x-n
	nodebeforedeletion:=head


	for i:=1;i<posdelete;i++{
		nodebeforedeletion=nodebeforedeletion.Next
	}
	nodebeforedeletion.Next=nodebeforedeletion.Next.Next
	
	return head
    
}
