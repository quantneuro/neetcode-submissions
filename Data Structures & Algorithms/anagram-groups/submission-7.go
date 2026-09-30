func groupAnagrams(strs []string) [][]string {
    groups:=make(map[[26]int][]string)
    
    for _,v:=range strs{
        freq:=[26]int{}
        for _,val:=range v{
            freq[val-'a']++
        }
        groups[freq]=append(groups[freq],v)
    }
    res:=[][]string{}

    for _,v:=range groups{
        res=append(res,v)
    }

    return res

}
