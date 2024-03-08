package main

// 2834. Find the Minimum Possible Sum of a Beautiful Array
// You are given positive integers n and target.
// An array nums is beautiful if it meets the following conditions:
//         nums.length == n.
//         nums consists of pairwise distinct positive integers.
//         There doesn't exist two distinct indices, i and j, in the range [0, n - 1], such that nums[i] + nums[j] == target.

// Return the minimum possible sum that a beautiful array could have modulo 10^9 + 7.

// Example 1:
// Input: n = 2, target = 3
// Output: 4
// Explanation: We can see that nums = [1,3] is beautiful.
// - The array nums has length n = 2.
// - The array nums consists of pairwise distinct positive integers.
// - There doesn't exist two distinct indices, i and j, with nums[i] + nums[j] == 3.
// It can be proven that 4 is the minimum possible sum that a beautiful array could have.

// Example 2:
// Input: n = 3, target = 3
// Output: 8
// Explanation: We can see that nums = [1,3,4] is beautiful.
// - The array nums has length n = 3.
// - The array nums consists of pairwise distinct positive integers.
// - There doesn't exist two distinct indices, i and j, with nums[i] + nums[j] == 3.
// It can be proven that 8 is the minimum possible sum that a beautiful array could have.

// Example 3:
// Input: n = 1, target = 1
// Output: 1
// Explanation: We can see, that nums = [1] is beautiful.
 
// Constraints:
//         1 <= n <= 10^9
//         1 <= target <= 10^9

import "fmt"

// 给你两个正整数：n 和 target
// 如果数组 nums 满足下述条件，则称其为 美丽数组 。
//         nums.length == n.
//         nums 由两两互不相同的正整数组成。
//         在范围 [0, n-1] 内，不存在 两个 不同 下标 i 和 j ，使得 nums[i] + nums[j] == target 。

func minimumPossibleSum(n int, target int) int {
    min := func (a, b int) int {
        if a <= b {
            return a
        }
        return b
    }
    var m = min(target/2, n)
	return (m * (m + 1) + (target * 2 + n - m - 1) * (n - m)) / 2 % (1e9 + 7)
}

// 构造一个大小为 n 的正整数数组，该数组由不同的数字组成，并且没有任意两个数字的和等于 target，
// 在满足这样的前提下，要保证数组的和最小。
// 为了让数组之和最小，我们按照 1,2,3,⋯  的顺序考虑，但添加了 x 之后，就不能添加 target − x，因此最大可以添加到 ⌊target / 2⌋，
// 如果个数还不够 n 个，就继续从 target, target+1, target + 2,⋯ 依次添加。由于添加的数字是连续的，所以可以用等差数列求和公式快速求解。
// 令 m = ⌊ target / 2 ⌋，然后可以分情况求解：
// 若 n ≤ m 最小数组和 为 ((1 + n) * n / 2)
// 否则 n>mn \gt mn>m，最小数组和为 (((1 + m) * m / 2) + (((target + target + (n - m) - 1) * (n - m) / 2)))
func minimumPossibleSum1(n int, target int) int {
    const mod = 1000000007
    m := target / 2
    if n <= m {
        return ((1 + n) * n / 2) % mod
    }
    return (((1 + m) * m / 2) + (((target + target + (n - m) - 1) * (n - m) / 2))) % mod
}

func main() {
    // We can see that nums = [1,3] is beautiful.
    // - The array nums has length n = 2.
    // - The array nums consists of pairwise distinct positive integers.
    // - There doesn't exist two distinct indices, i and j, with nums[i] + nums[j] == 3.
    // It can be proven that 4 is the minimum possible sum that a beautiful array could have.
    fmt.Println(minimumPossibleSum(2,3)) // 4
    // nums = [1,3,4] is beautiful
    fmt.Println(minimumPossibleSum(3,3)) // 8
    fmt.Println(minimumPossibleSum(1,1)) // 1

    fmt.Println(minimumPossibleSum1(2,3)) // 4
    fmt.Println(minimumPossibleSum1(3,3)) // 8
    fmt.Println(minimumPossibleSum1(1,1)) // 1
}