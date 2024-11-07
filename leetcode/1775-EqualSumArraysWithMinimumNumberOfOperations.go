package main

// 1775. Equal Sum Arrays With Minimum Number of Operations
// You are given two arrays of integers nums1 and nums2, possibly of different lengths. 
// The values in the arrays are between 1 and 6, inclusive.

// In one operation, you can change any integer's value in any of the arrays to any value between 1 and 6, inclusive.

// Return the minimum number of operations required to make the sum of values in nums1 equal to the sum of values in nums2. 
// Return -1​​​​​ if it is not possible to make the sum of the two arrays equal.

// Example 1:
// Input: nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
// Output: 3
// Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed.
// - Change nums2[0] to 6. nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2].
// - Change nums1[5] to 1. nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2].
// - Change nums1[2] to 2. nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2].

// Example 2:
// Input: nums1 = [1,1,1,1,1,1,1], nums2 = [6]
// Output: -1
// Explanation: There is no way to decrease the sum of nums1 or to increase the sum of nums2 to make them equal.

// Example 3:
// Input: nums1 = [6,6], nums2 = [1]
// Output: 3
// Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed. 
// - Change nums1[0] to 2. nums1 = [2,6], nums2 = [1].
// - Change nums1[1] to 2. nums1 = [2,2], nums2 = [1].
// - Change nums2[0] to 4. nums1 = [2,2], nums2 = [4].

// Constraints:
//     1 <= nums1.length, nums2.length <= 10^5
//     1 <= nums1[i], nums2[i] <= 6

import "fmt"

func minOperations(nums1 []int, nums2 []int) int {
    helper := func (nums []int) (int, int, []int) {
        sum, counts := 0, make([]int, 7)
        for _, v := range nums {
            sum += v
            counts[v]++
        }
        return len(nums), sum, counts
    }
    n1, sum1, counts1 := helper(nums1)
    n2, sum2, counts2 := helper(nums2)
    if sum1 == sum2 { return 0 }
    if n1 > n2 * 6 || n2 > n1 * 6 {  return -1 }
    if sum1 > sum2 {
        sum1, sum2 = sum2, sum1
        counts1, counts2 = counts2, counts1
    }
    res := 0
    for i := 1; i <= 6; i++ {
        dif, ci := 6 - i, counts1[i]
        difs := dif * ci
        if sum1 + difs >= sum2 {
            return res + (sum2 - sum1 - 1) / dif + 1
        } else {
            sum1 += difs
            res += ci
        }
        ci = counts2[7 - i]
        difs = dif * ci
        if sum2 - difs<= sum1 {
            return res + (sum2 - sum1 - 1) / dif + 1
        } else {
            sum2 -= difs
            res += ci
        }
    }
    return -1
}

func minOperations1(nums1 []int, nums2 []int) int {
    count1, count2 := make([]int, 7), make([]int, 7)
    for _, v := range nums1 { count1[v]++ }
    for _, v := range nums2 { count2[v]++ }

    res := 1 << 31
    helper := func(arr []int, target int) int {
        res, sum := 0, 0
        for i, v := range arr {
            sum += i * v
        }
        // i -> 6
        for i := 1; i < 6 && sum < target; i++ {
            ch := 6 - i
            if target-sum < ch * arr[i] {
                res += (target - sum + ch - 1) / ch
                sum = target
            } else {
                res += arr[i]
                sum += ch * arr[i]
            }
        }
        // i -> 1
        for i := 6; i > 1 && sum > target; i-- {
            ch := i - 1
            if sum-target < ch* arr[i] {
                res += (sum - target + ch - 1) / ch
                sum = target
            } else {
                res += arr[i]
                sum -= ch * arr[i]
            }
        }
        if sum != target { return 1 << 31 }
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < 1e6; i++ {
        res = min(res, helper(count1, i) + helper(count2, i))
    }
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
    // Output: 3
    // Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed.
    // - Change nums2[0] to 6. nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2].
    // - Change nums1[5] to 1. nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2].
    // - Change nums1[2] to 2. nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2].
    fmt.Println(minOperations([]int{1,2,3,4,5,6}, []int{1,1,2,2,2,2})) // 3
    // Example 2:
    // Input: nums1 = [1,1,1,1,1,1,1], nums2 = [6]
    // Output: -1
    // Explanation: There is no way to decrease the sum of nums1 or to increase the sum of nums2 to make them equal.
    fmt.Println(minOperations([]int{1,1,1,1,1,1,1}, []int{6})) // -1
    // Example 3:
    // Input: nums1 = [6,6], nums2 = [1]
    // Output: 3
    // Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed. 
    // - Change nums1[0] to 2. nums1 = [2,6], nums2 = [1].
    // - Change nums1[1] to 2. nums1 = [2,2], nums2 = [1].
    // - Change nums2[0] to 4. nums1 = [2,2], nums2 = [4].
    fmt.Println(minOperations([]int{6,6}, []int{1})) // 3

    fmt.Println(minOperations1([]int{1,2,3,4,5,6}, []int{1,1,2,2,2,2})) // 3
    fmt.Println(minOperations1([]int{1,1,1,1,1,1,1}, []int{6})) // -1
    fmt.Println(minOperations1([]int{6,6}, []int{1})) // 3
}