package main

// 3811. Number of Alternating XOR Partitions
// You are given an integer array nums and two distinct integers target1 and target2.

// A partition of nums splits it into one or more contiguous, non-empty blocks that cover the entire array without overlap.

// A partition is valid if the bitwise XOR of elements in its blocks alternates between target1 and target2, starting with target1.

// Formally, for blocks b1, b2, …:
//     1. XOR(b1) = target1
//     2. XOR(b2) = target2 (if it exists)
//     3. XOR(b3) = target1, and so on.

// Return the number of valid partitions of nums, modulo 10^9 + 7.

// Note: A single block is valid if its XOR equals target1.

// Example 1:
// Input: nums = [2,3,1,4], target1 = 1, target2 = 5
// Output: 1
// Explanation:​​​​​​​
// The XOR of [2, 3] is 1, which matches target1.
// The XOR of the remaining block [1, 4] is 5, which matches target2.
// This is the only valid alternating partition, so the answer is 1.

// Example 2:
// Input: nums = [1,0,0], target1 = 1, target2 = 0
// Output: 3
// Explanation:
// ​​​​​​​The XOR of [1, 0, 0] is 1, which matches target1.
// The XOR of [1] and [0, 0] are 1 and 0, matching target1 and target2.
// The XOR of [1, 0] and [0] are 1 and 0, matching target1 and target2.
// Thus, the answer is 3.​​​​​​​

// Example 3:
// Input: nums = [7], target1 = 1, target2 = 7
// Output: 0
// Explanation:
// The XOR of [7] is 7, which does not match target1, so no valid partition exists.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i], target1, target2 <= 10^5
//     target1 != target2

import "fmt"

func alternatingXOR(nums []int, target1, target2 int) int {
    prefix, mod := 0, 1_000_000_007
    f1, f2 := map[int]int{}, map[int]int{0: 1}
    for i, v := range nums {
        prefix ^= v
        last1 := f2[prefix ^ target1]
        last2 := f1[prefix ^ target2]
        if i == len(nums) - 1 {
            return (last1 + last2) % mod
        }
        f1[prefix] = (f1[prefix] + last1) % mod
        f2[prefix] = (f2[prefix] + last2) % mod
    }
    return -1
}

func alternatingXOR1(nums []int, target1 int, target2 int) int {
    a, b, n, mod := 0, 0, len(nums), 1_000_000_007
    xor := make([]int, n)
    c1, c2, c3, c4 := 0, 0, 0, 1
    for i := range xor {
        xor[i] = nums[i]
        if i > 0 {
            xor[i] ^= xor[i-1]
        }
        a, b = 0, 0
        // 当前能以 a 结尾
        if xor[i] == target1 || xor[i] == target2 {
            if xor[i] == target1 {
                a = c4
            } else {
                a = c2
            }
        }
        // 以 b 结尾的数量
        if xor[i] == target1^target2 || xor[i] == 0 {
            if xor[i] == target1^target2 {
                b = c1
            } else {
                b = c3
            }
        }
        if xor[i] == target1 {
            c1 += a
        }
        if xor[i] == target2 {
            c3 += a
        }
        if xor[i] == target1^target2 {
            c2 += b
        }
        if xor[i] == 0 {
            c4 += b
        }
        c1 %= mod
        c2 %= mod
        c3 %= mod
        c4 %= mod
    }
    return a + b
}

func main() {
    // Example 1:
    // Input: nums = [2,3,1,4], target1 = 1, target2 = 5
    // Output: 1
    // Explanation:​​​​​​​
    // The XOR of [2, 3] is 1, which matches target1.
    // The XOR of the remaining block [1, 4] is 5, which matches target2.
    // This is the only valid alternating partition, so the answer is 1.
    fmt.Println(alternatingXOR([]int{2,3,1,4}, 1, 5)) // 1
    // Example 2:
    // Input: nums = [1,0,0], target1 = 1, target2 = 0
    // Output: 3
    // Explanation:
    // ​​​​​​​The XOR of [1, 0, 0] is 1, which matches target1.
    // The XOR of [1] and [0, 0] are 1 and 0, matching target1 and target2.
    // The XOR of [1, 0] and [0] are 1 and 0, matching target1 and target2.
    // Thus, the answer is 3.​​​​​​​
    fmt.Println(alternatingXOR([]int{1,0,0}, 1, 0)) // 3
    // Example 3:
    // Input: nums = [7], target1 = 1, target2 = 7
    // Output: 0
    // Explanation:
    // The XOR of [7] is 7, which does not match target1, so no valid partition exists.
    fmt.Println(alternatingXOR([]int{7}, 1, 7)) // 0

    fmt.Println(alternatingXOR([]int{1,2,3,4,5,6,7,8,9}, 1, 5)) // 1
    fmt.Println(alternatingXOR([]int{9,8,7,6,5,4,3,2,1}, 1, 5)) // 1

    fmt.Println(alternatingXOR1([]int{2,3,1,4}, 1, 5)) // 1
    fmt.Println(alternatingXOR1([]int{1,0,0}, 1, 0)) // 3
    fmt.Println(alternatingXOR1([]int{7}, 1, 7)) // 0
    fmt.Println(alternatingXOR1([]int{1,2,3,4,5,6,7,8,9}, 1, 5)) // 1
    fmt.Println(alternatingXOR1([]int{9,8,7,6,5,4,3,2,1}, 1, 5)) // 1
}