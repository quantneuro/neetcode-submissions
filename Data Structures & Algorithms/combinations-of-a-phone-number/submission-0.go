var set = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}
func combinations(digits string,current string,res *[]string){
	if len(digits)==0{
		*res=append(*res,current)
		return
	}

	v:=set[digits[0]]
	for i:=0;i<len(v);i++{
		current+=string(v[i])
		combinations(digits[1:],current,res)
		current=current[:len(current)-1]
	}
	
}
func letterCombinations(digits string) []string {
	res:=[]string{}
	if len(digits)==0{return res}
	combinations(digits,"",&res)
	return res
	
	
}
