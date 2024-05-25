package main

// 1738. Find Kth Largest XOR Coordinate Value
// You are given a 2D matrix of size m x n, consisting of non-negative integers. 
// You are also given an integer k.
// The value of coordinate (a, b) of the matrix is the XOR of all matrix[i][j] where 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).
// Find the kth largest value (1-indexed) of all the coordinates of matrix.

// Example 1:
// Input: matrix = [[5,2],[1,6]], k = 1
// Output: 7
// Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.

// Example 2:
// Input: matrix = [[5,2],[1,6]], k = 2
// Output: 5
// Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.

// Example 3:
// Input: matrix = [[5,2],[1,6]], k = 3
// Output: 4
// Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 1000
//     0 <= matrix[i][j] <= 10^6
//     1 <= k <= m * n

import "fmt"
import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func kthLargestValue(matrix [][]int, k int) int {
    prefixMatrix := make([][]int, len(matrix)+1)
    for i := range prefixMatrix {
        prefixMatrix[i] = make([]int, len(matrix[0])+1)
    }
    h := IntHeap{}
    for i := 1; i < len(prefixMatrix); i++ {
        for j := 1; j < len(prefixMatrix[i]); j++ {
            prefixMatrix[i][j] = prefixMatrix[i-1][j] ^ prefixMatrix[i][j-1] ^ prefixMatrix[i-1][j-1] ^ matrix[i-1][j-1]
            heap.Push(&h, prefixMatrix[i][j])
            for h.Len() > k {
                heap.Pop(&h)
            }
        }
    }
    return heap.Pop(&h).(int)
}

// 二维前缀异或 + 堆排序
func kthLargestValue1(matrix [][]int, k int) int {
    heapLen,  m, n := 0, len(matrix), len(matrix[0])
    heap, pre := make([]int, k), make([][]int, m + 1)
    var down func (nums []int, i int, lens int) 
    down = func (nums []int, i int, lens int) {
        t := i
        if k := 2 * i + 1; k < lens && nums[k] < nums[t] { t = k; }
        if k := 2 * i + 2; k < lens && nums[k] < nums[t] { t = k; }
        if t != i {
            nums[t], nums[i] = nums[i], nums[t]
            down(nums, t, lens)
        }
    }
    for i := range pre {
        pre[i] = make([]int, n + 1)
    }
    for i, row := range matrix {
        for j, num := range row {
            pre[i+1][j+1] = pre[i][j+1] ^ pre[i+1][j] ^ pre[i][j] ^ num
            if heapLen < k {
                heap[heapLen] = pre[i+1][j+1]
                heap[0], heap[heapLen] = heap[heapLen], heap[0]
                heapLen++
                if heapLen == k {
                    for i := heapLen / 2; i >= 0; i-- {
                        down(heap, i, heapLen)
                    }
                }
            } else {
                if pre[i+1][j+1] > heap[0] {
                    heap[0] = pre[i+1][j+1]
                    down(heap, 0, heapLen)
                }
            }
        }
    }
    return heap[0]
}

func main() {
    // Example 1:
    // Input: matrix = [[5,2],[1,6]], k = 1
    // Output: 7
    // Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.
    fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 1)) // 7
    // Example 2:
    // Input: matrix = [[5,2],[1,6]], k = 2
    // Output: 5
    // Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.
    fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 2)) // 5
    // Example 3:
    // Input: matrix = [[5,2],[1,6]], k = 3
    // Output: 4
    // Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.
    fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 3)) // 4

    fmt.Println(kthLargestValue1([][]int{{5,2},{1,6}}, 1)) // 7
    fmt.Println(kthLargestValue1([][]int{{5,2},{1,6}}, 2)) // 5
    fmt.Println(kthLargestValue1([][]int{{5,2},{1,6}}, 3)) // 4
}