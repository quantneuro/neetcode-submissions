/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
//solved with two methods 1 creating new nodes and 2 reusing old nodes
// class Solution {
// public:
//     ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
//         ListNode dummy;//on stack
//         ListNode* current = &dummy;
//         // ListNode* temp = nullptr;

//         while(list1!=nullptr && list2 !=nullptr){
//             // temp = new ListNode;
//         if(list1->val > list2->val){
//             // temp->val=cl2->val;
//             current->next = new ListNode(list2->val);
//             list2=list2->next;
//         }
//         else {
//             // temp ->val =cl1->val;
//             // cl1=cl1->next;
//             current->next = new ListNode(list1->val);
//             list1=list1->next;
//         }
//         current=current->next;
//         }

//     //remaining values

//     while(list1 !=nullptr){
//         current->next = new ListNode(list1->val);
//         current=current->next;
//         list1 = list1->next;
//     }
//     while(list2!=nullptr){
//         current->next = new ListNode(list2->val);
//         current=current->next;
//         list2 =list2->next;
//     }
//     return dummy.next;
//     }
// };

//2nd way is resuing old node


class Solution{
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2){
        ListNode dummy;
        ListNode* current = &dummy;
        while(list1!=nullptr && list2!=nullptr){

            if(list1->val <= list2->val){
                current->next=list1;
                list1=list1->next;
            }
            else{
                current->next= list2;
                list2=list2->next;
            }
            current =current->next;
        }

        if(list1!=nullptr){
            current->next = list1;
        }
        else{
            current->next = list2;
        }

        return dummy.next;
    }
};



