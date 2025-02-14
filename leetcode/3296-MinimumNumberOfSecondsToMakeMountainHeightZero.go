package main

// 3296. Minimum Number of Seconds to Make Mountain Height Zero
// You are given an integer mountainHeight denoting the height of a mountain.

// You are also given an integer array workerTimes representing the work time of workers in seconds.

// The workers work simultaneously to reduce the height of the mountain. For worker i:
//     1. To decrease the mountain's height by x, it takes workerTimes[i] + workerTimes[i] * 2 + ... + workerTimes[i] * x seconds. For example:
//     2. To reduce the height of the mountain by 1, it takes workerTimes[i] seconds.
//     3. To reduce the height of the mountain by 2, it takes workerTimes[i] + workerTimes[i] * 2 seconds, and so on.

// Return an integer representing the minimum number of seconds required for the workers to make the height of the mountain 0.

// Example 1:
// Input: mountainHeight = 4, workerTimes = [2,1,1]
// Output: 3
// Explanation:
// One way the height of the mountain can be reduced to 0 is:
// Worker 0 reduces the height by 1, taking workerTimes[0] = 2 seconds.
// Worker 1 reduces the height by 2, taking workerTimes[1] + workerTimes[1] * 2 = 3 seconds.
// Worker 2 reduces the height by 1, taking workerTimes[2] = 1 second.
// Since they work simultaneously, the minimum time needed is max(2, 3, 1) = 3 seconds.

// Example 2:
// Input: mountainHeight = 10, workerTimes = [3,2,2,4]
// Output: 12
// Explanation:
// Worker 0 reduces the height by 2, taking workerTimes[0] + workerTimes[0] * 2 = 9 seconds.
// Worker 1 reduces the height by 3, taking workerTimes[1] + workerTimes[1] * 2 + workerTimes[1] * 3 = 12 seconds.
// Worker 2 reduces the height by 3, taking workerTimes[2] + workerTimes[2] * 2 + workerTimes[2] * 3 = 12 seconds.
// Worker 3 reduces the height by 2, taking workerTimes[3] + workerTimes[3] * 2 = 12 seconds.
// The number of seconds needed is max(9, 12, 12, 12) = 12 seconds.

// Example 3:
// Input: mountainHeight = 5, workerTimes = [1]
// Output: 15
// Explanation:
// There is only one worker in this example, so the answer is workerTimes[0] + workerTimes[0] * 2 + workerTimes[0] * 3 + workerTimes[0] * 4 + workerTimes[0] * 5 = 15.

// Constraints:
//     1 <= mountainHeight <= 10^5
//     1 <= workerTimes.length <= 10^4
//     1 <= workerTimes[i] <= 10^6

import "fmt"
import "container/heap"
import "slices"
import "sort"
import "math"

type Pair struct {
    timeTaken int
    workPerLap int
    i int 
} 
type PriorityQueue []Pair

func (h PriorityQueue) Len() int { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { 
    if h[i].timeTaken == h[j].timeTaken { return h[i].i > h[j].i }
    return h[i].timeTaken < h[j].timeTaken 
}
func (h PriorityQueue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x any) { *h = append(*h, x.(Pair)) }
func (h *PriorityQueue) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
    res := 0
    h := PriorityQueue{} 
    heap.Init(&h)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range workerTimes {
        heap.Push(&h, Pair{ timeTaken: v, workPerLap: v, i: 1 })
    }
    for mountainHeight > 0 {
        p := heap.Pop(&h).(Pair)
        res = max(res, p.timeTaken)
        mountainHeight--
        p.i++
        p.timeTaken += p.i * p.workPerLap
        heap.Push(&h, p)
    }
    return int64(res)
}

func minNumberOfSeconds1(mountainHeight int, workerTimes []int) int64 {
    mx, h := slices.Max(workerTimes), (mountainHeight - 1) / len(workerTimes) + 1
    res := 1 + sort.Search(mx * h * (h + 1) / 2 - 1, func(m int) bool {
        m++
        left := mountainHeight
        for _, t := range workerTimes {
            left -= (int(math.Sqrt(float64(m / t * 8 + 1))) - 1) / 2
            if left <= 0 {
                return true
            }
        }
        return false
    })
    return int64(res)
}

func main() {
    // Example 1:
    // Input: mountainHeight = 4, workerTimes = [2,1,1]
    // Output: 3
    // Explanation:
    // One way the height of the mountain can be reduced to 0 is:
    // Worker 0 reduces the height by 1, taking workerTimes[0] = 2 seconds.
    // Worker 1 reduces the height by 2, taking workerTimes[1] + workerTimes[1] * 2 = 3 seconds.
    // Worker 2 reduces the height by 1, taking workerTimes[2] = 1 second.
    // Since they work simultaneously, the minimum time needed is max(2, 3, 1) = 3 seconds.
    fmt.Println(minNumberOfSeconds(4, []int{2,1,1})) // 3
    // Example 2:
    // Input: mountainHeight = 10, workerTimes = [3,2,2,4]
    // Output: 12
    // Explanation:
    // Worker 0 reduces the height by 2, taking workerTimes[0] + workerTimes[0] * 2 = 9 seconds.
    // Worker 1 reduces the height by 3, taking workerTimes[1] + workerTimes[1] * 2 + workerTimes[1] * 3 = 12 seconds.
    // Worker 2 reduces the height by 3, taking workerTimes[2] + workerTimes[2] * 2 + workerTimes[2] * 3 = 12 seconds.
    // Worker 3 reduces the height by 2, taking workerTimes[3] + workerTimes[3] * 2 = 12 seconds.
    // The number of seconds needed is max(9, 12, 12, 12) = 12 seconds.
    fmt.Println(minNumberOfSeconds(10, []int{3,2,2,4})) // 3
    // Example 3:
    // Input: mountainHeight = 5, workerTimes = [1]
    // Output: 15
    // Explanation:
    // There is only one worker in this example, so the answer is workerTimes[0] + workerTimes[0] * 2 + workerTimes[0] * 3 + workerTimes[0] * 4 + workerTimes[0] * 5 = 15.
    fmt.Println(minNumberOfSeconds(5, []int{1})) // 15

    fmt.Println(minNumberOfSeconds(5, []int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(minNumberOfSeconds(5, []int{9,8,7,6,5,4,3,2,1})) // 4

    fmt.Println(minNumberOfSeconds1(4, []int{2,1,1})) // 3
    fmt.Println(minNumberOfSeconds1(10, []int{3,2,2,4})) // 3
    fmt.Println(minNumberOfSeconds1(5, []int{1})) // 15
    fmt.Println(minNumberOfSeconds1(5, []int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(minNumberOfSeconds1(5, []int{9,8,7,6,5,4,3,2,1})) // 4
}