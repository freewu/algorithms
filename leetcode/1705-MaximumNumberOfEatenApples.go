package main

// 1705. Maximum Number of Eaten Apples
// There is a special kind of apple tree that grows apples every day for n days. 
// On the ith day, the tree grows apples[i] apples that will rot after days[i] days, 
// that is on day i + days[i] the apples will be rotten and cannot be eaten. 
// On some days, the apple tree does not grow any apples, which are denoted by apples[i] == 0 and days[i] == 0.

// You decided to eat at most one apple a day (to keep the doctors away). 
// Note that you can keep eating after the first n days.

// Given two integer arrays days and apples of length n, return the maximum number of apples you can eat.

// Example 1:
// Input: apples = [1,2,3,5,2], days = [3,2,1,4,2]
// Output: 7
// Explanation: You can eat 7 apples:
// - On the first day, you eat an apple that grew on the first day.
// - On the second day, you eat an apple that grew on the second day.
// - On the third day, you eat an apple that grew on the second day. After this day, the apples that grew on the third day rot.
// - On the fourth to the seventh days, you eat apples that grew on the fourth day.

// Example 2:
// Input: apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
// Output: 5
// Explanation: You can eat 5 apples:
// - On the first to the third day you eat apples that grew on the first day.
// - Do nothing on the fouth and fifth days.
// - On the sixth and seventh days you eat apples that grew on the sixth day.

// Constraints:
//     n == apples.length == days.length
//     1 <= n <= 2 * 10^4
//     0 <= apples[i], days[i] <= 2 * 10^4
//     days[i] = 0 if and only if apples[i] = 0.

import "fmt"
import "container/heap"

type MinHeap [][2]int
func (h MinHeap) Len() int              { return len(h) }
func (h MinHeap) Less(i, j int) bool    { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{})   { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() (v interface{}) { v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]; return }

func eatenApples(apples []int, days []int) int {
    res, i, pq := 0, 0, MinHeap{}
    for i < len(apples) || len(pq) > 0 {
        if i < len(apples) && days[i] > 0 && apples[i] > 0 {
            heap.Push(&pq, [2]int{i + days[i], apples[i]})
        }
        if len(pq) > 0 {
            val := heap.Pop(&pq).([2]int)
            expiry, count := val[0], val[1]
            for expiry < i + 1 && len(pq) > 0 {
                val := heap.Pop(&pq).([2]int)
                expiry, count = val[0], val[1]
            }
            if expiry >= i + 1 {
                count-- // Eat one apple.
                res++
                if count > 0 {
                    heap.Push(&pq, [2]int{ expiry, count })
                }
            }
        }
        i += 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: apples = [1,2,3,5,2], days = [3,2,1,4,2]
    // Output: 7
    // Explanation: You can eat 7 apples:
    // - On the first day, you eat an apple that grew on the first day.
    // - On the second day, you eat an apple that grew on the second day.
    // - On the third day, you eat an apple that grew on the second day. After this day, the apples that grew on the third day rot.
    // - On the fourth to the seventh days, you eat apples that grew on the fourth day.
    fmt.Println(eatenApples([]int{1,2,3,5,2}, []int{3,2,1,4,2})) // 7
    // Example 2:
    // Input: apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
    // Output: 5
    // Explanation: You can eat 5 apples:
    // - On the first to the third day you eat apples that grew on the first day.
    // - Do nothing on the fouth and fifth days.
    // - On the sixth and seventh days you eat apples that grew on the sixth day.
    fmt.Println(eatenApples([]int{3,0,0,0,0,2}, []int{3,0,0,0,0,2})) // 5
}