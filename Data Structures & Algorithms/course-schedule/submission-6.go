//kahn's algo
func canFinish(numCourses int, prerequisites [][]int) bool{
	need:=make(map[int][]int)
	indegree:=make([]int,numCourses)
	queue:=[]int{}
	ans:=[]int{}
	//indegree of i mean the courses that need to taken before it
	for _,v:=range prerequisites {
		need[v[1]]=append(need[v[1]],v[0])
		indegree[v[0]]++
	}
	for i,v:=range indegree{
		if v==0{
			queue=append(queue,i)//i is index aka the course
		}

	}
	for len(queue)>0{
		course:=queue[0]
		queue=queue[1:]
		ans=append(ans,course)
		if v,ok:=need[course];ok{
			for i:=range v{
				indegree[v[i]]--
				if indegree[v[i]]==0{
					queue=append(queue,v[i])
				}
			}

		}
	}
	return len(ans)==numCourses



}
// func canFinish(numCourses int, prerequisites [][]int) bool {
//     need:=make(map[int][]int)
// 	check:=make(map[int]bool)//only for cycle check
// 	done:=make(map[int]bool)//to avoid recheking of something that is already seen


// 	for _,v:=range prerequisites{
// 			need[v[1]]=append(need[v[1]],v[0])
// 	}

// 	var dfs func(n int)bool

// 	dfs=func(n int)bool{
// 		if check[n]{
// 			return false
// 		}
// 		if done[n]{ return true}

// 		
// 		check[n]=true
// 		v,e:=need[n]

// 		if e{
// 			for i:=0;i<len(v);i++{
// 				result:=dfs(v[i])
// 				if !result { done[v[i]]=result;return false}
// 			}
// 		}
// // else{done[n]=true;check[n]=false ;return true }
// 		done[n]=true
// 		check[n]=false
		
// 		return true


// 	}

// 	for i:=0;i<numCourses;i++{
// 		result:=dfs(i)
// 		if !result { return false }
// 	}
// 	return true

// }
