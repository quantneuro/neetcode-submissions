
// func isPalindrome(s string) bool {
//     characters := []rune(strings.ToLower(s))
//     left :=0
//     right := len(characters)-1

//     for left < right{
        

//         if !unicode.IsLetter(characters[left])&&!unicode.IsDigit(characters[left]){
//             left++
//             continue
//         }
//         if !unicode.IsLetter(characters[right])&&!unicode.IsDigit(characters[right]){
//             right--
//             continue
//         }
//         if characters[left] != characters[right]{
//             return false
//         }

//         left++
//         right--

//         // leftCharacter :=unicode.Tolower(characters[left])
//         // rightCharacter := unicode.Tolower(characters[right])

//     }
//     return true
// }

func isPalindrome(s string) bool {
    s=strings.ToLower(s)

    cleaned := ""

    for _,character :=range s {
        if unicode.IsLetter(character) || unicode.IsDigit(character){
            cleaned += string(character)
        }
    }
    reversed := ""

    for index := len(cleaned)-1;index >=0;index--{
        reversed += string(cleaned[index])
    }
    return cleaned ==reversed 
}