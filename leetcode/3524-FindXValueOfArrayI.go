package main

// 3524. Find X Value of Array I
// You are given an array of positive integers nums, and a positive integer k.

// You are allowed to perform an operation once on nums, 
// where in each operation you can remove any non-overlapping prefix and suffix from nums such that nums remains non-empty.

// You need to find the x-value of nums, 
// which is the number of ways to perform this operation so that the product of the remaining elements leaves a remainder of x when divided by k.

// Return an array result of size k where result[x] is the x-value of nums for 0 <= x <= k - 1.

// A prefix of an array is a subarray that starts from the beginning of the array and extends to any point within it.

// A suffix of an array is a subarray that starts at any point within the array and extends to the end of the array.

// Note that the prefix and suffix to be chosen for the operation can be empty.

// Example 1:
// Input: nums = [1,2,3,4,5], k = 3
// Output: [9,2,4]
// Explanation:
// For x = 0, the possible operations include all possible ways to remove non-overlapping prefix/suffix that do not remove nums[2] == 3.
// For x = 1, the possible operations are:
//     Remove the empty prefix and the suffix [2, 3, 4, 5]. nums becomes [1].
//     Remove the prefix [1, 2, 3] and the suffix [5]. nums becomes [4].
// For x = 2, the possible operations are:
//     Remove the empty prefix and the suffix [3, 4, 5]. nums becomes [1, 2].
//     Remove the prefix [1] and the suffix [3, 4, 5]. nums becomes [2].
//     Remove the prefix [1, 2, 3] and the empty suffix. nums becomes [4, 5].
//     Remove the prefix [1, 2, 3, 4] and the empty suffix. nums becomes [5].

// Example 2:
// Input: nums = [1,2,4,8,16,32], k = 4
// Output: [18,1,2,0]
// Explanation:
// For x = 0, the only operations that do not result in x = 0 are:
//     Remove the empty prefix and the suffix [4, 8, 16, 32]. nums becomes [1, 2].
//     Remove the empty prefix and the suffix [2, 4, 8, 16, 32]. nums becomes [1].
//     Remove the prefix [1] and the suffix [4, 8, 16, 32]. nums becomes [2].
// For x = 1, the only possible operation is:
//     Remove the empty prefix and the suffix [2, 4, 8, 16, 32]. nums becomes [1].
// For x = 2, the possible operations are:
//     Remove the empty prefix and the suffix [4, 8, 16, 32]. nums becomes [1, 2].
//     Remove the prefix [1] and the suffix [4, 8, 16, 32]. nums becomes [2].
// For x = 3, there is no possible way to perform the operation.

// Example 3:
// Input: nums = [1,1,2,1,1], k = 2
// Output: [9,6]

// Constraints:
//     1 <= nums[i] <= 10^9
//     1 <= nums.length <= 10^5
//     1 <= k <= 5

import "fmt"

func resultArray(nums []int, k int) []int64 {
    res, arr1 := make([]int64, k), make([]int, k)
    for i := 0; i < len(nums); i++ {
        arr2 := make([]int, k)
        for j := 0; j < len(arr1); j++ {
            arr2[(nums[i] * j) % k] += arr1[j]
        }
        arr2[nums[i]%k]++
        arr1 = arr2
        for j := 0; j < len(arr1); j++ {
            res[j] += int64(arr1[j])
        } 
    }
    return res
}

func resultArray1(nums []int, k int) []int64 {
    res := make([]int64, k)
    dp, next := make([]int, k), make([]int, k)
    for _, v := range nums {
        v %= k
        clear(next)
        next[v]++
        res[v]++
        for i := range dp {
            if dp[i] > 0 {
                next[(i * v) % k] += dp[i]
                res[(i * v) % k] += int64(dp[i])
            }
        }
        dp, next = next, dp
    }
    return res
}

