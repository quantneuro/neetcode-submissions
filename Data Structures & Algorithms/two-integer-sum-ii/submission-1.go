func twoSum(numbers []int, target int) []int {
	
	r:=len(numbers)-1
	l:=0
	for l<r{
		if target<numbers[l]+numbers[r]{
			r--
		}else if target>numbers[l]+numbers[r]{
			l++
		}else{
		
			return [] int{l+1,r+1}
		}

	}
	return [] int{}
}
