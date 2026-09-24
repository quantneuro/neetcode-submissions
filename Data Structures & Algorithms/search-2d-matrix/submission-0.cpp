bool binarysearch(int row,int column,int t,vector<vector<int>> &matrix){
            int l =0;
            int r=column;
            while(l<=r){
                int mid=l+(r-l)/2;
                int value=matrix[row][mid];
                
                if (t<value){
                    r=mid-1;
                }else if(t>value){
                    l=mid+1;
                }
                else{
                    return true;
                }
        }
        return false;
        }

class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {

        
        int columns = matrix[0].size()-1;

        int rl=0;
        int rr = matrix.size()-1;
        while(rl<=rr){
            int mid = rl+(rr-rl)/2;
            
            
            if(target>matrix[mid][columns]){
                    rl=mid+1;
                }
            else if (target<matrix[mid][0]){
                    rr=mid-1;
                }

            else{
                return binarysearch(mid,columns,target,matrix);
            }
        
        }
        return false;
    }
};
