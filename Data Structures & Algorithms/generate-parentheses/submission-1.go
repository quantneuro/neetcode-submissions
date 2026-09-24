func generateParenthesis(n int) []string {
	result:=[]string{}

	var total func(o int, c int, s string)

	total=func(o int, c int , s string){
		if o==n&& c==n{
			result=append(result,s)
		}

		if c<o{
			s+=")"
			c++
			total(o,c,s)
			s=s[:len(s)-1]
			c--
		}
		if o<n{
			s+="("
			o++
			total(o,c,s)
			s=s[:len(s)-1]
			o--
		}

	}
	total(0,0,"")
	return result

}
