package main

// 3362. Zero Array Transformation III
// You are given an integer array nums of length n and a 2D array queries where queries[i] = [li, ri].

// Each queries[i] represents the following action on nums:
//     1.  Decrement the value at each index in the range [li, ri] in nums by at most 1.
//     2. The amount by which the value is decremented can be chosen independently for each index.

// A Zero Array is an array with all its elements equal to 0.

// Return the maximum number of elements that can be removed from queries, 
// such that nums can still be converted to a zero array using the remaining queries. 
// If it is not possible to convert nums to a zero array, return -1.

// Example 1:
// Input: nums = [2,0,2], queries = [[0,2],[0,2],[1,1]]
// Output: 1
// Explanation:
// After removing queries[2], nums can still be converted to a zero array.
//     Using queries[0], decrement nums[0] and nums[2] by 1 and nums[1] by 0.
//     Using queries[1], decrement nums[0] and nums[2] by 1 and nums[1] by 0.

// Example 2:
// Input: nums = [1,1,1,1], queries = [[1,3],[0,2],[1,3],[1,2]]
// Output: 2
// Explanation:
// We can remove queries[2] and queries[3].

// Example 3:
// Input: nums = [1,2,3,4], queries = [[0,3]]
// Output: -1
// Explanation:
// nums cannot be converted to a zero array even after using all the queries.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     0 <= li <= ri < nums.length

import "fmt"
import "container/heap"
import "sort"
import "slices"

type MaxHeap struct{ sort.IntSlice }
func (h MaxHeap)  Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *MaxHeap) Push(v any)         { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MaxHeap) Pop() any           { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func maxRemoval(nums []int, queries [][]int) int {
    slices.SortFunc(queries, func(x, y []int) int { // 把 queries 按照左端点排序
        return x[0] - y[0] 
    })
    hp := MaxHeap{} // 最大堆维护左端点 ≤i 的未选区间的右端点
    diff := make([]int, len(nums) + 1)
    sum, j := 0, 0
    for i, v := range nums {// 双指针遍历 nums 和左端点 ≤i 的区间
        sum += diff[i]
        for ; j < len(queries) && queries[j][0] <= i; j++ {
            heap.Push(&hp, queries[j][1])
        }
        for sum < v && hp.Len() > 0 && hp.IntSlice[0] >= i {
            sum++
            diff[heap.Pop(&hp).(int)+1]-- // 用差分数组去维护区间减一
        }
        if sum < v {
            return -1
        }
    }
    return hp.Len()
}

func main() {
    // Example 1:
    // Input: nums = [2,0,2], queries = [[0,2],[0,2],[1,1]]
    // Output: 1
    // Explanation:
    // After removing queries[2], nums can still be converted to a zero array.
    //     Using queries[0], decrement nums[0] and nums[2] by 1 and nums[1] by 0.
    //     Using queries[1], decrement nums[0] and nums[2] by 1 and nums[1] by 0.
    fmt.Println(maxRemoval([]int{2,0,2}, [][]int{{0,2},{0,2},{1,1}})) // 1
    // Example 2:
    // Input: nums = [1,1,1,1], queries = [[1,3],[0,2],[1,3],[1,2]]
    // Output: 2
    // Explanation:
    // We can remove queries[2] and queries[3].
    fmt.Println(maxRemoval([]int{1,1,1,1}, [][]int{{1,3},{0,2},{1,3},{1,2}})) // 2
    // Example 3:
    // Input: nums = [1,2,3,4], queries = [[0,3]]
    // Output: -1
    // Explanation:
    // nums cannot be converted to a zero array even after using all the queries.
    fmt.Println(maxRemoval([]int{1,2,3,4}, [][]int{{0,3}})) // -1
}