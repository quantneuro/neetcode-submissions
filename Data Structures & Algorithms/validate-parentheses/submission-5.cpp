class Solution {
public:
    bool isValid(string s) {
        stack<char>check;
        
        unordered_map<char,char> cm={
            {']','['},
            {'}','{'},
            {')','('}
        };
        for(char c : s){
            if(cm.count(c)){
                if(!stack.empty() && cm[c]==check.top()){
                    check.pop();
                }
                else{
                    return false;
                }
            }
            else{
                check.push(c);
            }

            // if(c=='(' || c=='[' || c=='{')
            //     check.push(c);
            // else{
            //     if(check.empty()) return false;
            //     if(c ==')' && check.top()!='(') return false;
            //     if(c==']' && check.top()!='[') return false;
            //     if(c=='}' && check.top() !='{') return false;

            //     check.pop();
            // }
            
        }
        return check.empty();
    }
};
