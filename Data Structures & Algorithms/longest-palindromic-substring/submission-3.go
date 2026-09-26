func longestPalindrome(s string) string {
	maxstring:=""
	// maxlength:=0
	cache:=make(map[[2]int]bool)
	dfscache:=make(map[[2]int]string)
	
	palindromecheck:=func(l int,r int)bool{
		for l<=r{
			if s[l]!=s[r]{
				return false
			}
			l++
			r--
		}
		return true

	}
	var dfs func(l int, r int) string 

	dfs=func(l int,r int)string{
		if l>r{
			return maxstring
		}
		val,ok:=cache[[2]int{l,r}]
		if ok {
			if val{
				temp:=s[l:r+1]
				if len(temp)>len(maxstring){
					maxstring=temp
				}
				return maxstring	
			}
		}else {
			   if palindromecheck(l,r){
					temp:=s[l:r+1]
					if len(temp)>len(maxstring){
						maxstring=temp
					}
					cache[[2]int{l,r}]=true
					return maxstring
				}else{
					cache[[2]int{l,r}]=false
				}
		}

		if v,e:=dfscache[[2]int{l,r}];e{
			return v
		}
		s1:=dfs(l,r-1)
		dfscache[[2]int{l,r-1}]=s1
		s2:=dfs(l+1,r)
		dfscache[[2]int{l+1,r}]=s2
		v1:=len(s1)
		v2:=len(s2)
		if v1>=v2{
			return s1
		}else{
		return s2
		}
	}

	return dfs(0,len(s)-1)
}
