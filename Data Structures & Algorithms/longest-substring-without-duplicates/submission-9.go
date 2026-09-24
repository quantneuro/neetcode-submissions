func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}
	if len(s)==1{
		return 1
	}

	seen:=make(map[rune]bool)
	
	maxlength:=0
	currentlength:=0
    i:=0
	j:=0
	for ;j<len(s);j++{

		for seen[rune(s[j])]{
			seen[rune(s[i])]=false
			i++
		}

	seen[rune(s[j])]=true
	currentlength = j-i+1
	maxlength=max(currentlength,maxlength)

	}
	return maxlength 

}
