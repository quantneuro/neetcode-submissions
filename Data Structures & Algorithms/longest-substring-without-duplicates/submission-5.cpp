class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        if(s.size()==0) return s.size();
        
        unordered_map<char,bool> m;
        string final="";
        int high=0;
        int l=0;
        int r=0;
        
        // for(int i=0;i<s.size();i++)
        while(r<s.size())
        {
            while(m.find(s[r])!=m.end()){
                if(high<final.size()){
                    high=final.size();
                }
                m.erase(s[l]);
                l++;
                final.erase(0,1);
            }
            m[s[r]]=true;
            final += s[r];
            r++;
        }
        return high;
    }
};
