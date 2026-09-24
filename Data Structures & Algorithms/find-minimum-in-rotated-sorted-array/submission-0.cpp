class Solution {
public:
    int findMin(vector<int> &nums) {
        int n= nums.size();
        
        int l=0;
        int r=n-1;
        int mid=0;
        int mini=INT_MAX;
        while(l<=r){
            mid=l+(r-l)/2;
            mini=min(nums[mid],mini);

            if(nums[mid]>nums[r]){
                l=mid+1;
            }
            else if(nums[mid]<nums[r]) {
                r=mid;

            }
            else if(nums[mid]==nums[r]){
                return nums[mid];
            }
        }
        return mini;
    }
};
