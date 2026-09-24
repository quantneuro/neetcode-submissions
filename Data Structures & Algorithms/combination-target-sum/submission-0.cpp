

class Solution {
public:
    vector<vector<int>> ac;
    vector<vector<int>> combinationSum(vector<int>& nums, int target) {
        vector<int> curr;
        
        allcombinations(0,target,curr,nums);

        return ac; 

    }

    void allcombinations(int i,int target, vector<int>&curr, vector<int>&nums){

        if(target==0){
            ac.push_back(curr);
            return;
        }
        if(target < 0 || i>=nums.size()){return ;}

        curr.push_back(nums[i]);//choose

        allcombinations(i,target-nums[i],curr,nums);//skip

        curr.pop_back();//un-choose

        allcombinations(i+1,target,curr,nums);//explore

    }

};
