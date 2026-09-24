// class Solution {
// public:
//     bool isAnagram(string s, string t) {
//         if(s.length()!=t.length())
//             return false;
//         sort(s.begin(),s.end());
//         sort(t.begin(),t.end());
//         return s==t;
//     }
// };

class Solution{
    public:
    bool isAnagram(string s, string t){
    if(s.length() != t.length())
        return false;

    vector<int> freq(26,0);//just make one no need for two

    for(int i = 0; i<s.length();i++){
        freq[s[i]-'a']++;//adds +ve numbers
        freq[t[i]-'a']--;//adds -ve numbers
    }
    //now just check if in freq all values are zero
    for(int i=0;i<26;i++){
        if(freq[i]!=0)
            return false;
    }
    return true;
    }
};
