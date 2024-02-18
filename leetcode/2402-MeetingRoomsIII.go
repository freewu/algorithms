package main

// 2402. Meeting Rooms III
// You are given an integer n. There are n rooms numbered from 0 to n - 1. 
// You are given a 2D integer array meetings where meetings[i] = [starti, endi] means that a meeting will be held during the half-closed time interval [starti, endi). All the values of starti are unique.

// Meetings are allocated to rooms in the following manner:
//         Each meeting will take place in the unused room with the lowest number.
//         If there are no available rooms, the meeting will be delayed until a room becomes free. The delayed meeting should have the same duration as the original meeting.
//         When a room becomes unused, meetings that have an earlier original start time should be given the room.

// Return the number of the room that held the most meetings. If there are multiple rooms, return the room with the lowest number.
// A half-closed interval [a, b) is the interval between a and b including a and not including b.

// Example 1:
// Input: n = 2, meetings = [[0,10],[1,5],[2,7],[3,4]]
// Output: 0
// Explanation:
// - At time 0, both rooms are not being used. The first meeting starts in room 0.
// - At time 1, only room 1 is not being used. The second meeting starts in room 1.
// - At time 2, both rooms are being used. The third meeting is delayed.
// - At time 3, both rooms are being used. The fourth meeting is delayed.
// - At time 5, the meeting in room 1 finishes. The third meeting starts in room 1 for the time period [5,10).
// - At time 10, the meetings in both rooms finish. The fourth meeting starts in room 0 for the time period [10,11).
// Both rooms 0 and 1 held 2 meetings, so we return 0. 

// Example 2:
// Input: n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]
// Output: 1
// Explanation:
// - At time 1, all three rooms are not being used. The first meeting starts in room 0.
// - At time 2, rooms 1 and 2 are not being used. The second meeting starts in room 1.
// - At time 3, only room 2 is not being used. The third meeting starts in room 2.
// - At time 4, all three rooms are being used. The fourth meeting is delayed.
// - At time 5, the meeting in room 2 finishes. The fourth meeting starts in room 2 for the time period [5,10).
// - At time 6, all three rooms are being used. The fifth meeting is delayed.
// - At time 10, the meetings in rooms 1 and 2 finish. The fifth meeting starts in room 1 for the time period [10,12).
// Room 0 held 1 meeting while rooms 1 and 2 each held 2 meetings, so we return 1. 
 
// Constraints:
//         1 <= n <= 100
//         1 <= meetings.length <= 10^5
//         meetings[i].length == 2
//         0 <= starti < endi <= 5 * 10^5
//         All the values of starti are unique.

import "fmt"
import "sort"
import "container/heap"

func mostBooked(n int, meetings [][]int) (ans int) {
	cnt := make([]int, n)
    // 用两个小顶堆模拟：
    // idle 维护在 start 时刻空闲的会议室的编号；
    // using  维护在 start 时刻使用中的会议室的结束时间和编号
	idle := hp{make([]int, n)}
	for i := 0; i < n; i++ {
		idle.IntSlice[i] = i
	}
	using := hp2{}
    // 对 meetings 按照开始时间排序，然后遍历 meetings
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	for _, m := range meetings {
		start, end := m[0], m[1]
		for len(using) > 0 && using[0].end <= start {
			heap.Push(&idle, heap.Pop(&using).(pair).i) // 维护在 start 时刻空闲的会议室
		}
		var i int
		if idle.Len() == 0 { // 没有可用的会议室
			p := heap.Pop(&using).(pair) // 那么弹出一个最早结束的会议室（若有多个同时结束的，会弹出下标最小的）
			end += p.end - start // 更新当前会议的结束时间 (延后)
			i = p.i
		} else {
			i = heap.Pop(&idle).(int)
		}
		cnt[i]++
		heap.Push(&using, pair{end, i}) // 使用一个会议室
	}
	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type pair struct{ end, i int }
type hp2 []pair
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { a, b := h[i], h[j]; return a.end < b.end || a.end == b.end && a.i < b.i }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
    fmt.Println(mostBooked(2,[][]int{[]int{0,10},[]int{1,5},[]int{2,7},[]int{3,4}})) // 0
    fmt.Println(mostBooked(3,[][]int{[]int{1,20},[]int{2,10},[]int{3,5},[]int{4,9},[]int{6,8}})) // 1
}