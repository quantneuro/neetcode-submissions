class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
    //     unordered_map<string,vector<string>> Stored;
    //     for(const auto& s:strs){
    //         string Storeds =s;
    //         sort(Storeds.begin(),Storeds.end());
    //         Stored[Storeds].push_back(s);
    //     }
    //     vector<vector<string>> result;
    //     for(auto& pair: Stored){
    //         result.push_back(pair.second);
    //     }
    //     return result;
    // }

    unordered_map<string,vector<string>> Stored;
    
    for(const auto& s:strs){
        
        vector<int>val(26,0);
        for(auto c:s){
            val[c-'a']++;
        }
        string key=to_string(val[0]);
        for(int i=0;i<26;i++){
            key+=','+to_string(val[i]);
        }
        Stored[key].push_back(s);

    }
    vector <vector<string>> result;
    for(const auto &pair:Stored){
        result.push_back(pair.second);

    }
    return result;

    }
};
