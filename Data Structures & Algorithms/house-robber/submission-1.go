func rob(nums []int) int {
    
	cache:=make(map[int]int)
	top :=len(nums)
	var recur func(n int)int

	recur=func(n int)int{
		if n>=top{
			return 0
		}

		if val,ok:=cache[n];ok{
			return val
		}

		robn:=nums[n]+recur(n+2)
		skipn:=recur(n+1)
		result:=max(skipn,robn)
		cache[n]=result
		return result

	}
	return recur(0) 

}
