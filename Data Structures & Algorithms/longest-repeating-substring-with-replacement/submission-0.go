func characterReplacement(s string, k int) int {
	high:=0//current longest substring length
	for i:=0;i<len(s);i++{
		check:=make(map[byte]int)
		maxf:=0
		for j:=i;j<len(s);j++{
			sl:=s[i:j+1]//I can use just i-j+1 to get the length of the current substrings length but via sl it is just easier 
			check[s[j]]++//map for the frequency of till the current char in this substring

			//update maxfreuency length till now 
			maxf=max(maxf,check[s[j]])

			//now check each substring 
			if len(sl)-maxf<=k{
				//update current longest substring length
				high =max(high,len(sl))
			}

		}
	}
	return high

}
