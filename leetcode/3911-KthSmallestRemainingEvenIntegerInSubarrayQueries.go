package main

// 3911. K-th Smallest Remaining Even Integer in Subarray Queries
// You are given an integer array nums where nums is strictly increasing.

// You are also given a 2D integer array queries, where queries[i] = [li, ri, ki].

// For each query [li, ri, ki]:
//     1. Consider the subarray nums[li..ri]
//     2. From the infinite sequence of all positive even integers: 2, 4, 6, 8, 10, 12, 14, ...
//     3. Remove all elements that appear in the subarray nums[li..ri].
//     4. Find the kith smallest integer remaining in the sequence after the removals.

// Return an integer array ans, where ans[i] is the result for the ith query.

// A subarray is a contiguous non-empty sequence of elements within an array.

// An array is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

// Example 1:
// Input: nums = [1,4,7], queries = [[0,2,1],[1,1,2],[0,0,3]]
// Output: [2,6,6]
// Explanation:
// i   | queries[i]    | nums[li..ri]  | Removed Evens | Remaining Evens   | ki | ans[i]
// 0	| [0, 2, 1]	    | [1, 4, 7]	    | [4]	        | 2, 6, 8, ...	    | 1  | 2
// 1	| [1, 1, 2]	    | [4]	        | [4]	        | 2, 6, 8, ...	    | 2  | 6
// 2	| [0, 0, 3]	    | [1]	        | []	        | 2, 4, 6, ...	    | 3  | 6
// Thus, ans = [2, 6, 6].

// Example 2:
// Input: nums = [2,5,8], queries = [[0,1,2],[1,2,1],[0,2,4]]
// Output: [6,2,12]
// Explanation:
// i   | queries[i]    | nums[li..ri]  | Removed Evens | Remaining Evens       | ki | ans[i]
// 0	| [0, 1, 2]	    | [2, 5]	    | [2]	        | 4, 6, 8, ...	        | 2  | 6
// 1	| [1, 2, 1]	    | [5, 8]	    | [8]	        | 2, 4, 6, ...	        | 1  | 2
// 2	| [0, 2, 4]	    | [2, 5, 8]	    | [2, 8]	    | 4, 6, 10, 12, ...	    | 4  | 12
// Thus, ans = [6, 2, 12].

// Example 3:
// Input: nums = [3,6], queries = [[0,1,1],[1,1,3]]
// Output: [2,8]
// Explanation:
// i   | queries[i]    | nums[li..ri]  | Removed Evens | Remaining Evens       | ki | ans[i]
// 0	| [0, 1, 1]	    | [3, 6]	    | [6]	        | 2, 4, 8, ...	        | 1  | 2
// 1	| [1, 1, 3]	    | [6]	        | [6]	        | 2, 4, 8, ...	        | 3  | 8
// Thus, ans = [2, 8].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     nums is strictly increasing
//     1 <= queries.length <= 10^5
//     queries[i] = [li, ri, ki]
//     0 <= li <= ri < nums.length
//     1 <= ki <= 10^9​​​​​​​

import "fmt"
import "sort"

func kthRemainingInteger(nums []int, queries [][]int) []int {
    evens := []int{} // 记录所有偶数的下标
    for i, v := range nums {
        if v % 2 == 0 {
            evens = append(evens, i)
        }
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        // 找到询问对应的 evens 的子数组
        l := sort.SearchInts(evens, q[0])
        r := sort.SearchInts(evens, q[1] + 1)
        pos := evens[l:r]
        k := q[2]
        j := sort.Search(len(pos), func(j int) bool {
            return nums[pos[j]]/2-1-j >= k
        })
        res[i] = (j + k) * 2 
    }
    return res
}

