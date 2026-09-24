class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int,int> TS;
        for(int i=0;i<nums.size();i++){
            cout<<"here";
            int diff=target-nums[i];
            cout<<diff;
            if(TS.find(diff)!=TS.end()){
                cout<<"out";
                return {TS[diff],i};}
            else{ 
                
                TS[nums[i]]=i;
                cout<<TS[diff];}
        }
        return {};
    }
};
