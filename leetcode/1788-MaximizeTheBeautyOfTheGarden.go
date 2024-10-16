package main

// 1788. Maximize the Beauty of the Garden
// There is a garden of n flowers, and each flower has an integer beauty value. 
// The flowers are arranged in a line. You are given an integer array flowers of size n and each flowers[i] represents the beauty of the ith flower.

// A garden is valid if it meets these conditions:
//     The garden has at least two flowers.
//     The first and the last flower of the garden have the same beauty value.

// As the appointed gardener, you have the ability to remove any (possibly none) flowers from the garden. 
// You want to remove flowers in a way that makes the remaining garden valid. 
// The beauty of the garden is the sum of the beauty of all the remaining flowers.

// Return the maximum possible beauty of some valid garden after you have removed any (possibly none) flowers.

// Example 1:
// Input: flowers = [1,2,3,1,2]
// Output: 8
// Explanation: You can produce the valid garden [2,3,1,2] to have a total beauty of 2 + 3 + 1 + 2 = 8.

// Example 2:
// Input: flowers = [100,1,1,-3,1]
// Output: 3
// Explanation: You can produce the valid garden [1,1,1] to have a total beauty of 1 + 1 + 1 = 3.

// Example 3:
// Input: flowers = [-1,-2,0,-1]
// Output: -2
// Explanation: You can produce the valid garden [-1,-1] to have a total beauty of -1 + -1 = -2.

// Constraints:
//     2 <= flowers.length <= 10^5
//     -10^4 <= flowers[i] <= 10^4
//     It is possible to create a valid garden by removing some (possibly none) flowers.

import "fmt"

func maximumBeauty(flowers []int) int {
    // 找到两个坐标，i和j，使得flowers[i] == flowers[j]并且他们之间的和最大
    // 最大子序列
    // 这两个数之间的负数也可以去除
    // 两个相等的数加上他们之间的正数和最大
    // 选择两个最远的数，然后求他们之间的正数的和即可
    // 两个数之间的和可以用前缀和
    sa := make([]int, len(flowers))
    for i, f := range flowers {
        if i == 0 {
            if f > 0 {
                sa[0] = f
            }
        } else {
            if f > 0 {
                sa[i] = sa[i - 1] + f
            } else {
                sa[i] = sa[i - 1]
            }
        }
    }
    mp := make(map[int][]int)
    for i, f := range flowers {
        if _, e := mp[f]; !e {
            mp[f] = []int{len(flowers), -1}
        }
        mp[f][0] = min(mp[f][0], i)
        mp[f][1] = max(mp[f][1], i)
    }
    res := -1 << 31
    for k, arr := range mp {
        if arr[0] == arr[1] { continue }
        sum := sa[arr[1]]
        if arr[0] > 0 {
            sum -= sa[arr[0] - 1]
        }
        if k < 0 {
            sum += 2 * k
        }
        res = max(res, sum)
    }
    return res
}

func main() {
    // Example 1:
    // Input: flowers = [1,2,3,1,2]
    // Output: 8
    // Explanation: You can produce the valid garden [2,3,1,2] to have a total beauty of 2 + 3 + 1 + 2 = 8.
    fmt.Println(maximumBeauty([]int{1,2,3,1,2})) // 8
    // Example 2:
    // Input: flowers = [100,1,1,-3,1]
    // Output: 3
    // Explanation: You can produce the valid garden [1,1,1] to have a total beauty of 1 + 1 + 1 = 3.
    fmt.Println(maximumBeauty([]int{100,1,1,-3,1})) // 3
    // Example 3:
    // Input: flowers = [-1,-2,0,-1]
    // Output: -2
    // Explanation: You can produce the valid garden [-1,-1] to have a total beauty of -1 + -1 = -2.
    fmt.Println(maximumBeauty([]int{-1,-2,0,-1})) // -2
}