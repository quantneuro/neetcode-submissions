func subsets(nums []int) [][]int {
	var result [][]int
    var choose []int
    helper(0,nums,choose,&result)
    return result

}

func helper(i int, nums []int, choose []int , result *[][]int ) {

//base case
if i==len(nums){
    temp :=make([]int, len(choose))
    copy(temp,choose)
    *result = append(*result,temp)
    return 
}


//skip
helper(i+1,nums,choose,result)

//choose
choose = append(choose,nums[i])
//explore
helper(i+1,nums,choose,result)

//unchoose
choose = choose[:len(choose)-1]


}
