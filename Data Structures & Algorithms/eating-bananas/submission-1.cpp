int totalhours(vector<int>&pile,int testspeed){
    int totalhourss=0;
    cout<<"current speed "<< testspeed<<endl;
    for(int val:pile){
        //the inner division make it the value lower integer
        //before ceil sees it , so be careful use double
        totalhourss += ceil((double)val/testspeed);
        cout<<"totalhours "<<totalhours<<endl;
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
             if(totalhourss<=h){
                currentmin=mid;
                r=mid-1;
                cout<<"firstif "<<mid<<endl;
             }
             else if(totalhourss>h){
                l=mid+1;
                cout<<"second if "<<mid<<endl;
             }
             
        }
        return currentmin;
    }
};
