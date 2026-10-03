func wordBreak(s string, wordDict []string) bool {


	var check func(s string)bool
	check = func(s string)bool{
		if len(s)==0{
			return true
		}

		val:=false
		for i:=0;i<len(wordDict);i++{
			flag:=true
			j:=0
			for ;j<len(wordDict[i]);j++{
				if len(s)<=j || s[j]!=wordDict[i][j]{
					flag=false
					break
				}
			}
			if flag {
				val=check(s[j:])
				if val{
					return true
				}
			}
		}
		return false

	}
    return check(s)
}
