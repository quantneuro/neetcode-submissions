/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

class Solution {
public:
int count=0;
int helper(TreeNode * root){
    if(!root) return -1;
        int ld = 1 + helper(root->left);
        int rd = 1 + helper(root->right);
        count=max(count,ld+rd);
        return max(rd,ld); 
}
    int diameterOfBinaryTree(TreeNode* root) {
        if(!root) return 0;
        helper(root);
        // int ld = 1 + diameterOfBinaryTree(root->left);
        // int rd = 1 + diameterOfBinaryTree(root->right);
        // count=max(count,ld+rd);
        return count; 
    }
};