func resultArray2(nums []int, k int) []int64 {
    n := len(nums)
    res, arr1, arr2 := make([]int64, k), make([]int64, k),  make([]int64, k)
    for i := 0; i < n; i++ {
        m := nums[i] % k
        for j := 0; j < k; j++ {
            arr2[j] = 0
        }
        for j := 0; j < k; j++ {
            arr2[(int((int64(j) * int64(m)) % int64(k)))] += arr1[j]
        }
        arr2[m]++
        for j := 0; j < k; j++ {
            res[j] += arr2[j]
        }
        for j := 0; j < k; j++ {
            arr1[j] = arr2[j]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], k = 3
    // Output: [9,2,4]
    // Explanation:
    // For x = 0, the possible operations include all possible ways to remove non-overlapping prefix/suffix that do not remove nums[2] == 3.
    // For x = 1, the possible operations are:
    //     Remove the empty prefix and the suffix [2, 3, 4, 5]. nums becomes [1].
    //     Remove the prefix [1, 2, 3] and the suffix [5]. nums becomes [4].
    // For x = 2, the possible operations are:
    //     Remove the empty prefix and the suffix [3, 4, 5]. nums becomes [1, 2].
    //     Remove the prefix [1] and the suffix [3, 4, 5]. nums becomes [2].
    //     Remove the prefix [1, 2, 3] and the empty suffix. nums becomes [4, 5].
    //     Remove the prefix [1, 2, 3, 4] and the empty suffix. nums becomes [5].
    fmt.Println(resultArray([]int{1,2,3,4,5}, 3)) // [9,2,4]
    // Example 2:
    // Input: nums = [1,2,4,8,16,32], k = 4
    // Output: [18,1,2,0]
    // Explanation:
    // For x = 0, the only operations that do not result in x = 0 are:
    //     Remove the empty prefix and the suffix [4, 8, 16, 32]. nums becomes [1, 2].
    //     Remove the empty prefix and the suffix [2, 4, 8, 16, 32]. nums becomes [1].
    //     Remove the prefix [1] and the suffix [4, 8, 16, 32]. nums becomes [2].
    // For x = 1, the only possible operation is:
    //     Remove the empty prefix and the suffix [2, 4, 8, 16, 32]. nums becomes [1].
    // For x = 2, the possible operations are:
    //     Remove the empty prefix and the suffix [4, 8, 16, 32]. nums becomes [1, 2].
    //     Remove the prefix [1] and the suffix [4, 8, 16, 32]. nums becomes [2].
    // For x = 3, there is no possible way to perform the operation.
    fmt.Println(resultArray([]int{1,2,4,8,16,32}, 4)) // [18,1,2,0]
    // Example 3:
    // Input: nums = [1,1,2,1,1], k = 2
    // Output: [9,6]
    fmt.Println(resultArray([]int{1,1,2,1,1}, 2)) // [9,6]

    fmt.Println(resultArray([]int{1,2,3,4,5,6,7,8,9}, 2)) // [40 5]
    fmt.Println(resultArray([]int{9,8,7,6,5,4,3,2,1}, 2)) // [40 5]

    fmt.Println(resultArray1([]int{1,2,3,4,5}, 3)) // [9,2,4]
    fmt.Println(resultArray1([]int{1,2,4,8,16,32}, 4)) // [18,1,2,0]
    fmt.Println(resultArray1([]int{1,1,2,1,1}, 2)) // [9,6]
    fmt.Println(resultArray1([]int{1,2,3,4,5,6,7,8,9}, 2)) // [40 5]
    fmt.Println(resultArray1([]int{9,8,7,6,5,4,3,2,1}, 2)) // [40 5]

    fmt.Println(resultArray2([]int{1,2,3,4,5}, 3)) // [9,2,4]
    fmt.Println(resultArray2([]int{1,2,4,8,16,32}, 4)) // [18,1,2,0]
    fmt.Println(resultArray2([]int{1,1,2,1,1}, 2)) // [9,6]
    fmt.Println(resultArray2([]int{1,2,3,4,5,6,7,8,9}, 2)) // [40 5]
    fmt.Println(resultArray2([]int{9,8,7,6,5,4,3,2,1}, 2)) // [40 5]
}