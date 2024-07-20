package main

// 1229. Meeting Scheduler
// Given the availability time slots arrays slots1 and slots2 of two people and a meeting duration duration, 
// return the earliest time slot that works for both of them and is of duration duration.

// If there is no common time slot that satisfies the requirements, return an empty array.
// The format of a time slot is an array of two elements [start, end] representing an inclusive time range from start to end.

// It is guaranteed that no two availability slots of the same person intersect with each other. 
// That is, for any two time slots [start1, end1] and [start2, end2] of the same person, either start1 > end2 or start2 > end1.

// Example 1:
// Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]], duration = 8
// Output: [60,68]

// Example 2:
// Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]], duration = 12
// Output: []

// Constraints:
//     1 <= slots1.length, slots2.length <= 10^4
//     slots1[i].length, slots2[i].length == 2
//     slots1[i][0] < slots1[i][1]
//     slots2[i][0] < slots2[i][1]
//     0 <= slots1[i][j], slots2[i][j] <= 10^9
//     1 <= duration <= 10^6

import "fmt"
import "sort"
import "container/heap"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
    sort.Slice(slots1, func(i, j int) bool { return slots1[i][0] < slots1[j][0] })
    sort.Slice(slots2, func(i, j int) bool { return slots2[i][0] < slots2[j][0] })
    pointer1, pointer2 := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for pointer1 < len(slots1) && pointer2 < len(slots2) {
        // 找出交集的边界，或者通用的时间段。
        intersectLeft := max(slots1[pointer1][0], slots2[pointer2][0])
        intersectRight := min(slots1[pointer1][1], slots2[pointer2][1])
        if intersectRight - intersectLeft >= duration {
            return []int{intersectLeft, intersectLeft + duration}
        }
        // 始终移动那个结束时间较早的时间段
        if slots1[pointer1][1] < slots2[pointer2][1] {
            pointer1++
        } else {
            pointer2++
        }
    }
    return []int{}
}

func minAvailableDuration1(slots1 [][]int, slots2 [][]int, duration int) []int {
    timeslots := &SlotHeap{}
    heap.Init(timeslots)
    for _, slot := range slots1 {
        if slot[1]-slot[0] >= duration {
            heap.Push(timeslots, slot)
        }
    }
    for _, slot := range slots2 {
        if slot[1]-slot[0] >= duration {
            heap.Push(timeslots, slot)
        }
    }
    for timeslots.Len() > 1 {
        slot1 := heap.Pop(timeslots).([]int)
        slot2 := (*timeslots)[0]
        if slot1[1] >= slot2[0]+duration {
            return []int{slot2[0], slot2[0] + duration}
        }
    }
    return []int{}
}

type SlotHeap [][]int
func (h SlotHeap) Len() int            { return len(h) }
func (h SlotHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h SlotHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *SlotHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *SlotHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    // Example 1:
    // Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]], duration = 8
    // Output: [60,68]
    fmt.Println(minAvailableDuration([][]int{{10,50},{60,120},{140,210}},[][]int{{0,15},{60,70}}, 8)) // [60,68]
    // Example 2:
    // Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]], duration = 12
    // Output: []
    fmt.Println(minAvailableDuration([][]int{{10,50},{60,120},{140,210}},[][]int{{0,15},{60,70}}, 12)) // []

    fmt.Println(minAvailableDuration1([][]int{{10,50},{60,120},{140,210}},[][]int{{0,15},{60,70}}, 8)) // [60,68]
    fmt.Println(minAvailableDuration1([][]int{{10,50},{60,120},{140,210}},[][]int{{0,15},{60,70}}, 12)) // []
}