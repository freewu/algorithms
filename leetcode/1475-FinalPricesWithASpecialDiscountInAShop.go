package main

// 1475. Final Prices With a Special Discount in a Shop
// You are given an integer array prices where prices[i] is the price of the ith item in a shop.

// There is a special discount for items in the shop. 
// If you buy the ith item, then you will receive a discount equivalent to prices[j] where j is the minimum index such that j > i and prices[j] <= prices[i]. 
// Otherwise, you will not receive any discount at all.

// Return an integer array answer where answer[i] is the final price you will pay for the ith item of the shop, considering the special discount.

// Example 1:
// Input: prices = [8,4,6,2,3]
// Output: [4,2,4,2,3]
// Explanation: 
// For item 0 with price[0]=8 you will receive a discount equivalent to prices[1]=4, therefore, the final price you will pay is 8 - 4 = 4.
// For item 1 with price[1]=4 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 4 - 2 = 2.
// For item 2 with price[2]=6 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 6 - 2 = 4.
// For items 3 and 4 you will not receive any discount at all.

// Example 2:
// Input: prices = [1,2,3,4,5]
// Output: [1,2,3,4,5]
// Explanation: In this case, for all items, you will not receive any discount at all.

// Example 3:
// Input: prices = [10,1,1,6]
// Output: [9,0,1,6]

// Constraints:
//     1 <= prices.length <= 500
//     1 <= prices[i] <= 1000

import "fmt"

// Monotonic Stack
func finalPrices(prices []int) []int {
    // Prices without a discount applied
    // Because we apply the first applicable discount, the values in this stack will be
    // in ascending order
    stack := make([]int, 0, len(prices))
    for i, discount := range prices {
        for len(stack) > 0 && discount <= prices[stack[len(stack)-1]] {
            prices[stack[len(stack)-1]] -= discount
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return prices
}

func main() {
    // Example 1:
    // Input: prices = [8,4,6,2,3]
    // Output: [4,2,4,2,3]
    // Explanation: 
    // For item 0 with price[0]=8 you will receive a discount equivalent to prices[1]=4, therefore, the final price you will pay is 8 - 4 = 4.
    // For item 1 with price[1]=4 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 4 - 2 = 2.
    // For item 2 with price[2]=6 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 6 - 2 = 4.
    // For items 3 and 4 you will not receive any discount at all.
    fmt.Println(finalPrices([]int{8,4,6,2,3})) // [4,2,4,2,3]
    // Example 2:
    // Input: prices = [1,2,3,4,5]
    // Output: [1,2,3,4,5]
    // Explanation: In this case, for all items, you will not receive any discount at all.
    fmt.Println(finalPrices([]int{1,2,3,4,5})) // [1,2,3,4,5]
    // Example 3:
    // Input: prices = [10,1,1,6]
    // Output: [9,0,1,6]
    fmt.Println(finalPrices([]int{10,1,1,6})) // [9,0,1,6]
}