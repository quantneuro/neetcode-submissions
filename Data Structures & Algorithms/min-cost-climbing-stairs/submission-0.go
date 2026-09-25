func minCostClimbingStairs(cost []int) int {
	
	cache:=make(map[int]int)
	var recur func(n int)int
	top:=len(cost)

	recur=func(n int)int{
		if n>top || n==top{
			return 0 
		}
		

		if val,ok:=cache[n];ok{
			return val
		}
		
		a:=recur(n+1)
		b:=recur(n+2)
		
		result:=cost[n]+min(a,b)
		cache[n]=result
		
		return result
		
	}

	a:=recur(0)
	b:=recur(1)

	return min(a,b)
}
