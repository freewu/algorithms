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
import "slices"

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

type Tuple struct{ x, i, j int }
type MinHeap []Tuple
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].x < h[j].x }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Tuple)) }
func (MinHeap) Pop() (_ any)         { return }

func smallestRange2(nums [][]int) []int {
    mx, mhp := -1 << 31, MinHeap{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, arr := range nums {
        heap.Push(&mhp, Tuple{arr[0], i, 0}) // 把每个列表的第一个元素入堆
        mx = max(mx, arr[0])
    }
    heap.Init(&mhp)
    left, right := mhp[0].x, mx
    for mhp[0].j + 1 < len(nums[mhp[0].i]) { // 堆顶列表有下一个元素
        x := nums[mhp[0].i][mhp[0].j + 1] // 堆顶列表的下一个元素
        mx = max(mx, x)
        mhp[0].x = x // 替换堆顶
        mhp[0].j++
        heap.Fix(&mhp, 0)
        if mx - mhp[0].x < right - left { // mhp[0].x 是当前合法区间的左端点
            left, right = mhp[0].x, mx
        }
    }
    return []int{ left, right }
}

func smallestRange3(nums [][]int) []int {
    type Pair struct{ x, i int }
    pairs := []Pair{}
    for i, arr := range nums {
        for _, x := range arr {
            pairs = append(pairs, Pair{x, i})
        }
    }
    slices.SortFunc(pairs, func(a, b Pair) int { 
        return a.x - b.x 
    })
    n, left := len(nums), 0
    res, count := [2]int{pairs[0].x, pairs[len(pairs)-1].x}, make([]int, n)
    for _, p := range pairs {
        r, i := p.x, p.i
        if count[i] == 0 { // 包含 nums[i] 的数字
            n--
        }
        count[i]++
        for n == 0 { // 每个列表都至少包含一个数
            l, i := pairs[left].x, pairs[left].i
            if r - l < res[1] - res[0] {
                res[0], res[1] = l, r
            }
            count[i]--
            if count[i] == 0 { // 不包含 nums[i] 的数字
                n++
            }
            left++
        }
    }
    return []int{ res[0], res[1] }
}

// func smallestRange4(nums [][]int) []int {
//     if len(nums) == 1 { return []int{nums[0][0], nums[0][0]} }
//     n, inf := len(nums), 1 << 31
//     ranges := [][2]int // intermediate ranges
//     minRange := [2]int{ -inf, inf } // resulting range
//     minRangeLen := inf // resulting range length
//     diff := func(a, b int) int {
//         if a > b { return a - b }
//         return b - a
//     }
//     for _, v := range nums[0] { // for each number of first input slice
//         ranges = ranges[:0]
//         il, ir := v, v // intermediate range borders
//         // find minimum range
//         for i := 1; i < len(nums); i++ { // for each next slice
//             // find numbers closest to num
//             l, r := math.MinInt, math.MaxInt
//             for _, num_ := range nums[i] {
//                 if num_ <= num {
//                     l = num_
//                 }
//                 if num_ >= num {
//                     r = num_
//                     break
//                 }
//             }
//             // try to grow borders
//             if l != r && l != math.MinInt && r != math.MaxInt {
//                 if (l < il || l > ir) && (r < il || r > ir) {
//                     // it's unclear yet which number to use
//                     ranges = append(ranges, [2]int{l, r})
//                 }
//             } else {
//                 // have only one number
//                 // grow range
//                 if l == math.MinInt {
//                     l = r
//                 }
//                 if l < il {
//                     il = l
//                 } else if l > ir {
//                     ir = l
//                 }
//             }
//         }
//         // merge remaining numbers
//         for idx := 0; idx < len(ranges); idx++ {
//             l, r := ranges[idx][0], ranges[idx][1]
//             if (l < il || l > ir) && (r < il || r > ir) {
//                 diffL := min(diff(l, il), diff(l, ir))
//                 diffR := min(diff(r, il), diff(r, ir))
//                 if diffL < diffR {
//                     // l is closer to range
//                     if il == ir {
//                         il = l
//                     } else if diff(l, il) < diff(l, ir) {
//                         il = l
//                     } else {
//                         ir = l
//                     }
//                 } else {
//                     // r is closer to range
//                     if il == ir {
//                         ir = r
//                     } else if diff(r, il) < diff(r, ir) {
//                         il = r
//                     } else {
//                         ir = r
//                     }
//                 }
//             }
//         }
//         // update result range
//         d := ir-il
//         if d < minRangeLen {
//             minRangeLen = d
//             minRange = [2]int{il, ir}
//         } else if d == minRangeLen && il < minRange[0] {
//             minRange = [2]int{il, ir}
//         }
//     }
//     return minRange[:]
// }

func main() {
    // Example 1:
    // Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
    // Output: [20,24]
    // Explanation: 
    // List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
    // List 2: [0, 9, 12, 20], 20 is in range [20,24].
    // List 3: [5, 18, 22, 30], 22 is in range [20,24].
    fmt.Println(smallestRange([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    // Example 2:
    // Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
    // Output: [1,1]
    fmt.Println(smallestRange([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]

    fmt.Println(smallestRange1([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    fmt.Println(smallestRange1([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]

    fmt.Println(smallestRange2([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    fmt.Println(smallestRange2([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]

    fmt.Println(smallestRange3([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    fmt.Println(smallestRange3([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]

    // fmt.Println(smallestRange4([][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}})) // [20,24]
    // fmt.Println(smallestRange4([][]int{{1,2,3},{1,2,3},{1,2,3}})) //  [1,1]
}