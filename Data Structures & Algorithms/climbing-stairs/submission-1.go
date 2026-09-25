func climbStairs(n int) int {

	cache:=make(map[int]int)
	
	var recur func(n int)int
	
	recur=func(n int)int{
	if n<0{
		return 0
	}
    if n==0{
		return 1
	}
	if val,ok:=cache[n];ok{
		return val
	}

	result:=recur(n-1)+recur(n-2)

	cache[n]=result

	return result

	}

	return recur(n)
	

}
