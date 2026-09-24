// // // import "math"
// // func minWindow(s string, t string) string {
// //     // tcount:=[26]int{}// cna't use since it can be mixed vlalues of lower and upper case values
// // 	//hence use map
// // 	tcount:=make(map[rune]int)
	

// // 	for _,value:=range t{
// // 		tcount[value]++
// // 	}

// // 	allversion:=[]string{}

// // 	for i:=0;i<len(s);i++{
// // 		for j:=i+1;j<=len(s);j++{
// // 			sl:=s[i:j]

// // 			if len(sl)>=len(t){
// // 				allversion = append(allversion,sl)
// // 		}
// // 		}
// // 	}
// // 	res:=[]string{}
// // 	for _,value:=range allversion{
// // 		temp:=make(map[rune]int)
// // 		valid:=true
// // 		for _,char:=range value{
// // 			temp[char]++
// // 		 }
		 
// // 		if valid {

// // 			// for i:=0;i<len(temp);i++{
		 
// // 		 	for char,needed:=range tcount{
// // 				if temp[char]>=needed {
// // 					continue
// // 				}else {
// // 					valid =false
// // 					}
// // 				}

// // 			}		
		

// // 	if valid {
// // 			res=append(res,value)
// // 		}

// // 	}
// // 	short:=math.MaxInt
// // 	index:=0
// // 	for i,val := range res{
// // 		if len(val)<short{
// // 			index=i
// // 			short =len(val)
// // 		}
		

// // 	}
// // 	if len(res)==0{return ""}
// // 	return res[index]
// // }



// // import "math"
// func minWindow(s string, t string) string {
//     // tcount:=[26]int{}// cna't use since it can be mixed vlalues of lower and upper case values
// 	//hence use mape
// 	tcount:=make(map[rune]int)
	

// 	for _,value:=range t{
// 		tcount[value]++
// 	}

// 	// allversion:=[]string{}
// 	res:=[]string{}
// 	for i:=0;i<len(s);i++{
// 		temp:=make(map[rune]int)
// 		for j:=i;j<len(s);j++{
// 			sl:=s[i:j+1]
// 			temp[rune(s[j])]++
// 			valid:=true
// 			if len(sl)>=len(t){
// 				// allversion = append(allversion,sl)
// 				for char,needed:=range tcount{
// 					if temp[char]>=needed {
// 						continue
// 					}else {
// 						valid =false
// 						break
// 					}
// 				}

// 				if valid {
// 					res=append(res,sl)
// 				}
// 			}
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

func minWindow(s string, t string) string {

	l:=0
	r:=0
	shortest:=math.MaxInt
	shortestsofar:=""

	tcount:=make(map[rune]int)
	//frequncy counter
	for _,val:=range t{
		tcount[val]++
	}
	
	testing:=make(map[rune]int)
	
	for r<len(s){
		valid:=true
		current:=s[l:r+1]

		testing[rune(s[r])]++
		
		//check the testing map has all values of t? IF yes then flag stays true
		for value,count:=range tcount{
			c,e:=testing[value]// e is boolean shows existence
			if !e || c<count {
				valid = false
				break
			}
		
		}
	
	
		//if valid is true shrink from left until not valid, also store the valid string as the shortest
		
		
		for valid{
			currentlen:=len(current)
			if currentlen<=shortest{
				shortestsofar=current
				shortest = currentlen 
			}

			testing[rune(s[l])]--
			l++
			current=s[l:r+1]//new string after moving l towards right
				
				for value,count := range tcount{
					c,e:=testing[value]// e is boolean shows existence
					
					if !e || c<count {
						valid = false
						break
					}
			}
			
		}
		
		r++
	}
	return shortestsofar
	
}





