import "slices"
func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    sCharacters:= []byte(s)//I can use rune here, when I use byte I assume all are letters only
    tCharacters:= []byte(t)

    // sort.Slice(sCharacters,func(i int , j int)){//this is old way and here sort picks two positions 
    // //to rest and those are indices passed to i and j then return compares them 
    //     return sCharacters[i] < sCharacters[j]
    // }
    slices.Sort(sCharacters)
    // sort.Slice(tCharacters, func(i int , j int)){
    //     return tCharacters[i] > tCharacters[j]
    // }
    slices.Sort(tCharacters)

    return string(sCharacters) == string(tCharacters)
}
