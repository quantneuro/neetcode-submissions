func generateParenthesis(n int) []string {
	res:=[]string{}


	var all func(s string,open int,closed int)
	all = func(s string,open int,closed int){
		if open==n && closed==n{
			a:=s
			res=append(res,a)
			return
		}

		if open<n{
			
			s+="("
			all(s,open+1,closed)
			s=s[:len(s)-1]
		}
		if closed<open{
			
			s+=")"
			all(s,open,closed+1)
			s=s[:len(s)-1]

		}
		
	}

	all("",0,0)

	return res


}
