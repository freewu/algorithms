package main

// 4014. Minimum Total Price After Applying Discounts
// You are given two integer arrays prices and discounts.

// The value prices[i] represents the price of the ith item, and discounts[j] represents a discount percentage.

// You may apply discounts subject to the following rules:
//     1. Each discount can be applied to at most one item.
//     2. Each item can receive at most one discount.
//     3. An item may also receive no discount.

// If a discount of d percent is applied to an item with price p, its final price becomes (p * (100 - d)) / 100. 
// The final price is not rounded.

// Return the minimum possible sum of final prices after assigning discounts optimally. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// Input: prices = [10,30,21], discounts = [50,60]
// Output: 32.50000
// Explanation:
// Apply discounts[1] = 60 to prices[1] = 30, thus 30 * (100 - 60) / 100 = 12.
// Apply discounts[0] = 50 to prices[2] = 21, thus 21 * (100 - 50) / 100 = 10.5.
// prices[0] = 10 receives no discount, so it stays 10.
// The total is 12 + 10.5 + 10 = 32.50000, which is the minimum possible.

// Example 2:
// Input: prices = [100,70], discounts = [10,40,50]
// Output: 92.00000
// Explanation:​​​​​​​
// Apply discounts[2] = 50 to prices[0] = 100, thus 100 * (100 - 50) / 100 = 50.
// Apply discounts[1] = 40 to prices[1] = 70, thus 70 * (100 - 40) / 100 = 42.
// The total is 50 + 42 = 92.00000, which is the minimum possible.

// Example 3:
// Input: prices = [7,3,9], discounts = [100,100]
// Output: 3.00000
// Explanation:
// Apply discounts[0] = 100 to prices[2] = 9, thus 9 * (100 - 100) / 100 = 0.
// Apply discounts[1] = 100 to prices[0] = 7, thus 7 * (100 - 100) / 100 = 0.
// prices[1] = 3 receives no discount, so it stays 3.
// The total is 0 + 0 + 3 = 3.00000, which is the minimum possible.

// Constraints:
//     1 <= prices.length, discounts.length <= 10^5
//     1 <= prices[i] <= 10^5
//     1 <= discounts[j] <= 100

import "fmt"
import "sort"

func minPrice(prices []int, discounts []int) float64 {
    sort.Slice(prices, func(i, j int) bool {
        return prices[i] > prices[j]
    })
    sort.Slice(discounts, func(i, j int) bool {
        return discounts[i] > discounts[j]
    })
    res := float64(0)
    for i := 0; i < len(prices); i++ {
        if i < len(discounts) {
            res += float64(prices[i]) * float64(100-discounts[i]) / float64(100)
        } else {
            res += float64(prices[i])
        }
    }
    return res
}

func minPrice1(prices []int, discounts []int) float64 {
    sort.Ints(discounts)
    sort.Ints(prices)
    res, i, j  := 0.0, len(prices) - 1, len(discounts) - 1
    for i >= 0 && j >= 0{
        res += (float64(prices[i]) * (100.0-float64(discounts[j])))  / 100.0 
        i--
        j--
    }
    for i>=0{
        res+=float64(prices[i])
        i--
    }
    return res
}

func main() {
    // Example 1:
    // Input: prices = [10,30,21], discounts = [50,60]
    // Output: 32.50000
    // Explanation:
    // Apply discounts[1] = 60 to prices[1] = 30, thus 30 * (100 - 60) / 100 = 12.
    // Apply discounts[0] = 50 to prices[2] = 21, thus 21 * (100 - 50) / 100 = 10.5.
    // prices[0] = 10 receives no discount, so it stays 10.
    // The total is 12 + 10.5 + 10 = 32.50000, which is the minimum possible.
    fmt.Println(minPrice([]int{10,30,21}, []int{50,60})) // 32.50000
    // Example 2:
    // Input: prices = [100,70], discounts = [10,40,50]
    // Output: 92.00000
    // Explanation:​​​​​​​
    // Apply discounts[2] = 50 to prices[0] = 100, thus 100 * (100 - 50) / 100 = 50.
    // Apply discounts[1] = 40 to prices[1] = 70, thus 70 * (100 - 40) / 100 = 42.
    // The total is 50 + 42 = 92.00000, which is the minimum possible.
    fmt.Println(minPrice([]int{100,70}, []int{10,40,50})) // 92.00000
    // Example 3:
    // Input: prices = [7,3,9], discounts = [100,100]
    // Output: 3.00000
    // Explanation:
    // Apply discounts[0] = 100 to prices[2] = 9, thus 9 * (100 - 100) / 100 = 0.
    // Apply discounts[1] = 100 to prices[0] = 7, thus 7 * (100 - 100) / 100 = 0.
    // prices[1] = 3 receives no discount, so it stays 3.
    // The total is 0 + 0 + 3 = 3.00000, which is the minimum possible.
    fmt.Println(minPrice([]int{7,3,9}, []int{100,100})) // 3.00000

    fmt.Println(minPrice([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 42.150000000000006
    fmt.Println(minPrice([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 42.150000000000006
    fmt.Println(minPrice([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 42.150000000000006
    fmt.Println(minPrice([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 42.150000000000006

    fmt.Println(minPrice1([]int{10,30,21}, []int{50,60})) // 32.50000
    fmt.Println(minPrice1([]int{100,70}, []int{10,40,50})) // 92.00000
    fmt.Println(minPrice1([]int{7,3,9}, []int{100,100})) // 3.00000
    fmt.Println(minPrice1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 42.150000000000006
    fmt.Println(minPrice1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 42.150000000000006
    fmt.Println(minPrice1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 42.150000000000006
    fmt.Println(minPrice1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 42.150000000000006
}