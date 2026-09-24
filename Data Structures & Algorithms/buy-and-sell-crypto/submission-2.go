func maxProfit(prices []int) int {
    
    if len(prices) ==0 {return 0}

    l:=0
    r:=1
    var maxp int = 0

    for r !=len(prices) {
        if prices[r] < prices[l]{
            l=r
        }else{
            maxp=max(maxp,prices[r]-prices[l])
            
        }
       
         r++
        
    }
    return maxp
}
