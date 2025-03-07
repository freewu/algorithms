package main

// 2234. Maximum Total Beauty of the Gardens
// Alice is a caretaker of n gardens and she wants to plant flowers to maximize the total beauty of all her gardens.

// You are given a 0-indexed integer array flowers of size n, where flowers[i] is the number of flowers already planted in the ith garden. 
// Flowers that are already planted cannot be removed. 
// You are then given another integer newFlowers, which is the maximum number of flowers that Alice can additionally plant. 
// You are also given the integers target, full, and partial.

// A garden is considered complete if it has at least target flowers. 
// The total beauty of the gardens is then determined as the sum of the following:
//     1. The number of complete gardens multiplied by full.
//     2. The minimum number of flowers in any of the incomplete gardens multiplied by partial. 
//        If there are no incomplete gardens, then this value will be 0.

// Return the maximum total beauty that Alice can obtain after planting at most newFlowers flowers.

// Example 1:
// Input: flowers = [1,3,1,1], newFlowers = 7, target = 6, full = 12, partial = 1
// Output: 14
// Explanation: Alice can plant
// - 2 flowers in the 0th garden
// - 3 flowers in the 1st garden
// - 1 flower in the 2nd garden
// - 1 flower in the 3rd garden
// The gardens will then be [3,6,2,2]. She planted a total of 2 + 3 + 1 + 1 = 7 flowers.
// There is 1 garden that is complete.
// The minimum number of flowers in the incomplete gardens is 2.
// Thus, the total beauty is 1 * 12 + 2 * 1 = 12 + 2 = 14.
// No other way of planting flowers can obtain a total beauty higher than 14.

// Example 2:
// Input: flowers = [2,4,5,3], newFlowers = 10, target = 5, full = 2, partial = 6
// Output: 30
// Explanation: Alice can plant
// - 3 flowers in the 0th garden
// - 0 flowers in the 1st garden
// - 0 flowers in the 2nd garden
// - 2 flowers in the 3rd garden
// The gardens will then be [5,4,5,5]. She planted a total of 3 + 0 + 0 + 2 = 5 flowers.
// There are 3 gardens that are complete.
// The minimum number of flowers in the incomplete gardens is 4.
// Thus, the total beauty is 3 * 2 + 4 * 6 = 6 + 24 = 30.
// No other way of planting flowers can obtain a total beauty higher than 30.
// Note that Alice could make all the gardens complete but in this case, she would obtain a lower total beauty.

// Constraints:
//     1 <= flowers.length <= 10^5
//     1 <= flowers[i], target <= 10^5
//     1 <= newFlowers <= 10^10
//     1 <= full, partial <= 10^5

import "fmt"
import "sort"

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
    sort.Ints(flowers)
    res, n := 0, len(flowers)
    prefix := make([]int, n + 1)
    for i, v := range flowers {
        prefix[i + 1] = prefix[i] + v
    }
    index := n - sort.SearchInts(flowers, target)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := index; i <= n; i++ {
        if i > 0 {
            newFlowers -= int64(max(target - flowers[n - i], 0))
        }
        if newFlowers < 0 { break }
        l, r := 0, n - i - 1
        for l < r {
            mid := (l + r + 1) >> 1
            if int64(flowers[mid] * (mid + 1) - prefix[mid + 1]) <= newFlowers {
                l = mid
            } else {
                r = mid - 1
            }
        }
        y := 0
        if r != -1 {
            cost := flowers[l] * (l + 1) - prefix[l + 1]
            y = min(flowers[l] + int((newFlowers - int64(cost)) / int64(l + 1)), target - 1)
        }
        res = max(res, i * full + y * partial)
    }
    return int64(res)
}

