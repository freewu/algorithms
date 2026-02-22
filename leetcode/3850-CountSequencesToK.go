package main

// 3850. Count Sequences to K
// You are given an integer array nums, and an integer k.

// Start with an initial value val = 1 and process nums from left to right. 
// At each index i, you must choose exactly one of the following actions:
//     1. Multiply val by nums[i].
//     2. Divide val by nums[i].
//     3. Leave val unchanged.

// After processing all elements, val is considered equal to k only if its final rational value exactly equals k.

// Return the count of distinct sequences of choices that result in val == k.

// Note: Division is rational (exact), not integer division. For example, 2 / 4 = 1 / 2.

// Example 1:
// Input: nums = [2,3,2], k = 6
// Output: 2
// Explanation:
// The following 2 distinct sequences of choices result in val == k:
// Sequence | Operation on nums[0]         | Operation on nums[1]          | Operation on nums[2]              | Final val 
// 1        | Multiply: val = 1 * 2 = 2    | Multiply: val = 2 * 3 = 6     | Leave val unchanged	            | 6
// 2        | Leave val unchanged          | Multiply: val = 1 * 3 = 3	    | Multiply: val = 3 * 2 = 6	        | 6   

// Example 2:
// Input: nums = [4,6,3], k = 2
// Output: 2
// Explanation:
// The following 2 distinct sequences of choices result in val == k:
// Sequence | Operation on nums[0]	        | Operation on nums[1]          | Operation on nums[2]              | Final val
// 1        | Multiply: val = 1 * 4 = 4    | Divide: val = 4 / 6 = 2 / 3   | Multiply: val = (2 / 3) * 3 = 2   | 2
// 2        | Leave val unchanged	        | Multiply: val = 1 * 6 = 6     | Divide: val = 6 / 3 = 2           | 2   

// Example 3:
// Input: nums = [1,5], k = 1
// Output: 3
// Explanation:
// The following 3 distinct sequences of choices result in val == k:
// Sequence    | Operation on nums[0]	    | Operation on nums[1]	| Final val
// 1           | Multiply: val = 1 * 1 = 1	| Leave val unchanged	| 1
// 2           | Divide: val = 1 / 1 = 1	| Leave val unchanged	| 1
// 3           | Leave val unchanged	    | Leave val unchanged	| 1

// Constraints:
//     1 <= nums.length <= 19
//     1 <= nums[i] <= 6
//     1 <= k <= 10^15

import "fmt"
import "math/bits"

// 最简分数
func countSequences(nums []int, k int64) int {
    type Args struct{ i, p, q int }
    memo := map[Args]int{}
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    var dfs func(int, int, int) int
    dfs = func(i, p, q int) int {
        if i < 0 {
            if p == int(k) && q == 1 {
                return 1
            }
            return 0
        }
        t := Args{i, p, q}
        if res, ok := memo[t]; ok {
            return res
        }
        x := nums[i]
        g := gcd(p, q*x)
        res1 := dfs(i-1, p/g, q*x/g) // 除以 nums[i]
        g = gcd(p*x, q)
        res2 := dfs(i-1, p*x/g, q/g) // 乘以 nums[i]
        res3 := dfs(i-1, p, q)       // 不变
        res := res1 + res2 + res3
        memo[t] = res
        return res
    }
    return dfs(len(nums)-1, 1, 1) // 从 1/1 开始，目标是变成 k/1
}

// 质因子分解
func countSequences1(nums []int, k int64) int {
    primeFactorization := func (k int) ([3]int, bool) {  // 返回 k 中的质因子 2,3,5 的个数，以及 k 是否只包含 <= 5 的质因子
        e2 := bits.TrailingZeros(uint(k))
        k >>= e2
        e3 := 0
        for k%3 == 0 {
            e3++
            k /= 3
        }
        e5 := 0
        for k%5 == 0 {
            e5++
            k /= 5
        }
        return [3]int{e2, e3, e5}, k == 1
    }
    e, ok := primeFactorization(int(k))
    if !ok { return 0 } // k 有大于 5 的质因子
    n := len(nums)
    es := make([][3]int, n)
    for i, x := range nums {
        es[i], _ = primeFactorization(x)
    }
    type Args struct{ i, e2, e3, e5 int }
    memo := map[Args]int{}
    var dfs func(int, int, int, int) int
    dfs = func(i, e2, e3, e5 int) int {
        if i < 0 {
            if e2 == 0 && e3 == 0 && e5 == 0 { // k 变成 1
                return 1
            }
            return 0
        }
        p := Args{i, e2, e3, e5}
        if res, ok := memo[p]; ok {
            return res
        }
        e := es[i]
        res1 := dfs(i-1, e2-e[0], e3-e[1], e5-e[2]) // k 除以 nums[i]
        res2 := dfs(i-1, e2+e[0], e3+e[1], e5+e[2]) // k 乘以 nums[i]
        res3 := dfs(i-1, e2, e3, e5)                // k 不变
        res := res1 + res2 + res3
        memo[p] = res
        return res
    }
    return dfs(n-1, e[0], e[1], e[2]) // 从 k 开始，目标是变成 1
}

