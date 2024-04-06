package main

// 875. Koko Eating Bananas
// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. 
// The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k. 
// Each hour, she chooses some pile of bananas and eats k bananas from that pile. 
// If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h hours.

// Example 1:
// Input: piles = [3,6,7,11], h = 8
// Output: 4

// Example 2:
// Input: piles = [30,11,23,4,20], h = 5
// Output: 30

// Example 3:
// Input: piles = [30,11,23,4,20], h = 6
// Output: 23
 
// Constraints:
//     1 <= piles.length <= 10^4
//     piles.length <= h <= 10^9
//     1 <= piles[i] <= 10^9

import "fmt"

func minEatingSpeed(piles []int, h int) int {
    howLong := func(speed int) int {
        time := 0
        for _, pile := range piles { time += pile / speed; if pile % speed != 0 { time++ } }
        return time
    }
    // 取到最大的值
    max := 0; 
    for _, pile := range piles { 
        if pile > max { 
            max = pile 
        } 
    }
    left, right, res := 1, max, max
    for left <= right {
        mid := left + (right-left) >> 1
        if howLong(mid) <= h { 
            res = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return res
}

func main() {
    fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8)) // 4
    fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5)) // 30
    fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6)) // 23
}