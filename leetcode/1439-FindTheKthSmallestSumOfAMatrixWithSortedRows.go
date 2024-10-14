package main

// 1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows
// You are given an m x n matrix mat that has its rows sorted in non-decreasing order and an integer k.

// You are allowed to choose exactly one element from each row to form an array.

// Return the kth smallest array sum among all possible arrays.

// Example 1:
// Input: mat = [[1,3,11],[2,4,6]], k = 5
// Output: 7
// Explanation: Choosing one element from each row, the first k smallest sum are:
// [1,2], [1,4], [3,2], [3,4], [1,6]. Where the 5th sum is 7.

// Example 2:
// Input: mat = [[1,3,11],[2,4,6]], k = 9
// Output: 17

// Example 3:
// Input: mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
// Output: 9
// Explanation: Choosing one element from each row, the first k smallest sum are:
// [1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]. Where the 7th sum is 9.  

// Constraints:
//     m == mat.length
//     n == mat.length[i]
//     1 <= m, n <= 40
//     1 <= mat[i][j] <= 5000
//     1 <= k <= min(200, nm)
//     mat[i] is a non-decreasing array.

import "fmt"
import "strconv"
import "strings"
import "container/heap"
import "sort"

type Node struct {
    Sum int
    Indexes []int
}

type MinHeap []Node
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Sum <= h[j].Sum }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Node)) }
func (h *MinHeap) Pop() interface{} {
    res := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return res
}

func kthSmallest(mat [][]int, k int) int {
    getHash := func (arr []int) string {
        b := strings.Builder{}
        b.WriteString(strconv.Itoa(arr[0]))
        for i := 1; i < len(arr); i++ {
            b.WriteByte(':')
            b.WriteString(strconv.Itoa(arr[i]))
        }
        return b.String()
    }
    getSum := func (mat [][]int, indexes []int) int {
        sum := 0
        for i, index := range indexes {
            sum += mat[i][index]
        }
        return sum
    }
    isVisited, indexes := map[string]bool{}, make([]int, len(mat))
    sum := getSum(mat, indexes)
    minHeap := &MinHeap{}
    heap.Push(minHeap,Node{sum, indexes})
    for k > 0 {
        now := heap.Pop(minHeap).(Node)
        hash := getHash(now.Indexes)
        if isVisited[hash] { continue }
        isVisited[hash] = true
        sum = now.Sum
        for i := 0; i < len(mat); i++ {
            if now.Indexes[i]+1 >= len(mat[0]) { continue }
            indexes = append([]int{}, now.Indexes...)
            indexes[i]++
            heap.Push(minHeap, Node{getSum(mat, indexes), indexes})
        }
        k--
    }
    return sum
}

func kthSmallest1(mat [][]int, k int) int {
    sl,sr := 0, 0
    for _, row := range mat {
        sl += row[0] // 累加每行开头
        sr += row[len(row)-1] // 累加每行结尾
    }
    return sl + sort.Search(sr - sl,func(s int) bool {
        leftk := k
        var dfs func(int, int)bool
        dfs = func(i,s int)bool{
            if i < 0 {
                leftk--
                return leftk == 0
            }
            for _, v  := range mat[i] {
                if v - mat[i][0] > s { break }
                if dfs(i-1, s - (v - mat[i][0])) {
                    return true
                }
            }
            return false
        }
        return dfs(len(mat) - 1, s)
    })
}

func main() {
    // Example 1:
    // Input: mat = [[1,3,11],[2,4,6]], k = 5
    // Output: 7
    // Explanation: Choosing one element from each row, the first k smallest sum are:
    // [1,2], [1,4], [3,2], [3,4], [1,6]. Where the 5th sum is 7.
    fmt.Println(kthSmallest([][]int{{1,3,11},{2,4,6}}, 5)) // 7
    // Example 2:
    // Input: mat = [[1,3,11],[2,4,6]], k = 9
    // Output: 17
    fmt.Println(kthSmallest([][]int{{1,3,11},{2,4,6}}, 9)) // 17
    // Example 3:
    // Input: mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
    // Output: 9
    // Explanation: Choosing one element from each row, the first k smallest sum are:
    // [1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]. Where the 7th sum is 9.  
    fmt.Println(kthSmallest([][]int{{1,10,10},{1,4,5},{2,3,6}}, 7)) // 9

    fmt.Println(kthSmallest1([][]int{{1,3,11},{2,4,6}}, 5)) // 7
    fmt.Println(kthSmallest1([][]int{{1,3,11},{2,4,6}}, 9)) // 17
    fmt.Println(kthSmallest1([][]int{{1,10,10},{1,4,5},{2,3,6}}, 7)) // 9
}