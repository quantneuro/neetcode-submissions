// import "slices"

// func groupAnagrams(strs []string) [][]string {
//     groups:=[][]string{}

//     for i:=0;i<len(strs);i++{

//         currentChar:=[]byte(strs[i])
//         slices.Sort(currentChar)
//         currentsorted:=string(currentChar)
//         j:=0
//         for ;j<len(groups);j++{
//             wordtocompare:=[]byte(groups[j][0])
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
//         }
//     }
//     return groups
// }


import "slices"

func groupAnagrams(strs [] string) [][]string{

    groups :=make(map[string][]string)

    for i:=0;i<len(strs);i++{
        characters:=[]byte(strs[i])

        slices.Sort(characters)
        key:=string(characters)

        groups[key] = append (groups[key],strs[i])
    }
    answer :=[][]string{}

    for _,group := range groups{
        answer=append(answer,group)
    }
    return answer


}