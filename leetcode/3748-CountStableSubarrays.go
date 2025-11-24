package main

// 3748. Count Stable Subarrays
// You are given an integer array nums.

// A subarray of nums is called stable if it contains no inversions, i.e., there is no pair of indices i < j such that nums[i] > nums[j].

// You are also given a 2D integer array queries of length q, where each queries[i] = [li, ri] represents a query. 
// For each query [li, ri], compute the number of stable subarrays that lie entirely within the segment nums[li..ri].

// Return an integer array ans of length q, where ans[i] is the answer to the ith query.​​​​​​​​​​​​​​

// Note: 
//     A single element subarray is considered stable.

// Example 1:
// Input: nums = [3,1,2], queries = [[0,1],[1,2],[0,2]]
// Output: [2,3,4]
// Explanation:​​​​​
// For queries[0] = [0, 1], the subarray is [nums[0], nums[1]] = [3, 1].
// The stable subarrays are [3] and [1]. The total number of stable subarrays is 2.
// For queries[1] = [1, 2], the subarray is [nums[1], nums[2]] = [1, 2].
// The stable subarrays are [1], [2], and [1, 2]. The total number of stable subarrays is 3.
// For queries[2] = [0, 2], the subarray is [nums[0], nums[1], nums[2]] = [3, 1, 2].
// The stable subarrays are [3], [1], [2], and [1, 2]. The total number of stable subarrays is 4.
// Thus, ans = [2, 3, 4].

// Example 2:
// Input: nums = [2,2], queries = [[0,1],[0,0]]
// Output: [3,1]
// Explanation:
// For queries[0] = [0, 1], the subarray is [nums[0], nums[1]] = [2, 2].
// The stable subarrays are [2], [2], and [2, 2]. The total number of stable subarrays is 3.
// For queries[1] = [0, 0], the subarray is [nums[0]] = [2].
// The stable subarray is [2]. The total number of stable subarrays is 1.
// Thus, ans = [3, 1].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i] = [li, ri]
//     0 <= li <= ri <= nums.length - 1

import "fmt"

func countStableSubarrays(nums []int, queries [][]int) []int64 {
    n := len(nums)
    prefix := make([]int64, n)
    prefix[0] = 1
    for i := 1; i < n; i++ {
        if nums[i-1] <= nums[i] {
            prefix[i] = prefix[i-1] + 1
        } else {
            prefix[i] = 1
        }
    }
    for i := 1; i < n; i++ {
        prefix[i] += prefix[i-1]
    }
    end := make([]int, n)
    end[n-1] = n - 1
    for i := n - 2; i >= 0; i-- {
        if nums[i] <= nums[i+1] {
            end[i] = end[i+1]
        } else {
            end[i] = i
        }
    }
    n = len(queries)
    res := make([]int64, n)
    for i, q := range queries {
        l, r := q[0], q[1]
        ll := min(r, end[l])    
        k := int64(ll - l + 1)
        res[i] = k*(k+1)/2 + prefix[r] - prefix[ll]
    }
    return res
}

func countStableSubarrays1(nums []int, queries [][]int) []int64 {
    count, n := 0, len(nums)
    sum := make([]int64, n + 1)   // 计算递增子数组的前缀和
    for i, v := range nums {
        if i > 0 && nums[i-1] > v {
            count = 0
        }
        count++ // 统计递增子数组的个数
        sum[i+1] = sum[i] + int64(count)
    }
    next := make([]int, n)
    next[n-1] = n
    for i := n - 2; i >= 0; i-- {
        if nums[i] <= nums[i+1] {
            next[i] = next[i+1]
        } else {
            next[i] = i + 1
        }
    }
    res := make([]int64, len(queries))
    for k, q := range queries {
        l, r := q[0], q[1]
        ll := next[l]
        if ll > r {  // 查询左右端点位于同一个递增区间中
            m := int64(r-l+1)
            res[k] = m * (m+1) / 2
            continue
        }
        m := int64(ll - l)
        res[k] = m*(m+1)/2 + sum[r+1] - sum[ll]   // l ~ ll 位于同一个递增区间中,  ll ~ r 可以用前缀和求出
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2], queries = [[0,1],[1,2],[0,2]]
    // Output: [2,3,4]
    // Explanation:​​​​​
    // For queries[0] = [0, 1], the subarray is [nums[0], nums[1]] = [3, 1].
    // The stable subarrays are [3] and [1]. The total number of stable subarrays is 2.
    // For queries[1] = [1, 2], the subarray is [nums[1], nums[2]] = [1, 2].
    // The stable subarrays are [1], [2], and [1, 2]. The total number of stable subarrays is 3.
    // For queries[2] = [0, 2], the subarray is [nums[0], nums[1], nums[2]] = [3, 1, 2].
    // The stable subarrays are [3], [1], [2], and [1, 2]. The total number of stable subarrays is 4.
    // Thus, ans = [2, 3, 4].
    fmt.Println(countStableSubarrays([]int{3,1,2}, [][]int{{0,1},{1,2},{0,2}})) // [2, 3, 4]
    // Example 2:
    // Input: nums = [2,2], queries = [[0,1],[0,0]]
    // Output: [3,1]
    // Explanation:
    // For queries[0] = [0, 1], the subarray is [nums[0], nums[1]] = [2, 2].
    // The stable subarrays are [2], [2], and [2, 2]. The total number of stable subarrays is 3.
    // For queries[1] = [0, 0], the subarray is [nums[0]] = [2].
    // The stable subarray is [2]. The total number of stable subarrays is 1.
    // Thus, ans = [3, 1].
    fmt.Println(countStableSubarrays([]int{2,2}, [][]int{{0,1},{0,0}})) // [3, 1]   

    fmt.Println(countStableSubarrays([]int{1,2,3,4,5,6,7,8,9}, [][]int{{0,1},{0,0}})) // [3 1]   
    fmt.Println(countStableSubarrays([]int{9,8,7,6,5,4,3,2,1}, [][]int{{0,1},{0,0}})) // [2 1]  

    fmt.Println(countStableSubarrays1([]int{3,1,2}, [][]int{{0,1},{1,2},{0,2}})) // [2, 3, 4]
    fmt.Println(countStableSubarrays1([]int{2,2}, [][]int{{0,1},{0,0}})) // [3, 1]
    fmt.Println(countStableSubarrays1([]int{1,2,3,4,5,6,7,8,9}, [][]int{{0,1},{0,0}})) // [3 1]
    fmt.Println(countStableSubarrays1([]int{9,8,7,6,5,4,3,2,1}, [][]int{{0,1},{0,0}})) // [2 1]
}