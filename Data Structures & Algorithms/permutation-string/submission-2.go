func checkInclusion(s1 string, s2 string) bool {
	s1count:=[26]int{}

	for _,value := range s1{
		s1count[value-'a']++
	}
	
	for sp:=0;sp<len(s2);sp++{
		ep:=sp+len(s1)
		var tempcheck [26]int 
		
		if ep<= len(s2){
			sl:=s2[sp:ep]
			for _,value := range sl{
				
				tempcheck[value-'a']++
		}
		if tempcheck ==s1count{return true}
	}
	}
	return false

}
