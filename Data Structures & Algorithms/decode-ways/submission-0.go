func numDecodings(s string) int {
	n:=len(s)
	cache :=make([]int,n+1)
	if s[0]=='0' ||len(s)==0{
		return 0
	}

	cache[0]=1
	cache[1]=1

	
	for i:=2;i<=n;i++{
		
		currentdigit := s[i-1]

		
		if currentdigit !='0'{
			waysbeforecurrent := cache[i-1]
			cache[i] = cache[i]+waysbeforecurrent
		}

		previousdigit := s[i-2]

		twodigitstring := string(previousdigit) + string(currentdigit)
		twodigitnum,_:= strconv.Atoi(twodigitstring)

		if twodigitnum>=10 && twodigitnum<=26{
			waysbeforetwodigitnum:=cache[i-2]
			cache[i]=cache[i]+waysbeforetwodigitnum
		}
	}

	return cache[n]

}


