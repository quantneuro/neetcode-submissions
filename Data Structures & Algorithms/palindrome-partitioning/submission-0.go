func palindrome(sp string)bool{
	l:=0
	r:=len(sp)-1
	for l<=r{
		if !(sp[l]==sp[r]){
			return false
		}
		l++
		r--
	}
	return true
}

func check(s string,current []string,result*[][]string){
	if len(s)==0 {
		temp := make([]string,len(current))
		copy(temp,current)
		*result=append(*result,temp)
	}
	for i:=0;i<len(s);i++{
		currentpiece:=s[:i+1]
		leftover:=s[i+1:]
		if palindrome(currentpiece){
			current=append(current,currentpiece)
			check(leftover,current,result)
			current=current[:len(current)-1]
		}
		
	}

}
func partition(s string) [][]string {
	result:=[][]string{}
	check(s,[]string{},&result)
	return result
	
}
