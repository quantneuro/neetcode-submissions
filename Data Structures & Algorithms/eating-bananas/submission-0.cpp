int totalhours(vector<int>&pile,int testspeed){
    int totalhourss=0;
    for(int val:pile){
        totalhourss += ceil(val/testspeed);
    }
    return totalhourss;
}



class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        auto it =max_element(piles.begin(),piles.end());
        int l=1;
        int r=*it;
        int currentmin=0;
        while(l<=r){
            int mid=l+(r-l)/2;
            int totalhourss=totalhours(piles,mid);
             if(totalhourss<h){
                currentmin=mid;
                r=mid-1;
             }
             else if(totalhourss>h){
                l=mid+1;
             }
        }
        return currentmin;
    }
};
