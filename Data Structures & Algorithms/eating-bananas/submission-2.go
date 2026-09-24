import "slices"
func minEatingSpeed(piles []int, h int) int {
	l:=1
	r:=slices.Max(piles)

	var mid int = 0
	var speed int = math.MaxInt
	var testsum float64 = 0

	for l<=r{
		mid=l+(r-l)/2
		testsum = 0
	   for  _,value := range piles{
			testsum += math.Ceil(float64(value)/float64(mid))
	   }
		
		if testsum<=float64(h) {
			r=mid-1
			speed=min(mid,speed)
		}else if testsum>float64(h) {
			l=mid+1
		}
		
	}
	return speed
}
