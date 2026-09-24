func help(n int,openN int , closeN int, current string, result *[]string){
	if n==openN && closeN==n {	
		temp:=""
		temp=current
		*result=append(*result,temp)
		return
	}

	if openN<n{
		current+="("
		help(n,openN+1,closeN,current,result)
		current=current[:len(current)-1]
	}
	if closeN<openN{
		current+=")"
		help(n,openN,closeN+1,current,result)
		current=current[:len(current)-1]
	}

}

func generateParenthesis(n int) []string {
	result:=[]string{}
	help(n,0,0,"",&result)//we don't need to double the n since open and close are saparate
	return result

}
