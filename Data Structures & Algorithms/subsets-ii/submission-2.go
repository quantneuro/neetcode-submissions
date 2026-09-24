import(
	"slices"
)
func help(nums []int, current []int , result *[][]int){
	if len(nums)==0{
		temp:=make([]int,len(current))
		copy(temp,current)
		*result=append(*result,temp)
		return
	}
	current=append(current,nums[0])
	help(nums[1:],current,result)

	for (len(nums)>1 && nums[0]==nums[1]){
		nums=nums[1:]
	}
	current=current[:len(current)-1]
	
	help(nums[1:],current,result)


}
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	result:=[][]int{}
	help(nums,[]int{},&result)
	return result

}