func countSequences2(nums []int, k int64) int {
    a, b, c, temp := 0, 0, 0, k
    for temp % 2 == 0 { a++; temp /= 2 }
    for temp % 3 == 0 { b++; temp /= 3 }
    for temp % 5 == 0 { c++; temp /= 5 }
    if temp != 1 { return 0 }
    if a > 38 || b > 19 || c > 19 { return 0 }
    n := len(nums)
    p2 := [7]int{0, 0, 1, 0, 2, 0, 1}
    p3 := [7]int{0, 0, 0, 1, 0, 0, 1}
    p5 := [7]int{0, 0, 0, 0, 0, 1, 0}
    mp2 := make([]int, n+1)
    mp3 := make([]int, n+1)
    mp5 := make([]int, n+1)
    for i := n - 1; i >= 0; i-- {
        v := nums[i]
        mp2[i] = mp2[i+1] + p2[v]
        mp3[i] = mp3[i+1] + p3[v]
        mp5[i] = mp5[i+1] + p5[v]
    }
    memo := make(map[int]int)
    var dfs func(i, c2, c3, c5 int) int
    dfs = func(i, c2, c3, c5 int) int {
        if i == n {
            if c2 == a && c3 == b && c5 == c { return 1 }
            return 0
        }
        if c2+mp2[i] < a || c2-mp2[i] > a { return 0 }
        if c3+mp3[i] < b || c3-mp3[i] > b { return 0 }
        if c5+mp5[i] < c || c5-mp5[i] > c { return 0 }
        key := (i * 130000) + ((c2 + 40) * 1600) + ((c3 + 20) * 40) + (c5 + 20)
        if val, exists := memo[key]; exists {
            return val
        }
        v := nums[i]
        res := 0
        res += dfs(i+1, c2, c3, c5)
        res += dfs(i+1, c2+p2[v], c3+p3[v], c5+p5[v])
        res += dfs(i+1, c2-p2[v], c3-p3[v], c5-p5[v])
        memo[key] = res
        return res
    }
    return dfs(0, 0, 0, 0)
}

func main() {
    // Example 1:
    // Input: nums = [2,3,2], k = 6
    // Output: 2
    // Explanation:
    // The following 2 distinct sequences of choices result in val == k:
    // Sequence | Operation on nums[0]         | Operation on nums[1]          | Operation on nums[2]              | Final val 
    // 1        | Multiply: val = 1 * 2 = 2    | Multiply: val = 2 * 3 = 6     | Leave val unchanged	            | 6
    // 2        | Leave val unchanged          | Multiply: val = 1 * 3 = 3	    | Multiply: val = 3 * 2 = 6	        | 6   
    fmt.Println(countSequences([]int{2,3,2}, 6)) // 2
    // Example 2:
    // Input: nums = [4,6,3], k = 2
    // Output: 2
    // Explanation:
    // The following 2 distinct sequences of choices result in val == k:
    // Sequence | Operation on nums[0]	        | Operation on nums[1]          | Operation on nums[2]              | Final val
    // 1        | Multiply: val = 1 * 4 = 4    | Divide: val = 4 / 6 = 2 / 3   | Multiply: val = (2 / 3) * 3 = 2   | 2
    // 2        | Leave val unchanged	        | Multiply: val = 1 * 6 = 6     | Divide: val = 6 / 3 = 2           | 2   
    fmt.Println(countSequences([]int{4,6,3}, 2)) // 2
    // Example 3:
    // Input: nums = [1,5], k = 1
    // Output: 3
    // Explanation:
    // The following 3 distinct sequences of choices result in val == k:
    // Sequence    | Operation on nums[0]	    | Operation on nums[1]	| Final val
    // 1           | Multiply: val = 1 * 1 = 1	| Leave val unchanged	| 1
    // 2           | Divide: val = 1 / 1 = 1	| Leave val unchanged	| 1
    // 3           | Leave val unchanged	    | Leave val unchanged	| 1 
    fmt.Println(countSequences([]int{1,5}, 1)) // 3

    fmt.Println(countSequences([]int{1,2,3,4,5,6,7,8,9}, 2)) // 45
    fmt.Println(countSequences([]int{9,8,7,6,5,4,3,2,1}, 2)) // 45

    fmt.Println(countSequences1([]int{2,3,2}, 6)) // 2
    fmt.Println(countSequences1([]int{4,6,3}, 2)) // 2
    fmt.Println(countSequences1([]int{1,5}, 1)) // 3
    fmt.Println(countSequences1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 135
    fmt.Println(countSequences1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 135

    fmt.Println(countSequences2([]int{2,3,2}, 6)) // 2
    fmt.Println(countSequences2([]int{4,6,3}, 2)) // 2
    fmt.Println(countSequences2([]int{1,5}, 1)) // 3
}