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
 ListNode* reverses(ListNode* head){
       ListNode* prev=nullptr;
       ListNode* temp=nullptr;
       ListNode* current= head;
       while(current!=nullptr){
        temp=current->next;
        current->next=prev;
        prev=current;
        current=temp;
       } 
       return prev;
    }

class Solution {
public:
   
    void reorderList(ListNode* head) {
        if(head==nullptr || head->next==nullptr) return ;
        ListNode* fast=head;
        ListNode* slow= head;
        ListNode* current=head;

        while(fast!=nullptr && fast->next!=nullptr){
            slow=slow->next;
            fast=fast->next->next;
        }
        ListNode* list2= slow->next;
        slow->next =nullptr;
        list2=reverses(list2);

        ListNode* list1next=nullptr;
        ListNode* list2next =nullptr;
        while(list2!=nullptr){
            list1next = current->next;
            list2next = list2->next;

            current->next=list2;
            list2->next=list1next;

            list2=list2next;
            current=list1next;
            
        }

    }
};
