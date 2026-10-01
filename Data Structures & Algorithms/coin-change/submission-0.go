func coinChange(coins []int, amount int) int {

	var recur func(coins []int,cur []int,total int)bool

	mini := 0
	recur = func(coins []int,cur []int,total int)bool{
		if total==amount {
			
			println(len(cur))
			if mini==0{
				mini=len(cur)
			}else if mini>len(cur){
				mini=len(cur)
			}
			return true
		}
		if len(coins)==0{
			return false
		}
		if total>amount{
			return false
		}

		cur=append(cur,coins[0])
		
		total+=coins[0]
		
		c:=recur(coins,cur,total)
		cur=cur[:len(cur)-1]
		total-=coins[0]
		
		nc:=recur(coins[1:],cur,total)

		return c || nc

	}

	a:=recur(coins,[]int{},0)
	if !a{return -1} 
	return mini
    
}
