package main

// 253. Meeting Rooms II
// Given an array of meeting time intervals intervals where intervals[i] = [starti, endi], 
// return the minimum number of conference rooms required.

// Example 1:
// Input: intervals = [[0,30],[5,10],[15,20]]
// Output: 2

// Example 2:
// Input: intervals = [[7,10],[2,4]]
// Output: 1
 
// Constraints:
//     1 <= intervals.length <= 10^4
//     0 <= starti < endi <= 10^6

import "fmt"
import "container/heap"
import "sort"

// 小根堆
type minHeap []int
func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
    v := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return v
}
func minMeetingRooms(intervals [][]int) int {
    // 检查基本情况。如果没有间隔，返回 0
    if len(intervals) == 0 {
        return 0
    }
    allocator := &minHeap{}
    heap.Init(allocator)
    // 根据开始时间排序会议
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0]})
    heap.Push(allocator, intervals[0][1])
    // 遍历剩余会议
    for i := 1; i < len(intervals); i++ {
        // 如果最早应该腾出的房间是空闲的，则将该房间分配给本次会议。
        if intervals[i][0] >= (*allocator)[0] {
            heap.Pop(allocator)
        }
        // 如果要分配一个新房间，那么我们也要添加到堆中，
        // 如果分配了一个旧房间，那么我们还必须添加到具有更新的结束时间的堆中。
        heap.Push(allocator, intervals[i][1])
    }
    // 堆的大小告诉我们所有会议所
    return len(*allocator)
}

func minMeetingRooms1(intervals [][]int) int {
    // 检查边界条件。如果没有间隔，返回 0
    if len(intervals) == 0 {
        return 0
    }
    usedRooms := 0
    var start, end []int // 存储会议开始和结束时间的数组
    for _, interval := range intervals {
        start = append(start, interval[0])
        end = append(end, interval[1])
    }
    sort.Ints(start) // 按照开始时间排序
    sort.Ints(end) // 按照结束时间排序
    endPointer := 0
    startPointer := 0
    for startPointer < len(intervals) {
        if start[startPointer] >= end[endPointer] {
            usedRooms-- // 释放一个会议室
            endPointer++
        }
        usedRooms++ // 占用一个会议室
        startPointer++
    }
    return usedRooms
}

func main() {
    fmt.Println(minMeetingRooms([][]int{{0,30}, {5,10}, {15,20}})) // 2
    fmt.Println(minMeetingRooms([][]int{{7,10}, {2,4}})) // 1

    fmt.Println(minMeetingRooms1([][]int{{0,30}, {5,10}, {15,20}})) // 2
    fmt.Println(minMeetingRooms1([][]int{{7,10}, {2,4}})) // 1
}