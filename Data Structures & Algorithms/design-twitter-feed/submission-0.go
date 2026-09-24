// import "container/heap"
type someheap []tweets

func(h someheap)Len()int{return len(h)}
func(h someheap)Less(i,j int)bool{return h[i].time<h[j].time}
func(h someheap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h *someheap)Push(x any){
	value:=x.(tweets)
	*h=append(*h,value)
}
func(h *someheap)Pop()any{
	old:=*h
	n:=len(old)
	value:=old[n-1]
	*h=old[:n-1]
	return value
}

type member struct{
	mytweets []tweets
	followers []int
	followees []int
}
type tweets struct{
	tweetid int
	time int
}
type Twitter struct {
    users map[int]member
	time int
}


func Constructor() Twitter {
    return Twitter{
		users: make(map[int]member),
		time:0,
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    m := this.users[userId]

	this.time++

	m.mytweets = append(m.mytweets, tweets{
		tweetid: tweetId,
		time:    this.time,
	})

	this.users[userId] = m
}


// func (this *Twitter) GetNewsFeed(userId int) []int {
	
// 	personaltweets:=this.users[userId].mytweets
// 	var finallist []tweets
// 	finallist=append(finallist,personaltweets...)
// 	for _,v:=range this.users[userId].followees{
// 		finallist=append(finallist,this.users[v].mytweets...)
// 	}

// 	h:=someheap(finallist)
// 	heap.Init(&h)
    
// }
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &someheap{}
	heap.Init(h)

	// user's own tweets
	for _, tweet := range this.users[userId].mytweets {
		heap.Push(h, tweet)

		if h.Len() > 10 {
			heap.Pop(h) // remove oldest
		}
	}

	// tweets from everyone this user follows
	for _, followeeId := range this.users[userId].followees {
		for _, tweet := range this.users[followeeId].mytweets {
			heap.Push(h, tweet)

			if h.Len() > 10 {
				heap.Pop(h) // remove oldest
			}
		}
	}

	// min-heap gives oldest first, but feed needs newest first
	result := make([]int, h.Len())

	for i := len(result) - 1; i >= 0; i-- {
		tweet := heap.Pop(h).(tweets)
		result[i] = tweet.tweetid
	}

	return result
}



func (this *Twitter) Follow(followerId int, followeeId int)  {
	//two objects
	if followerId == followeeId { return }
    follower:=this.users[followerId]
	followee:=this.users[followeeId]


	for _, id := range follower.followees {
		if id == followeeId {
			return
		}
	}

	follower.followees=append(follower.followees,followeeId)
	followee.followers = append(followee.followers, followerId)

	this.users[followerId] = follower
	this.users[followeeId] = followee
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    follower := this.users[followerId]
	followee := this.users[followeeId]

	for i, id := range follower.followees {
		if id == followeeId {
			follower.followees = append(
				follower.followees[:i],
				follower.followees[i+1:]...,
			)
			break
		}
	}

	for i, id := range followee.followers {
		if id == followerId {
			followee.followers = append(
				followee.followers[:i],
				followee.followers[i+1:]...,
			)
			break
		}
	}

	this.users[followerId] = follower
	this.users[followeeId] = followee
}
