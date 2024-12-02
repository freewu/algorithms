package main

// 2819. Minimum Relative Loss After Buying Chocolates
// You are given an integer array prices, 
// which shows the chocolate prices and a 2D integer array queries, where queries[i] = [ki, mi].

// Alice and Bob went to buy some chocolates, and Alice suggested a way to pay for them, and Bob agreed.

// The terms for each query are as follows:
//     If the price of a chocolate is less than or equal to ki, Bob pays for it.
//     Otherwise, Bob pays ki of it, and Alice pays the rest.

// Bob wants to select exactly mi chocolates such that his relative loss is minimized, 
// more formally, if, in total, Alice has paid ai and Bob has paid bi, Bob wants to minimize bi - ai.

// Return an integer array ans where ans[i] is Bob's minimum relative loss possible for queries[i].

// Example 1:
// Input: prices = [1,9,22,10,19], queries = [[18,4],[5,2]]
// Output: [34,-21]
// Explanation: For the 1st query Bob selects the chocolates with prices [1,9,10,22]. He pays 1 + 9 + 10 + 18 = 38 and Alice pays 0 + 0 + 0 + 4 = 4. So Bob's relative loss is 38 - 4 = 34.
// For the 2nd query Bob selects the chocolates with prices [19,22]. He pays 5 + 5 = 10 and Alice pays 14 + 17 = 31. So Bob's relative loss is 10 - 31 = -21.
// It can be shown that these are the minimum possible relative losses.

// Example 2:
// Input: prices = [1,5,4,3,7,11,9], queries = [[5,4],[5,7],[7,3],[4,5]]
// Output: [4,16,7,1]
// Explanation: For the 1st query Bob selects the chocolates with prices [1,3,9,11]. He pays 1 + 3 + 5 + 5 = 14 and Alice pays 0 + 0 + 4 + 6 = 10. So Bob's relative loss is 14 - 10 = 4.
// For the 2nd query Bob has to select all the chocolates. He pays 1 + 5 + 4 + 3 + 5 + 5 + 5 = 28 and Alice pays 0 + 0 + 0 + 0 + 2 + 6 + 4 = 12. So Bob's relative loss is 28 - 12 = 16.
// For the 3rd query Bob selects the chocolates with prices [1,3,11] and he pays 1 + 3 + 7 = 11 and Alice pays 0 + 0 + 4 = 4. So Bob's relative loss is 11 - 4 = 7.
// For the 4th query Bob selects the chocolates with prices [1,3,7,9,11] and he pays 1 + 3 + 4 + 4 + 4 = 16 and Alice pays 0 + 0 + 3 + 5 + 7 = 15. So Bob's relative loss is 16 - 15 = 1.
// It can be shown that these are the minimum possible relative losses.

// Example 3:
// Input: prices = [5,6,7], queries = [[10,1],[5,3],[3,3]]
// Output: [5,12,0]
// Explanation: For the 1st query Bob selects the chocolate with price 5 and he pays 5 and Alice pays 0. So Bob's relative loss is 5 - 0 = 5.
// For the 2nd query Bob has to select all the chocolates. He pays 5 + 5 + 5 = 15 and Alice pays 0 + 1 + 2 = 3. So Bob's relative loss is 15 - 3 = 12.
// For the 3rd query Bob has to select all the chocolates. He pays 3 + 3 + 3 = 9 and Alice pays 2 + 3 + 4 = 9. So Bob's relative loss is 9 - 9 = 0.
// It can be shown that these are the minimum possible relative losses.

// Constraints:
//     1 <= prices.length == n <= 10^5
//     1 <= prices[i] <= 10^9
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     1 <= ki <= 10^9
//     1 <= mi <= n

import "fmt"
import "sort"

func minimumRelativeLosses(prices []int, queries [][]int) []int64 {
    n := len(prices)
    sort.Ints(prices)
    prefix := make([]int, n + 1)
    for i, v := range prices {
        prefix[i+1] = prefix[i] + v
    }
    find := func(k, m int) int {
        l, r := 0, sort.Search(n, func(i int) bool { 
            return prices[i] > k 
        })
        if r > m {
            r = m
        }
        for l < r {
            mid := (l + r) >> 1
            right := m - mid
            if prices[mid] < 2 * k - prices[n - right] {
                l = mid + 1
            } else {
                r = mid
            }
        }
        return l
    }
    res := make([]int64, len(queries))
    for i, q := range queries {
        k, m := q[0], q[1]
        l := find(k, m)
        r := m - l
        res[i] = int64(prefix[l] + 2 * k * r - (prefix[n] - prefix[n - r]))
    }
    return res
}

