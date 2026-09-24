class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<string>st;
        
        for(string c:tokens){
            if(c == "+"||c == "-" || c == "/"||c == "*"){
                // int a=st.top()-'0';
                int a= stoi(st.top());
                st.pop();
                // int b=st.top()='0';
                int b = stoi(st.top());
                st.pop();
                int res = 0;
                switch(c[0]){//c[0 works becasue the the operand string is only one lenght]
                    case '+': res=a+b; break;
                    case '-': res=b-a;break;
                    case '*': res=a*b;break;
                    case '/': res=b/a;break;
                }
                st.push(to_string(res)); }
            else{
                st.push(c);
                }
            

        }
        return stoi(st.top());
    }
};
