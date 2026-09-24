
func isPalindrome(s string) bool {
    characters := []rune(strings.ToLower(s))
    left :=0
    right := len(characters)-1

    for left < right{
        

        if !unicode.IsLetter(characters[left])&&!unicode.IsDigit(characters[left]){
            left++
            continue
        }
        if !unicode.IsLetter(characters[right])&&!unicode.IsDigit(characters[right]){
            right--
            continue
        }
        if characters[left] != characters[right]{
            return false
        }

        left++
        right--

        // leftCharacter :=unicode.Tolower(characters[left])
        // rightCharacter := unicode.Tolower(characters[right])

    }
    return true
}
