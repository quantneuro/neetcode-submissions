func help(nums [] int,current []int,result *[][]int ){
	if len(nums)==0{
		temp:=make([]int,len(current))
		copy(temp,current)
		*result=append(*result,temp)
		return
		}

	for i:=0;i<len(nums);i++{
	current=append(current,nums[i])
	tempnums:=[]int{}
	tempnums=append(tempnums,nums[:i]...)
	tempnums=append(tempnums,nums[i+1:]...)
	help(tempnums,current,result)
	current=current[:len(current)-1]
	}

}
func permute(nums []int) [][]int {
	result:=[][]int{}
	help(nums,[]int{},&result)

	return result

}
