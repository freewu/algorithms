package main

// 2532. Time to Cross a Bridge
// There are k workers who want to move n boxes from the right (old) warehouse to the left (new) warehouse. 
// You are given the two integers n and k, and a 2D integer array time of size k x 4 where time[i] = [righti, picki, lefti, puti].

// The warehouses are separated by a river and connected by a bridge. 
// Initially, all k workers are waiting on the left side of the bridge. 
// To move the boxes, the ith worker can do the following:
//     1. Cross the bridge to the right side in righti minutes.
//     2. Pick a box from the right warehouse in picki minutes.
//     3. Cross the bridge to the left side in lefti minutes.
//     4. Put the box into the left warehouse in puti minutes.

// The ith worker is less efficient than the jth worker if either condition is met:
//     1. lefti + righti > leftj + rightj
//     2. lefti + righti == leftj + rightj and i > j

// The following rules regulate the movement of the workers through the bridge:
//     1. Only one worker can use the bridge at a time.
//     2. When the bridge is unused prioritize the least efficient worker (who have picked up the box) on the right side to cross. 
//        If not, prioritize the least efficient worker on the left side to cross.
//     3. If enough workers have already been dispatched from the left side to pick up all the remaining boxes, 
//        no more workers will be sent from the left side.

// Return the elapsed minutes at which the last box reaches the left side of the bridge.

// Example 1:
// Input: n = 1, k = 3, time = [[1,1,2,1],[1,1,3,1],[1,1,4,1]]
// Output: 6
// Explanation:
// From 0 to 1 minutes: worker 2 crosses the bridge to the right.
// From 1 to 2 minutes: worker 2 picks up a box from the right warehouse.
// From 2 to 6 minutes: worker 2 crosses the bridge to the left.
// From 6 to 7 minutes: worker 2 puts a box at the left warehouse.
// The whole process ends after 7 minutes. We return 6 because the problem asks for the instance of time at which the last worker reaches the left side of the bridge.

// Example 2:
// Input: n = 3, k = 2, time = [[1,5,1,8],[10,10,10,10]]
// Output: 37
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/21/378539249-c6ce3c73-40e7-4670-a8b5-7ddb9abede11.png" />
// The last box reaches the left side at 37 seconds. Notice, how we do not put the last boxes down, as that would take more time, and they are already on the left with the workers.

// Constraints:
//     1 <= n, k <= 10^4
//     time.length == k
//     time[i].length == 4
//     1 <= lefti, picki, righti, puti <= 1000

import "fmt"
import "sort"
import "container/heap"

type MinHeap struct{ sort.IntSlice }
func (h MinHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *MinHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MinHeap) Pop() any {
    a := h.IntSlice
    v := a[len(a)-1]
    h.IntSlice = a[:len(a)-1]
    return v
}

type Pair struct{ t, i int }
type MaxHeap []Pair
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MaxHeap) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func findCrossingTime(n int, k int, time [][]int) int {
    sort.SliceStable(time, func(i, j int) bool { 
        return time[i][0]+time[i][2] < time[j][0]+time[j][2] 
    })
    waitInLeft, waitInRight := MinHeap{}, MinHeap{}
    workInLeft, workInRight := MaxHeap{}, MaxHeap{}
    for i := range time {
        heap.Push(&waitInLeft, i)
    }
    res := 0
    for {
        for len(workInLeft) > 0 {
            if workInLeft[0].t > res { break }
            heap.Push(&waitInLeft, heap.Pop(&workInLeft).(Pair).i)
        }
        for len(workInRight) > 0 {
            if workInRight[0].t > res { break }
            heap.Push(&waitInRight, heap.Pop(&workInRight).(Pair).i)
        }
        leftToGo, rightToGo := n > 0 && waitInLeft.Len() > 0,  waitInRight.Len() > 0
        if !leftToGo && !rightToGo {
            next := 1 << 30
            if len(workInLeft) > 0 {
                next = min(next, workInLeft[0].t)
            }
            if len(workInRight) > 0 {
                next = min(next, workInRight[0].t)
            }
            res = next
            continue
        }
        if rightToGo {
            i := heap.Pop(&waitInRight).(int)
            res += time[i][2]
            if n == 0 && waitInRight.Len() == 0 && len(workInRight) == 0 {
                return res
            }
            heap.Push(&workInLeft, Pair{res + time[i][3], i})
        } else {
            i := heap.Pop(&waitInLeft).(int)
            res += time[i][0]
            n--
            heap.Push(&workInRight, Pair{res + time[i][1], i})
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 1, k = 3, time = [[1,1,2,1],[1,1,3,1],[1,1,4,1]]
    // Output: 6
    // Explanation:
    // From 0 to 1 minutes: worker 2 crosses the bridge to the right.
    // From 1 to 2 minutes: worker 2 picks up a box from the right warehouse.
    // From 2 to 6 minutes: worker 2 crosses the bridge to the left.
    // From 6 to 7 minutes: worker 2 puts a box at the left warehouse.
    // The whole process ends after 7 minutes. We return 6 because the problem asks for the instance of time at which the last worker reaches the left side of the bridge.
    fmt.Println(findCrossingTime(1, 3, [][]int{{1,1,2,1},{1,1,3,1},{1,1,4,1}})) // 6
    // Example 2:
    // Input: n = 3, k = 2, time = [[1,5,1,8],[10,10,10,10]]
    // Output: 37
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/21/378539249-c6ce3c73-40e7-4670-a8b5-7ddb9abede11.png" />
    // The last box reaches the left side at 37 seconds. Notice, how we do not put the last boxes down, as that would take more time, and they are already on the left with the workers.
    fmt.Println(findCrossingTime(3, 2, [][]int{{1,5,1,8},{10,10,10,10}})) // 37
}