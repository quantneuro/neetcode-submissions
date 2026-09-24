// func twoSum(nums []int, target int) []int {
//     for i:=0;i<len(nums);i++{
//         for j:=i+1;j<len(nums);j++{
//             if nums[i]+nums[j] == target {
//                 return []int{i,j}
//             }
//         }
//     }
//     return []int{}
// }


func twoSums(nums [] int, target int) []int{
    seen :=make(map[int]int)
    //NOTE THAT IN THIS KEY - VALUE pair key is the value and value is the index in the numes array
    for i:=0; i<lens(nums);i++ {
        needed := target -nums[i]
        //when a key is searched in a map in Go , it can return upto two values -> value and boolean
        index,exists :=seen[needed]
        //index is the value and exists get the bool value
        if exists {
            return []int{index,i}
        }
        seen[nums[i]] = i
    }
    return []int{}
}
