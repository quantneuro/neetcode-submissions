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
        int i=0;
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
        int i=1;
        if(n<=len){
         while(i<n){
            cout<<temp->val<<endl;
            temp=temp->next;
            i++;
         }
        }
        temp->next=temp->next->next;
        return head;
        
    }
};
