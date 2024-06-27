package main

// 683. K Empty Slots
// You have n bulbs in a row numbered from 1 to n. 
// Initially, all the bulbs are turned off. 
// We turn on exactly one bulb every day until all bulbs are on after n days.

// You are given an array bulbs of length n where bulbs[i] = x means that on the (i+1)th day, 
// we will turn on the bulb at position x where i is 0-indexed and x is 1-indexed.

// Given an integer k, return the minimum day number such 
// that there exists two turned on bulbs that have exactly k bulbs between them that are all turned off. 
// If there isn't such day, return -1.

// Example 1:
// Input: bulbs = [1,3,2], k = 1
// Output: 2
// Explanation:
// On the first day: bulbs[0] = 1, first bulb is turned on: [1,0,0]
// On the second day: bulbs[1] = 3, third bulb is turned on: [1,0,1]
// On the third day: bulbs[2] = 2, second bulb is turned on: [1,1,1]
// We return 2 because on the second day, there were two on bulbs with one off bulb between them.

// Example 2:
// Input: bulbs = [1,2,3], k = 1
// Output: -1

// Constraints:
//     n == bulbs.length
//     1 <= n <= 2 * 10^4
//     1 <= bulbs[i] <= n
//     bulbs is a permutation of numbers from 1 to n.
//     0 <= k <= 2 * 10^4

import "fmt"

func kEmptySlots(bulbs []int, k int) int {
    inf, l, r, n := 1 << 32 - 1, 0, k+1, len(bulbs)
    res, days := inf, make([]int, n)
    for i := 0; i < n; i++ {
        days[bulbs[i]-1] = i + 1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; r < n; i++ {
        if days[i] < days[l] || days[i] <= days[r] {
            if i == r {
                t := max(days[l], days[r])
                if t < res {
                    res = t
                }
            }
            l, r = i, k + i + 1
        }
    }
    if res == inf {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: bulbs = [1,3,2], k = 1
    // Output: 2
    // Explanation:
    // On the first day: bulbs[0] = 1, first bulb is turned on: [1,0,0]
    // On the second day: bulbs[1] = 3, third bulb is turned on: [1,0,1]
    // On the third day: bulbs[2] = 2, second bulb is turned on: [1,1,1]
    // We return 2 because on the second day, there were two on bulbs with one off bulb between them.
    fmt.Println(kEmptySlots([]int{1,3,2}, 1)) // 2
    // Example 2:
    // Input: bulbs = [1,2,3], k = 1
    // Output: -1
    fmt.Println(kEmptySlots([]int{1,2,3}, 1)) // -1
}