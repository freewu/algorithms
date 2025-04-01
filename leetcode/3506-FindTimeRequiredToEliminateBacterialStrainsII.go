package main

// 3506. Find Time Required to Eliminate Bacterial Strains II
// You are given an integer array timeReq and an integer splitTime.

// In the microscopic world of the human body, the immune system faces an extraordinary challenge: 
// combatting a rapidly multiplying bacterial colony that threatens the body's survival.

// Initially, only one white blood cell (WBC) is deployed to eliminate the bacteria. 
// However, the lone WBC quickly realizes it cannot keep up with the bacterial growth rate.

// The WBC devises a clever strategy to fight the bacteria:
//     1. The ith bacterial strain takes timeReq[i] units of time to be eliminated.
//     2. A single WBC can eliminate only one bacterial strain. 
//        Afterwards, the WBC is exhausted and cannot perform any other tasks.
//     3. A WBC can split itself into two WBCs, but this requires splitTime units of time. 
//        Once split, the two WBCs can work in parallel on eliminating the bacteria.
//     4. Only one WBC can work on a single bacterial strain. Multiple WBCs cannot attack one strain in parallel.

// You must determine the minimum time required to eliminate all the bacterial strains.

// Note that the bacterial strains can be eliminated in any order.

// Example 1:
// Input: timeReq = [10,4,5], splitTime = 2
// Output: 12
// Explanation:
// The elimination process goes as follows:
// Initially, there is a single WBC. The WBC splits into 2 WBCs after 2 units of time.
// One of the WBCs eliminates strain 0 at a time t = 2 + 10 = 12. The other WBC splits again, using 2 units of time.
// The 2 new WBCs eliminate the bacteria at times t = 2 + 2 + 4 and t = 2 + 2 + 5.

// Example 2:
// Input: timeReq = [10,4], splitTime = 5
// Output:15
// Explanation:
// The elimination process goes as follows:
// Initially, there is a single WBC. The WBC splits into 2 WBCs after 5 units of time.
// The 2 new WBCs eliminate the bacteria at times t = 5 + 10 and t = 5 + 4.

// Constraints:
//     2 <= timeReq.length <= 10^5
//     1 <= timeReq[i] <= 10^9
//     1 <= splitTime <= 10^9

import "fmt"
import "container/heap"
import "sort"

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minEliminationTime(timeReq []int, splitTime int) int64 {
    h := &MinHeap{}
    heap.Init(h)
    for _, t := range timeReq {
        heap.Push(h, t)
    }
    for h.Len() > 1 {
        _ = heap.Pop(h).(int)
        v := heap.Pop(h).(int)
        newTime := v + splitTime
        heap.Push(h, newTime)
    }
    return int64(heap.Pop(h).(int))
}

func minEliminationTime1(timeReq []int, splitTime int) int64 {
    sort.Slice(timeReq, func(i, j int) bool {
        return timeReq[i] > timeReq[j]
    })
    helper := func(arr []int, k, target int64) bool {
        index, count, timeVal, n := 0, int64(1), int64(0), len(arr)
        for timeVal <= target && index < n && count > 0 {
            if count >= int64(n - index) {
                timeVal += int64(arr[index])
                index = n
                break
            }
            for index < n && count > 0 && (target - int64(arr[index]) < timeVal + k) {
                count--
                index++
            }
            count *= 2
            timeVal += k
        }
        return timeVal <= target && index == n
    }
    l, r := int64(0), int64(1 << 61)
    res := r
    for l <= r {
        mid := l + (r - l) / 2
        if helper(timeReq, int64(splitTime), mid) {
            res, r = mid, mid - 1
        } else {
            l = mid + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: timeReq = [10,4,5], splitTime = 2
    // Output: 12
    // Explanation:
    // The elimination process goes as follows:
    // Initially, there is a single WBC. The WBC splits into 2 WBCs after 2 units of time.
    // One of the WBCs eliminates strain 0 at a time t = 2 + 10 = 12. The other WBC splits again, using 2 units of time.
    // The 2 new WBCs eliminate the bacteria at times t = 2 + 2 + 4 and t = 2 + 2 + 5.
    fmt.Println(minEliminationTime([]int{10,4,5}, 2)) // 12
    // Example 2:
    // Input: timeReq = [10,4], splitTime = 5
    // Output: 15
    // Explanation:
    // The elimination process goes as follows:
    // Initially, there is a single WBC. The WBC splits into 2 WBCs after 5 units of time.
    // The 2 new WBCs eliminate the bacteria at times t = 5 + 10 and t = 5 + 4.
    fmt.Println(minEliminationTime([]int{10,4}, 5)) // 15

    fmt.Println(minEliminationTime([]int{1,2,3,4,5,6,7,8,9}, 2)) // 13
    fmt.Println(minEliminationTime([]int{9,8,7,6,5,4,3,2,1}, 2)) // 13

    fmt.Println(minEliminationTime1([]int{10,4,5}, 2)) // 12
    fmt.Println(minEliminationTime1([]int{10,4}, 5)) // 15
    fmt.Println(minEliminationTime1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 13
    fmt.Println(minEliminationTime1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 13
}

