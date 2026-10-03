func wordBreak(s string, wordDict []string) bool {

	memo:=make(map[string]bool)
	var check func(s string)bool
	check = func(s string)bool{
		if len(s)==0{
			return true
		}
		if val,ok:= memo[s];ok{
			return val
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
				memo[s[j:]]=val
				if val{

					return true
				}
			}
		}
		return false

	}
    return check(s)
}
