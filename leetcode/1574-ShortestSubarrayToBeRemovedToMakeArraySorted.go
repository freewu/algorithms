package main

// 1574. Shortest Subarray to be Removed to Make Array Sorted
// Given an integer array arr, remove a subarray (can be empty) from arr such that the remaining elements in arr are non-decreasing.

// Return the length of the shortest subarray to remove.

// A subarray is a contiguous subsequence of the array.

// Example 1:
// Input: arr = [1,2,3,10,4,2,3,5]
// Output: 3
// Explanation: The shortest subarray we can remove is [10,4,2] of length 3. The remaining elements after that will be [1,2,3,3,5] which are sorted.
// Another correct solution is to remove the subarray [3,10,4].

// Example 2:
// Input: arr = [5,4,3,2,1]
// Output: 4
// Explanation: Since the array is strictly decreasing, we can only keep a single element. Therefore we need to remove a subarray of length 4, either [5,4,3,2] or [4,3,2,1].

// Example 3:
// Input: arr = [1,2,3]
// Output: 0
// Explanation: The array is already non-decreasing. We do not need to remove any elements.

// Constraints:
//     1 <= arr.length <= 10^5
//     0 <= arr[i] <= 10^9

import "fmt"

func findLengthOfShortestSubarray(arr []int) int {
    n := len(arr)
    left, right := 0, n - 1
    for ; left < n - 1; left++ {
        if arr[left] > arr[left + 1] { break }
    }
    if left == n -1 { return 0 } // Base case for a non-decreasing sequence
    for ; right > 0; right-- {
        if arr[right] < arr[right - 1] { break }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // Possibilities 1 or 2 as mentioned above
    res := min(n - left -1, right)
    for ; left >= 0; left-- {
        for i := right; i < n; i++ {
            if arr[left] > arr[i] { continue }
            res = min(res, i - left - 1)
            break
        }
    }
    return res
}

func findLengthOfShortestSubarray1(arr []int) int {
    n := len(arr)
    right := n - 1
    for right >= 1 && arr[right] >= arr[right - 1] {
        right--
    }
    if right == 0 {
        return 0
    }
    res := n - right
    for left := 0 ; left == 0 || arr[left - 1] <= arr[left] ; left++ {
        for right < n && arr[right] < arr[left]{
            right++
        }
        res = max(res, left + 1 + n - right)
    }
    return n - res
}

func main() {
    // Example 1:
    // Input: arr = [1,2,3,10,4,2,3,5]
    // Output: 3
    // Explanation: The shortest subarray we can remove is [10,4,2] of length 3. The remaining elements after that will be [1,2,3,3,5] which are sorted.
    // Another correct solution is to remove the subarray [3,10,4].
    fmt.Println(findLengthOfShortestSubarray([]int{1,2,3,10,4,2,3,5})) // 3
    // Example 2:
    // Input: arr = [5,4,3,2,1]
    // Output: 4
    // Explanation: Since the array is strictly decreasing, we can only keep a single element. Therefore we need to remove a subarray of length 4, either [5,4,3,2] or [4,3,2,1].
    fmt.Println(findLengthOfShortestSubarray([]int{5,4,3,2,1})) // 4
    // Example 3:
    // Input: arr = [1,2,3]
    // Output: 0
    // Explanation: The array is already non-decreasing. We do not need to remove any elements.
    fmt.Println(findLengthOfShortestSubarray([]int{1,2,3})) // 0

    fmt.Println(findLengthOfShortestSubarray1([]int{1,2,3,10,4,2,3,5})) // 3
    fmt.Println(findLengthOfShortestSubarray1([]int{5,4,3,2,1})) // 4
    fmt.Println(findLengthOfShortestSubarray1([]int{1,2,3})) // 0
}