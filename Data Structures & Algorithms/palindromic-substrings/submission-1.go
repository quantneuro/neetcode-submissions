func countSubstrings(s string) int {
	// cache:=make(map[[2]int]bool)
	//using a 2d slice instead of a map might prevent overhead
	n:=len(s)

	cache:=make([][]bool,n)

	for i:=0;i<n;i++{
		cache[i]=make([]bool,n)
	}
    
	count:=0
	for length:=1;length<=n;length++{
		for left:=0;left<=n-length;left++{
			right:=left+length-1
			if right<=n-1{
				if length==1 || length==2{
					if s[left]==s[right]{
						// cache[[2]int{left,right}]=true
						cache[left][right]=true
						count++
					}
				}else{
					// if cache[[2]int{left+1,right-1}]{
					if cache[left+1][right-1]{
						if s[left]==s[right]{
						// cache[[2]int{left,right}]=true
						cache[left][right]=true
						count++
					}
					}
				}
			}
		}
	}

	return count
}
