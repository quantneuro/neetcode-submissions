class MinStack {
private:
    vector<int>minstack;
    vector<int>mintracker;
public:
    
    MinStack() {
        
    }
    
    void push(int val) {
    if(mintracker.empty() || val<mintracker.back()) mintracker.push_back(val);
    else{  mintracker.push_back(mintracker.back());}
        minstack.push_back(val);

    }
    
    void pop() {
        minstack.pop_back();
        mintracker.pop_back();
    }
    
    int top() {

        int topval=minstack.back();
        return topval;
    }
    
    int getMin() {
        return mintracker.back();
    }
};
