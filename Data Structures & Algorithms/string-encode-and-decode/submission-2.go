type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded_string :=""
	for _,val := range strs{
		encoded_string +=strconv.Itoa(len(val))+"#" + val
	}
	
	return encoded_string	
}

func (s *Solution) Decode(encoded string) []string {
	decoded_strings := []string {}
	j:=0
	for i:=0;i<len(encoded);{
		for encoded[j]!='#'{
			j++
		}
		length,_:=strconv.Atoi(encoded[i:j])
		j++//skip the #
		i=j
		decoded:=encoded[i:i+length]
		decoded_strings = append(decoded_strings,decoded)

		i=i+length
		j=j+length
	}
	return decoded_strings


	
}
