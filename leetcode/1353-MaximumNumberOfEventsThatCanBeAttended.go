package main

// 1353. Maximum Number of Events That Can Be Attended
// You are given an array of events where events[i] = [startDayi, endDayi]. 
// Every event i starts at startDayi and ends at endDayi.

// You can attend an event i at any day d where startTimei <= d <= endTimei. 
// You can only attend one event at any time d.

// Return the maximum number of events you can attend.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/05/e1.png" />
// Input: events = [[1,2],[2,3],[3,4]]
// Output: 3
// Explanation: You can attend all the three events.
// One way to attend them all is as shown.
// Attend the first event on day 1.
// Attend the second event on day 2.
// Attend the third event on day 3.

// Example 2:
// Input: events= [[1,2],[2,3],[3,4],[1,2]]
// Output: 4

// Constraints:
//     1 <= events.length <= 10^5
//     events[i].length == 2
//     1 <= startDayi <= endDayi <= 10^5

import "fmt"
import "sort"
import "container/heap"

// Priority Queue 
type PriorityQueue []int

func (this PriorityQueue) isEmpty() bool {
    return this.Len() == 0
}

func (this PriorityQueue) Len() int {
    return len(this)
}

// Min-heap
func (this PriorityQueue) Less(i, j int) bool {
    return this[i] < this[j]
}

func (this PriorityQueue) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func (this *PriorityQueue) Push(num interface{}) {
	*this = append(*this, num.(int))
}

func (this *PriorityQueue) Pop() interface{} {
    n := len(*this)
    num := (*this)[n - 1]
    *this = (*this)[:n - 1]
    return num
}

func (this PriorityQueue) Peek() interface{} {
    return this[0]
}

func maxEvents(events [][]int) int {
    sort.SliceStable(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    pq := new(PriorityQueue)
    eventsIndex, res := 0, 0
    for currentDay := 1; currentDay <= 100000; currentDay++ {
        for !pq.isEmpty() && pq.Peek().(int) < currentDay {
            heap.Pop(pq)
        }
        for eventsIndex < len(events) && events[eventsIndex][0] == currentDay {
            heap.Push(pq, events[eventsIndex][1])
            eventsIndex++
        }
        if !pq.isEmpty() {
            heap.Pop(pq)
            res++
        }
    }
    return res
}

func maxEvents1(events [][]int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    n := len(events)
    h := &MinHeap{ items: make([]int, n),  tail: 0, }
    time, idx, res := events[0][0], 0, 0
    for  {
        for idx < n && events[idx][0] <= time { // 在 time 天的时候，把能参加的会议入堆
            heap.Push(h, events[idx][1])
            idx++
        }  
        for h.Len() > 0 && h.top() < time { // 删除过期的会议
            heap.Pop(h)
        }
        if h.Len() > 0 { // 在这些会议中，选择时间最紧迫的（结束时间最早的）
            heap.Pop(h)
            res++
        } else if idx >= n {
            break
        }
        time++
    }
    return res
}

type MinHeap struct {
    items []int
    tail int
}

func (h MinHeap) top() int {
    return h.items[0]
}

func (h MinHeap) Len() int {
    return h.tail
}
func (h MinHeap) Swap(i, j int) {
    h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h MinHeap) Less(i, j int) bool {
    return h.items[i] < h.items[j]
}

func (h *MinHeap) Pop() any {
    h.tail--
    return h.items[h.tail]
}

func (h *MinHeap) Push(x any) {
    h.items[h.tail] = x.(int)
    h.tail++
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/02/05/e1.png" />
    // Input: events = [[1,2],[2,3],[3,4]]
    // Output: 3
    // Explanation: You can attend all the three events.
    // One way to attend them all is as shown.
    // Attend the first event on day 1.
    // Attend the second event on day 2.
    // Attend the third event on day 3.
    fmt.Println(maxEvents([][]int{{1,2},{2,3},{3,4}})) // 3
    // Example 2:
    // Input: events= [[1,2],[2,3],[3,4],[1,2]]
    // Output: 4
    fmt.Println(maxEvents([][]int{{1,2},{2,3},{3,4},{1,2}})) // 4

    fmt.Println(maxEvents1([][]int{{1,2},{2,3},{3,4}})) // 3
    fmt.Println(maxEvents1([][]int{{1,2},{2,3},{3,4},{1,2}})) // 4
}