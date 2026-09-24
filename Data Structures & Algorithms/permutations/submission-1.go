func perm(nums[]int,current []int, res *[][]int){
	if len(nums)==0{
		temp:=make([]int,len(current))
		copy(temp,current)
		*res=append(*res,temp)
		return
	}

	for i:=0;i<len(nums);i++{

	current=append(current,nums[i])
	tempnums:=[]int{}
	tempnums=append(tempnums,nums[:i]...)
	tempnums=append(tempnums,nums[i+1:]...)
	perm(tempnums,current,res)
	current=current[:len(current)-1]

	}

}

func permute(nums []int) [][]int {
res:=[][]int{}
perm(nums,[]int{},&res)
return res
}
