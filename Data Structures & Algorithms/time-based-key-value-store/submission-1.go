type MapValue struct{
	val string
	ts int
}

type TimeMap struct {
	tm map[string][]*MapValue

}

func Constructor() TimeMap {
	return TimeMap{
		tm:make(map[string][]*MapValue),
	}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	this.tm[key]=append(this.tm[key],
		&MapValue{
			val:value,
			ts:timestamp,
		},
	)

}

func (this *TimeMap) Get(key string, timestamp int) string {
	_,e:=this.tm[key]
	
	if !e{return ""}

	r := len(this.tm[key])-1
	l := 0
	mid:=0
	// found:=false

	for l<=r {
		mid=l+(r-l)/2
		if this.tm[key][mid].ts<=timestamp{
			l=mid+1
		}else if this.tm[key][mid].ts>timestamp{
			r=mid-1
		}
		// else{
		// 	found=true
		// 	break
		// }
	}
	if r>=0 {
		return this.tm[key][r].val
	}else{
		return ""
	}




}
