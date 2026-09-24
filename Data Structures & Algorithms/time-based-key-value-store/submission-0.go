type timevalue struct{
	value string
	timestamp int
}
type TimeMap struct {
	//string is the key
	//time value is the object
	//[] before the timevalue means slice of timevalue objects
	data map[string][]timevalue

}

func Constructor() TimeMap {
	return TimeMap{
		data : make(map[string][]timevalue),
	}
}

//return by pointer
// func Constructor() *TimeMap {
// 	return &TimeMap{
		
// 	}
// }

func (this *TimeMap) Set(key string, value string, timestamp int) {
	item:= timevalue{
		value:value,
		timestamp:timestamp,
	}
	this.data[key]=append(this.data[key],item)
}

func (this *TimeMap) Get(key string, timestamp int) string {

	values:=this.data[key]//values is a slice of timevalue objects

	l:=0
	r:=len(values)-1
	mid:=0
	answer:=""
	for l<=r{
		mid=l+(r-l)/2

		if values[mid].timestamp <=timestamp{
			answer =values[mid].value//save the current value
			l=mid+1//move up the l so to find out any new larger timestamp
		}else {
			r=mid-1
		}
	}
	return answer


}
