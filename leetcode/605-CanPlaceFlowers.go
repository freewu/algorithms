package main

// 605. Can Place Flowers
// You have a long flowerbed in which some of the plots are planted, and some are not. 
// However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, 
// return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

// Example 1:
// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true

// Example 2:
// Input: flowerbed = [1,0,0,0,1], n = 2
// Output: false
 
// Constraints:
//     1 <= flowerbed.length <= 2 * 10^4
//     flowerbed[i] is 0 or 1.
//     There are no two adjacent flowers in flowerbed.
//     0 <= n <= flowerbed.length

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed); i++ {
        // flowerbed[i] == 0 时
        // 判断前后是否已种 flowerbed[i + 1] == 0  && flowerbed[i - 1] == 0
        // 边界判断 i == len(flowerbed) - 1 && i == 0
        if (flowerbed[i] == 0 &&
            (i == len(flowerbed) - 1 || flowerbed[i + 1] == 0) &&
            (i == 0 || flowerbed[i - 1] == 0)) {
                n-- // 能中 目标 -1
                i++ // 能中的话不需要判断下位了 减少一次 判断
            }
    }
    return n <= 0
}

func canPlaceFlowers1(flowerbed []int, n int) bool {
    before, length := 0, len(flowerbed)
    for i := 0;i < length;i ++ {
        if flowerbed[i] == 1 {
            before++
        }
    }
    c, dp := make([]int,length), make([]int,length + 1)// dp[i]代表以i结尾最多多少朵花
    copy(c,flowerbed)
    flowerbed = make([]int,length + 1)
    flowerbed[0] = 0
    for i := 1;i < length+1;i++ {
        flowerbed[i] = c[i-1]
    }
    for i := 1;i < length;i ++ {
        if flowerbed[i-1] == 0 {
            if flowerbed[i+1] == 1 {
                dp[i] = dp[i-1]
            } else {
                dp[i] = dp[i-1] + 1
                flowerbed[i] = 1
            }
        } else {
            dp[i] = dp[i-1]
        }
    }
    if flowerbed[length] == 1 {
        return (dp[length-1] + 1) >= (n + before)
    } else {
        if flowerbed[length-1] != 1 {
            return (dp[length-1] + 1) >= (n + before)
        }
    }
    return dp[length-1] >= (n + before)
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
    l := len(flowerbed)
    for i := 0; i < l && n > 0; i++ {
        if flowerbed[i] != 0 { continue }
        if (i == 0 || flowerbed[i-1] == 0) && (i == l - 1 || flowerbed[i+1] == 0) {
            i++
            n--
        }
    }
    return n == 0
}

func main() {
    // Example 1:
    // Input: flowerbed = [1,0,0,0,1], n = 1
    // Output: true
    fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 1)) // true
    // Example 2:
    // Input: flowerbed = [1,0,0,0,1], n = 2
    // Output: false
    fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 2)) // false

    fmt.Println(canPlaceFlowers1([]int{1,0,0,0,1}, 1)) // true
    fmt.Println(canPlaceFlowers1([]int{1,0,0,0,1}, 2)) // false

    fmt.Println(canPlaceFlowers2([]int{1,0,0,0,1}, 1)) // true
    fmt.Println(canPlaceFlowers2([]int{1,0,0,0,1}, 2)) // false
}