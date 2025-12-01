package main

// 3757. Number of Effective Subsequences
// You are given an integer array nums.

// The strength of the array is defined as the bitwise OR of all its elements.

// A subsequence is considered effective if removing that subsequence strictly decreases the strength of the remaining elements.

// Return the number of effective subsequences in nums. Since the answer may be large, return it modulo 10^9 + 7.

// The bitwise OR of an empty array is 0.

// Example 1:
// Input: nums = [1,2,3]
// Output: 3
// Explanation:
// The Bitwise OR of the array is 1 OR 2 OR 3 = 3.
// Subsequences that are effective are:
// [1, 3]: The remaining element [2] has a Bitwise OR of 2.
// [2, 3]: The remaining element [1] has a Bitwise OR of 1.
// [1, 2, 3]: The remaining elements [] have a Bitwise OR of 0.
// Thus, the total number of effective subsequences is 3.

// Example 2:
// Input: nums = [7,4,6]
// Output: 4
// Explanation:​​​​​​​
// The Bitwise OR of the array is 7 OR 4 OR 6 = 7.
// Subsequences that are effective are:
// [7]: The remaining elements [4, 6] have a Bitwise OR of 6.
// [7, 4]: The remaining element [6] has a Bitwise OR of 6.
// [7, 6]: The remaining element [4] has a Bitwise OR of 4.
// [7, 4, 6]: The remaining elements [] have a Bitwise OR of 0.
// Thus, the total number of effective subsequences is 4.

// Example 3:
// Input: nums = [8,8]
// Output: 1
// Explanation:
// The Bitwise OR of the array is 8 OR 8 = 8.
// Only the subsequence [8, 8] is effective since removing it leaves [] which has a Bitwise OR of 0.
// Thus, the total number of effective subsequences is 1.

// Example 4:
// Input: nums = [2,2,1]
// Output: 5
// Explanation:
// The Bitwise OR of the array is 2 OR 2 OR 1 = 3.
// Subsequences that are effective are:
// [1]: The remaining elements [2, 2] have a Bitwise OR of 2.
// [2, 1] (using nums[0], nums[2]): The remaining element [2] has a Bitwise OR of 2.
// [2, 1] (using nums[1], nums[2]): The remaining element [2] has a Bitwise OR of 2.
// [2, 2]: The remaining element [1] has a Bitwise OR of 1.
// [2, 2, 1]: The remaining elements [] have a Bitwise OR of 0.
// Thus, the total number of effective subsequences is 5.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"
import "math/bits"

const MOD = 1_000_000_007
const MX = 100_001

var pow2 = [MX]int{1}

func init() {
    for i := 1; i < MX; i++ {
        pow2[i] = pow2[i-1] * 2 % MOD
    }
}

func countEffective(nums []int) int {
    or, same := 0, true
    for _, x := range nums {
        or |= x
        if x != nums[0] {
            same = false
        }
    }
    if same {
        return 1
    }
    w := bits.Len(uint(or))
    f := make([]int, 1<<w)
    for _, x := range nums {
        f[x]++
    }
    for i := range w {
        if or>>i&1 == 0 {
            continue
        }
        for s := 0; s < 1<<w; s++ {
            s |= 1 << i
            f[s] += f[s^1<<i]
        }
    }
    res := pow2[len(nums)]
    for sub, ok := or, true; ok; ok = sub != or {
        sign := 1 - bits.OnesCount(uint(or^sub))%2*2
        res -= sign * pow2[f[sub]]
        sub = (sub - 1) & or
    }
    return (res % MOD + MOD) % MOD
}