func maximumBeauty1(flowers []int, newFlowers int64, target, full, partial int) int64 {
    sort.Ints(flowers)
    res, n := 0, len(flowers)
    if flowers[0] >= target { // 剪枝，此时所有花园都是完善的
        return int64(n * full)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    leftFlowers := int(newFlowers) - target * n // 填充后缀后，剩余可以种植的花
    for i, f := range flowers {
        flowers[i] = min(f, target) // 去掉多余的花
        leftFlowers += flowers[i]   // 补上已有的花
    }
    for i, x, sumFlowers := 0, 0, 0; i <= n; i++ { // 枚举后缀长度 n-i
        if leftFlowers >= 0 {
            for ; x < i && flowers[x]*x-sumFlowers <= leftFlowers; x++ { // 计算最长前缀的长度
                sumFlowers += flowers[x] // 注意 x 只增不减，二重循环的时间复杂度为 O(n)
            }
            beauty := (n - i) * full // 计算总美丽值
            if x > 0 {
                beauty += min((leftFlowers+sumFlowers) / x, target - 1) * partial
            }
            res = max(res, beauty)
        }
        if i < n {
            leftFlowers += target - flowers[i]
        }
    }
    return int64(res)
}

func maximumBeauty2(flowers []int, newFlowers int64, target, full, partial int) int64 {
    n := len(flowers)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 如果全部种满，还剩下多少朵花？
    leftFlowers := int(newFlowers) - target*n // 先减掉
    for i, flower := range flowers {
        flowers[i] = min(flower, target)
        leftFlowers += flowers[i] // 把已有的加回来
    }
    // 没有种花，所有花园都已种满
    if leftFlowers == int(newFlowers) {
        return int64(n * full) // 答案只能是 n*full（注意不能减少花的数量）
    }
    // 可以全部种满
    if leftFlowers >= 0 {
        // 两种策略取最大值：留一个花园种 target-1 朵花，其余种满；或者，全部种满
        return int64(max((target - 1) * partial + (n - 1) * full, n * full))
    }
    sort.Ints(flowers) // 时间复杂度的瓶颈在这，尽量写在后面
    res, prefix, j := 0, 0, 0
    // 枚举 i，表示后缀 [i, n-1] 种满（i=0 的情况上面已讨论）
    for i := 1; i <= n; i++ {
        // 撤销，flowers[i-1] 不变成 target
        leftFlowers += target - flowers[i-1]
        if leftFlowers < 0 { // 花不能为负数，需要继续撤销
            continue
        }
        // 满足以下条件说明 [0, j] 都可以种 flowers[j] 朵花
        for j < i && flowers[j]*j <= prefix + leftFlowers {
            prefix += flowers[j]
            j++
        }
        // 计算总美丽值
        // 在前缀 [0, j-1] 中均匀种花，这样最小值最大
        avg := (leftFlowers + prefix) / j // 由于上面特判了，这里 avg 一定小于 target
        res = max(res, avg * partial + (n - i) * full)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: flowers = [1,3,1,1], newFlowers = 7, target = 6, full = 12, partial = 1
    // Output: 14
    // Explanation: Alice can plant
    // - 2 flowers in the 0th garden
    // - 3 flowers in the 1st garden
    // - 1 flower in the 2nd garden
    // - 1 flower in the 3rd garden
    // The gardens will then be [3,6,2,2]. She planted a total of 2 + 3 + 1 + 1 = 7 flowers.
    // There is 1 garden that is complete.
    // The minimum number of flowers in the incomplete gardens is 2.
    // Thus, the total beauty is 1 * 12 + 2 * 1 = 12 + 2 = 14.
    // No other way of planting flowers can obtain a total beauty higher than 14.
    fmt.Println(maximumBeauty([]int{1,3,1,1}, 7, 6, 12, 1)) // 14
    // Example 2:
    // Input: flowers = [2,4,5,3], newFlowers = 10, target = 5, full = 2, partial = 6
    // Output: 30
    // Explanation: Alice can plant
    // - 3 flowers in the 0th garden
    // - 0 flowers in the 1st garden
    // - 0 flowers in the 2nd garden
    // - 2 flowers in the 3rd garden
    // The gardens will then be [5,4,5,5]. She planted a total of 3 + 0 + 0 + 2 = 5 flowers.
    // There are 3 gardens that are complete.
    // The minimum number of flowers in the incomplete gardens is 4.
    // Thus, the total beauty is 3 * 2 + 4 * 6 = 6 + 24 = 30.
    // No other way of planting flowers can obtain a total beauty higher than 30.
    // Note that Alice could make all the gardens complete but in this case, she would obtain a lower total beauty.
    fmt.Println(maximumBeauty([]int{2,4,5,3}, 10, 5, 2, 6)) // 30

    fmt.Println(maximumBeauty1([]int{1,3,1,1}, 7, 6, 12, 1)) // 14
    fmt.Println(maximumBeauty1([]int{2,4,5,3}, 10, 5, 2, 6)) // 30

    fmt.Println(maximumBeauty2([]int{1,3,1,1}, 7, 6, 12, 1)) // 14
    fmt.Println(maximumBeauty2([]int{2,4,5,3}, 10, 5, 2, 6)) // 30
}