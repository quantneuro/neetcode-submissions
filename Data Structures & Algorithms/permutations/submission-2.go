func permute(nums []int) [][]int {
res:=[][]int{}
var total func(nums []int,current []int)

total=func(nums []int,current []int){
	if len(nums)==0{
		temp:=make([]int,len(current))
		copy(temp,current)
		res=append(res,temp)
		return
	}

	for i:=0;i<len(nums);i++{
		current=append(current,nums[i])
		tempnums:=[]int{}
		tempnums=append(tempnums,nums[:i]...)
		tempnums=append(tempnums,nums[i+1:]...)
		total(tempnums,current)
		current=current[:len(current)-1]
	}

}
total(nums,[]int{})
return res
}
