// func hasDuplicate(nums []int) bool {
    
//     for i:=0; i<len(nums);i++{
//         for j:=i+1;j<len(nums);j++{
//             if nums[i] ==nums[j] {
//                 return true
//             }
//         }
//     }

//     return false
// }

func hasDuplicate(nums [] int) bool{
    seen := make(map[int]bool)
    for i:=0 ;i<len(nums);i++{
        if seen[nums[i]]{
            return true
        }
        seen[nums[i]]=true
    }
    return false
}