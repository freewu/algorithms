package main

// 2517. Maximum Tastiness of Candy Basket
// You are given an array of positive integers price where price[i] denotes the price of the ith candy and a positive integer k.

// The store sells baskets of k distinct candies. 
// The tastiness of a candy basket is the smallest absolute difference of the prices of any two candies in the basket.

// Return the maximum tastiness of a candy basket.

// Example 1:
// Input: price = [13,5,1,8,21,2], k = 3
// Output: 8
// Explanation: Choose the candies with the prices [13,5,21].
// The tastiness of the candy basket is: min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8.
// It can be proven that 8 is the maximum tastiness that can be achieved.

// Example 2:
// Input: price = [1,3,1], k = 2
// Output: 2
// Explanation: Choose the candies with the prices [1,3].
// The tastiness of the candy basket is: min(|1 - 3|) = min(2) = 2.
// It can be proven that 2 is the maximum tastiness that can be achieved.

// Example 3:
// Input: price = [7,7,7,7], k = 2
// Output: 0
// Explanation: Choosing any two distinct candies from the candies we have will result in a tastiness of 0.

// Constraints:
//     2 <= k <= price.length <= 10^5
//     1 <= price[i] <= 10^9

import "fmt"
import "sort"

func maximumTastiness(price []int, k int) int {
    n := len(price)
    sort.Ints(price)
    maxDiff := price[n-1] - price[0]
    if maxDiff == 0 { return 0 }
    return sort.Search(maxDiff, func(tastiness int) bool {
        minNextCandy := price[0] + tastiness
        for i, j := 1, n - k + 1; i <= j; i++ {
            if price[i] > minNextCandy {
                minNextCandy = price[i] + tastiness
                j++
                if j == n {
                    return false
                }
            }
        }
        return true
    })
}

func maximumTastiness1(price []int, k int) int {
    sort.Ints(price)
    mx := price[len(price) - 1] - price[0]
    return sort.Search(mx, func(v int) bool {
        selected, pre := 1, price[0]
        for i := 1; i < len(price); i++ {
            if price[i] - pre >= v + 1 {
                selected++
                pre = price[i]
            }
        }
        return selected < k
    })
}

func maximumTastiness2(price []int, k int) int {
    sort.Ints(price)
    left, right := 0, (price[len(price) - 1] - price[0]) / (k - 1) + 1
    check := func(diff int) bool {
        count, pre := 1, price[0]
        for _, v := range price {
            if pre + diff <= v {
                count++
                pre = v
            }
        }
        return count < k
    }
    for left < right {
        mid := left + (right - left) / 2
        if check(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left - 1
}

func main() {
    // Example 1:
    // Input: price = [13,5,1,8,21,2], k = 3
    // Output: 8
    // Explanation: Choose the candies with the prices [13,5,21].
    // The tastiness of the candy basket is: min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8.
    // It can be proven that 8 is the maximum tastiness that can be achieved.
    fmt.Println(maximumTastiness([]int{13,5,1,8,21,2}, 3)) // 8
    // Example 2:
    // Input: price = [1,3,1], k = 2
    // Output: 2
    // Explanation: Choose the candies with the prices [1,3].
    // The tastiness of the candy basket is: min(|1 - 3|) = min(2) = 2.
    // It can be proven that 2 is the maximum tastiness that can be achieved.
    fmt.Println(maximumTastiness([]int{1,3,1}, 2)) // 2
    // Example 3:
    // Input: price = [7,7,7,7], k = 2
    // Output: 0
    // Explanation: Choosing any two distinct candies from the candies we have will result in a tastiness of 0.
    fmt.Println(maximumTastiness([]int{7,7,7,7}, 2)) // 0

    fmt.Println(maximumTastiness([]int{1,2,3,4,5,6,7,8,9}, 2)) // 8
    fmt.Println(maximumTastiness([]int{9,8,7,6,5,4,3,2,1}, 2)) // 8

    fmt.Println(maximumTastiness1([]int{13,5,1,8,21,2}, 3)) // 8
    fmt.Println(maximumTastiness1([]int{1,3,1}, 2)) // 2
    fmt.Println(maximumTastiness1([]int{7,7,7,7}, 2)) // 0
    fmt.Println(maximumTastiness1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 8
    fmt.Println(maximumTastiness1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 8

    fmt.Println(maximumTastiness2([]int{13,5,1,8,21,2}, 3)) // 8
    fmt.Println(maximumTastiness2([]int{1,3,1}, 2)) // 2
    fmt.Println(maximumTastiness2([]int{7,7,7,7}, 2)) // 0
    fmt.Println(maximumTastiness2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 8
    fmt.Println(maximumTastiness2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 8
}