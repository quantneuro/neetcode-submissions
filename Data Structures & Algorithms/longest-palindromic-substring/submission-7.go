func longestPalindrome(s string) string {
	cache:=make(map[[2]int]bool)
	n:=len(s)
	longestpalin:=""

	for length:=1;length<=n;length++{
		
		

		for left:=0;left<=n-length;left++{
			right:=left+length-1

			if right<=n-1{

			if length == 1{
				cache[[2]int{left,right}]=true
				if len(s[left:right+1]) > len(longestpalin) {
					longestpalin = s[left:right+1]
				}
			}else if length == 2{
				if s[left] == s[right] {
					cache[[2]int{left, right}] = true
					if len(s[left:right+1]) > len(longestpalin) {
						longestpalin = s[left:right+1]
				}
				}
			}else{
				if cache[[2]int{left+1, right-1}]{
					if s[left]==s[right]{

						cache[[2]int{left, right}]=true
						if len(s[left:right+1]) > len(longestpalin) {
							longestpalin = s[left:right+1]
						}
					}
				}

			}


			}
		}
	}

	return longestpalin
}

// func longestPalindrome(s string) string {
// 	maxstring:=""
// 	cache:=make(map[[2]int]bool)
// 	dfscache:=make(map[[2]int]bool)
	
// 	var palindromecheck func(l int,r int)bool
// 	palindromecheck=func(l int,r int)bool{

// 		for l<=r{

// 			if value,exist:=cache[[2]int{l,r}]; exist {
// 				return value
// 			}

// 			result:= s[l]==s[r] && palindromecheck(l+1,r-1)
				
// 			cache[[2]int{l, r}] = result
// 			return result
// 		}
// 		return true

// 	}
// 	var dfs func(l int, r int)
// 	dfs = func(l int, r int){
// 		if l>r{return}
// 		if dfscache[[2]int{l,r}]{return}
// 		dfscache[[2]int{l,r}]=true
// 		if palindromecheck(l,r){
// 			temp:=s[l:r+1]
// 			if len(temp)>len(maxstring){
// 				maxstring=temp
// 			}
// 		}

// 	dfs(l,r-1)
// 	dfs(l+1,r)
// 	}

// 	dfs(0,len(s)-1)
// 	return maxstring
// }

	// var dfs func(l int, r int) string 

	// dfs=func(l int,r int)string{
	// 	if l>r{
	// 		return maxstring
	// 	}
	// 	val,ok:=cache[[2]int{l,r}]
	// 	if ok {
	// 		if val{
	// 			temp:=s[l:r+1]
	// 			if len(temp)>len(maxstring){
	// 				maxstring=temp
	// 			}
	// 			return maxstring	
	// 		}
	// 	}else {
	// 		   if palindromecheck(l,r){
	// 				temp:=s[l:r+1]
	// 				if len(temp)>len(maxstring){
	// 					maxstring=temp
	// 				}
	// 				cache[[2]int{l,r}]=true
	// 				return maxstring
	// 			}else{
	// 				cache[[2]int{l,r}]=false
	// 			}
	// 	}

	// 	if v,e:=dfscache[[2]int{l,r}];e{
	// 		return v
	// 	}
	// 	s1:=dfs(l,r-1)
	// 	dfscache[[2]int{l,r-1}]=s1
	// 	s2:=dfs(l+1,r)
	// 	dfscache[[2]int{l+1,r}]=s2
	// 	v1:=len(s1)
	// 	v2:=len(s2)
	// 	if v1>=v2{
	// 		return s1
	// 	}else{
	// 	return s2
	// 	}
	// }


