func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}	
	if len(s)==1{
		return 1
	}

	l:=0
	r:=0
	currentcount:=0
	maxcount:=0

	seen:=make(map[byte]bool)
	for r<len(s) {
		if seen[s[r]]{
			seen[s[l]]=false
			l++
			
		}else{
			currentcount=r-l+1
			maxcount=max(maxcount,currentcount)
			seen[s[r]]=true
			r++
			
		}
		
	}
	return maxcount

}
