package main

// 4059. Lexicographically Largest Power Array
// You are given an integer array nums of length n. 
// You may rearrange its elements to form any permutation perm.

// Define an array power of length 15. 
// For each 0 <= i < 15, power[i] is the largest integer j, where 0 <= j <= n, such that the first j elements of perm all have the (14 - i)th bit set.

// Bit positions are indexed from right to left, starting with the 0th bit.

// Return the lexicographically largest possible power array.

// Example 1:
// Input: nums = [7,5]
// Output: [0,0,0,0,0,0,0,0,0,0,0,0,2,1,2]
// Explanation:
// Choose perm = [7, 5].
// Both elements have bit 2 set, so power[12] = 2.
// The first element has bit 1 set, but the second does not, so power[13] = 1.
// Both elements have bit 0 set, so power[14] = 2.
// All higher bits are unset in the first element, so the remaining entries are 0.

// Example 2:
// Input: nums = [3,1,7]
// Output: [0,0,0,0,0,0,0,0,0,0,0,0,1,2,3]
// Explanation:
// Choose perm = [7, 3, 1].
// The first element has bit 2 set, but the second does not, so power[12] = 1.
// The first two elements have bit 1 set, but the third does not, so power[13] = 2.
// All three elements have bit 0 set, so power[14] = 3.
// All higher bits are unset in the first element, so the remaining entries are 0.

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     0 <= nums[i] < 2^15

import "fmt"
import "slices"
import "math/bits"

func largestPower(nums []int) []int {
    res := make([]int, 15)
    slices.SortFunc(nums, func(a, b int) int { 
        return b - a 
    })
    n, mx := len(nums), bits.Len(uint(nums[0]))
    for i := mx - 1; i >= 0; i-- {
        // 找最长前缀连续 1
        j := 0
        for j < n && nums[j]>>i&1 > 0 {
            j++
        }
        res[14-i] = j
        // [0, j-1] 这一位都是 1，其余元素无关紧要，为方便排序，全置为 0
        for ; j < n; j++ {
            nums[j] &^= 1 << i
        }
        slices.SortFunc(nums, func(a, b int) int { 
            return b - a
        })
    }
    return res
}

func largestPower1(nums []int) []int {
    n := len(nums)
    perm := append([]int{}, nums...)
    bounds := make([]int, 2, 17)
    bounds[0], bounds[1] = 0, n
    res := make([]int, 15)
    for bit := 14; bit >= 0; bit-- {
        prefix := 0
        for i := 0; i + 1 < len(bounds); i++ {
            start, end := bounds[i], bounds[i+1]
            split := start
            for j := start; j < end; j++ {
                if 0 != perm[j]&(1<<bit) {
                    perm[split], perm[j] = perm[j], perm[split]
                    split++
                }
            }
            prefix += split - start
            if split == end {
                continue
            }
            if split > start {
                bounds = append(bounds, 0)
                copy(bounds[i+2:], bounds[i+1:])
                bounds[i+1] = split
            }
            break
        }
        res[14-bit] = prefix
    }
    return res    
}

func main() {
    // Example 1:
    // Input: nums = [7,5]
    // Output: [0,0,0,0,0,0,0,0,0,0,0,0,2,1,2]
    // Explanation:
    // Choose perm = [7, 5].
    // Both elements have bit 2 set, so power[12] = 2.
    // The first element has bit 1 set, but the second does not, so power[13] = 1.
    // Both elements have bit 0 set, so power[14] = 2.
    // All higher bits are unset in the first element, so the remaining entries are 0.
    fmt.Println(largestPower([]int{7,5})) // [0,0,0,0,0,0,0,0,0,0,0,0,2,1,2]
    // Example 2:
    // Input: nums = [3,1,7]
    // Output: [0,0,0,0,0,0,0,0,0,0,0,0,1,2,3]
    // Explanation:
    // Choose perm = [7, 3, 1].
    // The first element has bit 2 set, but the second does not, so power[12] = 1.
    // The first two elements have bit 1 set, but the third does not, so power[13] = 2.
    // All three elements have bit 0 set, so power[14] = 3.
    // All higher bits are unset in the first element, so the remaining entries are 0.
    fmt.Println(largestPower([]int{3,1,7})) // [0,0,0,0,0,0,0,0,0,0,0,0,1,2,3]

    fmt.Println(largestPower([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0 0 0 2 0 0 1]
    fmt.Println(largestPower([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 0 0 0 0 0 0 0 2 0 0 1]

    fmt.Println(largestPower1([]int{7,5})) // [0,0,0,0,0,0,0,0,0,0,0,0,2,1,2]
    fmt.Println(largestPower1([]int{3,1,7})) // [0,0,0,0,0,0,0,0,0,0,0,0,1,2,3]
    fmt.Println(largestPower1([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0 0 0 2 0 0 1]
    fmt.Println(largestPower1([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 0 0 0 0 0 0 0 2 0 0 1]
}