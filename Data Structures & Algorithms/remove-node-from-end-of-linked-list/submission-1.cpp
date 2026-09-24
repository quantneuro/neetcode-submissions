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
    int Count(ListNode*head){
        ListNode* temp = head;
        int i=1;
        while(temp->next){
            i++;
            temp=temp->next;
        }
        return i;

    }
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode*temp = head;
        if(!head || n==0){return head;}
        if(n==1){
            head=head->next;
            temp->next=nullptr;
            return head;
        }
        int len=Count(head);
        cout<<len<<endl;
        cout<<"dsf"<<endl;
        int ne=len-n;

        int i=1;
        if(ne>=0){
         while(i<ne){
            cout<<temp->val<<endl;
            temp=temp->next;
            cout<<"ith value"<<i<<endl;
            i++;
         }
        }
        ListNode* t1=temp->next;
        temp->next=t1->next;
        return head;
        
    }
};
