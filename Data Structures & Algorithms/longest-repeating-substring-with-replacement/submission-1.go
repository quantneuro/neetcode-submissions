// func characterReplacement(s string, k int) int {
// 	high:=0//current longest substring length
// 	for i:=0;i<len(s);i++{
// 		check:=make(map[byte]int)
// 		maxf:=0
// 		for j:=i;j<len(s);j++{
// 			sl:=s[i:j+1]//I can use just i-j+1 to get the length of the current substrings length but via sl it is just easier 
// 			check[s[j]]++//map for the frequency of till the current char in this substring

// 			//update maxfreuency length till now 
// 			maxf=max(maxf,check[s[j]])

// 			//now check each substring 
// 			if len(sl)-maxf<=k{
// 				//update current longest substring length
// 				high =max(high,len(sl))
// 			}

// 		}
// 	}
// 	return high

// }
func characterReplacement(s string, k int) int {
	l:=0
	r:=0
	high:=0//current longest substring length
	check:=make(map[byte]int)
	maxf:=0
	for r<len(s){
		
		// sl:=s[l:r+1] 

		check[s[r]]++ //map for the frequency of till the current char in this substring
			//s[r] is current characters length in the substring  and maxf is last varible maximum length
		//now check each substring 
		maxf=max(maxf,check[s[r]])//update maxfreuency length till now 

		for r-l+1-maxf>k{
			check[s[l]]--
			l++
			// sl:=s[l:r+1] 
		}
		//	^
		//	|
		//	|
		// too much confusin in the if else part I reduce l until the window is valid SIMPLE!
		//keep updating sl so the window updates, FYI I am using l,r, directly to calc window no need sl
			// if len(sl)-maxf<=k{
				 
				
			// 	//update current longest substring length
			// 	r++
				
			// }else{
			// 	if maxf==check[s[l]]{
			// 		maxf--
			// 	}
			// 	check[s[l]]--
			// 	l++
			// }
			//only once calc the window lenght and compare it with last highest window length
			high =max(high,r-l+1)
			r++
		}
	
	return high

}