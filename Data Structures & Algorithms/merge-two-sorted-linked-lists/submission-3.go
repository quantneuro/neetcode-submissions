/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    //this created a new node
    dummy := &ListNode{}

    current :=dummy
    
    for list1!=nil && list2!=nil {

        if list1.Val <= list2.Val {
            //current.Next = &ListNode{Val: list1.Val}
            current.Next=list1
            list1=list1.Next
        }else{
            //current.Next = &ListNode{Val: list2.Val}//this creates a new node
            current.Next=list2
            list2=list2.Next
        }
        current=current.Next
    }

    for list1!=nil{
        current.Next = &ListNode{Val: list1.Val}
        current=current.Next
        list1 = list1.Next
    }
    for list2 !=nil{
        current.Next = &ListNode{Val: list2.Val}
        current=current.Next
        list2 = list2.Next
        
    }

    if list1!=nil{
        current.Next = list1;
    }else{
        current.Next = list2;
    }

    return dummy.Next
}
