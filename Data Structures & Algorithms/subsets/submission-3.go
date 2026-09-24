func help(nums []int,subset *[][]int,currentset []int){
    if len(nums)==0{//base case
        temp:=make([]int,len(currentset))
        copy(temp, currentset)
        *subset=append(*subset,temp)

        return
    }
    currentset=append(currentset,nums[0])//choose the first value

    help(nums[1:],subset,currentset)//explore
    currentset=currentset[:len(currentset)-1]//unchoose
    help(nums[1:],subset,currentset)//explore


}
func subsets(nums []int) [][]int {
    subset:=[][]int{}//to collect all subsets from backtracking
    help(nums,&subset,[]int{})
    return subset

}
