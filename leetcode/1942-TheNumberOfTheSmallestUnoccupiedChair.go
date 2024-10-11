package main

// 1942. The Number of the Smallest Unoccupied Chair
// There is a party where n friends numbered from 0 to n - 1 are attending. 
// There is an infinite number of chairs in this party that are numbered from 0 to infinity. 
// When a friend arrives at the party, they sit on the unoccupied chair with the smallest number.
//     For example, if chairs 0, 1, and 5 are occupied when a friend comes, they will sit on chair number 2.

// When a friend leaves the party, their chair becomes unoccupied at the moment they leave. 
// If another friend arrives at that same moment, they can sit in that chair.

// You are given a 0-indexed 2D integer array times where times[i] = [arrivali, leavingi], 
// indicating the arrival and leaving times of the ith friend respectively, and an integer targetFriend.
//  All arrival times are distinct.

// Return the chair number that the friend numbered targetFriend will sit on.

// Example 1:
// Input: times = [[1,4],[2,3],[4,6]], targetFriend = 1
// Output: 1
// Explanation: 
// - Friend 0 arrives at time 1 and sits on chair 0.
// - Friend 1 arrives at time 2 and sits on chair 1.
// - Friend 1 leaves at time 3 and chair 1 becomes empty.
// - Friend 0 leaves at time 4 and chair 0 becomes empty.
// - Friend 2 arrives at time 4 and sits on chair 0.
// Since friend 1 sat on chair 1, we return 1.

// Example 2:
// Input: times = [[3,10],[1,5],[2,6]], targetFriend = 0
// Output: 2
// Explanation: 
// - Friend 1 arrives at time 1 and sits on chair 0.
// - Friend 2 arrives at time 2 and sits on chair 1.
// - Friend 0 arrives at time 3 and sits on chair 2.
// - Friend 1 leaves at time 5 and chair 0 becomes empty.
// - Friend 2 leaves at time 6 and chair 1 becomes empty.
// - Friend 0 leaves at time 10 and chair 2 becomes empty.
// Since friend 0 sat on chair 2, we return 2.

// Constraints:
//     n == times.length
//     2 <= n <= 10^4
//     times[i].length == 2
//     1 <= arrivali < leavingi <= 10^5
//     0 <= targetFriend <= n - 1
//     Each arrivali time is distinct.

import "fmt"
import "math/bits"
import "container/heap"
import "sort"

// bitset
// 用 0 表示空椅子，1 表示椅子被占用，这样可以在 bitset 内暴力找第一个空椅子
func smallestChair(times [][]int, targetFriend int) int {
    wLog, wMask := 5 + bits.UintSize >> 6, bits.UintSize - 1
    // 按时间顺序，记录每个到达事件和离开事件相对应的朋友编号
    events := make([][2][]int, 1e5+1)
    for i, t := range times {
        l, r := t[0], t[1]
        events[l][1] = append(events[l][1], i) // 朋友到达
        events[r][0] = append(events[r][0], i) // 朋友离开
    }
    // 初始化未被占据的椅子
    n := len(times)
    bitset := make([]uint, n >> wLog + 1)
    flip := func(p int) { bitset[p>>wLog] ^= 1 << (p & wMask) }

    // 按时间顺序扫描每个事件
    belong := make([]int, n)
    for _, e := range events {
        for _, id := range e[0] { // 朋友离开
            flip(belong[id]) // 返还椅子
        }
        for _, id := range e[1] { // 朋友到达
            for i, mask := range bitset { // 暴力找未被占据的椅子
                if ^mask != 0 {
                    belong[id] = i<<wLog | bits.TrailingZeros(^mask) // 记录占据该椅子的朋友编号
                    break
                }
            }
            if id == targetFriend {
                return belong[id]
            }
            flip(belong[id])
        }
    }
    return 0
}

// heap
func smallestChair1(times [][]int, targetFriend int) int {
    // 按时间顺序，记录每个到达事件和离开事件相对应的朋友编号
    events := make([][2][]int, 1e5+1)
    for i, t := range times {
        l, r := t[0], t[1]
        events[l][1] = append(events[l][1], i) // 朋友到达
        events[r][0] = append(events[r][0], i) // 朋友离开
    }
    // 初始化未被占据的椅子
    n := len(times)
    unoccupied := hp{make([]int, n)}
    for i := range unoccupied.IntSlice {
        unoccupied.IntSlice[i] = i
    }
    // 按时间顺序扫描每个事件
    belong := make([]int, n)
    for _, e := range events {
        for _, id := range e[0] { // 朋友离开
            heap.Push(&unoccupied, belong[id]) // 返还椅子
        }
        for _, id := range e[1] { // 朋友到达
            belong[id] = heap.Pop(&unoccupied).(int) // 记录占据该椅子的朋友编号
            if id == targetFriend {
                return belong[id]
            }
        }
    }
    return 0
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }


func main() {
    // Example 1:
    // Input: times = [[1,4],[2,3],[4,6]], targetFriend = 1
    // Output: 1
    // Explanation: 
    // - Friend 0 arrives at time 1 and sits on chair 0.
    // - Friend 1 arrives at time 2 and sits on chair 1.
    // - Friend 1 leaves at time 3 and chair 1 becomes empty.
    // - Friend 0 leaves at time 4 and chair 0 becomes empty.
    // - Friend 2 arrives at time 4 and sits on chair 0.
    // Since friend 1 sat on chair 1, we return 1.
    fmt.Println(smallestChair([][]int{{1,4},{2,3},{4,6}}, 1)) // 1
    // Example 2:
    // Input: times = [[3,10],[1,5],[2,6]], targetFriend = 0
    // Output: 2
    // Explanation: 
    // - Friend 1 arrives at time 1 and sits on chair 0.
    // - Friend 2 arrives at time 2 and sits on chair 1.
    // - Friend 0 arrives at time 3 and sits on chair 2.
    // - Friend 1 leaves at time 5 and chair 0 becomes empty.
    // - Friend 2 leaves at time 6 and chair 1 becomes empty.
    // - Friend 0 leaves at time 10 and chair 2 becomes empty.
    // Since friend 0 sat on chair 2, we return 2.
    fmt.Println(smallestChair([][]int{{3,10},{1,5},{2,6}}, 0)) // 2

    fmt.Println(smallestChair1([][]int{{1,4},{2,3},{4,6}}, 1)) // 1
    fmt.Println(smallestChair1([][]int{{3,10},{1,5},{2,6}}, 0)) // 2
}