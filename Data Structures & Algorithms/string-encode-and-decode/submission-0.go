type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded_string:=""
	for _,value:=range strs{
		encoded_string+=value + ";"

	}
	return encoded_string
}

func (s *Solution) Decode(encoded string) []string {
	decoded_strs:=[]string{}
	currentchar:=""
	for _,value := range encoded{
		//value is a rune and ; is string so either make value a string or ; a rune 
		if string(value) == ";"{
			decoded_strs=append(decoded_strs,currentchar)
			currentchar=""
			continue
		}
		//value is a rune there is no automatic type conversion so convert first
		currentchar+=string(value)
		
	}
	return decoded_strs
}
