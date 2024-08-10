package main

// 1348. Tweet Counts Per Frequency
// A social media company is trying to monitor activity on their site by analyzing the number of tweets that occur in select periods of time. 
// These periods can be partitioned into smaller time chunks based on a certain frequency (every minute, hour, or day).

// For example, the period [10, 10000] (in seconds) would be partitioned into the following time chunks with these frequencies:
//     Every minute (60-second chunks): [10,69], [70,129], [130,189], ..., [9970,10000]
//     Every hour (3600-second chunks): [10,3609], [3610,7209], [7210,10000]
//     Every day (86400-second chunks): [10,10000]

// Notice that the last chunk may be shorter than the specified frequency's chunk size and will always end with the end time of the period (10000 in the above example).

// Design and implement an API to help the company with their analysis.

// Implement the TweetCounts class:
//     TweetCounts() 
//         Initializes the TweetCounts object.
//     void recordTweet(String tweetName, int time) 
//         Stores the tweetName at the recorded time (in seconds).
//     List<Integer> getTweetCountsPerFrequency(String freq, String tweetName, int startTime, int endTime) 
//         Returns a list of integers representing the number of tweets with tweetName in each time chunk for the given period of time [startTime, endTime] (in seconds) and frequency freq.
//             freq is one of "minute", "hour", or "day" representing a frequency of every minute, hour, or day respectively.

// Example:
// Input
// ["TweetCounts","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency"]
// [[],["tweet3",0],["tweet3",60],["tweet3",10],["minute","tweet3",0,59],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]
// Output
// [null,null,null,null,[2],[2,1],null,[4]]
// Explanation
// TweetCounts tweetCounts = new TweetCounts();
// tweetCounts.recordTweet("tweet3", 0);                              // New tweet "tweet3" at time 0
// tweetCounts.recordTweet("tweet3", 60);                             // New tweet "tweet3" at time 60
// tweetCounts.recordTweet("tweet3", 10);                             // New tweet "tweet3" at time 10
// tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 59); // return [2]; chunk [0,59] had 2 tweets
// tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 60); // return [2,1]; chunk [0,59] had 2 tweets, chunk [60,60] had 1 tweet
// tweetCounts.recordTweet("tweet3", 120);                            // New tweet "tweet3" at time 120
// tweetCounts.getTweetCountsPerFrequency("hour", "tweet3", 0, 210);  // return [4]; chunk [0,210] had 4 tweets

// Constraints:
//     0 <= time, startTime, endTime <= 10^9
//     0 <= endTime - startTime <= 10^4
//     There will be at most 10^4 calls in total to recordTweet and getTweetCountsPerFrequency.

import "fmt"

type TweetCounts struct {
    Tweet map[string][]int // key: name val: []int times
}

func Constructor() TweetCounts {
    return TweetCounts{  Tweet: map[string][]int{}, }
}

func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
    if _, ok := this.Tweet[tweetName]; !ok {
        this.Tweet[tweetName] = []int{time}
    } else {
        this.Tweet[tweetName] = append(this.Tweet[tweetName], time)
    }
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
    interval := GetInterval(freq)
    if _, ok := this.Tweet[tweetName]; !ok {
        return []int{}
    }
    res := make([]int, (endTime - startTime + interval) / interval)
    for _, time := range this.Tweet[tweetName] {
        if time > endTime || time < startTime {
            continue
        }
        thisInterval := (time-startTime) / interval
        res[thisInterval]++
    }
    return res
}

func GetInterval(freq string) int {
    if freq == "minute" { return 60 }
    if freq == "hour" { return 3600 }
    if freq == "day" { return 86400 }
    return 0 // never reach here
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */

func main() {
    // TweetCounts tweetCounts = new TweetCounts();
    obj := Constructor()
    fmt.Println(obj)
    // tweetCounts.recordTweet("tweet3", 0);                              // New tweet "tweet3" at time 0
    obj.RecordTweet("tweet3",0)
    fmt.Println(obj)
    // tweetCounts.recordTweet("tweet3", 60);                             // New tweet "tweet3" at time 60
    obj.RecordTweet("tweet3",60)
    fmt.Println(obj)
    // tweetCounts.recordTweet("tweet3", 10);                             // New tweet "tweet3" at time 10
    obj.RecordTweet("tweet3",10)
    fmt.Println(obj)
    // tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 59); // return [2]; chunk [0,59] had 2 tweets
    fmt.Println(obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)) // [2]
    // tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 60); // return [2,1]; chunk [0,59] had 2 tweets, chunk [60,60] had 1 tweet
    fmt.Println(obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)) // [2,1]
    // tweetCounts.recordTweet("tweet3", 120);                            // New tweet "tweet3" at time 120
    obj.RecordTweet("tweet3",120)
    fmt.Println(obj)
    // tweetCounts.getTweetCountsPerFrequency("hour", "tweet3", 0, 210);  // return [4]; chunk [0,210] had 4 tweets
    fmt.Println(obj.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210)) // [4]
}