static void allsubsets(int i, vector<int>&nums,vector<int>&cur,vector<vector<int>>&allsubset ){
        
        if(i==nums.size()) {
            allsubset.push_back(cur);
            return;
        }

        allsubsets(i+1,nums,cur,allsubset);//the skip part

        cur.push_back(nums[i]);//the choose part

        allsubsets(i+1,nums,cur,allsubset);// the explore part

        cur.pop_back();//remove part

        
    }

class Solution {
public:
    
    
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> allsubset;
        vector<int> curr;
        allsubsets(0,nums,curr,allsubset);

        return allsubset;

    }
};
