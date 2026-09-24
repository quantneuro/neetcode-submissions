/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var previous *ListNode=nil

    current :=head

    for current !=nil {
        temp:=current.Next
        current.Next= previous
        previous = current
        current=temp
    }
    return previous
}
