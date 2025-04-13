package main

// 3509. Maximum Product of Subsequences With an Alternating Sum Equal to K
// You are given an integer array nums and two integers, k and limit. Your task is to find a non-empty subsequence of nums that:
//     1. Has an alternating sum equal to k.
//     2. Maximizes the product of all its numbers without the product exceeding limit.

// Return the product of the numbers in such a subsequence. If no subsequence satisfies the requirements, return -1.

// The alternating sum of a 0-indexed array is defined as the sum of the elements at even indices minus the sum of the elements at odd indices.

// Example 1:
// Input: nums = [1,2,3], k = 2, limit = 10
// Output: 6
// Explanation:
// The subsequences with an alternating sum of 2 are:
// [1, 2, 3]
// Alternating Sum: 1 - 2 + 3 = 2
// Product: 1 * 2 * 3 = 6
// [2]
// Alternating Sum: 2
// Product: 2
// The maximum product within the limit is 6.

// Example 2:
// Input: nums = [0,2,3], k = -5, limit = 12
// Output: -1
// Explanation:
// A subsequence with an alternating sum of exactly -5 does not exist.

// Example 3:
// Input: nums = [2,2,3,3], k = 0, limit = 9
// Output: 9
// Explanation:
// The subsequences with an alternating sum of 0 are:
// [2, 2]
// Alternating Sum: 2 - 2 = 0
// Product: 2 * 2 = 4
// [3, 3]
// Alternating Sum: 3 - 3 = 0
// Product: 3 * 3 = 9
// [2, 2, 3, 3]
// Alternating Sum: 2 - 2 + 3 - 3 = 0
// Product: 2 * 2 * 3 * 3 = 36
// The subsequence [2, 2, 3, 3] has the greatest product with an alternating sum equal to k, but 36 > 9. The next greatest product is 9, which is within the limit.

// Constraints:
//     1 <= nums.length <= 150
//     0 <= nums[i] <= 12
//     -10^5 <= k <= 10^5
//     1 <= limit <= 5000

import "fmt"

