package main

// 2275. Largest Combination With Bitwise AND Greater Than Zero
// The bitwise AND of an array nums is the bitwise AND of all integers in nums.
//     For example, for nums = [1, 5, 3], the bitwise AND is equal to 1 & 5 & 3 = 1.
//     Also, for nums = [7], the bitwise AND is 7.

// You are given an array of positive integers candidates. 
// Evaluate the bitwise AND of every combination of numbers of candidates. 
// Each number in candidates may only be used once in each combination.

// Return the size of the largest combination of candidates with a bitwise AND greater than 0.

// Example 1:
// Input: candidates = [16,17,71,62,12,24,14]
// Output: 4
// Explanation: The combination [16,17,62,24] has a bitwise AND of 16 & 17 & 62 & 24 = 16 > 0.
// The size of the combination is 4.
// It can be shown that no combination with a size greater than 4 has a bitwise AND greater than 0.
// Note that more than one combination may have the largest size.
// For example, the combination [62,12,24,14] has a bitwise AND of 62 & 12 & 24 & 14 = 8 > 0.

// Example 2:
// Input: candidates = [8,8]
// Output: 2
// Explanation: The largest combination [8,8] has a bitwise AND of 8 & 8 = 8 > 0.
// The size of the combination is 2, so we return 2.

// Constraints:
//     1 <= candidates.length <= 10^5
//     1 <= candidates[i] <= 10^7

import "fmt"
import "slices"
import "math/bits"

func largestCombination(candidates []int) int {
    res := 0
    // 如果组合中所有数字的在同一个bit位都是1, 那么 这组数的 & > 0,
    // 可以按bit 位统计.
    // int32 (根据题目中的数字范围)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < 32; i++ {
        count := 0
        for _, v := range candidates {
            if v >> i & 1 == 1 {
                count++
            }
        }
        res = max(res, count)
    }
    return res
}

func largestCombination1(candidates []int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < 30; i++ {
        t := 0
        for _, v := range candidates {
            t += v >> i & 1
        }
        res = max(res, t)
    }
    return res
}

func largestCombination2(candidates []int) int {
    count := [24]int{}
    for _, v := range candidates {
        for i := 0; v > 0; i++ {
            count[i] += v & 1
            v >>= 1
        }
    }
    return slices.Max(count[:])
}

func largestCombination3(candidates []int) int {
    mx := slices.Max(candidates)
    res, n := 0, bits.Len(uint(mx))
    for i := 0; i < n; i++ {
        count := 0
        for _, v := range candidates {
            count += (v >> i) & 1
        }
        res = max(res, count)
    }
    return res
}

func main() {
    // Example 1:
    // Input: candidates = [16,17,71,62,12,24,14]
    // Output: 4
    // Explanation: The combination [16,17,62,24] has a bitwise AND of 16 & 17 & 62 & 24 = 16 > 0.
    // The size of the combination is 4.
    // It can be shown that no combination with a size greater than 4 has a bitwise AND greater than 0.
    // Note that more than one combination may have the largest size.
    // For example, the combination [62,12,24,14] has a bitwise AND of 62 & 12 & 24 & 14 = 8 > 0.
    fmt.Println(largestCombination([]int{16,17,71,62,12,24,14})) // 4
    // Example 2:
    // Input: candidates = [8,8]
    // Output: 2
    // Explanation: The largest combination [8,8] has a bitwise AND of 8 & 8 = 8 > 0.
    // The size of the combination is 2, so we return 2.s
    fmt.Println(largestCombination([]int{8,8})) // 2
    fmt.Println(largestCombination([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(largestCombination([]int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(largestCombination1([]int{16,17,71,62,12,24,14})) // 4
    fmt.Println(largestCombination1([]int{8,8})) // 2
    fmt.Println(largestCombination1([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(largestCombination1([]int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(largestCombination2([]int{16,17,71,62,12,24,14})) // 4
    fmt.Println(largestCombination2([]int{8,8})) // 2
    fmt.Println(largestCombination2([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(largestCombination2([]int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(largestCombination3([]int{16,17,71,62,12,24,14})) // 4
    fmt.Println(largestCombination3([]int{8,8})) // 2
    fmt.Println(largestCombination3([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(largestCombination3([]int{9,8,7,6,5,4,3,2,1})) // 5
}