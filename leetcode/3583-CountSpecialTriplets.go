package main

// 3583. Count Special Triplets
// You are given an integer array nums.

// A special triplet is defined as a triplet of indices (i, j, k) such that:
//     1. 0 <= i < j < k < n, where n = nums.length
//     2. nums[i] == nums[j] * 2
//     3. nums[k] == nums[j] * 2

// Return the total number of special triplets in the array.

// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [6,3,6]
// Output: 1
// Explanation:
// The only special triplet is (i, j, k) = (0, 1, 2), where:
// nums[0] = 6, nums[1] = 3, nums[2] = 6
// nums[0] = nums[1] * 2 = 3 * 2 = 6
// nums[2] = nums[1] * 2 = 3 * 2 = 6

// Example 2:
// Input: nums = [0,1,0,0]
// Output: 1
// Explanation:
// The only special triplet is (i, j, k) = (0, 2, 3), where:
// nums[0] = 0, nums[2] = 0, nums[3] = 0
// nums[0] = nums[2] * 2 = 0 * 2 = 0
// nums[3] = nums[2] * 2 = 0 * 2 = 0

// Example 3:
// Input: nums = [8,4,2,8,4]
// Output: 2
// Explanation:
// There are exactly two special triplets:
// (i, j, k) = (0, 1, 3)
// nums[0] = 8, nums[1] = 4, nums[3] = 8
// nums[0] = nums[1] * 2 = 4 * 2 = 8
// nums[3] = nums[1] * 2 = 4 * 2 = 8
// (i, j, k) = (1, 2, 4)
// nums[1] = 4, nums[2] = 2, nums[4] = 4
// nums[1] = nums[2] * 2 = 2 * 2 = 4
// nums[4] = nums[2] * 2 = 2 * 2 = 4
 
// Constraints:
//     3 <= n == nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"

func specialTriplets(nums []int) int {
    res, left, right := 0, make(map[int]int), make(map[int]int)
    left[nums[0]]++
    for i := 2; i < len(nums); i++ {
        right[nums[i]]++
    }
    for i := 1; i < len(nums)-1; i++ {
        res += left[nums[i]*2] * right[nums[i]*2]
        left[nums[i]]++
        right[nums[i+1]]--
    }
    return res % 1_000_000_007
}

func specialTriplets1(nums []int) int {
    res, n := 0, len(nums)
    l, r, mp := make([]int, n),make([]int, n), make(map[int]int)
    for i, v := range nums {
        l[i] = mp[v * 2]
        mp[v]++
    }
    mp = make(map[int]int)
    for i := n - 1; i >= 0;i-- {
        v := nums[i]
        r[i] = mp[v * 2]
        mp[v]++
    }
    for i := 1; i < n - 1;i++ {
        res += l[i] * r[i]
        res %= 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [6,3,6]
    // Output: 1
    // Explanation:
    // The only special triplet is (i, j, k) = (0, 1, 2), where:
    // nums[0] = 6, nums[1] = 3, nums[2] = 6
    // nums[0] = nums[1] * 2 = 3 * 2 = 6
    // nums[2] = nums[1] * 2 = 3 * 2 = 6
    fmt.Println(specialTriplets([]int{6,3,6})) // 1
    // Example 2:
    // Input: nums = [0,1,0,0]
    // Output: 1
    // Explanation:
    // The only special triplet is (i, j, k) = (0, 2, 3), where:
    // nums[0] = 0, nums[2] = 0, nums[3] = 0
    // nums[0] = nums[2] * 2 = 0 * 2 = 0
    // nums[3] = nums[2] * 2 = 0 * 2 = 0
    fmt.Println(specialTriplets([]int{0,1,0,0})) // 1
    // Example 3:
    // Input: nums = [8,4,2,8,4]
    // Output: 2
    // Explanation:
    // There are exactly two special triplets:
    // (i, j, k) = (0, 1, 3)
    // nums[0] = 8, nums[1] = 4, nums[3] = 8
    // nums[0] = nums[1] * 2 = 4 * 2 = 8
    // nums[3] = nums[1] * 2 = 4 * 2 = 8
    // (i, j, k) = (1, 2, 4)
    // nums[1] = 4, nums[2] = 2, nums[4] = 4
    // nums[1] = nums[2] * 2 = 2 * 2 = 4
    // nums[4] = nums[2] * 2 = 2 * 2 = 4
    fmt.Println(specialTriplets([]int{8,4,2,8,4})) // 2
    fmt.Println(specialTriplets([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(specialTriplets([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(specialTriplets1([]int{6,3,6})) // 1
    fmt.Println(specialTriplets1([]int{0,1,0,0})) // 1
    fmt.Println(specialTriplets1([]int{8,4,2,8,4})) // 2
    fmt.Println(specialTriplets1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(specialTriplets1([]int{9,8,7,6,5,4,3,2,1})) // 0
}