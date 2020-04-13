package main

import "fmt"

func main() {
	c := Constructor()
	c.PostTweet(1, 5)
	c.PostTweet(2, 3)
	c.PostTweet(1, 101)
	c.PostTweet(2, 13)
	c.PostTweet(2, 10)
	c.PostTweet(1, 2)
	c.PostTweet(1, 94)
	c.PostTweet(2, 505)
	c.PostTweet(1, 333)
	c.PostTweet(2, 22)
	c.PostTweet(1, 11)
	c.PostTweet(1, 205)
	c.PostTweet(2, 203)
	c.PostTweet(1, 201)
	c.PostTweet(2, 213)
	c.PostTweet(1, 200)
	c.PostTweet(2, 202)
	c.PostTweet(1, 204)
	c.PostTweet(2, 208)
	c.PostTweet(2, 233)
	fmt.Println(c.tuiwens)
	c.PostTweet(1, 222)
	c.PostTweet(2, 211)
	fmt.Println(c.GetNewsFeed(1))

}

type tuiwen struct {
	id     int
	userid int
}
type Twitter struct {
	attention map[int][]int
	tuiwens   []tuiwen
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	tw := Twitter{}
	tw.attention = make(map[int][]int)
	tw.tuiwens = []tuiwen{}
	return tw
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	tui := tuiwen{tweetId, userId}
	this.tuiwens = append(this.tuiwens, tui)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	a := this.attention[userId]
	a = append(a, userId)
	att := map[int]bool{}
	for _, u := range a {
		att[u] = true
	}
	ret := []int{}
	i := 0
	l := len(this.tuiwens)
	fmt.Println(l, att)
	for index := l - 1; index >= 0; index-- {
		fmt.Println(this.tuiwens[index].userid, 1)
		if att[this.tuiwens[index].userid] {
			ret = append(ret, this.tuiwens[index].id)
			i++
		}

		if i == 10 {
			break
		}
	}
	return ret
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.attention[followerId] = append(this.attention[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for i, id := range this.attention[followerId] {
		if id == followeeId {
			this.attention[followerId] = append(this.attention[followerId][:i], this.attention[followerId][i+1:]...)
			break
		}

	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
