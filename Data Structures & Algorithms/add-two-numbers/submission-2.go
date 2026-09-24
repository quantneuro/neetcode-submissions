/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverse (head *ListNode)*ListNode{
// 	if head==nil || head.Next ==nil{
// 		return head
// 	} 
// 	temp:=reverse(head.Next)
// 	head.Next.Next=head
// 	head.Next=nil
// 	return temp
// }
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     n1:= ""
// 	n2:=""
// 	unreversedl1:=reverse(l1)
// 	unreversedl2:=reverse(l2)

// 	for unreversedl1 !=nil{
// 		n1+=strconv.Itoa(unreversedl1.Val)
// 		unreversedl1 =unreversedl1.Next
// 	}
// 	for unreversedl2 != nil {
// 		n2+=strconv.Itoa(unreversedl2.Val)
// 		unreversedl2 = unreversedl2.Next
// 	}
// 	num1,_ := strconv.Atoi(n1)
// 	num2,_ := strconv.Atoi(n2)
// 	println(num1)
// 	println(num2)
// 	num3:=num1+num2
// 	println(num3)

// 	sumlist:=&ListNode{}

// 	parse:=sumlist
// 	n3:=strconv.Itoa(num3)
// 	println(n3)
// 	for _,s:=range n3{
// 		temp:=&ListNode{}
		
// 		temp.Val=int(s-'0')
// 		parse.Next=temp
// 		parse=temp
// 	}

// 	finalreversedlist:=reverse(sumlist.Next)

// 	return finalreversedlist 

// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newlist:=&ListNode{}
	parse:=newlist
	carry:=0
	v1,v2:=0,0
	for l1!=nil || l2!=nil {
		if l1==nil{
			v1=0
		}else{
			v1=l1.Val
		}
		if l2 ==nil{
			v2=0
		}else{
			v2=l2.Val
		}
		v3:=v1+v2+carry
		carry=v3/10
		temp:=&ListNode{}
		temp.Val=v3%10
		parse.Next=temp
		parse=temp
		if l1!=nil{
			l1=l1.Next
		} 
		if l2!=nil{
			l2=l2.Next

		}


	}
	if carry>0{
		temp:=&ListNode{}
		temp.Val=carry
		parse.Next=temp
		parse=temp
	}
	return newlist.Next

}

