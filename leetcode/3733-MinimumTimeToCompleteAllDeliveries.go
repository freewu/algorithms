package main

// 3733. Minimum Time to Complete All Deliveries
// You are given two integer arrays of size 2: d = [d1, d2] and r = [r1, r2].

// Two delivery drones are tasked with completing a specific number of deliveries. Drone i must complete di deliveries.

// Each delivery takes exactly one hour and only one drone can make a delivery at any given hour.

// Additionally, both drones require recharging at specific intervals during which they cannot make deliveries. 
// Drone i must recharge every ri hours (i.e. at hours that are multiples of ri).

// Return an integer denoting the minimum total time (in hours) required to complete all deliveries.

// Example 1:
// Input: d = [3,1], r = [2,3]
// Output: 5
// Explanation:
// The first drone delivers at hours 1, 3, 5 (recharges at hours 2, 4).
// The second drone delivers at hour 2 (recharges at hour 3).

// Example 2:
// Input: d = [1,3], r = [2,2]
// Output: 7
// Explanation:
// The first drone delivers at hour 3 (recharges at hours 2, 4, 6).
// The second drone delivers at hours 1, 5, 7 (recharges at hours 2, 4, 6).

// Example 3:
// Input: d = [2,1], r = [3,4]
// Output: 3
// Explanation:
// The first drone delivers at hours 1, 2 (recharges at hour 3).
// The second drone delivers at hour 3.

// Constraints:
//     d = [d1, d2]
//     1 <= di <= 10^9
//     r = [r1, r2]
//     2 <= ri <= 3 * 10^4

import "fmt"
import "sort"

func minimumTime(d, r []int) int64 {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    d1, d2 := d[0], d[1]
    r1, r2 := r[0], r[1]
    l := lcm(r1, r2)
    // 库函数是左闭右开区间
    left := d1 + d2
    right := (d1+d2)*2 - 1
    res := left + sort.Search(right-left, func(t int) bool {
        t += left
        return d1 <= t-t/r1 && d2 <= t-t/r2 && d1+d2 <= t-t/l
    })
    return int64(res)
}

func minimumTime1(d, r []int) int64 {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    f := func(d, r int) int { return d + (d-1) / (r-1) }
    d1, d2 := d[0], d[1]
    r1, r2 := r[0], r[1]
    l := lcm(r1, r2)
    return int64(max(f(d1, r1), f(d2, r2), f(d1+d2, l)))
}


func main() {
    // Example 1:
    // Input: d = [3,1], r = [2,3]
    // Output: 5
    // Explanation:
    // The first drone delivers at hours 1, 3, 5 (recharges at hours 2, 4).
    // The second drone delivers at hour 2 (recharges at hour 3).
    fmt.Println(minimumTime([]int{3,1}, []int{2,3})) // 5
    // Example 2:
    // Input: d = [1,3], r = [2,2]
    // Output: 7
    // Explanation:
    // The first drone delivers at hour 3 (recharges at hours 2, 4, 6).
    // The second drone delivers at hours 1, 5, 7 (recharges at hours 2, 4, 6).
    fmt.Println(minimumTime([]int{1,3}, []int{2,2})) // 7
    // Example 3:
    // Input: d = [2,1], r = [3,4]
    // Output: 3
    // Explanation:
    // The first drone delivers at hours 1, 2 (recharges at hour 3).
    // The second drone delivers at hour 3.
    fmt.Println(minimumTime([]int{2,1}, []int{3,4})) // 3

    fmt.Println(minimumTime([]int{1,2}, []int{1,2})) // 5
    fmt.Println(minimumTime([]int{1,2}, []int{9,8})) // 3
    fmt.Println(minimumTime([]int{9,8}, []int{1,2})) // 33
    fmt.Println(minimumTime([]int{9,8}, []int{9,8})) // 17

    fmt.Println(minimumTime1([]int{3,1}, []int{2,3})) // 5
    fmt.Println(minimumTime1([]int{1,3}, []int{2,2})) // 7
    // fmt.Println(minimumTime1([]int{2,1}, []int{3,4})) // 3
    // fmt.Println(minimumTime1([]int{1,2}, []int{1,2})) // 5
    // fmt.Println(minimumTime1([]int{1,2}, []int{9,8})) // 3
    // fmt.Println(minimumTime1([]int{9,8}, []int{1,2})) // 33
    // fmt.Println(minimumTime1([]int{9,8}, []int{9,8})) // 17
}