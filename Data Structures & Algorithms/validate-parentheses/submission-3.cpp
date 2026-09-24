class Solution {
public:
    bool isValid(string s) {
        stack<char>check;
        

        for(char c : s){
            if(c=='(' || c=='[' || c=='{')
                check.push(c);
            else{
                if(check.empty()) return false;
                if(c ==')' && check.top()!='(') return false;
                if(c==']' && check.top()!='[') return false;
                if(c=='}' && check.top() !='{') return false;

                check.pop();
            }
            
        }
        return check.empty();
    }
};
