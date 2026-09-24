/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    //when using := this you don't need define *ListNode it infers automatically of the type
    current :=head
    fast :=current
    slow  :=current

    for fast!= nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
        if slow == fast {
            return true
        }

    } 
    return false;
}
