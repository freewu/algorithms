package main

// 2107. Number of Unique Flavors After Sharing K Candies
// You are given a 0-indexed integer array candies, where candies[i] represents the flavor of the ith candy. 
// Your mom wants you to share these candies with your little sister by giving her k consecutive candies, 
// but you want to keep as many flavors of candies as possible.

// Return the maximum number of unique flavors of candy you can keep after sharing with your sister. 

// Example 1:
// Input: candies = [1,2,2,3,4,3], k = 3
// Output: 3
// Explanation: 
// Give the candies in the range [1, 3] (inclusive) with flavors [2,2,3].
// You can eat candies with flavors [1,4,3].
// There are 3 unique flavors, so return 3.

// Example 2:
// Input: candies = [2,2,2,2,3,3], k = 2
// Output: 2
// Explanation: 
// Give the candies in the range [3, 4] (inclusive) with flavors [2,3].
// You can eat candies with flavors [2,2,2,3].
// There are 2 unique flavors, so return 2.
// Note that you can also share the candies with flavors [2,2] and eat the candies with flavors [2,2,3,3].

// Example 3:
// Input: candies = [2,4,5], k = 0
// Output: 3
// Explanation: 
// You do not have to give any candies.
// You can eat the candies with flavors [2,4,5].
// There are 3 unique flavors, so return 3.

// Constraints:
//     1 <= candies.length <= 10^5
//     1 <= candies[i] <= 10^5
//     0 <= k <= candies.length

import "fmt"

func shareCandies(candies []int, k int) int {
    mp, count := make([]int,100005), 0 
    for _, v := range candies { // 统计有多少个不同的数字
        mp[v]++
        if mp[v] == 1 {
            count++
        }
    }
    if k == 0 { return count }
    for i := 0; i < k; i++ { // 初始化窗口
        mp[candies[i]]--
        if mp[candies[i]] == 0 {
            mp[candies[i]] = -1
            count--
        }
    }
    res, j := count, 0
    for i := k; i < len(candies); i++ {
        // 滑出窗口
        mp[candies[j]]++ 
        if mp[candies[j]] == 0 {
            mp[candies[j]] = 1
            count++
        }
        // 滑入窗口
        mp[candies[i]]--
        if mp[candies[i]] == 0 {
            mp[candies[i]] = -1
            count--
        }
        if count > res {
            res = count
        }
        j++
    }
    return res
}

func shareCandies1(candies []int, k int) int {
    res, n, mp := 0, len(candies), make(map[int]int)
    for i := k; i < n; i++ {
        mp[candies[i]]++
    }
    for i := k; i <= n; i++ {
        if res < len(mp) {
            res = len(mp)
        }
        if i == n { break }
        mp[candies[i-k]]++
        mp[candies[i]]--
        if mp[candies[i]] == 0 {
            delete(mp, candies[i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: candies = [1,2,2,3,4,3], k = 3
    // Output: 3
    // Explanation: 
    // Give the candies in the range [1, 3] (inclusive) with flavors [2,2,3].
    // You can eat candies with flavors [1,4,3].
    // There are 3 unique flavors, so return 3.
    fmt.Println(shareCandies([]int{1,2,2,3,4,3}, 3)) // 3
    // Example 2:
    // Input: candies = [2,2,2,2,3,3], k = 2
    // Output: 2
    // Explanation: 
    // Give the candies in the range [3, 4] (inclusive) with flavors [2,3].
    // You can eat candies with flavors [2,2,2,3].
    // There are 2 unique flavors, so return 2.
    // Note that you can also share the candies with flavors [2,2] and eat the candies with flavors [2,2,3,3].
    fmt.Println(shareCandies([]int{2,2,2,2,3,3}, 2)) // 2
    // Example 3:
    // Input: candies = [2,4,5], k = 0
    // Output: 3
    // Explanation: 
    // You do not have to give any candies.
    // You can eat the candies with flavors [2,4,5].
    // There are 3 unique flavors, so return 3.
    fmt.Println(shareCandies([]int{2,4,5}, 0)) // 3

    fmt.Println(shareCandies1([]int{1,2,2,3,4,3}, 3)) // 3
    fmt.Println(shareCandies1([]int{2,2,2,2,3,3}, 2)) // 2
    fmt.Println(shareCandies1([]int{2,4,5}, 0)) // 3
}