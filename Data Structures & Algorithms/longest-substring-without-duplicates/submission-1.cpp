class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        if(s.size()<1) return s.size();
        
        unordered_map<char,bool> m;
        string final="";
        int max=0;
        int l=0;
        int r=1;
        
        // for(int i=0;i<s.size();i++)
        while(r<s.size())
        {
            if(m.find(s[r])!=m.end()){
                if(max<final.size()){
                    max=final.size();
                }
                l++;
                final="";
            }
            m[s[r]]=true;
            final += s[r];
            r++;
        }
        return max;
    }
};
