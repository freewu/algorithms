package main

// 480. Sliding Window Median
// The median is the middle value in an ordered integer list. 
// If the size of the list is even, there is no middle value. So the median is the mean of the two middle values.
//     For examples, if arr = [2,3,4], the median is 3.
//     For examples, if arr = [1,2,3,4], the median is (2 + 3) / 2 = 2.5.

// You are given an integer array nums and an integer k. 
// There is a sliding window of size k which is moving from the very left of the array to the very right. 
// You can only see the k numbers in the window. Each time the sliding window moves right by one position.

// Return the median array for each window in the original array. 
// Answers within 10-5 of the actual value will be accepted.

// Example 1:
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
// Explanation: 
// Window position                Median
// ---------------                -----
// [1  3  -1] -3  5  3  6  7        1
//  1 [3  -1  -3] 5  3  6  7       -1
//  1  3 [-1  -3  5] 3  6  7       -1
//  1  3  -1 [-3  5  3] 6  7        3
//  1  3  -1  -3 [5  3  6] 7        5
//  1  3  -1  -3  5 [3  6  7]       6

// Example 2:
// Input: nums = [1,2,3,4,2,3,1,4,2], k = 3
// Output: [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     -2^31 <= nums[i] <= 2^31 - 1

import "fmt"
import "sort"
import "container/heap"

// Time Limit Exceeded 42 / 43
func medianSlidingWindow(nums []int, k int) []float64 {
    res := make([]float64,len(nums) - k + 1)
    // 获取数组中位数据
    getMedian := func(arr []int) float64 {
        // fmt.Println(arr)
        sort.Ints(arr)
        l := len(arr)
        if 0 == l % 2 {
            return (float64(arr[l/2]) + float64(arr[l/2 - 1])) / 2.0
        } else {
            return float64(arr[l/2])
        }
    }
    for i := 0; i < len(nums) - k + 1; i++ {
        arr,count := make([]int,k), 0
        // 滑动 k 步
        for count < k {
            arr[count] = nums[i + count]
            count++
        }
        // 做个取中位数的函数解决
        res[i] = getMedian(arr)
        //res[i] = float64(sum) / float64(k)
    }
    return res
}

// 使用 大小堆
func medianSlidingWindow1(nums []int, k int) []float64 {
    medians, smallNumbers, largeNumbers := []float64{}, new(MaxHeap), new(MinHeap)
    heap.Init(smallNumbers)
    heap.Init(largeNumbers)
    hashMap := make(map[int]int)
    i := 0
    for i < k {
        heap.Push(smallNumbers, nums[i])
        i++
    }
    for j := 0; j < k/2; j++ {
        heap.Push(largeNumbers, smallNumbers.Top())
        heap.Pop(smallNumbers)
    }
    for {
        if k % 2 == 1{ 
            medians = append(medians,float64(smallNumbers.Top()))
        } else {
            temp := (float64(smallNumbers.Top()) + float64(largeNumbers.Top())) * 0.5
            medians = append(medians,float64(temp))
        }
        if i >= len(nums) {
            break
        }  
        income, outcome := nums[i], nums[i-k]
        i++
        balance := 0  
        if outcome <= smallNumbers.Top(){
            balance -= 1
        } else {
            balance += 1
        }
        if _, ok := hashMap[outcome]; ok {  
            hashMap[outcome]++
        } else{
            hashMap[outcome] = 1
        }
        if !smallNumbers.Empty() && income <= smallNumbers.Top() {
            balance++
            heap.Push(smallNumbers, income)
        } else {
            balance--
            heap.Push(largeNumbers, income)
        }
        if balance < 0 {
            heap.Push(smallNumbers, largeNumbers.Top())
            heap.Pop(largeNumbers)
            balance++
        }
        if balance > 0 {
            heap.Push(largeNumbers, smallNumbers.Top())
            heap.Pop(smallNumbers)
            balance--
        }
        num, _ := hashMap[smallNumbers.Top()]
        for num > 0 {
            hashMap[smallNumbers.Top()]--
            heap.Pop(smallNumbers)
            num, _ = hashMap[smallNumbers.Top()]
        }

        num, _ = hashMap[largeNumbers.Top()]
        for !largeNumbers.Empty() && num > 0 {
            hashMap[largeNumbers.Top()]--
            heap.Pop(largeNumbers)
            num, _ = hashMap[largeNumbers.Top()]
        }
    }
    return medians
}

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface {}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface {} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
func (h MinHeap) Top() int {
    if h.Empty() {
        return 0
    } else {
        return h[0]
    }
}
func (h MinHeap) Empty() bool {
    return len(h) == 0
}

type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface {}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface {} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
func (h MaxHeap) Top() int {
    if h.Empty() {
        return 0
    } else {
        return h[0]
    }
}
func (h MaxHeap) Empty() bool {
    return len(h) == 0
}

func main() {
    // Window position                Median
    // ---------------                -----
    // [1  3  -1] -3  5  3  6  7        1
    //  1 [3  -1  -3] 5  3  6  7       -1
    //  1  3 [-1  -3  5] 3  6  7       -1
    //  1  3  -1 [-3  5  3] 6  7        3
    //  1  3  -1  -3 [5  3  6] 7        5
    //  1  3  -1  -3  5 [3  6  7]       6
    fmt.Println(medianSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3)) // [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
    fmt.Println(medianSlidingWindow([]int{1,2,3,4,2,3,1,4,2},3)) // [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]

    fmt.Println(medianSlidingWindow1([]int{1,3,-1,-3,5,3,6,7},3)) // [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
    fmt.Println(medianSlidingWindow1([]int{1,2,3,4,2,3,1,4,2},3)) // [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]
}