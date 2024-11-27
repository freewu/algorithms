package main

// 2054. Two Best Non-Overlapping Events
// You are given a 0-indexed 2D integer array of events where events[i] = [startTimei, endTimei, valuei]. 
// The ith event starts at startTimei and ends at endTimei, and if you attend this event, you will receive a value of valuei. 
// You can choose at most two non-overlapping events to attend such that the sum of their values is maximized.

// Return this maximum sum.

// Note that the start time and end time is inclusive: that is, you cannot attend two events where one of them starts and the other ends at the same time. 
// More specifically, if you attend an event with end time t, the next event must start at or after t + 1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/picture5.png" />
// Input: events = [[1,3,2],[4,5,2],[2,4,3]]
// Output: 4
// Explanation: Choose the green events, 0 and 1 for a sum of 2 + 2 = 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/picture1.png" />
// Example 1 Diagram
// Input: events = [[1,3,2],[4,5,2],[1,5,5]]
// Output: 5
// Explanation: Choose event 2 for a sum of 5.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/09/21/picture3.png" />
// Input: events = [[1,5,3],[1,5,1],[6,6,5]]
// Output: 8
// Explanation: Choose events 0 and 2 for a sum of 3 + 5 = 8.

// Constraints:
//     2 <= events.length <= 10^5
//     events[i].length == 3
//     1 <= startTimei <= endTimei <= 10^9
//     1 <= valuei <= 10^6

import "fmt"
import "sort"
import "container/heap"

func maxTwoEvents(events [][]int) int {
    if len(events) == 0 { return 0 }
    typed := make([][]int, 2*len(events))
    for i, v := range events {
        typed[2*i] = []int{ v[0], 1, v[2] }
        typed[2*i+1] = []int{ v[1] + 1, -1, v[2] }
    }
    sort.Slice(typed, func(i, j int) bool {
        if typed[i][0] == typed[j][0] {
            if typed[i][1] == typed[j][1] { return typed[i][2] < typed[j][2] }
            return typed[i][1] < typed[j][1]
        }
        return typed[i][0] < typed[j][0]
    })
    currOptimal, res := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(typed); i++  {
        switch typed[i][1] {
        case 1:
            res = max(res, currOptimal + typed[i][2])
        case -1:
            currOptimal = max(currOptimal, typed[i][2])
        }
    }
    return res
}

type Pair struct {
    end, cost int
}

type MinHeap []Pair
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].end < h[j].end }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() any {
    old := *h 
    n := len(old)
    res := old[n-1]
    *h = old[:n-1]
    return res
}

func maxTwoEvents1(events [][]int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
    })
    mhp := &MinHeap{}
    res, mx := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range events {
        l, r, c := v[0], v[1], v[2]
        for mhp.Len() > 0 && (*mhp)[0].end < l {
            t := heap.Pop(mhp).(Pair)
            mx = max(mx, t.cost)
        }
        res = max(res, c + mx)
        heap.Push(mhp, Pair{r, c})
    }
    return res
}

func maxTwoEvents2(events [][]int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    res, n := 0, len(events)
    facts := make([]int, n + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        facts[i] = max(facts[i+1], events[i][2])
    }
    for _, e := range events {
        v, left, right := e[2], 0, n
        for left < right {
            mid := (left + right) >> 1
            if events[mid][0] > e[1] {
                right = mid
            } else {
                left = mid + 1
            }
        }
        if left < n {
            v += facts[left]
        }
        res = max(res, v)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/picture5.png" />
    // Input: events = [[1,3,2],[4,5,2],[2,4,3]]
    // Output: 4
    // Explanation: Choose the green events, 0 and 1 for a sum of 2 + 2 = 4.
    fmt.Println(maxTwoEvents([][]int{{1,3,2},{4,5,2},{2,4,3}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/picture1.png" />
    // Example 1 Diagram
    // Input: events = [[1,3,2],[4,5,2],[1,5,5]]
    // Output: 5
    // Explanation: Choose event 2 for a sum of 5.
    fmt.Println(maxTwoEvents([][]int{{1,3,2},{4,5,2},{1,5,5}})) // 5
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/09/21/picture3.png" />
    // Input: events = [[1,5,3],[1,5,1],[6,6,5]]
    // Output: 8
    // Explanation: Choose events 0 and 2 for a sum of 3 + 5 = 8.
    fmt.Println(maxTwoEvents([][]int{{1,5,3},{1,5,1},{6,6,5}})) // 8

    fmt.Println(maxTwoEvents1([][]int{{1,3,2},{4,5,2},{2,4,3}})) // 4
    fmt.Println(maxTwoEvents1([][]int{{1,3,2},{4,5,2},{1,5,5}})) // 5
    fmt.Println(maxTwoEvents1([][]int{{1,5,3},{1,5,1},{6,6,5}})) // 8

    fmt.Println(maxTwoEvents2([][]int{{1,3,2},{4,5,2},{2,4,3}})) // 4
    fmt.Println(maxTwoEvents2([][]int{{1,3,2},{4,5,2},{1,5,5}})) // 5
    fmt.Println(maxTwoEvents2([][]int{{1,5,3},{1,5,1},{6,6,5}})) // 8
}