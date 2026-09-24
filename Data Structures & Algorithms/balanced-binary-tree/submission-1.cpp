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

// class Solution {
// public:
//     int height(TreeNode* root){
//         if(!root) return 0;

//         return 1+max(height(root->right), height(root->left));
    
//     }
//     bool isBalanced(TreeNode* root) {
//         if(!root) return true;
//         bool bl = isBalanced(root->left);
//         bool br = isBalanced(root->right);
//         if(!bl || !br)return false;
//         int l=height(root->left);
//         int r=height(root->right);

//         if(abs(l-r)>1){
//             return false;
//         }
//         return true;
//     }
// };



class Solution {
public:
    int height(TreeNode* root){
        if(!root) return 0;
        int l=height(root->left);
        if (l == -1) return -1;

        int r=height(root->right);
        if (r == -1) return -1;

        if(abs(l-r)>1){
            return -1;
        }
        return 1+max(height(root->right), height(root->left));
    
    }
    bool isBalanced(TreeNode* root) {
        return height(root)!=-1; 
    }
};
