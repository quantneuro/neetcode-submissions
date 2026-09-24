import("slices")
func subsetsWithDup(nums []int) [][]int {
slices.Sort(nums)
res:=[][]int{}
var final func(nums []int, current[]int)

final=func(nums []int, current []int){
	if len(nums)==0{
		temp:=make([]int,len(current))
		copy(temp,current)
		res=append(res,temp)
		return
	}

	current=append(current,nums[0])
	final(nums[1:],current)
	for len(nums)>1 && nums[0]==nums[1]{
		nums=nums[1:]
	}
	current=current[:len(current)-1]
	final(nums[1:],current)
}
final(nums,[]int{})
return res


}
