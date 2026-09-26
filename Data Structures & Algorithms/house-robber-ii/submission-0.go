type state struct {
    n     int
    start bool
}

func rob(nums []int) int {
	cache:=make(map[state]int)
	top:=len(nums)-1
	var recur func(n int,start bool)int

	recur=func(n int,start bool)int{
		

		if n > top {
			return 0
		}

		if n==top{
			if start{
				return 0
			}else{
				return nums[n]
			}
		}

		key := state{n: n, start: start}

		if val, ok := cache[key]; ok {
			return val
		}

		robflag:=start
		skipflag:=start
		if n==0{
			robflag=true
			skipflag=false
		}

		robn:=nums[n]+recur(n+2,robflag)
		skipn:=recur(n+1,skipflag)
		result:=max(robn,skipn)

		cache[key]=result

		return result
	}

	return recur(0,false)

	
}
