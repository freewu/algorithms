package main

// 933. Number of Recent Calls
// You have a RecentCounter class which counts the number of recent requests within a certain time frame.
// Implement the RecentCounter class:
//     RecentCounter() Initializes the counter with zero recent requests.
//     int ping(int t) 
//          Adds a new request at time t, where t represents some time in milliseconds, 
//          and returns the number of requests that has happened in the past 3000 milliseconds (including the new request). Specifically, 
//          return the number of requests that have happened in the inclusive range [t - 3000, t].

// It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.

// Example 1:
// Input
// ["RecentCounter", "ping", "ping", "ping", "ping"]
// [[], [1], [100], [3001], [3002]]
// Output
// [null, 1, 2, 3, 3]
// Explanation
// RecentCounter recentCounter = new RecentCounter();
// recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
// recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
// recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
// recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3
 
// Constraints:
//     1 <= t <= 109
//     Each test case will call ping with strictly increasing values of t.
//     At most 10^4 calls will be made to ping.

import "fmt"

type RecentCounter struct {
    requests []int
}

func Constructor() RecentCounter {
    return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
    // 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，
    this.requests = append(this.requests, t)
    // Remove timestamps that are outside the 3000 milliseconds window
    for this.requests[0] < t - 3000 {
        this.requests = this.requests[1:]
    }
    // 并返回过去 3000 毫秒内发生的所有请求数（包括新请求）
    return len(this.requests)
}


type RecentCounter1 []int

func Constructor1() (_ RecentCounter1) { return }

func (q *RecentCounter1) Ping(t int) int {
    *q = append(*q, t)
    for (*q)[0] < t-3000 {
        *q = (*q)[1:]
    }
    return len(*q)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {
    // RecentCounter recentCounter = new RecentCounter();
    obj := Constructor()
    // recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
    fmt.Println(obj.Ping(1)) // 1
    // recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
    fmt.Println(obj.Ping(100)) // 2
    // recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
    fmt.Println(obj.Ping(3001)) // 3
    // recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3
    fmt.Println(obj.Ping(3002)) // 3

    obj1 := Constructor1()
    fmt.Println(obj1.Ping(1)) // 1
    fmt.Println(obj1.Ping(100)) // 2
    fmt.Println(obj1.Ping(3001)) // 3
    fmt.Println(obj1.Ping(3002)) // 3
}