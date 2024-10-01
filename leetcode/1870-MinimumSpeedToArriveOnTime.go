package main

// 1870. Minimum Speed to Arrive on Time
// You are given a floating-point number hour, representing the amount of time you have to reach the office. 
// To commute to the office, you must take n trains in sequential order. 

// You are also given an integer array dist of length n, where dist[i] describes the distance (in kilometers) of the ith train ride.
// Each train can only depart at an integer hour, so you may need to wait in between each train ride.
//     For example, if the 1st train ride takes 1.5 hours, 
//     you must wait for an additional 0.5 hours before you can depart on the 2nd train ride at the 2 hour mark.

// Return the minimum positive integer speed (in kilometers per hour) that all the trains must travel at for you to reach the office on time, or -1 if it is impossible to be on time.
// Tests are generated such that the answer will not exceed 107 and hour will have at most two digits after the decimal point.

// Example 1:
// Input: dist = [1,3,2], hour = 6
// Output: 1
// Explanation: At speed 1:
// - The first train ride takes 1/1 = 1 hour.
// - Since we are already at an integer hour, we depart immediately at the 1 hour mark. The second train takes 3/1 = 3 hours.
// - Since we are already at an integer hour, we depart immediately at the 4 hour mark. The third train takes 2/1 = 2 hours.
// - You will arrive at exactly the 6 hour mark.

// Example 2:
// Input: dist = [1,3,2], hour = 2.7
// Output: 3
// Explanation: At speed 3:
// - The first train ride takes 1/3 = 0.33333 hours.
// - Since we are not at an integer hour, we wait until the 1 hour mark to depart. The second train ride takes 3/3 = 1 hour.
// - Since we are already at an integer hour, we depart immediately at the 2 hour mark. The third train takes 2/3 = 0.66667 hours.
// - You will arrive at the 2.66667 hour mark.

// Example 3:
// Input: dist = [1,3,2], hour = 1.9
// Output: -1
// Explanation: It is impossible because the earliest the third train can depart is at the 2 hour mark.
 
// Constraints:
//     n == dist.length
//     1 <= n <= 10^5
//     1 <= dist[i] <= 10^5
//     1 <= hour <= 10^9
//     There will be at most two digits after the decimal point in hour.

import "fmt"
import "math"
import "sort"

func minSpeedOnTime(dist []int, hour float64) int {
    res, low, high := -1, 0, int(10e9)
    check := func(dist []int, mid int) float64 {
        res := float64(0)
        for i:= 0; i < len(dist); i++ {
            if i == len(dist) - 1 {
                res += float64(dist[i]) / float64(mid)
            } else {
                res += math.Ceil(float64(dist[i]) / float64(mid))
            }
        }
        return res
    }
    for low <= high {
        mid := low + (high - low) / 2
        if check(dist, mid) <= hour {
            high = mid - 1
            res = mid
        } else {
            low = mid + 1
        }
    }
    return res
}

func minSpeedOnTime1(dist []int, hour float64) int {
    n, mx := len(dist), -1
    if hour <= float64(n) - 1 { return -1 }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ { // 找出最大距离
        mx = max(mx, dist[i])
    }
    valid := func(speed int) bool {
        time := 0.0
        for i := 0; i < n - 1; i++ {
            time += math.Ceil(float64(dist[i]) / float64(speed))
        }
        return time+(float64(dist[n-1])/float64(speed)) <= hour
    }
    final := math.Ceil(float64(dist[n-1]) / (hour - float64(n-1)))
    return sort.Search(max(int(final), mx), valid)
}

func main() {
    // Explanation: At speed 1:
    // - The first train ride takes 1/1 = 1 hour.
    // - Since we are already at an integer hour, we depart immediately at the 1 hour mark. The second train takes 3/1 = 3 hours.
    // - Since we are already at an integer hour, we depart immediately at the 4 hour mark. The third train takes 2/1 = 2 hours.
    // - You will arrive at exactly the 6 hour mark.
    fmt.Println(minSpeedOnTime([]int{1,3,2}, 6)) // 1
    // Explanation: At speed 3:
    // - The first train ride takes 1/3 = 0.33333 hours.
    // - Since we are not at an integer hour, we wait until the 1 hour mark to depart. The second train ride takes 3/3 = 1 hour.
    // - Since we are already at an integer hour, we depart immediately at the 2 hour mark. The third train takes 2/3 = 0.66667 hours.
    // - You will arrive at the 2.66667 hour mark.
    fmt.Println(minSpeedOnTime([]int{1,3,2}, 2.7)) // 3
    // Explanation: It is impossible because the earliest the third train can depart is at the 2 hour mark.
    fmt.Println(minSpeedOnTime([]int{1,3,2}, 1.9)) // -1

    fmt.Println(minSpeedOnTime1([]int{1,3,2}, 6)) // 1
    fmt.Println(minSpeedOnTime1([]int{1,3,2}, 2.7)) // 3
    fmt.Println(minSpeedOnTime1([]int{1,3,2}, 1.9)) // -1
}
