class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size()==0) return 0;
        int maxprofit=0;
        int l=0;
        int r=1;
        while(!(r==prices.size())){
            if(prices[r]<prices[l]){
                l=r;
            }else{
                maxprofit=max(maxprofit, prices[r]-prices[l]);
            }
            r++;
        }
        return maxprofit;

        
    }
};
