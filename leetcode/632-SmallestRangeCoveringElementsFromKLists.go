package main

// 632. Smallest Range Covering Elements from K Lists
// You have k lists of sorted integers in non-decreasing order. 
// Find the smallest range that includes at least one number from each of the k lists.

// We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.

// Example 1:
// Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
// Output: [20,24]
// Explanation: 
// List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
// List 2: [0, 9, 12, 20], 20 is in range [20,24].
// List 3: [5, 18, 22, 30], 22 is in range [20,24].

// Example 2:
// Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
// Output: [1,1]
 
// Constraints:
//     nums.length == k
//     1 <= k <= 3500
//     1 <= nums[i].length <= 50
//     -10^5 <= nums[i][j] <= 10^5
//     nums[i] is sorted in non-decreasing order.

import "fmt"
import "math"
import "container/heap"

// heap item: the num, the array index, the index of num in the array
func smallestRange(nums [][]int) []int {
    var heap [][3]int
    minVal, maxVal := 1<<63-1, -1<<63
    for i := 0; i < len(nums); i++ {
        if nums[i][0] > maxVal {
            maxVal = nums[i][0]
        }
        if nums[i][0] < minVal {
            minVal = nums[i][0]
        }
        heap = append(heap, [3]int{nums[i][0], i, 0})
        heapUp(heap, len(heap)-1)
    }
    start, end := minVal, maxVal
    for len(heap) > 0 {
        if maxVal - heap[0][0] < end - start {
            end = maxVal
            start = heap[0][0]
        }
        arrIndex := heap[0][1]
        currentIndexInTheSubArr := heap[0][2]
        if currentIndexInTheSubArr+1 < len(nums[arrIndex]) {
            heap[0] = [3]int{nums[arrIndex][currentIndexInTheSubArr+1], arrIndex, currentIndexInTheSubArr+1}
            if nums[arrIndex][currentIndexInTheSubArr+1] > maxVal {
                maxVal = nums[arrIndex][currentIndexInTheSubArr+1]
            }  
            heapDown(heap, 0, len(heap)-1)
        } else {
            break
        }
    }
    return []int{start, end}
}

func heapDown(heap [][3]int, p, limit int) {
    l, r := 2*p+1, 2*p+2
    smaller := p
    if l <= limit && heap[l][0] < heap[smaller][0] {
        smaller = l
    }
    if r <= limit && heap[r][0] < heap[smaller][0] {
        smaller = r
    }
    if smaller != p {
        heap[smaller], heap[p] = heap[p], heap[smaller]
        heapDown(heap, smaller, limit)
    }
}

func heapUp(heap [][3]int, p int) {
    parent := (p-1) / 2
    if parent >= 0 && heap[p][0] < heap[parent][0] {
        heap[parent], heap[p] = heap[p], heap[parent]
        heapUp(heap, parent)
    }
}


func smallestRange1(nums [][]int) []int {
    k, h := len(nums), new(MyHeap)
    heap.Init(h)
    left, right, distance := 0, 0, math.MaxInt32
    maxVal := math.MinInt32
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < k; i++ {
        heap.Push(h, Info{nums[i][0], i, 0})
        maxVal = max(maxVal, nums[i][0])
    }
    for h.Len() == k {
        curInfo := heap.Pop(h).(Info)
        val := curInfo.Val
        from := curInfo.From
        index := curInfo.Index
        if maxVal-val < distance {
            distance = maxVal - val
            left = val
            right = maxVal
        }
        if index+1 < len(nums[from]) {
            heap.Push(h, Info{nums[from][index+1], from, index+1})
            maxVal = max(maxVal, nums[from][index+1])
        }
    }

    return []int{left, right}
}

type Info struct {
    Val int
    From int
    Index int
}

type MyHeap []Info
func (h MyHeap) Len() int {
    return len(h)
}
func (h MyHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}
func (h MyHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MyHeap) Push(x any) {
    *h = append(*h, x.(Info))
}
func (h *MyHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func main() {
    // List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
    // List 2: [0, 9, 12, 20], 20 is in range [20,24].
    // List 3: [5, 18, 22, 30], 22 is in range [20,24].
    fmt.Println(smallestRange([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    fmt.Println(smallestRange([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]

    fmt.Println(smallestRange1([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    fmt.Println(smallestRange1([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]
}