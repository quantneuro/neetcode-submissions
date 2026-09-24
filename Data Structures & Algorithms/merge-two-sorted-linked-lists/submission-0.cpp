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

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode dummy;//on stack
        ListNode* current = &dummy;
        // ListNode* temp = nullptr;

        while(list1!=nullptr && list2 !=nullptr){
            // temp = new ListNode;
        if(list1->val > list2->val){
            // temp->val=cl2->val;
            current->next = new ListNode(list2->val);
            list2=list2->next;
        }
        else {
            // temp ->val =cl1->val;
            // cl1=cl1->next;
            current->next = new ListNode(list1->val);
            list1=list1->next;
        }
        current=current->next;
        }

    //remaining values

    while(list1 !=nullptr){
        current->next = new ListNode(list1->val);
        current=current->next;
        list1 = list1->next;
    }
    while(list2!=nullptr){
        current->next = new ListNode(list2->val);
        current=current->next;
        list2 =list2->next;
    }
    return dummy.next;
    }
};
