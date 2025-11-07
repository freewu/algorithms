package main

// 2528. Maximize the Minimum Powered City
// You are given a 0-indexed integer array stations of length n, 
// where stations[i] represents the number of power stations in the ith city.

// Each power station can provide power to every city in a fixed range. 
// In other words, if the range is denoted by r, 
// then a power station at city i can provide power to all cities j such that |i - j| <= r and 0 <= i, j <= n - 1.
//     Note that |x| denotes absolute value. 
//     For example, |7 - 5| = 2 and |3 - 10| = 7.

// The power of a city is the total number of power stations it is being provided power from.

// The government has sanctioned building k more power stations, 
// each of which can be built in any city, and have the same range as the pre-existing ones.

// Given the two integers r and k, 
// return the maximum possible minimum power of a city, if the additional power stations are built optimally.

// Note that you can build the k power stations in multiple cities.

// Example 1:
// Input: stations = [1,2,4,5,0], r = 1, k = 2
// Output: 5
// Explanation: 
// One of the optimal ways is to install both the power stations at city 1. 
// So stations will become [1,4,4,5,0].
// - City 0 is provided by 1 + 4 = 5 power stations.
// - City 1 is provided by 1 + 4 + 4 = 9 power stations.
// - City 2 is provided by 4 + 4 + 5 = 13 power stations.
// - City 3 is provided by 5 + 4 = 9 power stations.
// - City 4 is provided by 5 + 0 = 5 power stations.
// So the minimum power of a city is 5.
// Since it is not possible to obtain a larger power, we return 5.

// Example 2:
// Input: stations = [4,4,4,4], r = 0, k = 3
// Output: 4
// Explanation: 
// It can be proved that we cannot make the minimum power of a city greater than 4.

// Constraints:
//     n == stations.length
//     1 <= n <= 10^5
//     0 <= stations[i] <= 10^5
//     0 <= r <= n - 1
//     0 <= k <= 10^9

import "fmt"
import "slices"

func maxPower(stations []int, r int, k int) int64 {
    n := len(stations)
    d, s := make([]int, n + 1), make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range stations {
        left, right := max(0, i-r), min(i + r, n - 1)
        d[left] += v
        d[right + 1] -= v
    }
    s[0] = d[0]
    for i := 1; i < n+1; i++ {
        s[i] = s[i-1] + d[i]
    }
    check := func(x, k int) bool {
        t, d := 0, make([]int, n + 1)
        for i := range stations {
            t += d[i]
            dist := x - (s[i] + t)
            if dist > 0 {
                if k < dist { return false }
                k -= dist
                j := min(i + r, n - 1)
                left, right := max(0, j - r), min(j + r, n - 1)
                d[left] += dist
                d[right + 1] -= dist
                t += dist
            }
        }
        return true
    }
    left, right := 0, 1 << 40
    for left < right {
        mid := (left + right + 1) >> 1
        if check(mid, k) {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return int64(left)
}

func maxPower1(stations []int, r int, k int) int64 {
    n, mn := len(stations), 1 << 31
    sum := make([]int, n + 1)
    for i, v := range stations {
        sum[i + 1] = sum[i] + v
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range stations {
        stations[i] = sum[min(i + r + 1, n)] - sum[max(i - r, 0)]
        mn = min(mn, stations[i])
    }
    check := func(stations []int, r int, k int, t int) bool {
        s, diff := 0, make([]int, len(stations)+1)
        for i := 0; i < len(diff)-1; i++ {
            s = s + diff[i]
            v := stations[i] + s
            if v >= t {
                continue
            } else {
                k = k - (t - v)
                if k < 0 {
                    return false
                }
                diff[min(len(stations), i+2*r+1)] = diff[min(len(stations), i+2*r+1)] - (t - v)
                s = s + (t - v)
            }
        }
        return true
    }
    left, right := slices.Min(stations), slices.Max(stations) * (r + 1) + k
    res := left
    for left <= right {
        mid := (left + right) / 2
        if check(stations, r, k, mid) {
            res, left = mid, mid + 1
        } else {
            right = mid - 1
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: stations = [1,2,4,5,0], r = 1, k = 2
    // Output: 5
    // Explanation: 
    // One of the optimal ways is to install both the power stations at city 1. 
    // So stations will become [1,4,4,5,0].
    // - City 0 is provided by 1 + 4 = 5 power stations.
    // - City 1 is provided by 1 + 4 + 4 = 9 power stations.
    // - City 2 is provided by 4 + 4 + 5 = 13 power stations.
    // - City 3 is provided by 5 + 4 = 9 power stations.
    // - City 4 is provided by 5 + 0 = 5 power stations.
    // So the minimum power of a city is 5.
    // Since it is not possible to obtain a larger power, we return 5.
    fmt.Println(maxPower([]int{1,2,4,5,0}, 1, 2)) // 5
    // Example 2:
    // Input: stations = [4,4,4,4], r = 0, k = 3
    // Output: 4
    // Explanation: 
    // It can be proved that we cannot make the minimum power of a city greater than 4.
    fmt.Println(maxPower([]int{4,4,4,4}, 0, 3)) // 4

    fmt.Println(maxPower([]int{1,2,3,4,5,6,7,8,9}, 1, 2)) // 5
    fmt.Println(maxPower([]int{9,8,7,6,5,4,3,2,1}, 1, 2)) // 5

    fmt.Println(maxPower1([]int{1,2,4,5,0}, 1, 2)) // 5
    fmt.Println(maxPower1([]int{4,4,4,4}, 0, 3)) // 4
    fmt.Println(maxPower1([]int{1,2,3,4,5,6,7,8,9}, 1, 2)) // 5
    fmt.Println(maxPower1([]int{9,8,7,6,5,4,3,2,1}, 1, 2)) // 5
}