type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded_string:=""
	for _,value:=range strs{
		encoded_string+=strconv.Itoa(len(value)) +"#"+ value  

	}
	return encoded_string
}

func (s *Solution) Decode(encoded string) []string {
	decoded_strs:=[]string{}
	
	for i:=0;i<len(encoded);{
		j:=i
		for encoded[j]!='#'{
			j++
		}
		//here j is on #
		length,_:=strconv.Atoi(encoded[i:j]) //from ith position to just before # and not #
		j++//skips the #
		
		decoded_strs=append(decoded_strs,encoded[j:j+length])
		i=j+length
		
	}
	return decoded_strs
}
