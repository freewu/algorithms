package main 

// 355. Design Twitter
// Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, 
// and is able to see the 10 most recent tweets in the user's news feed.

// Implement the Twitter class:
//     Twitter() 
//         nitializes your twitter object.
//     void postTweet(int userId, int tweetId) 
//         Composes a new tweet with ID tweetId by the user userId. 
//         Each call to this function will be made with a unique tweetId.
//     List<Integer> getNewsFeed(int userId) 
//         Retrieves the 10 most recent tweet IDs in the user's news feed. 
//         Each item in the news feed must be posted by users who the user followed or by the user themself. 
//         Tweets must be ordered from most recent to least recent.
//     void follow(int followerId, int followeeId) 
//         The user with ID followerId started following the user with ID followeeId.
//     void unfollow(int followerId, int followeeId) 
//         The user with ID followerId started unfollowing the user with ID followeeId.
    
// Example 1:
// Input
// ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
// [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
// Output
// [null, null, [5], null, null, [6, 5], null, [5]]
// Explanation
// Twitter twitter = new Twitter();
// twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
// twitter.follow(1, 2);    // User 1 follows user 2.
// twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
// twitter.unfollow(1, 2);  // User 1 unfollows user 2.
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
 
// Constraints:
//     1 <= userId, followerId, followeeId <= 500
//     0 <= tweetId <= 10^4
//     All the tweets have unique IDs.
//     At most 3 * 10^4 calls will be made to postTweet, getNewsFeed, follow, and unfollow.

import "fmt"
import "container/heap"

type Twitter struct {
    tweets    map[int]*Tweet
    follows   map[int]map[int]bool
    requestId int
}

type Tweet struct {
    tweetId   int
    previous  *Tweet
    timestamp int
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(1)
func Constructor() Twitter {
    return Twitter{
        tweets:    make(map[int]*Tweet),
        follows:   make(map[int]map[int]bool),
        requestId: 0,
    }
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) PostTweet(userId int, tweetId int) {
    twitter.InitUser(userId)
    twitter.requestId += 1
    twitter.tweets[userId] = &Tweet{
        tweetId:   tweetId,
        previous:  twitter.tweets[userId],
        timestamp: twitter.requestId,
    }
}

// Worst time complexity:   O(f log f) where f is the number of followers of the user.
// Average time complexity: O(a log a) where a is the average number of followers a user has.
// Space complexity:        O(f) where f is the number of followers of the user.
func (twitter *Twitter) GetNewsFeed(userId int) []int {
    twitter.InitUser(userId)
    newsFeed, tweetHeap := []int{}, &TweetMaxHeap{}
    heap.Init(tweetHeap)
    for folleweeId, isFollowing := range twitter.follows[userId] {
        if isFollowing {
            twitter.InitUser(folleweeId)
            heap.Push(tweetHeap, twitter.tweets[folleweeId])
        }
    }
    for len(*tweetHeap) > 0 && len(newsFeed) < 10 {
        nextTweet := heap.Pop(tweetHeap).(*Tweet)
        if nextTweet.timestamp == 0 {
            break
        }
        newsFeed = append(newsFeed, nextTweet.tweetId)
        heap.Push(tweetHeap, nextTweet.previous)
    }
    return newsFeed
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) Follow(followerId int, followeeId int) {
    twitter.InitUser(followerId)
    twitter.follows[followerId][followeeId] = true
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) Unfollow(followerId int, followeeId int) {
    twitter.InitUser(followerId)
    twitter.follows[followerId][followeeId] = false
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) InitUser(userId int) {
    if _, hasTweet := twitter.tweets[userId]; hasTweet {
        return
    }
    twitter.tweets[userId] = &Tweet{}
    twitter.follows[userId] = make(map[int]bool)
    twitter.follows[userId][userId] = true
}

type TweetMaxHeap []*Tweet

func (tweetHeap TweetMaxHeap) Len() int {
    return len(tweetHeap)
}

func (tweetHeap TweetMaxHeap) Less(i, j int) bool {
    return tweetHeap[i].timestamp > tweetHeap[j].timestamp
}

func (tweetHeap TweetMaxHeap) Swap(i, j int) {
    tweetHeap[i], tweetHeap[j] = tweetHeap[j], tweetHeap[i]
}

func (tweetHeap *TweetMaxHeap) Push(x any) {
    *tweetHeap = append(*tweetHeap, x.(*Tweet))
}

func (tweetHeap *TweetMaxHeap) Pop() any {
    current := *tweetHeap
    value := current[len(current)-1]
    *tweetHeap = current[:len(current)-1]
    return value
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {
    // Twitter twitter = new Twitter();
    obj := Constructor()
    fmt.Println(obj)
    // twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
    obj.PostTweet(1,5)
    fmt.Println(obj)
    // twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
    fmt.Println(obj.GetNewsFeed(1)) // [5]
    // twitter.follow(1, 2);    // User 1 follows user 2.
    obj.Follow(1,2)
    fmt.Println(obj)
    // twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
    obj.PostTweet(2,6)
    fmt.Println(obj)
    // twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
    fmt.Println(obj.GetNewsFeed(1)) // [6, 5]
    // twitter.unfollow(1, 2);  // User 1 unfollows user 2.
    obj.Unfollow(1,2)
    fmt.Println(obj)
    // twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
    fmt.Println(obj.GetNewsFeed(1)) // [5]
}