package main

// 2367. Number of Arithmetic Triplets
// You are given a 0-indexed, strictly increasing integer array nums and a positive integer diff. 
// A triplet (i, j, k) is an arithmetic triplet if the following conditions are met:
//     1. i < j < k,
//     2. nums[j] - nums[i] == diff, and
//     3. nums[k] - nums[j] == diff.

// Return the number of unique arithmetic triplets.

// Example 1:
// Input: nums = [0,1,4,6,7,10], diff = 3
// Output: 2
// Explanation:
// (1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
// (2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3. 

// Example 2:
// Input: nums = [4,5,6,7,8,9], diff = 2
// Output: 2
// Explanation:
// (0, 2, 4) is an arithmetic triplet because both 8 - 6 == 2 and 6 - 4 == 2.
// (1, 3, 5) is an arithmetic triplet because both 9 - 7 == 2 and 7 - 5 == 2.

// Constraints:
//     3 <= nums.length <= 200
//     0 <= nums[i] <= 200
//     1 <= diff <= 50
//     nums is strictly increasing.

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {
    res := 0
    for i, mp := 0, make(map[int]bool, 200); i < len(nums); i++ {
        if mp[nums[i]] = true; nums[i] >= diff * 2 && mp[nums[i] - diff] && mp[nums[i] - diff * 2] {
            res++
        }
    }
    return res
}

func arithmeticTriplets1(nums []int, diff int) int {
    mp := make(map[int]bool, len(nums))
    for _, v := range nums {
        mp[v] = true
    }
    res := 0
    for _, v := range nums {
        if mp[v + diff] && mp[v + 2 * diff] {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,4,6,7,10], diff = 3
    // Output: 2
    // Explanation:
    // (1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
    // (2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3. 
    fmt.Println(arithmeticTriplets([]int{0,1,4,6,7,10}, 3))
    // Example 2:
    // Input: nums = [4,5,6,7,8,9], diff = 2
    // Output: 2
    // Explanation:
    // (0, 2, 4) is an arithmetic triplet because both 8 - 6 == 2 and 6 - 4 == 2.
    // (1, 3, 5) is an arithmetic triplet because both 9 - 7 == 2 and 7 - 5 == 2.
    fmt.Println(arithmeticTriplets([]int{4,5,6,7,8,9}, 2)) // 2
    fmt.Println(arithmeticTriplets1([]int{0,1,4,6,7,10}, 3)) // 2

    fmt.Println(arithmeticTriplets1([]int{0,1,4,6,7,10}, 3))
    fmt.Println(arithmeticTriplets1([]int{4,5,6,7,8,9}, 2)) // 2

    fmt.Println(arithmeticTriplets1([]int{0,1,4,6,7,10}, 3)) // 2
}