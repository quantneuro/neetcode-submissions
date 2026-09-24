class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string,vector<string>> Stored;
        for(const auto& s:strs){
            string Storeds =s;
            sort(Storeds.begin(),Storeds.end());
            Stored[Storeds].push_back(s);
        }
        vector<vector<string>> result;
        for(auto& pair: Stored){
            result.push_back(pair.second);
        }
        return result;
    }
};
