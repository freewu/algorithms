package main

// 2382. Maximum Segment Sum After Removals
// You are given two 0-indexed integer arrays nums and removeQueries, both of length n. 
// For the ith query, the element in nums at the index removeQueries[i] is removed, splitting nums into different segments.

// A segment is a contiguous sequence of positive integers in nums. 
// A segment sum is the sum of every element in a segment.

// Return an integer array answer, of length n, where answer[i] is the maximum segment sum after applying the ith removal.

// Note: The same index will not be removed more than once.

// Example 1:
// Input: nums = [1,2,5,6,1], removeQueries = [0,3,2,4,1]
// Output: [14,7,2,2,0]
// Explanation: Using 0 to indicate a removed element, the answer is as follows:
// Query 1: Remove the 0th element, nums becomes [0,2,5,6,1] and the maximum segment sum is 14 for segment [2,5,6,1].
// Query 2: Remove the 3rd element, nums becomes [0,2,5,0,1] and the maximum segment sum is 7 for segment [2,5].
// Query 3: Remove the 2nd element, nums becomes [0,2,0,0,1] and the maximum segment sum is 2 for segment [2]. 
// Query 4: Remove the 4th element, nums becomes [0,2,0,0,0] and the maximum segment sum is 2 for segment [2]. 
// Query 5: Remove the 1st element, nums becomes [0,0,0,0,0] and the maximum segment sum is 0, since there are no segments.
// Finally, we return [14,7,2,2,0].

// Example 2:
// Input: nums = [3,2,11,1], removeQueries = [3,2,1,0]
// Output: [16,5,3,0]
// Explanation: Using 0 to indicate a removed element, the answer is as follows:
// Query 1: Remove the 3rd element, nums becomes [3,2,11,0] and the maximum segment sum is 16 for segment [3,2,11].
// Query 2: Remove the 2nd element, nums becomes [3,2,0,0] and the maximum segment sum is 5 for segment [3,2].
// Query 3: Remove the 1st element, nums becomes [3,0,0,0] and the maximum segment sum is 3 for segment [3].
// Query 4: Remove the 0th element, nums becomes [0,0,0,0] and the maximum segment sum is 0, since there are no segments.
// Finally, we return [16,5,3,0].

// Constraints:
//     n == nums.length == removeQueries.length
//     1 <= n <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= removeQueries[i] < n
//     All the values of removeQueries are unique.

import "fmt"

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
    n, mx := len(nums), 0
    res, segs, data := make([]int64, n),  make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        data[i] = -1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    find := func(data []int, pos int) int {
        for data[pos] != pos {
            pos = data[pos]
        }
        return pos
    }
    for i := n - 1; i >= 0; i-- {
        query := removeQueries[i]
        // initial segment size
        segs[query], data[query] = nums[query], query
        // check if we can join our segment from left side
        if query > 0 {
            l := data[query-1]
            if l >= 0 {
                leftSeg := find(data, query - 1)
                data[leftSeg] = query // union
                segs[query] += segs[leftSeg] // increase segment
            }
        }
        // check if we can join our segment from right side
        if query < n - 1 {
            r := data[query + 1]
            if r >= 0 {
                rightSeg := find(data, query + 1)
                data[rightSeg] = query // union
                segs[query] += segs[rightSeg] // increase segment
            }
        }
        res[i] = int64(mx)
        mx = max(mx, segs[query]) // increase max segment if possible
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,5,6,1], removeQueries = [0,3,2,4,1]
    // Output: [14,7,2,2,0]
    // Explanation: Using 0 to indicate a removed element, the answer is as follows:
    // Query 1: Remove the 0th element, nums becomes [0,2,5,6,1] and the maximum segment sum is 14 for segment [2,5,6,1].
    // Query 2: Remove the 3rd element, nums becomes [0,2,5,0,1] and the maximum segment sum is 7 for segment [2,5].
    // Query 3: Remove the 2nd element, nums becomes [0,2,0,0,1] and the maximum segment sum is 2 for segment [2]. 
    // Query 4: Remove the 4th element, nums becomes [0,2,0,0,0] and the maximum segment sum is 2 for segment [2]. 
    // Query 5: Remove the 1st element, nums becomes [0,0,0,0,0] and the maximum segment sum is 0, since there are no segments.
    // Finally, we return [14,7,2,2,0].
    fmt.Println(maximumSegmentSum([]int{1,2,5,6,1}, []int{0,3,2,4,1})) // [14,7,2,2,0]
    // Example 2:
    // Input: nums = [3,2,11,1], removeQueries = [3,2,1,0]
    // Output: [16,5,3,0]
    // Explanation: Using 0 to indicate a removed element, the answer is as follows:
    // Query 1: Remove the 3rd element, nums becomes [3,2,11,0] and the maximum segment sum is 16 for segment [3,2,11].
    // Query 2: Remove the 2nd element, nums becomes [3,2,0,0] and the maximum segment sum is 5 for segment [3,2].
    // Query 3: Remove the 1st element, nums becomes [3,0,0,0] and the maximum segment sum is 3 for segment [3].
    // Query 4: Remove the 0th element, nums becomes [0,0,0,0] and the maximum segment sum is 0, since there are no segments.
    // Finally, we return [16,5,3,0].
    fmt.Println(maximumSegmentSum([]int{3,2,11,1}, []int{3,2,1,0})) // [16,5,3,0]
}