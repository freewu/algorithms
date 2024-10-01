package main

// 1497. Check If Array Pairs Are Divisible by k
// Given an array of integers arr of even length n and an integer k.

// We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.

// Return true If you can find a way to do that or false otherwise.

// Example 1:
// Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
// Output: true
// Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).

// Example 2:
// Input: arr = [1,2,3,4,5,6], k = 7
// Output: true
// Explanation: Pairs are (1,6),(2,5) and(3,4).

// Example 3:
// Input: arr = [1,2,3,4,5,6], k = 10
// Output: false
// Explanation: You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.

// Constraints:
//     arr.length == n
//     1 <= n <= 10^5
//     n is even.
//     -10^9 <= arr[i] <= 10^9
//     1 <= k <= 10^5

import "fmt"

func canArrange(arr []int, k int) bool {
    mp := make(map[int]int)
    for _, v := range arr {
        mp[(v % k + k) % k]++
    }
    for i := 1; i <= k / 2; i++ {
        if mp[i] != mp[k - i] { // 没有找到可以配对的
            return false
        }
    }
    return mp[0] % 2 == 0
}

func canArrange1(arr []int, k int) bool {
    mp := make([]int,k)
    for _, v := range arr {
        mp[(v % k + k) % k]++
    }
    for i := 1;i < k;i++ {
        // we can match all pairings with remaineder==j to remainder==k-j
        // eg. k=5 [1, 2,3,4, 12, 13, 22, 23]
        //     mp array looks like [0, 1, 3, 3, 1] => true
        if mp[i] != mp[k-i] {
            return false
        }
    }
    // for remainder == 0, we must have an even number of pairings
    // eg. k=5 [5,10,15,20]
    //     mp array looks like [4, 0, 0, 0, 0] => true
    return mp[0] % 2 == 0
}

func main() {
    // Example 1:
    // Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
    // Output: true
    // Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
    fmt.Println(canArrange([]int{1,2,3,4,5,10,6,7,8,9}, 5)) // true
    // Example 2:
    // Input: arr = [1,2,3,4,5,6], k = 7
    // Output: true
    // Explanation: Pairs are (1,6),(2,5) and(3,4).
    fmt.Println(canArrange([]int{1,2,3,4,5,6}, 7)) // true
    // Example 3:
    // Input: arr = [1,2,3,4,5,6], k = 10
    // Output: false
    // Explanation: You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.
    fmt.Println(canArrange([]int{1,2,3,4,5,6}, 10)) // false

    fmt.Println(canArrange1([]int{1,2,3,4,5,10,6,7,8,9}, 5)) // true
    fmt.Println(canArrange1([]int{1,2,3,4,5,6}, 7)) // true
    fmt.Println(canArrange1([]int{1,2,3,4,5,6}, 10)) // false
}