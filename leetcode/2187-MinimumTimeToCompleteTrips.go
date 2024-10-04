package main

// 2187. Minimum Time to Complete Trips
// You are given an array time where time[i] denotes the time taken by the ith bus to complete one trip.

// Each bus can make multiple trips successively; that is, the next trip can start immediately after completing the current trip. 
// Also, each bus operates independently; that is, the trips of one bus do not influence the trips of any other bus.

// You are also given an integer totalTrips, which denotes the number of trips all buses should make in total. 
// Return the minimum time required for all buses to complete at least totalTrips trips.

// Example 1:
// Input: time = [1,2,3], totalTrips = 5
// Output: 3
// Explanation:
// - At time t = 1, the number of trips completed by each bus are [1,0,0]. 
//   The total number of trips completed is 1 + 0 + 0 = 1.
// - At time t = 2, the number of trips completed by each bus are [2,1,0]. 
//   The total number of trips completed is 2 + 1 + 0 = 3.
// - At time t = 3, the number of trips completed by each bus are [3,1,1]. 
//   The total number of trips completed is 3 + 1 + 1 = 5.
// So the minimum time needed for all buses to complete at least 5 trips is 3.

// Example 2:
// Input: time = [2], totalTrips = 1
// Output: 2
// Explanation:
// There is only one bus, and it will complete its first trip at t = 2.
// So the minimum time needed to complete 1 trip is 2.

// Constraints:
//     1 <= time.length <= 10^5
//     1 <= time[i], totalTrips <= 10^7

import "fmt"
import "slices"

// binary search
func minimumTime(time []int, totalTrips int) int64 {
    l, r := int64(0), int64(totalTrips) * int64(time[0])
    for l <= r {
        trips, mid := 0, (l + r) / 2
        for _, t := range time { 
            trips += int(mid / int64(t)) 
        }
        if trips >= totalTrips {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}

func minimumTime1(time []int, totalTrips int) int64 {
    mn := slices.Min(time)
    left, right := mn - 1, totalTrips * mn + 1 // 循环不变量：sum >= totalTrips 恒为 false, 循环不变量：sum >= totalTrips 恒为 true
    for left + 1 < right { // 开区间 (left, right) 不为空
        sum, mid := 0, (left + right) / 2
        for _, t := range time {
            sum += mid / t
        }
        if sum >= totalTrips {
            right = mid // 缩小二分区间为 (left, mid)
        } else {
            left = mid // 缩小二分区间为 (mid, right)
        }
    }
    return int64(right) // 最小的 true
}

func minimumTime2(time []int, totalTrips int) int64 {
    left := int64(1)
    mn := 10000000 + 1
    for _, t := range time { // 找出最小的时间
        if t < mn { 
            mn = t
        }
    }
    calculateTrips := func(time []int, spendTime int64) int64 {
        res := int64(0)
        for _, t := range(time) {
            res += spendTime/int64(t)
        }
        return res
    }
    right := int64(mn) * int64(totalTrips)
    for left + 1 < right {
        current := (left + right) / 2
        trips := calculateTrips(time, current)
        if trips < int64(totalTrips) {
            left = current + 1
        } else {
            right = current
        }
    }
    // fmt.Printf("l: %d r: %d, calcL: %d \n", left, right, calculateTrips(time, left))
    if int64(totalTrips) <= calculateTrips(time, left) {
        return left
    }
    return right
}

func main() {
    // Example 1:
    // Input: time = [1,2,3], totalTrips = 5
    // Output: 3
    // Explanation:
    // - At time t = 1, the number of trips completed by each bus are [1,0,0]. 
    //   The total number of trips completed is 1 + 0 + 0 = 1.
    // - At time t = 2, the number of trips completed by each bus are [2,1,0]. 
    //   The total number of trips completed is 2 + 1 + 0 = 3.
    // - At time t = 3, the number of trips completed by each bus are [3,1,1]. 
    //   The total number of trips completed is 3 + 1 + 1 = 5.
    // So the minimum time needed for all buses to complete at least 5 trips is 3.
    fmt.Println(minimumTime([]int{1,2,3}, 5)) // 3
    // Example 2:
    // Input: time = [2], totalTrips = 1
    // Output: 2
    // Explanation:
    // There is only one bus, and it will complete its first trip at t = 2.
    // So the minimum time needed to complete 1 trip is 2.
    fmt.Println(minimumTime([]int{2}, 1)) // 2

    fmt.Println(minimumTime1([]int{1,2,3}, 5)) // 3
    fmt.Println(minimumTime1([]int{2}, 1)) // 2

    fmt.Println(minimumTime2([]int{1,2,3}, 5)) // 3
    fmt.Println(minimumTime2([]int{2}, 1)) // 2
}