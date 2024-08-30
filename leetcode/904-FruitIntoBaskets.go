package main

// 904. Fruit Into Baskets
// You are visiting a farm that has a single row of fruit trees arranged from left to right. 
// The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

// You want to collect as much fruit as possible.
// However, the owner has some strict rules that you must follow:
//     1. You only have two baskets, and each basket can only hold a single type of fruit. 
//        There is no limit on the amount of fruit each basket can hold.
//     2. Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. 
//        The picked fruits must fit in one of your baskets.
//     3. Once you reach a tree with fruit that cannot fit in your baskets, you must stop.

// Given the integer array fruits, return the maximum number of fruits you can pick.

// Example 1:
// Input: fruits = [1,2,1]
// Output: 3
// Explanation: We can pick from all 3 trees.

// Example 2:
// Input: fruits = [0,1,2,2]
// Output: 3
// Explanation: We can pick from trees [1,2,2].
// If we had started at the first tree, we would only pick from trees [0,1].

// Example 3:
// Input: fruits = [1,2,3,2,2]
// Output: 4
// Explanation: We can pick from trees [2,3,2,2].
// If we had started at the first tree, we would only pick from trees [1,2].

// Constraints:
//     1 <= fruits.length <= 10^5
//     0 <= fruits[i] < fruits.length

import "fmt"

func totalFruit(fruits []int) int {
    res, i, unique, n := 0, 0, 0, len(fruits)
    mp := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := 0; j < n; j++ {
        mp[fruits[j]]++ 
        if mp[fruits[j]] == 1 {
            unique++
        }
        for unique > 2 {
            mp[fruits[i]]--
            if mp[fruits[i]] == 0 {
                unique--
            }
            i++
        }
        res = max(res, j - i + 1)
    }
    return res
}

func totalFruit1(fruits []int) int {
    res, s := 0, 0 
    ts, te := fruits[0],fruits[0] // 起始位置, 终止位置
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(fruits); i++{
        // 当 i 下标对应的值不再两种连续类型中，则 te 更新成i 小标的值
        // ts 的值是 te 前一位相同值的起始位置 例如 1 2 3 2 2 4 当 i 下标是 4 是 s 下标应该跑到 3 后面的那个 2
        if fruits[i] != ts && fruits[i] != te {
            te = fruits[i]
            s = i - 1 
            ts = fruits[s]
            for s >= 0 && fruits[s] == ts {
                s--
            }
            s++
        }
        res = max(res, i - s + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: fruits = [1,2,1]
    // Output: 3
    // Explanation: We can pick from all 3 trees.
    fmt.Println(totalFruit([]int{1,2,1})) // 3 {1,2,1}
    // Example 2:
    // Input: fruits = [0,1,2,2]
    // Output: 3
    // Explanation: We can pick from trees [1,2,2].
    // If we had started at the first tree, we would only pick from trees [0,1].
    fmt.Println(totalFruit([]int{0,1,2,2})) // 3 {1,2,2}
    // Example 3:
    // Input: fruits = [1,2,3,2,2]
    // Output: 4
    // Explanation: We can pick from trees [2,3,2,2].
    // If we had started at the first tree, we would only pick from trees [1,2].
    fmt.Println(totalFruit([]int{1,2,3,2,2})) // 4 {2,3,2,2}

    fmt.Println(totalFruit([]int{3,3,3,1,2,1,1,2,3,3,4})) // 5 {1,2,1,1,2}

    fmt.Println(totalFruit1([]int{1,2,1})) // 3 {1,2,1}
    fmt.Println(totalFruit1([]int{0,1,2,2})) // 3 {1,2,2}
    fmt.Println(totalFruit1([]int{1,2,3,2,2})) // 4 {2,3,2,2}
    fmt.Println(totalFruit1([]int{3,3,3,1,2,1,1,2,3,3,4})) // 5 {1,2,1,1,2}
}