package main

// 4007. Widest Possible Fence
// You are given an integer array planks, where planks[i] represents the height of the ith wooden plank. Each plank has a width of 1 unit.

// You want to build a fence consisting of planks that all have the same height.

// You may either use a plank as is, or combine exactly two distinct original planks into a single plank whose height equals the sum of their heights. 
// Each original plank can be used at most once, and not all original planks need to be used.

// Return the maximum possible width of the fence that can be built.

// Example 1:
// Input: planks = [1,3,2,5,7,5,4,2,1]
// Output: 4
// Explanation:
// We can have four planks of height 5.
// planks[3] = 5
// planks[5] = 5
// planks[0] + planks[6] = 1 + 4 = 5
// planks[1] + planks[2] = 3 + 2 = 5
// Hence, the maximum width is 4.

// Example 2:
// Input: planks = [2,3,7]
// Output: 1
// Explanation:
// It is impossible to form two planks of the same height, even after combining two distinct original planks.
// Since not all original planks need to be used, we can choose any one plank as the fence.
// Therefore, the maximum possible width is 1.

// Constraints:
//     1 <= planks.length <= 1000
//     1 <= planks[i] <= 10^9

import "fmt"

func maximumWidth(planks []int) int {
    // 统计 planks 的元素出现次数
    res, count := 0, map[int]int{}
    for _, v := range planks {
        count[v]++
    }
    // 枚举所有高度对 (x,y)
    countPair := map[int]int{}
    for k, c := range count {
        countPair[k] += c // 方便后面统计
        countPair[k*2] += c / 2 // 高为 x 的木板内部配对
        for y, c2 := range count {
            if y > k { // 避免 x+y 和 y+x 重复统计
                countPair[k+y] += min(c, c2)
            }
        }
    }
    // 枚举最终木板高度
    for _, v := range countPair {
        res = max(res, v)
    }
    return res
}

func maximumWidth1(planks []int) int {
    res, count := 0, make(map[int]int)
    for _, v := range planks {
        count[v]++
    }
    values := make([]int, 0, len(planks))
    for v := range count {
        values = append(values, v)
    }
    m := len(values)
    pairs := make(map[int]int)
    for i := 0; i < m; i++ {
        for j := i; j < m; j++ {
            sum := values[i] + values[j]
            if i == j {
                pairs[sum] += count[values[i]] / 2
            } else {
                a, b := count[values[i]], count[values[j]]
                if a < b {
                    pairs[sum] += a
                } else {
                    pairs[sum] += b
                }
            }
        }
    }
    for h, v := range count {
        if v + pairs[h] > res {
            res = v + pairs[h]
        }    
    }
    for _, v := range pairs {
        res = max(res, v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: planks = [1,3,2,5,7,5,4,2,1]
    // Output: 4
    // Explanation:
    // We can have four planks of height 5.
    // planks[3] = 5
    // planks[5] = 5
    // planks[0] + planks[6] = 1 + 4 = 5
    // planks[1] + planks[2] = 3 + 2 = 5
    // Hence, the maximum width is 4.
    fmt.Println(maximumWidth([]int{1,3,2,5,7,5,4,2,1})) // 4
    // Example 2:
    // Input: planks = [2,3,7]
    // Output: 1
    // Explanation:
    // It is impossible to form two planks of the same height, even after combining two distinct original planks.
    // Since not all original planks need to be used, we can choose any one plank as the fence.
    // Therefore, the maximum possible width is 1.
    fmt.Println(maximumWidth([]int{2,3,7})) // 1

    fmt.Println(maximumWidth([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(maximumWidth([]int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(maximumWidth1([]int{1,3,2,5,7,5,4,2,1})) // 4
    fmt.Println(maximumWidth1([]int{2,3,7})) // 1
    fmt.Println(maximumWidth1([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(maximumWidth1([]int{9,8,7,6,5,4,3,2,1})) // 5
}