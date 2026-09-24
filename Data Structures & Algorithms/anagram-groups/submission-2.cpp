class Solution {
public:
    
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> groups;

        for(int i=0;i<strs.size();i++){

        
        string currentword=strs[i];
        sort(currentword.begin(),currentword.end());
        int j=0;
        for(;j < groups.size();j++){
            string wordtocompare = groups[j][0];
            sort(wordtocompare.begin(),wordtocompare.end());

            if(currentword==wordtocompare){
                groups[j].push_back(strs[i]);
                break;//stop the inner loop and goes out 
            }
        }
        if(j == groups.size()){
            groups.push_back({strs[i]});
        }
    }

    return groups;
    
}
};
