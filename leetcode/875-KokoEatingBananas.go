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
    mx := 0; 
    for _, pile := range piles {  
        if pile > mx { // 取到最大的值
            mx = pile 
        } 
    }
    left, right, res := 1, mx, mx
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

func minEatingSpeed1(piles []int, h int) int {
    r, l := 0, 1
    for _, pile := range piles {
        if pile > r {
            r = pile
        }
    }
    res := r
    for l <= r {
        mid, hour := l + (r-l) / 2, 0
        for _, pile := range piles {
            hour += (pile + mid - 1) / mid
        }
        if hour <= h {
            res = min(res, mid)
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: piles = [3,6,7,11], h = 8
    // Output: 4
    fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8)) // 4
    // Example 2:
    // Input: piles = [30,11,23,4,20], h = 5
    // Output: 30
    fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5)) // 30
    // Example 3:
    // Input: piles = [30,11,23,4,20], h = 6
    // Output: 23
    fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6)) // 23

    fmt.Println(minEatingSpeed([]int{1,2,3,4,5,6,7,8,9}, 6)) // 9
    fmt.Println(minEatingSpeed([]int{9,8,7,6,5,4,3,2,1}, 6)) // 9

    fmt.Println(minEatingSpeed1([]int{3,6,7,11}, 8)) // 4
    fmt.Println(minEatingSpeed1([]int{30,11,23,4,20}, 5)) // 30
    fmt.Println(minEatingSpeed1([]int{30,11,23,4,20}, 6)) // 23
    fmt.Println(minEatingSpeed1([]int{1,2,3,4,5,6,7,8,9}, 6)) // 9
    fmt.Println(minEatingSpeed1([]int{9,8,7,6,5,4,3,2,1}, 6)) // 9
}