func minimumRelativeLosses1(prices []int, queries [][]int) []int64 {
    n := len(prices)
    sort.Ints(prices)
    sum := make([]int64, n + 1)
    for i := 0; i < n; i++ {
        sum[i + 1] = sum[i] + int64(prices[i])
    }
    get := func(k, dist int) int {
        left := sort.SearchInts(prices, k - dist + 1)
        right := n - sort.SearchInts(prices, k + dist)
        return left + right
    }
    calc := func(k, dist int) int64 {
        left := sort.SearchInts(prices, k - dist + 1)
        right := sort.SearchInts(prices, k + dist)
        return sum[left] + int64(k)*(int64(n) - int64(right)) * 2 - (sum[n] - sum[right])
    }
    res := make([]int64, len(queries))
    for i, q := range queries {
        k, m, maxdist := q[0], q[1], 0
        low, high := 0, int(1e9)
        for low <= high {
            mid := (high + low) >> 1
            if get(k, mid) >= m {
                low = mid + 1
                maxdist = mid
            } else {
                high = mid - 1
            }
        }
        total, cur := get(k, maxdist), calc(k, maxdist)
        if total >= m { // 减去多余的贡献值
            cur -= int64(total - m) * int64(k - maxdist)
        }
        res[i] = cur
    }
    return res
}

func main() {
    // Example 1:
    // Input: prices = [1,9,22,10,19], queries = [[18,4],[5,2]]
    // Output: [34,-21]
    // Explanation: For the 1st query Bob selects the chocolates with prices [1,9,10,22]. He pays 1 + 9 + 10 + 18 = 38 and Alice pays 0 + 0 + 0 + 4 = 4. So Bob's relative loss is 38 - 4 = 34.
    // For the 2nd query Bob selects the chocolates with prices [19,22]. He pays 5 + 5 = 10 and Alice pays 14 + 17 = 31. So Bob's relative loss is 10 - 31 = -21.
    // It can be shown that these are the minimum possible relative losses.
    fmt.Println(minimumRelativeLosses([]int{1,9,22,10,19}, [][]int{{18,4},{5,2}})) // [34,-21]
    // Example 2:
    // Input: prices = [1,5,4,3,7,11,9], queries = [[5,4],[5,7],[7,3],[4,5]]
    // Output: [4,16,7,1]
    // Explanation: For the 1st query Bob selects the chocolates with prices [1,3,9,11]. He pays 1 + 3 + 5 + 5 = 14 and Alice pays 0 + 0 + 4 + 6 = 10. So Bob's relative loss is 14 - 10 = 4.
    // For the 2nd query Bob has to select all the chocolates. He pays 1 + 5 + 4 + 3 + 5 + 5 + 5 = 28 and Alice pays 0 + 0 + 0 + 0 + 2 + 6 + 4 = 12. So Bob's relative loss is 28 - 12 = 16.
    // For the 3rd query Bob selects the chocolates with prices [1,3,11] and he pays 1 + 3 + 7 = 11 and Alice pays 0 + 0 + 4 = 4. So Bob's relative loss is 11 - 4 = 7.
    // For the 4th query Bob selects the chocolates with prices [1,3,7,9,11] and he pays 1 + 3 + 4 + 4 + 4 = 16 and Alice pays 0 + 0 + 3 + 5 + 7 = 15. So Bob's relative loss is 16 - 15 = 1.
    // It can be shown that these are the minimum possible relative losses.
    fmt.Println(minimumRelativeLosses([]int{1,5,4,3,7,11,9}, [][]int{{5,4},{5,7},{7,3},{4,5}})) // [4,16,7,1]
    // Example 3:
    // Input: prices = [5,6,7], queries = [[10,1],[5,3],[3,3]]
    // Output: [5,12,0]
    // Explanation: For the 1st query Bob selects the chocolate with price 5 and he pays 5 and Alice pays 0. So Bob's relative loss is 5 - 0 = 5.
    // For the 2nd query Bob has to select all the chocolates. He pays 5 + 5 + 5 = 15 and Alice pays 0 + 1 + 2 = 3. So Bob's relative loss is 15 - 3 = 12.
    // For the 3rd query Bob has to select all the chocolates. He pays 3 + 3 + 3 = 9 and Alice pays 2 + 3 + 4 = 9. So Bob's relative loss is 9 - 9 = 0.
    // It can be shown that these are the minimum possible relative losses.
    fmt.Println(minimumRelativeLosses([]int{5,6,7}, [][]int{{10,1},{5,3},{3,3}})) // [5,12,0]

    fmt.Println(minimumRelativeLosses1([]int{1,9,22,10,19}, [][]int{{18,4},{5,2}})) // [34,-21]
    fmt.Println(minimumRelativeLosses1([]int{1,5,4,3,7,11,9}, [][]int{{5,4},{5,7},{7,3},{4,5}})) // [4,16,7,1]
    fmt.Println(minimumRelativeLosses1([]int{5,6,7}, [][]int{{10,1},{5,3},{3,3}})) // [5,12,0]
}