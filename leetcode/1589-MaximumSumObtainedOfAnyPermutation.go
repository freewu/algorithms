package main

// 1589. Maximum Sum Obtained of Any Permutation
// We have an array of integers, nums, and an array of requests where requests[i] = [starti, endi]. 
// The ith request asks for the sum of nums[starti] + nums[starti + 1] + ... + nums[endi - 1] + nums[endi]. 
// Both starti and endi are 0-indexed.

// Return the maximum total sum of all requests among all permutations of nums.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3,4,5], requests = [[1,3],[0,1]]
// Output: 19
// Explanation: One permutation of nums is [2,1,3,4,5] with the following result: 
// requests[0] -> nums[1] + nums[2] + nums[3] = 1 + 3 + 4 = 8
// requests[1] -> nums[0] + nums[1] = 2 + 1 = 3
// Total sum: 8 + 3 = 11.
// A permutation with a higher total sum is [3,5,4,2,1] with the following result:
// requests[0] -> nums[1] + nums[2] + nums[3] = 5 + 4 + 2 = 11
// requests[1] -> nums[0] + nums[1] = 3 + 5  = 8
// Total sum: 11 + 8 = 19, which is the best that you can do.

// Example 2:
// Input: nums = [1,2,3,4,5,6], requests = [[0,1]]
// Output: 11
// Explanation: A permutation with the max total sum is [6,5,4,3,2,1] with request sums [11].

// Example 3:
// Input: nums = [1,2,3,4,5,10], requests = [[0,2],[1,3],[1,1]]
// Output: 47
// Explanation: A permutation with the max total sum is [4,10,5,3,2,1] with request sums [19,18,10].

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     0 <= nums[i] <= 10^5
//     1 <= requests.length <= 10^5
//     requests[i].length == 2
//     0 <= starti <= endi < n

import "fmt"
import "sort"

type BITree struct {
    t []int
    l int
}

func (b BITree) sum(n int) int {
    res := 0
    for n > 0 {
        res += b.t[n]
        n = n & (n - 1)
    }
    return res
}

func (b BITree) add(n int, a int) { 
    for n < b.l {
        b.t[n] += a
        n = n + n - n & (n - 1)
    }
}

func maxSumRangeQuery(nums []int, requests [][]int) int {
    res, n := 0, len(nums)
    tree := BITree{ make([]int, n + 1), n + 1 }
    for _, v := range requests {
        tree.add(v[0] + 1, 1) 
        tree.add(v[1] + 2, -1)
    }
    times := make([]int, n)
    for i := 0; i < len(nums); i++ {
        times[i] = tree.sum(i + 1)
    }
    sort.Ints(nums)
    sort.Ints(times)    
    for i := 0; i < len(times); i++ {
        res = (res + times[i] * nums[i]) % 1_000_000_007
    }
    return res
}

func maxSumRangeQuery1(nums []int, requests [][]int) int {
    res, rightMost, n := 0, -1, len(nums) 
    arr := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, request := range requests {
        start, end := request[0], request[1] 
        arr[start] += 1;
        if end + 1 < n { 
            arr[end + 1] -= 1;
        }
        rightMost = max(rightMost, end) // 统计最大右端点 
    }
    freq := make([]int, 0) 
    if arr[0] > 0 {
        freq = append(freq, arr[0]) 
    }
    
    for i := 1; i <= rightMost; i += 1 { // 统计查询次数最多的频数
        arr[i] = arr[i - 1] + arr[i] 
        if arr[i] > 0 {
            freq = append(freq, arr[i]) 
        }
    }
    helper := func(nums []int) {
        mn, mx := nums[0], nums[0] 
        for _, v := range nums {
            mn = min(mn, v)
            mx = max(mx, v)
        }
        counter := make([]int, 1 + mx) 
        for _, v := range nums {
            counter[v] += 1 
        }
        index := 0 
        for i := mn; i <= mx; i++ {
            for counter[i] > 0 {
                nums[index] = i
                index++
                counter[i]-- 
            }
        }
    }
    helper(freq) 
    helper(nums)
    index := (n - 1)
    for j := len(freq) - 1; j >= 0; j-- {
        res = (res + nums[index] * freq[j] ) % 1_000_000_007 
        index -= 1 
    }
    return res 
}


func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], requests = [[1,3],[0,1]]
    // Output: 19
    // Explanation: One permutation of nums is [2,1,3,4,5] with the following result: 
    // requests[0] -> nums[1] + nums[2] + nums[3] = 1 + 3 + 4 = 8
    // requests[1] -> nums[0] + nums[1] = 2 + 1 = 3
    // Total sum: 8 + 3 = 11.
    // A permutation with a higher total sum is [3,5,4,2,1] with the following result:
    // requests[0] -> nums[1] + nums[2] + nums[3] = 5 + 4 + 2 = 11
    // requests[1] -> nums[0] + nums[1] = 3 + 5  = 8
    // Total sum: 11 + 8 = 19, which is the best that you can do.
    fmt.Println(maxSumRangeQuery([]int{1,2,3,4,5}, [][]int{{1,3},{0,1}})) // 19
    // Example 2:
    // Input: nums = [1,2,3,4,5,6], requests = [[0,1]]
    // Output: 11
    // Explanation: A permutation with the max total sum is [6,5,4,3,2,1] with request sums [11].
    fmt.Println(maxSumRangeQuery([]int{1,2,3,4,5,6}, [][]int{{0,1}})) // 11
    // Example 3:
    // Input: nums = [1,2,3,4,5,10], requests = [[0,2],[1,3],[1,1]]
    // Output: 47
    // Explanation: A permutation with the max total sum is [4,10,5,3,2,1] with request sums [19,18,10].
    fmt.Println(maxSumRangeQuery([]int{1,2,3,4,5,10}, [][]int{{0,2},{1,3},{1,1}})) // 47

    fmt.Println(maxSumRangeQuery1([]int{1,2,3,4,5}, [][]int{{1,3},{0,1}})) // 19
    fmt.Println(maxSumRangeQuery1([]int{1,2,3,4,5,6}, [][]int{{0,1}})) // 11
    fmt.Println(maxSumRangeQuery1([]int{1,2,3,4,5,10}, [][]int{{0,2},{1,3},{1,1}})) // 47
}