// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

// func solve(head * ListNode) {
//     if head==nil ||head.Next==nil{
//         return 
//         }
//     slow:= head
//     fast:= head.Next
    
//     for fast!=nil && fast.Next!=nil{
//         slow=slow.Next
//         fast=fast.Next.Next
//     }
    
//     //revere from slow 
    
//     second:= slow.Next
//     slow.Next=nil
//     var prev *ListNode
    

    
//     for second!=nil{
//         nextNode:=second.Next
//         second.Next=prev
//         prev=second
//         second=nextNode
//     }

//     first:=head
//     second = prev
//     for second!=nil {
//         tmp1,tmp2:=first.Next,second.Next
//         first.Next=second
//         second.Next=tmp1
//         first=tmp1
//         second =tmp2
//     } 
    


// }
// func reorderList(head *ListNode) {
//     if head ==nil || head.Next ==nil{
//         return
//     }
    
//     solve(head)
// }



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode)*ListNode{
	if head == nil || head.Next==nil{
		return head
	}
	temp:=reverse(head.Next)//temp need not be touched head is already going back
	head.Next.Next=head    // make the next of the just one after the head to point back at head 1->2->3->4 so now if head is at3
						  //  head Next is 4 so 4's next to point at 3 , aka reversed
	head.Next=nil		  //putting current head Next as nil
	return temp 
}

func reorderList(head *ListNode) {
    if head ==nil || head.Next ==nil{
        return
    }
	fast:=head
	slow:=head

	for fast!=nil && fast.Next != nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	//Now slow is at midpoint
	list2:=slow.Next
	slow.Next=nil
	list2=reverse(list2)

	list1 :=head

	for list2 !=nil{
		temp1,temp2:=list1.Next,list2.Next
		list1.Next=list2
		list2.Next=temp1
		list1 = temp1
		list2 =temp2
	}
}

