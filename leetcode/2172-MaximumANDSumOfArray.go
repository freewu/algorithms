package main

// 2172. Maximum AND Sum of Array
// You are given an integer array nums of length n and an integer numSlots such that 2 * numSlots >= n. 
// There are numSlots slots numbered from 1 to numSlots.

// You have to place all n integers into the slots such that each slot contains at most two numbers. 
// The AND sum of a given placement is the sum of the bitwise AND of every number with its respective slot number.
//     For example, the AND sum of placing the numbers [1, 3] into slot 1 
//     and [4, 6] into slot 2 is equal to (1 AND 1) + (3 AND 1) + (4 AND 2) + (6 AND 2) = 1 + 1 + 0 + 2 = 4.

// Return the maximum possible AND sum of nums given numSlots slots.

// Example 1:
// Input: nums = [1,2,3,4,5,6], numSlots = 3
// Output: 9
// Explanation: One possible placement is [1, 4] into slot 1, [2, 6] into slot 2, and [3, 5] into slot 3. 
// This gives the maximum AND sum of (1 AND 1) + (4 AND 1) + (2 AND 2) + (6 AND 2) + (3 AND 3) + (5 AND 3) = 1 + 0 + 2 + 2 + 3 + 1 = 9.

// Example 2:
// Input: nums = [1,3,10,4,7,1], numSlots = 9
// Output: 24
// Explanation: One possible placement is [1, 1] into slot 1, [3] into slot 3, [4] into slot 4, [7] into slot 7, and [10] into slot 9.
// This gives the maximum AND sum of (1 AND 1) + (1 AND 1) + (3 AND 3) + (4 AND 4) + (7 AND 7) + (10 AND 9) = 1 + 1 + 3 + 4 + 7 + 8 = 24.
// Note that slots 2, 5, 6, and 8 are empty which is permitted.

// Constraints:
//     n == nums.length
//     1 <= numSlots <= 9
//     1 <= n <= 2 * numSlots
//     1 <= nums[i] <= 15

import "fmt"
import "math/bits"

func maximumANDSum(nums []int, numSlots int) int {
    pow := func (b int, p int) int { // 求幂
        res := 1
        for i := 0; i < p; i++ {
            res *= b
        }
        return res
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    mask := pow(3, numSlots) - 1
    memo := make([]int, mask + 1)
    var dfs func(index, mask int) int
    dfs = func(index, mask int) int {
        if memo[mask] > 0 { return memo[mask] }
        if index < 0 { return 0 }
        for slot, bit := 1, 1; slot <= numSlots; slot, bit = slot + 1, bit * 3 {
            if mask / bit % 3 > 0 {
                memo[mask] = max(memo[mask], (nums[index] & slot) + dfs(index - 1, mask - bit))
            }
        }
        return memo[mask]
    }
    return dfs(len(nums) - 1, mask)
}

func maximumANDSum1(nums []int, numSlots int) int {
    res, n := 0, len(nums)
    facts := make([]int, 1 << (numSlots << 1))
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range facts {
        c := bits.OnesCount(uint(i))
        if c >= n { continue }
        for j := 0; j < numSlots * 2; j++ {
            if (i >> j) & 1 == 0 {
                s := i | (1 << j)
                facts[s] = max(facts[s], v + (j / 2 + 1) & nums[c])
                res = max(res, facts[s])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,6], numSlots = 3
    // Output: 9
    // Explanation: One possible placement is [1, 4] into slot 1, [2, 6] into slot 2, and [3, 5] into slot 3. 
    // This gives the maximum AND sum of (1 AND 1) + (4 AND 1) + (2 AND 2) + (6 AND 2) + (3 AND 3) + (5 AND 3) = 1 + 0 + 2 + 2 + 3 + 1 = 9.
    fmt.Println(maximumANDSum([]int{1,2,3,4,5,6}, 3)) // 9
    // Example 2:
    // Input: nums = [1,3,10,4,7,1], numSlots = 9
    // Output: 24
    // Explanation: One possible placement is [1, 1] into slot 1, [3] into slot 3, [4] into slot 4, [7] into slot 7, and [10] into slot 9.
    // This gives the maximum AND sum of (1 AND 1) + (1 AND 1) + (3 AND 3) + (4 AND 4) + (7 AND 7) + (10 AND 9) = 1 + 1 + 3 + 4 + 7 + 8 = 24.
    // Note that slots 2, 5, 6, and 8 are empty which is permitted.
    fmt.Println(maximumANDSum([]int{1,3,10,4,7,1}, 9)) // 24

    fmt.Println(maximumANDSum1([]int{1,2,3,4,5,6}, 3)) // 9
    fmt.Println(maximumANDSum1([]int{1,3,10,4,7,1}, 9)) // 24
}