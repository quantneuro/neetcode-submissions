class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size()!=t.size()){
            return false;
        }
        // sort(s.begin(),s.end());
        // sort(t.begin(),t.end());
        unordered_map<char,int> countS;
        unordered_map<char,int> countT;
        for(int i=0;i<s.size();i++){
            countS[s[i]]++;
            countT[t[i]]++;
        }


        // for(int i=0;i<s.size();i++){
        //     if(s[i]!=t[i]){
        //         return false;
        //     }
        // }
        // return true;
        return countS==countT;

    }
};
