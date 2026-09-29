func subsets(nums []int) [][]int {

    res:=[][]int{}
    var backtrack func(nums []int,cur []int)

    backtrack=func(nums []int, cur []int){
        if len(nums)==0{
            temp:=make([]int,len(cur))
            copy(temp,cur)
            res=append(res,temp)
            return
        }

        cur=append(cur,nums[0])
        nums=nums[1:]
        backtrack(nums,cur)
        cur=cur[:len(cur)-1]
        backtrack(nums,cur)
    }
    backtrack(nums,[]int{})

    return res
}
