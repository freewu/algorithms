package main

// 3647. Maximum Weight in Two Bags
// You are given an integer array weights and two integers w1 and w2 representing the maximum capacities of two bags.

// Each item may be placed in at most one bag such that:
//     1. Bag 1 holds at most w1 total weight.
//     2. Bag 2 holds at most w2 total weight.

// Return the maximum total weight that can be packed into the two bags.

// Example 1:
// Input: weights = [1,4,3,2], w1 = 5, w2 = 4
// Output: 9
// Explanation:
// Bag 1: Place weights[2] = 3 and weights[3] = 2 as 3 + 2 = 5 <= w1
// Bag 2: Place weights[1] = 4 as 4 <= w2
// Total weight: 5 + 4 = 9

// Example 2:
// Input: weights = [3,6,4,8], w1 = 9, w2 = 7
// Output: 15
// Explanation:
// Bag 1: Place weights[3] = 8 as 8 <= w1
// Bag 2: Place weights[0] = 3 and weights[2] = 4 as 3 + 4 = 7 <= w2
// Total weight: 8 + 7 = 15

// Example 3:
// Input: weights = [5,7], w1 = 2, w2 = 3
// Output: 0
// Explanation:
// No weight fits in either bag, thus the answer is 0.

// Constraints:
//     1 <= weights.length <= 100
//     1 <= weights[i] <= 100
//     1 <= w1, w2 <= 300

import "fmt"

func maxWeight(weights []int, w1 int, w2 int) int {
    set := make(map[int]struct{}) // 使用map存储可能的重量组合，键为"100*x + y"形式的字符串
    set[w1 * 1000 + w2] = struct{}{} // 初始状态
    for _, w := range weights {
        // 遍历当前所有状态的副本，避免修改迭代中的map
        curr := make([]int, 0, len(set))
        for k := range set {
            curr = append(curr, k)
        }
        for _, v := range curr {
            x, y := v / 1000, v % 1000
            // 尝试将当前重量放入第一个容器
            if x >= w {
                newX, newY := x - w, y
                set[newX * 1000 + newY] = struct{}{}
            }
            // 尝试将当前重量放入第二个容器
            if y >= w {
                newX, newY := x, y - w
                set[newX * 1000 + newY] = struct{}{}
            }
        }
    }
    // 找到剩余总重量最小的状态，初始值设为一个较大的数
    minSum := w1 + w2 + 1
    for v := range set {
        x, y := v / 1000, v % 1000
        sum := x + y
        if sum < minSum {
            minSum = sum
        }
    }
    // 最大装载重量 = 初始总容量 - 最小剩余总容量
    return w1 + w2 - minSum
}

func main() {
    // Example 1:
    // Input: weights = [1,4,3,2], w1 = 5, w2 = 4
    // Output: 9
    // Explanation:
    // Bag 1: Place weights[2] = 3 and weights[3] = 2 as 3 + 2 = 5 <= w1
    // Bag 2: Place weights[1] = 4 as 4 <= w2
    // Total weight: 5 + 4 = 9
    fmt.Println(maxWeight([]int{1,4,3,2}, 5, 4)) // 9
    // Example 2:
    // Input: weights = [3,6,4,8], w1 = 9, w2 = 7
    // Output: 15
    // Explanation:
    // Bag 1: Place weights[3] = 8 as 8 <= w1
    // Bag 2: Place weights[0] = 3 and weights[2] = 4 as 3 + 4 = 7 <= w2
    // Total weight: 8 + 7 = 15
    fmt.Println(maxWeight([]int{3,6,4,8}, 9, 7)) // 15
    // Example 3:
    // Input: weights = [5,7], w1 = 2, w2 = 3
    // Output: 0
    // Explanation:
    // No weight fits in either bag, thus the answer is 0.
    fmt.Println(maxWeight([]int{5,7}, 2, 3)) // 0

    fmt.Println(maxWeight([]int{1,2,3,4,5,6,7,8,9}, 2, 3)) // 5
    fmt.Println(maxWeight([]int{9,8,7,6,5,4,3,2,1}, 2, 3)) // 5
    fmt.Println(maxWeight([]int{17,3,7,3,4,12,15,25,30,6,10,9,7,13,7,10,6,5,17,8,6,18,15,3,19,7,10,25,18,6}, 23, 102)) // 125
}