func countEffective1(nums []int) int {
    const MOD = 1_000_000_007
    curr, n := 0, len(nums)
    for _, v := range nums {
        curr |= v
    }
    bitVals := []int{}
    for b := 0; b < 31; b++ {
        if (curr >> b) & 1 == 1 {
            bitVals = append(bitVals, b)
        }
    }
    k := len(bitVals)
    blkVals := make([]int, 1<<k)
    for _, v := range nums {
        val := 0
        for i, b := range bitVals {
            if (v>>b)&1 == 1 {
                val |= 1 << i
            }
        }
        blkVals[val]++
    }
    for i := 0; i < k; i++ {
        for v := 0; v < (1 << k); v++ {
            if (v & (1 << i)) != 0 {
                blkVals[v] += blkVals[v^(1<<i)]
            }
        }
    }
    pwVals := make([]int64, n+1)
    pwVals[0] = 1
    for i := 1; i <= n; i++ {
        pwVals[i] = (pwVals[i-1] << 1) % MOD
    }
    res, mx := int64(0), (1 << k) - 1
    for val := mx; val > 0; val = (val - 1) & mx {
        pwVal := pwVals[blkVals[mx^val]]
        if bits.OnesCount(uint(val))&1 == 1 {
            res = (res + pwVal) % MOD
        } else {
            res = (res - pwVal + MOD) % MOD
        }
    }
    return int(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 3
    // Explanation:
    // The Bitwise OR of the array is 1 OR 2 OR 3 = 3.
    // Subsequences that are effective are:
    // [1, 3]: The remaining element [2] has a Bitwise OR of 2.
    // [2, 3]: The remaining element [1] has a Bitwise OR of 1.
    // [1, 2, 3]: The remaining elements [] have a Bitwise OR of 0.
    // Thus, the total number of effective subsequences is 3.
    fmt.Println(countEffective([]int{1,2,3})) // 3
    // Example 2:
    // Input: nums = [7,4,6]
    // Output: 4
    // Explanation:​​​​​​​
    // The Bitwise OR of the array is 7 OR 4 OR 6 = 7.
    // Subsequences that are effective are:
    // [7]: The remaining elements [4, 6] have a Bitwise OR of 6.
    // [7, 4]: The remaining element [6] has a Bitwise OR of 6.
    // [7, 6]: The remaining element [4] has a Bitwise OR of 4.
    // [7, 4, 6]: The remaining elements [] have a Bitwise OR of 0.
    // Thus, the total number of effective subsequences is 4.
    fmt.Println(countEffective([]int{7,4,6})) // 4
    // Example 3:
    // Input: nums = [8,8]
    // Output: 1
    // Explanation:
    // The Bitwise OR of the array is 8 OR 8 = 8.
    // Only the subsequence [8, 8] is effective since removing it leaves [] which has a Bitwise OR of 0.
    // Thus, the total number of effective subsequences is 1.
    fmt.Println(countEffective([]int{8,8})) // 1    
    // Example 4:
    // Input: nums = [2,2,1]
    // Output: 5
    // Explanation:
    // The Bitwise OR of the array is 2 OR 2 OR 1 = 3.
    // Subsequences that are effective are:
    // [1]: The remaining elements [2, 2] have a Bitwise OR of 2.
    // [2, 1] (using nums[0], nums[2]): The remaining element [2] has a Bitwise OR of 2.
    // [2, 1] (using nums[1], nums[2]): The remaining element [2] has a Bitwise OR of 2.
    // [2, 2]: The remaining element [1] has a Bitwise OR of 1.
    // [2, 2, 1]: The remaining elements [] have a Bitwise OR of 0.
    // Thus, the total number of effective subsequences is 5.
    fmt.Println(countEffective([]int{2,2,1})) // 5

    fmt.Println(countEffective([]int{1,2,3,4,5,6,7,8,9})) // 175
    fmt.Println(countEffective([]int{9,8,7,6,5,4,3,2,1})) // 175

    fmt.Println(countEffective1([]int{1,2,3})) // 3
    fmt.Println(countEffective1([]int{7,4,6})) // 4
    fmt.Println(countEffective1([]int{8,8})) // 1    
    fmt.Println(countEffective1([]int{2,2,1})) // 5
    fmt.Println(countEffective1([]int{1,2,3,4,5,6,7,8,9})) // 175
    fmt.Println(countEffective1([]int{9,8,7,6,5,4,3,2,1})) // 175
}