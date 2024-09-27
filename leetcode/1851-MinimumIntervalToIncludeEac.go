package main

// 1851. Minimum Interval to Include Eac
// You are given a 2D integer array intervals, 
// where intervals[i] = [lefti, righti] describes the ith interval starting at lefti and ending at righti (inclusive). 
// The size of an interval is defined as the number of integers it contains, or more formally righti - lefti + 1.

// You are also given an integer array queries. 
// The answer to the jth query is the size of the smallest interval i such that lefti <= queries[j] <= righti. 
// If no such interval exists, the answer is -1.

// Return an array containing the answers to the queries.

// Example 1:
// Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
// Output: [3,3,1,4]
// Explanation: The queries are processed as follows:
// - Query = 2: The interval [2,4] is the smallest interval containing 2. The answer is 4 - 2 + 1 = 3.
// - Query = 3: The interval [2,4] is the smallest interval containing 3. The answer is 4 - 2 + 1 = 3.
// - Query = 4: The interval [4,4] is the smallest interval containing 4. The answer is 4 - 4 + 1 = 1.
// - Query = 5: The interval [3,6] is the smallest interval containing 5. The answer is 6 - 3 + 1 = 4.

// Example 2:
// Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
// Output: [2,-1,4,6]
// Explanation: The queries are processed as follows:
// - Query = 2: The interval [2,3] is the smallest interval containing 2. The answer is 3 - 2 + 1 = 2.
// - Query = 19: None of the intervals contain 19. The answer is -1.
// - Query = 5: The interval [2,5] is the smallest interval containing 5. The answer is 5 - 2 + 1 = 4.
// - Query = 22: The interval [20,25] is the smallest interval containing 22. The answer is 25 - 20 + 1 = 6.

// Constraints:
//     1 <= intervals.length <= 10^5
//     1 <= queries.length <= 10^5
//     intervals[i].length == 2
//     1 <= lefti <= righti <= 10^7
//     1 <= queries[j] <= 10^7

import "fmt"
import "sort"
import "container/heap"
import "slices"

type IntHeap [][2]int

func (h IntHeap)  Len() int           { return len(h) }
func (h IntHeap)  Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {  *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func minInterval(intervals [][]int, queries []int) []int {
    sort.Slice(intervals, func(i, j int) bool { // sorted array by start
        return intervals[i][0] < intervals[j][0]
    })
    sorted_queries := make([]int, len(queries)) // store unsorted queries to return results
    copy(sorted_queries, queries)
    sort.Ints(sorted_queries)
    mapQueries := map[int]int{} // mapping result to unsorted_queries
    h := IntHeap{}
    heap.Init(&h)
    i := 0
    for j := 0; j < len(sorted_queries); j++ {
        for j + 1 < len(sorted_queries) && sorted_queries[j] == sorted_queries[j+1] { j++ }
        qj := sorted_queries[j]
        for i < len(intervals) && (intervals[i][0] <= qj) {
            end := intervals[i][1]
            size := end - intervals[i][0] + 1
            heap.Push(&h, [2]int{size, end}) // O(logn)
            i++
        }
        for h.Len() > 0 && qj > h[0][1] { // pop intervals is oubounded
            heap.Pop(&h)
        }
        if len(h) > 0 {
            mapQueries[qj] = h[0][0]
        } else {
            mapQueries[qj] = -1
        }
    }
    res := []int{}
    for _, q := range queries {
        res = append(res, mapQueries[q])
    }
    return res
}

func minInterval1(intervals [][]int, queries []int) []int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return (a[1] - a[0]) - (b[1] - b[0])
    })
    n := len(queries)
    res := make([]int, n)
    for i := range res {
        res[i] = -1
    }
    fa := make([]int, n + 1)
    for i := range fa {
        fa[i] = i
    }
    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }
    type pair struct{ x, i int }
    pairs := make([]pair, n)
    for i, x := range queries {
        pairs[i] = pair{x, i}
    }
    slices.SortFunc(pairs, func(a, b pair) int { return a.x - b.x })
    for _, interval := range intervals {
        l, r := interval[0], interval[1]
        sz := r - l + 1
        i := sort.Search(n, func(i int) bool { return pairs[i].x >= l })
        for i = find(i); i < n && pairs[i].x <= r; i = find(i + 1) {
            res[pairs[i].i] = sz
            fa[i] = i + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
    // Output: [3,3,1,4]
    // Explanation: The queries are processed as follows:
    // - Query = 2: The interval [2,4] is the smallest interval containing 2. The answer is 4 - 2 + 1 = 3.
    // - Query = 3: The interval [2,4] is the smallest interval containing 3. The answer is 4 - 2 + 1 = 3.
    // - Query = 4: The interval [4,4] is the smallest interval containing 4. The answer is 4 - 4 + 1 = 1.
    // - Query = 5: The interval [3,6] is the smallest interval containing 5. The answer is 6 - 3 + 1 = 4.
    fmt.Println(minInterval([][]int{{1,4},{2,4},{3,6},{4,4}}, []int{2,3,4,5})) // [3,3,1,4]
    // Example 2:
    // Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
    // Output: [2,-1,4,6]
    // Explanation: The queries are processed as follows:
    // - Query = 2: The interval [2,3] is the smallest interval containing 2. The answer is 3 - 2 + 1 = 2.
    // - Query = 19: None of the intervals contain 19. The answer is -1.
    // - Query = 5: The interval [2,5] is the smallest interval containing 5. The answer is 5 - 2 + 1 = 4.
    // - Query = 22: The interval [20,25] is the smallest interval containing 22. The answer is 25 - 20 + 1 = 6.
    fmt.Println(minInterval([][]int{{2,3},{2,5},{1,8},{20,25}}, []int{2,19,5,22})) // [2,-1,4,6]

    fmt.Println(minInterval1([][]int{{1,4},{2,4},{3,6},{4,4}}, []int{2,3,4,5})) // [3,3,1,4]
    fmt.Println(minInterval1([][]int{{2,3},{2,5},{1,8},{20,25}}, []int{2,19,5,22})) // [2,-1,4,6]
}