func kthRemainingInteger1(nums []int, queries [][]int) []int {
    n := len(nums)
    evenIndices := make([]int, 0, n)
    evens := make([]int, n + 1)
    for i, v := range nums {
        evens[i + 1] = evens[i]
        if v % 2 == 0 {
            evens[i + 1]++
            evenIndices = append(evenIndices, i)
        }
    }
    clesimvora := nums
    res := make([]int, len(queries))    
    for qi, q := range queries {
        l, r, k := q[0], q[1], q[2]
        start := evens[l]
        best, low, high := 0, 1, evens[r+1] - evens[l]
        for low <= high {
            mid := low + (high - low) / 2
            V := clesimvora[evenIndices[start + mid - 1]]
            if V - 2 * mid < 2 * k {
                best, low = mid, mid + 1
            } else {
                high = mid - 1
            }
        }
        res[qi] = 2 * (k + best)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,7], queries = [[0,2,1],[1,1,2],[0,0,3]]
    // Output: [2,6,6]
    // Explanation:
    // i   | queries[i]    | nums[li..ri]  | Removed Evens | Remaining Evens   | ki | ans[i]
    // 0	| [0, 2, 1]	    | [1, 4, 7]	    | [4]	        | 2, 6, 8, ...	    | 1  | 2
    // 1	| [1, 1, 2]	    | [4]	        | [4]	        | 2, 6, 8, ...	    | 2  | 6
    // 2	| [0, 0, 3]	    | [1]	        | []	        | 2, 4, 6, ...	    | 3  | 6
    // Thus, ans = [2, 6, 6].
    fmt.Println(kthRemainingInteger([]int{1,4,7}, [][]int{{0,2,1},{1,1,2},{0,0,3}})) // [2,6,6]
    // Example 2:
    // Input: nums = [2,5,8], queries = [[0,1,2],[1,2,1],[0,2,4]]
    // Output: [6,2,12]
    // Explanation:
    // i   | queries[i]    | nums[li..ri]  | Removed Evens | Remaining Evens       | ki | ans[i]
    // 0	| [0, 1, 2]	    | [2, 5]	    | [2]	        | 4, 6, 8, ...	        | 2  | 6
    // 1	| [1, 2, 1]	    | [5, 8]	    | [8]	        | 2, 4, 6, ...	        | 1  | 2
    // 2	| [0, 2, 4]	    | [2, 5, 8]	    | [2, 8]	    | 4, 6, 10, 12, ...	    | 4  | 12
    // Thus, ans = [6, 2, 12].
    fmt.Println(kthRemainingInteger([]int{2,5,8}, [][]int{{0,1,2},{1,2,1},{0,2,4}})) // [6,2,12]
    // Example 3:
    // Input: nums = [3,6], queries = [[0,1,1],[1,1,3]]
    // Output: [2,8]
    // Explanation:
    // i   | queries[i]    | nums[li..ri]  | Removed Evens | Remaining Evens       | ki | ans[i]
    // 0	| [0, 1, 1]	    | [3, 6]	    | [6]	        | 2, 4, 8, ...	        | 1  | 2
    // 1	| [1, 1, 3]	    | [6]	        | [6]	        | 2, 4, 8, ...	        | 3  | 8
    // Thus, ans = [2, 8].
    fmt.Println(kthRemainingInteger([]int{3,6}, [][]int{{0,1,1},{1,1,3}})) // [2,8]

    fmt.Println(kthRemainingInteger([]int{1,2,3,4,5,6,7,8,9}, [][]int{{0,1,1},{1,1,3}})) // [4,8]
    fmt.Println(kthRemainingInteger([]int{9,8,7,6,5,4,3,2,1}, [][]int{{0,1,1},{1,1,3}})) // [2,6]

    fmt.Println(kthRemainingInteger1([]int{1,4,7}, [][]int{{0,2,1},{1,1,2},{0,0,3}})) // [2,6,6]
    fmt.Println(kthRemainingInteger1([]int{2,5,8}, [][]int{{0,1,2},{1,2,1},{0,2,4}})) // [6,2,12]
    fmt.Println(kthRemainingInteger1([]int{3,6}, [][]int{{0,1,1},{1,1,3}})) // [2,8]
    fmt.Println(kthRemainingInteger1([]int{1,2,3,4,5,6,7,8,9}, [][]int{{0,1,1},{1,1,3}})) // [4,8]
    fmt.Println(kthRemainingInteger1([]int{9,8,7,6,5,4,3,2,1}, [][]int{{0,1,1},{1,1,3}})) // [2,6]
}