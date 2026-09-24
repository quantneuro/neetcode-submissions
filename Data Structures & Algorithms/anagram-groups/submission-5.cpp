// class Solution {
// public:
    
//     vector<vector<string>> groupAnagrams(vector<string>& strs) {
//         vector<vector<string>> groups;

//         for(int i=0;i<strs.size();i++){

        
//         string currentword=strs[i];
//         sort(currentword.begin(),currentword.end());
//         int j=0;
//         for(;j < groups.size();j++){
//             string wordtocompare = groups[j][0];
//             sort(wordtocompare.begin(),wordtocompare.end());

//             if(currentword==wordtocompare){
//                 groups[j].push_back(strs[i]);
//                 break;//stop the inner loop and goes out 
//             }
//         }
//         if(j == groups.size()){
//             groups.push_back({strs[i]});
//         }
//     }

//     return groups;
    
// }
// };

// class Solution{
// public:
//     vector<vector<string>> groupAnagrams(vector<string>& strs){
//         unordered_map<string,vector<string>> groups;
//         for(int i=0;i<strs.size();i++){
//             string key =strs[i];
//             sort(key.begin(),key.end());
//             //if same key hence same word so only word will be appended !
//             groups[key].push_back(strs[i]);
//         }
//         vector<vector<string>> answer;
//         //second value in map is the list of words
//         for(auto & entry :groups){
//             answer.push_back(entry.second);
//         }
//         return answer;
//     }

// };

class Solution{
    public:
    vector<vector<string>> groupAnagrams(vector<string>& strs){

    // unordedered_map<vector<int>(26,0),vector<string>> groups;
    //can do that but vector as a key needs it's own hashin
    unordered_map<string ,vector<string>> groups;
    vector<vector<string>> result;
    for(string val :strs){
        vector<int> freq(26,0);//26 values index from 0 1..24 25
        for(char c: val){
            freq[c -'a']++;
        }
        //convert that to a key typed string

        string key=to_string(freq[0]);//first value is started

        for(int i=1;i<26;++i){
            key+=','+to_string(freq[i]);//make the string included each value from the freq vector
        }
        
        groups[key].push_back(val);
    }
    for(const auto &pair : groups){
        result.push_back(pair.second);
    }
    return result;

}
};
















