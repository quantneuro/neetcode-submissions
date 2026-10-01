func partition(s string) [][]string {

	res:=[][]string{}

	checkpalindrome:=func(s string)bool{
		l:=0
		r:=len(s)-1
		for l<=r{
			if !(s[l]==s[r]){
				return false
			}
			l++
			r--
		}
		return true
	}
	var all func(s string, cur []string )

	all = func(s string,cur []string){
		if len(s)==0{
			temp:=make([]string,len(cur))
			copy(temp,cur)
			res=append(res,cur)
		}


		for i:=0;i<len(s);i++{
			currentpiece:=s[:i+1]
			leftover:=s[i+1:]
			if checkpalindrome(currentpiece){
				cur=append(cur,currentpiece)
				all(leftover,cur)
				cur=cur[:len(cur)-1]
			}
		}
	}

	all(s,[]string{})

	return res

}
