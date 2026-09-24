// // import "math"
// func minWindow(s string, t string) string {
//     // tcount:=[26]int{}// cna't use since it can be mixed vlalues of lower and upper case values
// 	//hence use mape
// 	tcount:=make(map[rune]int)
	

// 	for _,value:=range t{
// 		tcount[value]++
// 	}

// 	allversion:=[]string{}

// 	for i:=0;i<len(s);i++{
// 		for j:=i+1;j<=len(s);j++{
// 			sl:=s[i:j]

// 			if len(sl)>=len(t){
// 				allversion = append(allversion,sl)
// 		}
// 		}
// 	}
// 	res:=[]string{}
// 	for _,value:=range allversion{
// 		temp:=make(map[rune]int)
// 		valid:=true
// 		for _,char:=range value{
// 			temp[char]++
// 		 }
		 
// 		if valid {

// 			// for i:=0;i<len(temp);i++{
		 
// 		 	for char,needed:=range tcount{
// 				if temp[char]>=needed {
// 					continue
// 				}else {
// 					valid =false
// 					}
// 				}

// 			}		
		

// 	if valid {
// 			res=append(res,value)
// 		}

// 	}
// 	short:=math.MaxInt
// 	index:=0
// 	for i,val := range res{
// 		if len(val)<short{
// 			index=i
// 			short =len(val)
// 		}
		

// 	}
// 	if len(res)==0{return ""}
// 	return res[index]
// }



// import "math"
func minWindow(s string, t string) string {
    // tcount:=[26]int{}// cna't use since it can be mixed vlalues of lower and upper case values
	//hence use mape
	tcount:=make(map[rune]int)
	

	for _,value:=range t{
		tcount[value]++
	}

	// allversion:=[]string{}
	res:=[]string{}
	for i:=0;i<len(s);i++{
		temp:=make(map[rune]int)
		for j:=i;j<len(s);j++{
			sl:=s[i:j+1]
			temp[rune(s[j])]++
			valid:=true
			if len(sl)>=len(t){
				// allversion = append(allversion,sl)
				for char,needed:=range tcount{
					if temp[char]>=needed {
						continue
					}else {
						valid =false
						break
					}
				}

				if valid {
					res=append(res,sl)
				}
			}
		}
	}

	
	
	short:=math.MaxInt
	index:=0
	for i,val := range res{
		if len(val)<short{
			index=i
			short =len(val)
		}
		

	}
	if len(res)==0{return ""}
	return res[index]
}



