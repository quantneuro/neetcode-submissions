func productExceptSelf(nums []int) []int {
 prefix:=make([]int ,len(nums))
 suffix:=make([]int ,len(nums))

for i:=range prefix{
	prefix[i]=1
}

for i:=range suffix{
	suffix[i]=1
}

p:=1
s:=1

 for i:=range nums{
	if i-1<0{
		p=p*1
	}else{
	p=p*nums[i-1]
	}
	prefix[i] = p
 }



 for i:=len(nums)-1;i>=0;i--{
	if i+1>len(nums)-1{	
		s=s*1

	}else{
	s = s*nums[i+1]
	}
	
	suffix[i]=s

 }


 ans:=make([]int, len(nums))

 for i:=range ans{
	ans[i]=suffix[i]*prefix[i]
 }
 return ans


}
