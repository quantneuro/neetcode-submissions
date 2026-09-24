class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        unordered_map <int,int> Countn;
        for(int i=0 ;i<nums.size();i++){
            cout<<"Here :"<<i<<", ";
            if(Countn[nums[i]]>0){
                
                cout<<"out";
                return true;
            }
            Countn[nums[i]]++;

        }
        return false;
    }
};