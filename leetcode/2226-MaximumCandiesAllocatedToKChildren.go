package main

// 2226. Maximum Candies Allocated to K Children
// You are given a 0-indexed integer array candies. 
// Each element in the array denotes a pile of candies of size candies[i]. 
// You can divide each pile into any number of sub piles, but you cannot merge two piles together.

// You are also given an integer k. 
// You should allocate piles of candies to k children such that each child gets the same number of candies. 
// Each child can take at most one pile of candies and some piles of candies may go unused.

// Return the maximum number of candies each child can get.

// Example 1:
// Input: candies = [5,8,6], k = 3
// Output: 5
// Explanation: We can divide candies[1] into 2 piles of size 5 and 3, and candies[2] into 2 piles of size 5 and 1. We now have five piles of candies of sizes 5, 5, 3, 5, and 1. We can allocate the 3 piles of size 5 to 3 children. It can be proven that each child cannot receive more than 5 candies.

// Example 2:
// Input: candies = [2,5], k = 11
// Output: 0
// Explanation: There are 11 children but only 7 candies in total, so it is impossible to ensure each child receives at least one candy. Thus, each child gets no candy and the answer is 0.

// Constraints:
//     1 <= candies.length <= 10^5
//     1 <= candies[i] <= 10^7
//     1 <= k <= 10^12

import "fmt"

func maximumCandies(candies []int, k int64) int {
    res, n, low, high := 0, len(candies), 1, -1 << 31
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range candies {
        high = max(high, v)
    }
    for low <= high {
        val, mid := int64(0), low + (high - low) / 2
        for i := 0; i < n; i++ {
            val += int64(candies[i] / mid)
        }
        if val >= k {
            res, low = mid, mid + 1
        } else{
            high = mid - 1
        }
    }
    return res
}

func maximumCandies1(candies []int, k int64) int {
    sum, mx := int64(0), 0
    for _, v := range candies {
        sum += int64(v)
        mx = max(mx, v)
    }
    if sum < k { return 0 }
    check := func(mid int) bool {
        if k * int64(mid) > sum { return false }
        count := 0
        for _, v := range candies  {
            count += (v / mid)
        }
        return int64(count) >= k
    }
    left, right := 1, mx
    for left <= right {
        mid := (left + right) / 2
        if check(mid) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right
}

func main() {
    // Example 1:
    // Input: candies = [5,8,6], k = 3
    // Output: 5
    // Explanation: We can divide candies[1] into 2 piles of size 5 and 3, and candies[2] into 2 piles of size 5 and 1. We now have five piles of candies of sizes 5, 5, 3, 5, and 1. We can allocate the 3 piles of size 5 to 3 children. It can be proven that each child cannot receive more than 5 candies.
    fmt.Println(maximumCandies([]int{5,8,6}, 3)) // 5
    // Example 2:
    // Input: candies = [2,5], k = 11
    // Output: 0
    // Explanation: There are 11 children but only 7 candies in total, so it is impossible to ensure each child receives at least one candy. Thus, each child gets no candy and the answer is 0.
    fmt.Println(maximumCandies([]int{2,5}, 11)) // 0

    fmt.Println(maximumCandies1([]int{5,8,6}, 3)) // 5
    fmt.Println(maximumCandies1([]int{2,5}, 11)) // 0
}