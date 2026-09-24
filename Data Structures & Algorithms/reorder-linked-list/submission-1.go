/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func solve(head * ListNode) {
    if head==nil ||head.Next==nil{
        return 
        }
    slow:= head
    fast:= head.Next
    
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    }
    
    //revere from slow 
    
    second:= slow.Next
    slow.Next=nil
    var prev *ListNode
    

    
    for second!=nil{
        nextNode:=second.Next
        second.Next=prev
        prev=second
        second=nextNode
    }

    first:=head
    second = prev
    for second!=nil {
        tmp1,tmp2:=first.Next,second.Next
        first.Next=second
        second.Next=tmp1
        first=tmp1
        second =tmp2
    } 
    


}
func reorderList(head *ListNode) {
    if head ==nil || head.Next ==nil{
        return
    }
    
    solve(head)
}