func maxProduct(nums []int, k int, limit int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if k > sum || k < -sum {
        return -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    getHash := func(idx int, sign int8, sum, prod int) int {
        hash := sum
        hash *= 150
        hash += idx
        hash *= 5000
        hash += (prod - 1)
        hash *= 2
        hash += max(int(sign), 0)
        return hash % 1_000_000_007
    }
    dp := make(map[int]int)
    var dfs func(k, limit int, idx int, sign int8, sum, prod int) int
    dfs = func(k, limit int, idx int, sign int8, sum, prod int) int {
        if idx == len(nums) {
            if sum == k && prod <= limit {
                return prod
            }
            return -1
        }
        key := getHash(idx, sign, sum, prod)
        if v, ok := dp[key]; ok {
            return v
        }
        dp[key] = -1
        // skip
        dp[key] = max(dp[key], dfs(k, limit, idx+1, sign, sum, min(prod, limit + 1)))
        // take
        if prod == -1 {
            prod = nums[idx]
        } else {
            prod *= nums[idx]
        }
        dp[key] = max(dp[key], dfs(k, limit, idx+1, -sign, sum+int(sign)*nums[idx], min(limit + 1, prod)))
        return dp[key]
    }
    return dfs(k, limit, 0, 1, 0, -1)
}

func maxProduct1(nums []int, k int, limit int) int {
    type Pair struct { parity , prod int}
    type set map[int]struct{}
    dp := map[Pair]set{}
    res, inf := -1, limit + 1
    for _, x := range nums {
        newDp := map[Pair]set{}
        for key, s := range dp {
            newDp[key] = make(set)
            for v := range s {
                newDp[key][v] = struct{}{}
            }
        }
        if x == 0 {
            key := Pair{1, 0}
            if _, ok := newDp[key]; !ok {
                newDp[key] = make(set)
            }
            newDp[key][0] = struct{}{}
        } else {
            prod := x
            if x > limit {
                prod = inf
            }
            key := Pair{ 1, prod} 
            if _, ok := newDp[key]; !ok {
                newDp[key] = make(set)
            }
            newDp[key][x] = struct{}{}
        }
        for key, s := range dp {
            for alt := range s {
                newParity, newAlt := 1 - key.parity, 0
                if key.parity == 0 {
                    newAlt = alt + x
                } else {
                    newAlt = alt - x
                }
                newProd := 0
                if x == 0 {
                    newProd = 0
                } else {
                    if key.prod == inf {
                        newProd = inf
                    } else {
                        candidate := key.prod * x
                        if candidate <= limit {
                            newProd = candidate
                        } else {
                            newProd = inf
                        }
                    }
                }
                newKey := Pair{ newParity, newProd }
                if _, ok := newDp[newKey]; !ok {
                    newDp[newKey] = make(set)
                }
                newDp[newKey][newAlt] = struct{}{}
            }
        }
        dp = newDp
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for key, s := range dp {
        if key.prod != inf {
            if _, ok := s[k]; ok {
                res = max(res, key.prod)
            }
        }
    }
    return res
}

func maxProduct2(nums []int, k, limit int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    if sum < abs(k) { return -1 } // 如果数组和小于 |k|，则返回 -1
    // s -> {m}
    oddS, evenS := map[int]map[int]struct{}{}, map[int]map[int]struct{}{}
    add := func(mp map[int]map[int]struct{}, key, val int) {
        if _, ok := mp[key]; !ok { mp[key] = map[int]struct{}{} }
        mp[key][val] = struct{}{}
    }
    for _, v := range nums {
        // 长为偶数的子序列的计算结果 newEvenS
        newEvenS := map[int]map[int]struct{}{}
        for s, set := range oddS {
            newEvenS[s - v] = map[int]struct{}{}
            for m := range set {
                if m * v <= limit {
                    newEvenS[s - v][m * v] = struct{}{}
                }
            }
        }
        // 长为奇数的子序列的计算结果 oddS
        for s, set := range evenS {
            if _, ok := oddS[s + v]; !ok {
                oddS[s + v] = map[int]struct{}{}
            }
            for m := range set {
                if m * v <= limit {
                    oddS[s + v][m * v] = struct{}{}
                }
            }
            if v == 0 {
                add(oddS, s, 0)
            }
        }
        // 更新 evenS
        for s, set := range newEvenS {
            if eSet, ok := evenS[s]; ok {
                for m := range set {
                    eSet[m] = struct{}{}
                }
            } else {
                evenS[s] = set
            }
            if v == 0 {
                add(evenS, s, 0)
            }
        }
        // 子序列只有一个数的情况
        if v <= limit {
            add(oddS, v, v)
        }
        if set, ok := oddS[k]; ok {
            if _, ok := set[limit]; ok {
                return limit // 提前返回
            }
        }
        if set, ok := evenS[k]; ok {
            if _, ok := set[limit]; ok {
                return limit // 提前返回
            }
        }
    }
    calcMax := func(m map[int]struct{}) int {
        maxVal := -1
        if m != nil {
            for v := range m {
                maxVal = max(maxVal, v)
            }
        }
        return maxVal
    }
    return max(calcMax(oddS[k]), calcMax(evenS[k]))
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], k = 2, limit = 10
    // Output: 6
    // Explanation:
    // The subsequences with an alternating sum of 2 are:
    // [1, 2, 3]
    // Alternating Sum: 1 - 2 + 3 = 2
    // Product: 1 * 2 * 3 = 6
    // [2]
    // Alternating Sum: 2
    // Product: 2
    // The maximum product within the limit is 6.
    fmt.Println(maxProduct([]int{1,2,3}, 2, 10)) // 6
    // Example 2:
    // Input: nums = [0,2,3], k = -5, limit = 12
    // Output: -1
    // Explanation:
    // A subsequence with an alternating sum of exactly -5 does not exist.
    fmt.Println(maxProduct([]int{0,2,3}, -5, 12)) // -1
    // Example 3:
    // Input: nums = [2,2,3,3], k = 0, limit = 9
    // Output: 9
    // Explanation:
    // The subsequences with an alternating sum of 0 are:
    // [2, 2]
    // Alternating Sum: 2 - 2 = 0
    // Product: 2 * 2 = 4
    // [3, 3]
    // Alternating Sum: 3 - 3 = 0
    // Product: 3 * 3 = 9
    // [2, 2, 3, 3]
    // Alternating Sum: 2 - 2 + 3 - 3 = 0
    // Product: 2 * 2 * 3 * 3 = 36
    // The subsequence [2, 2, 3, 3] has the greatest product with an alternating sum equal to k, but 36 > 9. The next greatest product is 9, which is within the limit.
    fmt.Println(maxProduct([]int{2,2,3,3}, 0, 9)) // 9

    fmt.Println(maxProduct([]int{1,2,3,4,5,6,7,8,9}, 2, 10)) // 6
    fmt.Println(maxProduct([]int{9,8,7,6,5,4,3,2,1}, 2, 10)) // 8

    fmt.Println(maxProduct1([]int{1,2,3}, 2, 10)) // 6
    fmt.Println(maxProduct1([]int{0,2,3}, -5, 12)) // -1
    fmt.Println(maxProduct1([]int{2,2,3,3}, 0, 9)) // 9
    fmt.Println(maxProduct1([]int{1,2,3,4,5,6,7,8,9}, 2, 10)) // 6
    fmt.Println(maxProduct1([]int{9,8,7,6,5,4,3,2,1}, 2, 10)) // 8

    fmt.Println(maxProduct2([]int{1,2,3}, 2, 10)) // 6
    fmt.Println(maxProduct2([]int{0,2,3}, -5, 12)) // -1
    fmt.Println(maxProduct2([]int{2,2,3,3}, 0, 9)) // 9
    fmt.Println(maxProduct2([]int{1,2,3,4,5,6,7,8,9}, 2, 10)) // 6
    fmt.Println(maxProduct2([]int{9,8,7,6,5,4,3,2,1}, 2, 10)) // 8
}