package main

// 2411. Smallest Subarrays With Maximum Bitwise OR
// You are given a 0-indexed array nums of length n, consisting of non-negative integers. 
// For each index i from 0 to n - 1, you must determine the size of the minimum sized non-empty subarray of nums starting at i (inclusive) that has the maximum possible bitwise OR.

// In other words, let Bij be the bitwise OR of the subarray nums[i...j]. 
// You need to find the smallest subarray starting at i, such that bitwise OR of this subarray is equal to max(Bik) where i <= k <= n - 1.
//     The bitwise OR of an array is the bitwise OR of all the numbers in it.

// Return an integer array answer of size n where answer[i] is the length of the minimum sized subarray starting at i with maximum bitwise OR.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,0,2,1,3]
// Output: [3,3,2,2,1]
// Explanation:
// The maximum possible bitwise OR starting at any index is 3. 
// - Starting at index 0, the shortest subarray that yields it is [1,0,2].
// - Starting at index 1, the shortest subarray that yields the maximum bitwise OR is [0,2,1].
// - Starting at index 2, the shortest subarray that yields the maximum bitwise OR is [2,1].
// - Starting at index 3, the shortest subarray that yields the maximum bitwise OR is [1,3].
// - Starting at index 4, the shortest subarray that yields the maximum bitwise OR is [3].
// Therefore, we return [3,3,2,2,1]. 

// Example 2:
// Input: nums = [1,2]
// Output: [2,1]
// Explanation:
// Starting at index 0, the shortest subarray that yields the maximum bitwise OR is of length 2.
// Starting at index 1, the shortest subarray that yields the maximum bitwise OR is of length 1.
// Therefore, we return [2,1].

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"

func smallestSubarrays(nums []int) []int {
    bitmap := make([]int, 30)
    getMax := func(arr []int, i int) int {
        res := i
        for _, v := range arr {
            if v > res {
                res = v
            }
        }
        return res
    }
    for i := len(nums) - 1; i >= 0; i-- {
        for j := 0; (1 << j) <= nums[i]; j++ {
            if nums[i] & (1 << j) != 0 {
                bitmap[j] = i
            }
        }
        nums[i] = getMax(bitmap[:], i) - i + 1
    }
    return nums
}

func smallestSubarrays1(nums []int) []int {
    res := make([]int, len(nums))
    for i, v := range nums {
        res[i] = 1
        for j := i - 1; j >= 0 && nums[j] | v != nums[j]; j-- {
            nums[j] |= v
            res[j] = i - j + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,0,2,1,3]
    // Output: [3,3,2,2,1]
    // Explanation:
    // The maximum possible bitwise OR starting at any index is 3. 
    // - Starting at index 0, the shortest subarray that yields it is [1,0,2].
    // - Starting at index 1, the shortest subarray that yields the maximum bitwise OR is [0,2,1].
    // - Starting at index 2, the shortest subarray that yields the maximum bitwise OR is [2,1].
    // - Starting at index 3, the shortest subarray that yields the maximum bitwise OR is [1,3].
    // - Starting at index 4, the shortest subarray that yields the maximum bitwise OR is [3].
    // Therefore, we return [3,3,2,2,1]. 
    fmt.Println(smallestSubarrays([]int{1,0,2,1,3})) // [3,3,2,2,1]
    // Example 2:
    // Input: nums = [1,2]
    // Output: [2,1]
    // Explanation:
    // Starting at index 0, the shortest subarray that yields the maximum bitwise OR is of length 2.
    // Starting at index 1, the shortest subarray that yields the maximum bitwise OR is of length 1.
    // Therefore, we return [2,1].
    fmt.Println(smallestSubarrays([]int{1,2})) // [2,1]

    fmt.Println(smallestSubarrays([]int{1,2,3,4,5,6,7,8,9})) // [8 7 6 5 4 3 2 2 1]
    fmt.Println(smallestSubarrays([]int{9,8,7,6,5,4,3,2,1})) // [3 2 1 2 3 2 1 2 1]

    fmt.Println(smallestSubarrays1([]int{1,0,2,1,3})) // [3,3,2,2,1]
    fmt.Println(smallestSubarrays1([]int{1,2})) // [2,1]
    fmt.Println(smallestSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // [8 7 6 5 4 3 2 2 1]
    fmt.Println(smallestSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // [3 2 1 2 3 2 1 2 1]
}