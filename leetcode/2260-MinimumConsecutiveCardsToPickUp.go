package main

// 2260. Minimum Consecutive Cards to Pick Up
// You are given an integer array cards where cards[i] represents the value of the ith card. 
// A pair of cards are matching if the cards have the same value.

// Return the minimum number of consecutive cards you have to pick up to have a pair of matching cards among the picked cards. 
// If it is impossible to have matching cards, return -1.

// Example 1:
// Input: cards = [3,4,2,3,4,7]
// Output: 4
// Explanation: We can pick up the cards [3,4,2,3] which contain a matching pair of cards with value 3. Note that picking up the cards [4,2,3,4] is also optimal.

// Example 2:
// Input: cards = [1,0,5,3]
// Output: -1
// Explanation: There is no way to pick up a set of consecutive cards that contain a pair of matching cards.

// Constraints:
//     1 <= cards.length <= 10^5
//     0 <= cards[i] <= 10^6

import "fmt"

// Sliding Window
func minimumCardPickup(cards []int) int {
    res, left := 1 << 31, 0
    mp := make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for right := 0; right < len(cards); right++ {
        mp[cards[right]]++
        for mp[cards[right]] > 1 {
            res = min(res, right - left + 1)
            mp[cards[left]]--
            left++
        }
    }
    if res == 1 << 31 {
        return -1
    }
    return res
}

func minimumCardPickup1(cards []int) int {
    res := 1 << 31 
    mp := make(map[int]int,len(cards)) // key cards[i], value index
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, c := range cards {
        if v, ok := mp[c]; ok {
            res = min(res, i - v + 1)
        }
        mp[c] = i 
    }    
    if res == 1 << 31 {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: cards = [3,4,2,3,4,7]
    // Output: 4
    // Explanation: We can pick up the cards [3,4,2,3] which contain a matching pair of cards with value 3. Note that picking up the cards [4,2,3,4] is also optimal.
    fmt.Println(minimumCardPickup([]int{3,4,2,3,4,7})) // 4
    // Example 2:
    // Input: cards = [1,0,5,3]
    // Output: -1
    // Explanation: There is no way to pick up a set of consecutive cards that contain a pair of matching cards.
    fmt.Println(minimumCardPickup([]int{1,0,5,3})) // -1

    fmt.Println(minimumCardPickup1([]int{3,4,2,3,4,7})) // 4
    fmt.Println(minimumCardPickup1([]int{1,0,5,3})) // -1
}