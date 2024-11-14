package main

// 2064. Minimized Maximum of Products Distributed to Any Store
// You are given an integer n indicating there are n specialty retail stores. 
// There are m product types of varying amounts, which are given as a 0-indexed integer array quantities, 
// where quantities[i] represents the number of products of the ith product type.

// You need to distribute all products to the retail stores following these rules:
//     1. A store can only be given at most one product type but can be given any amount of it.
//     2. After distribution, each store will have been given some number of products (possibly 0). 
//        Let x represent the maximum number of products given to any store. 
//        You want x to be as small as possible, i.e., you want to minimize the maximum number of products that are given to any store.

// Return the minimum possible x.

// Example 1:
// Input: n = 6, quantities = [11,6]
// Output: 3
// Explanation: One optimal way is:
// - The 11 products of type 0 are distributed to the first four stores in these amounts: 2, 3, 3, 3
// - The 6 products of type 1 are distributed to the other two stores in these amounts: 3, 3
// The maximum number of products given to any store is max(2, 3, 3, 3, 3, 3) = 3.

// Example 2:
// Input: n = 7, quantities = [15,10,10]
// Output: 5
// Explanation: One optimal way is:
// - The 15 products of type 0 are distributed to the first three stores in these amounts: 5, 5, 5
// - The 10 products of type 1 are distributed to the next two stores in these amounts: 5, 5
// - The 10 products of type 2 are distributed to the last two stores in these amounts: 5, 5
// The maximum number of products given to any store is max(5, 5, 5, 5, 5, 5, 5) = 5.

// Example 3:
// Input: n = 1, quantities = [100000]
// Output: 100000
// Explanation: The only optimal way is:
// - The 100000 products of type 0 are distributed to the only store.
// The maximum number of products given to any store is max(100000) = 100000.

// Constraints:
//     m == quantities.length
//     1 <= m <= n <= 10^5
//     1 <= quantities[i] <= 10^5

import "fmt"
import "math"

func minimizedMaximum(n int, quantities []int) int {
    right := 0
    for _, v := range quantities { // 寻找数组 quantities 中的最大值，当作返回值x可能的最大值，
        if v > right {
            right = v
        }
    }
    left := 1 // x的最小值则设为1.
    canDistribute := func(k int) bool {
        res := 0
        for _, v := range quantities {
            res += v / k
            if !(v % k == 0) {
                res++
            }
        }
        return res <= n
    }
    for left < right {
        mid := left + (right - left)/2
        if canDistribute(mid) { // 来判断当商店中分配商品数目为k时，是否能将所有产品都分发到商店中
            right = mid
        } else {
            left = mid + 1 // 使用二分查找来判断返回的x的最小值
        }
    }
    return left
}

func minimizedMaximum1(n int, quantities []int) int {
    left, right := 1, 0
    for _, v := range quantities {
        if v > right {
            right = v
        }
    }
    for left < right {
        mid := (left + right) / 2
        shop := 0
        for _, v := range quantities {
            shop += int(math.Ceil(float64(v) / float64(mid)))
        }
        if shop > n {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: n = 6, quantities = [11,6]
    // Output: 3
    // Explanation: One optimal way is:
    // - The 11 products of type 0 are distributed to the first four stores in these amounts: 2, 3, 3, 3
    // - The 6 products of type 1 are distributed to the other two stores in these amounts: 3, 3
    // The maximum number of products given to any store is max(2, 3, 3, 3, 3, 3) = 3.
    fmt.Println(minimizedMaximum(6, []int{11,6})) // 3
    // Example 2:
    // Input: n = 7, quantities = [15,10,10]
    // Output: 5
    // Explanation: One optimal way is:
    // - The 15 products of type 0 are distributed to the first three stores in these amounts: 5, 5, 5
    // - The 10 products of type 1 are distributed to the next two stores in these amounts: 5, 5
    // - The 10 products of type 2 are distributed to the last two stores in these amounts: 5, 5
    // The maximum number of products given to any store is max(5, 5, 5, 5, 5, 5, 5) = 5.
    fmt.Println(minimizedMaximum(7, []int{15,10,10})) // 5
    // Example 3:
    // Input: n = 1, quantities = [100000]
    // Output: 100000
    // Explanation: The only optimal way is:
    // - The 100000 products of type 0 are distributed to the only store.
    // The maximum number of products given to any store is max(100000) = 100000.
    fmt.Println(minimizedMaximum(1, []int{100000})) // 100000

    fmt.Println(minimizedMaximum1(6, []int{11,6})) // 3
    fmt.Println(minimizedMaximum1(7, []int{15,10,10})) // 5
    fmt.Println(minimizedMaximum1(1, []int{100000})) // 100000
}