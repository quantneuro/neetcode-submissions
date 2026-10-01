func coinChange(coins []int, amount int) int {

	var recur func(i int ,cur []int,total int)int
	
	
	check := make([][]int, len(coins))
	for i := range check {
		check[i] = make([]int, amount+1)
		for j := range check[i] {
			check[i][j] = -2
		}
	}
	
	recur = func(i int,cur []int,total int)int{
		
		
		if total==amount {
			return 0
		}

		if i>len(coins)-1 || total>amount{
			return -1
		}
		if v:=check[i][total];v!=-2{
			return v
		}
		
		cur=append(cur,coins[i])
		
		total+=coins[i]
		
		c:=recur(i,cur,total)
		cur=cur[:len(cur)-1]
		total-=coins[i]
		
		nc:=recur(i+1,cur,total)

		best := -1
		if c != -1 {
			best = c + 1
		}
		if nc != -1 && (best == -1 || nc < best) {
			best = nc
		}
		check[i][total]=best
		

		return best

	}

	a:=recur(0,[]int{},0)
	
	return a
    
}
