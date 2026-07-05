package main

// 3984. Divisible Game
// You are given an integer array nums of length n.

// Alice and Bob are playing a game. Alice chooses:
//     1. An integer k such that k > 1.
//     2. Two integers l and r such that 0 <= l <= r < n.

// Initially, both Alice's and Bob's scores are 0.

// For each index i in the range [l, r] (inclusive):
//     1. If nums[i] is divisible by k, Alice's score increases by nums[i].
//     2. Otherwise, Bob's score increases by nums[i].

// The score difference is Alice's score minus Bob's score.

// Alice wants to maximize the score difference. 
// If there are multiple values of k that achieve the maximum score difference, she chooses the smallest such k.

// Return the product of the maximum score difference and the chosen value of k. 
// Since the result can be large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,4,6,8]
// Output: 36
// Explanation:
// Alice can choose k = 2, l = 1, and r = 3.
// All values in nums[1..3] are divisible by 2, so Alice's score is 4 + 6 + 8 = 18, while Bob's score is 0.
// The score difference is 18, which is the maximum possible. Among all values of k that achieve this score difference, the smallest is 2.
// Therefore, the answer is 18 * 2 = 36.

// Example 2:
// Input: nums = [2,1,2]
// Output: 6
// Explanation:
// Alice can choose k = 2, l = 0, and r = 2.
// The values nums[0] and nums[2] are divisible by 2, so Alice's score is 2 + 2 = 4. The value nums[1] is not divisible by 2, so Bob's score is 1.
// The score difference is 4 - 1 = 3, which is the maximum possible. Among all values of k that achieve this score difference, the smallest is 2.
// Therefore, the answer is 3 * 2 = 6.

// Example 3:
// Input: nums = [1]
// Output: 1000000005
// Explanation:
// Alice must choose some k > 1. The smallest possible choice is k = 2.
// Since nums[0] is not divisible by 2, Alice's score is 0, while Bob's score is 1.
// The score difference is -1, which is the maximum possible.
// Therefore, the answer is -1 * 2 = -2. Modulo 109 + 7, this equals 1000000005.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^6

import "fmt"
import "slices"

const MX = 1_000_001
var primeDivisors [MX][]int32

// 预处理每个数的质因子
func init() {
    for i := int32(2); i < MX; i++ {
        if primeDivisors[i] == nil { // i 是质数
            for j := i; j < MX; j += i { // 枚举 i 的倍数 j 
                primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
            }
        }
    }
}

func divisibleGame(nums []int) int {
    const MOD = 1_000_000_007
    // 收集所有质因子
    allPrimeDivisors := []int32{}
    for _, x := range nums {
        allPrimeDivisors = append(allPrimeDivisors, primeDivisors[x]...)
    }
    if len(allPrimeDivisors) == 0 {
        // 每个数都是 1
        // 最优是只选一个 1（分数差为 -1），最小 k 为 2
        return MOD - 2
    }
    maxSubArray := func(nums []int, k int) int { // 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
        res, f := -1 << 61, 0
        for _, v := range nums {
            if v % k != 0 {
                v = -v
            }
            f = max(f, 0) + v
            res = max(res, f)
        }
        return res
    }
    // 排序去重
    slices.Sort(allPrimeDivisors)
    allPrimeDivisors = slices.Compact(allPrimeDivisors)
    maxDiff, bestK := -1 << 61, 0
    // 枚举质因子作为 k，计算最大子数组和
    for _, d := range allPrimeDivisors {
        k := int(d)
        diff := maxSubArray(nums, k)
        if diff > maxDiff {
            maxDiff, bestK = diff, k
        }
    }
    // 保证结果非负
    return (maxDiff * bestK % MOD + MOD) % MOD
}

func divisibleGame1(nums []int) int {
    n, mod := len(nums), int64(1_000_000_007)
    prefix := make([]int64, n + 1)
    for i, v := range nums{
        prefix[i + 1] = prefix[i] + int64(v)
    }
    pid, mem := make(map[int][]int), make(map[int][]int)
    for i, v := range nums {
        if v == 1 { continue }
        if _, ok := mem[v]; !ok {
            var fact []int
            temp := v
            for j := 2; j * j <= temp; j++ {
                if temp % j == 0{
                    fact = append(fact, j)
                    for temp % j == 0{
                        temp /= j
                    }
                }
            }
            if temp > 1 {
                fact = append(fact, temp)
            }
            mem[v] = fact
        }
        for _, p := range mem[v] {
            pid[p] = append(pid[p], i)
        }
    }
    if len(pid) == 0 {
        a := nums[0]
        for _, v := range nums {
            a = min(a, v)
        }
        res := (-int64(a) * 2 + mod) % mod
        return int(res)
    }
    b, c := int64(-1e10), -1
    for i, v := range pid {
        curr, temp := int64(0), int64(-1e18)
        for j, id := range v {
            val := int64(nums[id])
            if j == 0{
                curr = val
            } else{
                prev := v[j - 1]
                gap := prefix[id] - prefix[prev + 1]
                d := curr - gap + val
                curr = max(val, d)
            }
            temp = max(temp, curr)
        }
        if temp > b {
            b = temp
            c = i
        } else if temp == b && i < c {
            c = i
        }
    }
    res := (b % mod) * int64(c) % mod
    if res < 0 {
        res = (res + mod) % mod
    }
    return int(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,4,6,8]
    // Output: 36
    // Explanation:
    // Alice can choose k = 2, l = 1, and r = 3.
    // All values in nums[1..3] are divisible by 2, so Alice's score is 4 + 6 + 8 = 18, while Bob's score is 0.
    // The score difference is 18, which is the maximum possible. Among all values of k that achieve this score difference, the smallest is 2.
    // Therefore, the answer is 18 * 2 = 36.
    fmt.Println(divisibleGame([]int{1,4,6,8})) // 36
    // Example 2:
    // Input: nums = [2,1,2]
    // Output: 6
    // Explanation:
    // Alice can choose k = 2, l = 0, and r = 2.
    // The values nums[0] and nums[2] are divisible by 2, so Alice's score is 2 + 2 = 4. The value nums[1] is not divisible by 2, so Bob's score is 1.
    // The score difference is 4 - 1 = 3, which is the maximum possible. Among all values of k that achieve this score difference, the smallest is 2.
    // Therefore, the answer is 3 * 2 = 6.
    fmt.Println(divisibleGame([]int{2,1,2})) // 6    
    // Example 3:
    // Input: nums = [1]
    // Output: 1000000005
    // Explanation:
    // Alice must choose some k > 1. The smallest possible choice is k = 2.
    // Since nums[0] is not divisible by 2, Alice's score is 0, while Bob's score is 1.
    // The score difference is -1, which is the maximum possible.
    // Therefore, the answer is -1 * 2 = -2. Modulo 109 + 7, this equals 1000000005.
    fmt.Println(divisibleGame([]int{1})) // 1000000005

    fmt.Println(divisibleGame([]int{1,2,3,4,5,6,7,8,9})) // 27
    fmt.Println(divisibleGame([]int{9,8,7,6,5,4,3,2,1})) // 27

    fmt.Println(divisibleGame1([]int{1,4,6,8})) // 36
    fmt.Println(divisibleGame1([]int{2,1,2})) // 6    
    fmt.Println(divisibleGame1([]int{1})) // 1000000005
    fmt.Println(divisibleGame1([]int{1,2,3,4,5,6,7,8,9})) // 27
    fmt.Println(divisibleGame1([]int{9,8,7,6,5,4,3,2,1})) // 27
}