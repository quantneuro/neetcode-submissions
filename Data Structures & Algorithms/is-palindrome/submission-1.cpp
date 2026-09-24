// class Solution {
// public:
//     bool isPalindrome(string s) {
//         string cleaned="";
//         for(char character:s){
            
//             auto unsignedCharacter = static_cast<unsigned char>(character);
//             //why this? becuase function like isalnum and to lower intakes only letter and number(only positive)
//             // so that can be a issue here if number is negative:: static cast removes the the sign (-ve)
//             if(isalnum(unsignedCharacter)){
//                 cleaned +=tolower(unsignedCharacter);
//             }
//         }

//         string reversed="";

//         for(int index = cleaned.size()-1;index>=0;index--){
//             reversed +=cleaned[index];
//         }

//         return cleaned ==reversed;
//     }
// };

class Solution{
public:
    bool isPalindrome(string s){
        int left=0;
        int right =static_cast<int>(s.size())-1;

        while(left<right){
            while(left < right && !isalnum(s[left])){
                left++;
            }
            while(left<right && !isalnum(s[right])){
                right--;
            }
            if (tolower(s[left]) != tolower(s[right])){
                return false;
            }
        left++;
        right--;

        }
        return true;
        
    }

};
