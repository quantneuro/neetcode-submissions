// import "slices"

// func groupAnagrams(strs []string) [][]string {
//     groups:=[][]string{}

//     for i:=0;i<len(strs);i++{

//         currentChar:=[]byte(strs[i])
//         slices.Sort(currentChar)
//         currentsorted:=string(currentChar)
//         j:=0
//         for ;j<len(groups);j++{
//             wordtocompare:=[]byte(groups[j][0])//need type conversion to byte since a string can't be sorted directly and also the j is row and 0 is column since in every row all stirng once sorted are same so just take the zeorth value
//             slices.Sort(wordtocompare)
//             wordtocomparesorted:=string(wordtocompare)

//             if currentsorted == wordtocomparesorted{
//                 //groups[j] is a single row, which is a slice of strings ([]string).strs[i] is a 
//                 //single string (string).The append function expects its first argument to be a slice, 
//                 //and its second argument to be the individual element you want to put inside that slice.
//                 groups[j] =append(groups[j],strs[i])
//                 break
//             }
//         }
//         if j==len(groups){
//             groups=append(groups,[]string{strs[i]})
                //To add something to a 2D slice, you must wrap your string inside a new slice container.
//         }
//     }
//     return groups
// }


// import "slices"

// func groupAnagrams(strs [] string) [][]string{

//     groups :=make(map[string][]string)

//     for i:=0;i<len(strs);i++{
//         characters:=[]byte(strs[i])

//         slices.Sort(characters)
//         key:=string(characters)

//         groups[key] = append (groups[key],strs[i])
//     }
//     answer :=[][]string{}

//     for _,group := range groups{
//         answer=append(answer,group)
//     }
//     return answer


// }

func groupAnagrams(strs []string)[][]string{
    answer := [][]string{}
    groups := make(map[[26]int][]string)//slices can't be keys so used array
    

    for _,val := range strs{
        freq := [26]int{}
        for i:=0;i<len(val);i++{
            freq[val[i]-'a']++
        }

       groups[freq]=append(groups[freq],val)
    }
    for _,val:=range groups{
        answer=append(answer,val)
    }
    return answer

}







