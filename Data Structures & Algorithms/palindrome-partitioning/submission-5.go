func partition(s string) [][]string {

	check:=make(map[string]bool)

	ispal := func(s string)bool{
		if val,ok:=check[s];ok{
			return val
		}
		l:=0
		r:=len(s)-1
		for l<=r {
			if s[l]!=s[r] {
				check[s]=false
				return false}
			l++
			r--
		}
		check[s]=true
		return true
	}

	res:=[][]string{}
	
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
