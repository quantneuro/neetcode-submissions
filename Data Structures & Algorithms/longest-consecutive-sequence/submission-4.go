// func longestConsecutive(nums []int) int {
// 	check:=make(map[int]bool)

	

// 	for _,val:=range nums{
// 		check[val]=true
// 	}
// 	maxx:=0
	
// 	for _,value:= range nums{
// 		count:=0
//         //check if a lower values exist if so then keep moving forward
// 		if !check[value-1]{
//             //if no lower value is there 
//             //make count increase for that value 
// 			count++
// 			for check[value +1]{//keep checking for continous values 
// 				value++//for checking continous values 
// 				count++//current window lenght
// 			}
// 		}
// 		maxx=max(count,maxx)//current length vs maximum so far
// 		}
// 		return maxx

// }

func longestConsecutive(nums []int) int {
	check:=make(map[int]bool)

	

	for _,val:=range nums{
		check[val]=true
	}
	maxx:=0
	
	for _,value:= range nums{
		count:=0
        //check if a lower values exist if so then keep moving forward
		if check[value+1]{
			for check[value]{//keep checking for continous values 
				value--//for checking continous values 
				count++//current window lenght
			}
        }
		
		maxx=max(count,maxx)//current length vs maximum so far
    }
		return maxx

}



