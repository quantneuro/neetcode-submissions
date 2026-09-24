func binarySearch(row int, target int, mat [][] int) bool {
    l:=0
    r:=len(mat[row])-1
    mid:=0

    for l<=r{
        mid=l+((r-l)/2)
        if target<mat[row][mid]{
            r=mid-1
        }else if target>mat[row][mid]{
            l=mid+1
        }else{
            return true
        }
    }
    return false
}

func searchMatrix(matrix [][]int, target int) bool {
    
    l:=0
    lastcolumn:=len(matrix[0])-1
    r:=len(matrix)-1
    mid:=0
    for l<=r{
        mid=l+((r-l)/2)
        if target <matrix[mid][0] {
            r=mid-1
        }else if target>matrix[mid][lastcolumn]{
            l=mid+1
        }else{
            return binarySearch(mid,target,matrix)
        }

    }
    return false;


}
