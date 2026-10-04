func letterCombinations(digits string) []string {
	mapped:=map[rune][]string{
		'2':{"a","b","c"},
		'3':{"d","e","f"},
		'4':{"g","h","i"},
		'5':{"j","k","l"},
		'6':{"m","n","o"},
		'7':{"p","q","r","s"},
		'8':{"t","u","v"},
		'9':{"w","x","y","z"},
	}
	res:=[]string{}

	var comb func(d string,curs string)

	comb = func(d string,curs string){
		if len(d)==0{
			if len(curs)>0{
			res=append(res,curs)
			}
			return
		}

		current:=d[0]
		currentlist:=mapped[rune(current)]

		for i:=0;i<len(currentlist);i++{
			curs+=currentlist[i]
			comb(d[1:],curs)
			curs=curs[:len(curs)-1]
		}
	}

	comb(digits,"")
	return res

}
