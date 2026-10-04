func partition(s string) [][]string {

	ispal := func(s string)bool{
		l:=0
		r:=len(s)-1
		for l<=r {
			if s[l]!=s[r] {return false}
			l++
			r--
		}
		return true
	}

	res:=[][]string{}
	// check:=make(map[string]bool)
	var all func(s string, cur []string)

	all=func(s string,cur []string){
		if len(s)==0{
			temp:=make([]string,len(cur))
			copy(temp,cur)
			res=append(res,temp)
			return
		}
		

		for i:=0;i<len(s);i++{
			current:=s[:i+1]
			left:=s[i+1:]
			if ispal(current) {
				cur=append(cur,current)
				all(left,cur)
				cur=cur[:len(cur)-1]
			}
			
			
		}

	}
	all(s,[]string{})

	return res

}
