package main

// 2541. Minimum Operations to Make Array Equal II
// You are given two integer arrays nums1 and nums2 of equal length n and an integer k. 
// You can perform the following operation on nums1:
//     1. Choose two indexes i and j and increment nums1[i] by k and decrement nums1[j] by k. 
//        In other words, nums1[i] = nums1[i] + k and nums1[j] = nums1[j] - k.

// nums1 is said to be equal to nums2 if for all indices i such that 0 <= i < n, nums1[i] == nums2[i].

// Return the minimum number of operations required to make nums1 equal to nums2. 
// If it is impossible to make them equal, return -1.

// Example 1:
// Input: nums1 = [4,3,1,4], nums2 = [1,3,7,1], k = 3
// Output: 2
// Explanation: In 2 operations, we can transform nums1 to nums2.
// 1st operation: i = 2, j = 0. After applying the operation, nums1 = [1,3,4,4].
// 2nd operation: i = 2, j = 3. After applying the operation, nums1 = [1,3,7,1].
// One can prove that it is impossible to make arrays equal in fewer operations.

// Example 2:
// Input: nums1 = [3,8,5,2], nums2 = [2,4,1,6], k = 1
// Output: -1
// Explanation: It can be proved that it is impossible to make the two arrays equal.

// Constraints:
//     n == nums1.length == nums2.length
//     2 <= n <= 10^5
//     0 <= nums1[i], nums2[j] <= 10^9
//     0 <= k <= 10^5

import "fmt"

func minOperations(nums1 []int, nums2 []int, k int) int64 {
    if k == 0 {
        for i := range nums1 {
            if nums1[i] != nums2[i] { return -1 }
        }
        return 0
    }
    // Compute k-diff, store it in nums1
    for i := range nums1 {
        diff := nums2[i] - nums1[i]
        if diff % k != 0 { return -1 }
        nums1[i] = diff / k
    }
    positivesSum, negativesSum := 0, 0
    for _, v := range nums1 {
        if v > 0 {
            positivesSum += v
        } else if v < 0 {
            negativesSum += v
        }
    }
    if positivesSum != -negativesSum { return -1 }
    return int64(positivesSum)
}

func minOperations1(nums1 []int, nums2 []int, k int) int64 {
    positive, negative := 0, 0
    for i := 0; i < len(nums1); i++ {
        diff := nums1[i] - nums2[i]
        if diff != 0 && (k == 0 || diff % k != 0) { return -1 }
        if diff > 0 {
            positive += diff
        } else if diff < 0 {
            negative += diff
        }
    }
    if positive == 0 && negative == 0 { return 0 }
    res := positive / k
    if res != 0 && -negative / k == res { return int64(res) }
    return -1
}

func main() {
    // Example 1:
    // Input: nums1 = [4,3,1,4], nums2 = [1,3,7,1], k = 3
    // Output: 2
    // Explanation: In 2 operations, we can transform nums1 to nums2.
    // 1st operation: i = 2, j = 0. After applying the operation, nums1 = [1,3,4,4].
    // 2nd operation: i = 2, j = 3. After applying the operation, nums1 = [1,3,7,1].
    // One can prove that it is impossible to make arrays equal in fewer operations.
    fmt.Println(minOperations([]int{4,3,1,4}, []int{1,3,7,1}, 3)) // 2
    // Example 2:
    // Input: nums1 = [3,8,5,2], nums2 = [2,4,1,6], k = 1
    // Output: -1
    // Explanation: It can be proved that it is impossible to make the two arrays equal.
    fmt.Println(minOperations([]int{3,8,5,2}, []int{2,4,1,6}, 1)) // -1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 1)) // 20
    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 1)) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 1)) // 20
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 1)) // 0

    fmt.Println(minOperations1([]int{4,3,1,4}, []int{1,3,7,1}, 3)) // 2
    fmt.Println(minOperations1([]int{3,8,5,2}, []int{2,4,1,6}, 1)) // -1
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 1)) // 20
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 1)) // 0
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 1)) // 20
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 1)) // 0
}