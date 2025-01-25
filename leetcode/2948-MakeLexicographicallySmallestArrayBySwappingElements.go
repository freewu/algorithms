package main

// 2948. Make Lexicographically Smallest Array by Swapping Elements
// You are given a 0-indexed array of positive integers nums and a positive integer limit.

// In one operation, you can choose any two indices i and j and swap nums[i] and nums[j] if |nums[i] - nums[j]| <= limit.

// Return the lexicographically smallest array that can be obtained by performing the operation any number of times.

// An array a is lexicographically smaller than an array b if in the first position where a and b differ, 
// array a has an element that is less than the corresponding element in b. 
// For example, the array [2,10,3] is lexicographically smaller than the array [10,2,3] because they differ at index 0 and 2 < 10.

// Example 1:
// Input: nums = [1,5,3,9,8], limit = 2
// Output: [1,3,5,8,9]
// Explanation: Apply the operation 2 times:
// - Swap nums[1] with nums[2]. The array becomes [1,3,5,9,8]
// - Swap nums[3] with nums[4]. The array becomes [1,3,5,8,9]
// We cannot obtain a lexicographically smaller array by applying any more operations.
// Note that it may be possible to get the same result by doing different operations.

// Example 2:
// Input: nums = [1,7,6,18,2,1], limit = 3
// Output: [1,6,7,18,1,2]
// Explanation: Apply the operation 3 times:
// - Swap nums[1] with nums[2]. The array becomes [1,6,7,18,2,1]
// - Swap nums[0] with nums[4]. The array becomes [2,6,7,18,1,1]
// - Swap nums[0] with nums[5]. The array becomes [1,6,7,18,1,2]
// We cannot obtain a lexicographically smaller array by applying any more operations.

// Example 3:
// Input: nums = [1,7,28,19,10], limit = 3
// Output: [1,7,28,19,10]
// Explanation: [1,7,28,19,10] is the lexicographically smallest array we can obtain because we cannot apply the operation on any two indices.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= limit <= 10^9

import "fmt"
import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
    n := len(nums)
    index := make([]int, n)
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(i, j int) int { 
        return nums[i] - nums[j] 
    })
    res := make([]int, n)
    for i := 0; i < n; {
        j := i + 1
        for j < n && nums[index[j]] - nums[index[j - 1]] <= limit {
            j++
        }
        arr := slices.Clone(index[i:j])
        slices.Sort(arr)
        for k := i; k < j; k++ {
            res[arr[k - i]] = nums[index[k]]
        }
        i = j
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,5,3,9,8], limit = 2
    // Output: [1,3,5,8,9]
    // Explanation: Apply the operation 2 times:
    // - Swap nums[1] with nums[2]. The array becomes [1,3,5,9,8]
    // - Swap nums[3] with nums[4]. The array becomes [1,3,5,8,9]
    // We cannot obtain a lexicographically smaller array by applying any more operations.
    // Note that it may be possible to get the same result by doing different operations.
    fmt.Println(lexicographicallySmallestArray([]int{1,5,3,9,8}, 2)) // [1,3,5,8,9]
    // Example 2:
    // Input: nums = [1,7,6,18,2,1], limit = 3
    // Output: [1,6,7,18,1,2]
    // Explanation: Apply the operation 3 times:
    // - Swap nums[1] with nums[2]. The array becomes [1,6,7,18,2,1]
    // - Swap nums[0] with nums[4]. The array becomes [2,6,7,18,1,1]
    // - Swap nums[0] with nums[5]. The array becomes [1,6,7,18,1,2]
    // We cannot obtain a lexicographically smaller array by applying any more operations.
    fmt.Println(lexicographicallySmallestArray([]int{1,7,6,18,2,1}, 3)) // [1,6,7,18,1,2]
    // Example 3:
    // Input: nums = [1,7,28,19,10], limit = 3
    // Output: [1,7,28,19,10]
    // Explanation: [1,7,28,19,10] is the lexicographically smallest array we can obtain because we cannot apply the operation on any two indices.
    fmt.Println(lexicographicallySmallestArray([]int{1,7,28,19,10}, 3)) // [1,7,28,19,10]

    fmt.Println(lexicographicallySmallestArray([]int{1,2,3,4,5,6,7,8,9}, 3)) // [1,2,3,4,5,6,7,8,9]
    fmt.Println(lexicographicallySmallestArray([]int{9,8,7,6,5,4,3,2,1}, 3)) // [1,2,3,4,5,6,7,8,9]
}