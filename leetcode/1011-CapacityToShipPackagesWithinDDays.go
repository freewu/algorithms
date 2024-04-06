package main 

// 1011. Capacity To Ship Packages Within D Days
// A conveyor belt has packages that must be shipped from one port to another within days days.
// The ith package on the conveyor belt has a weight of weights[i]. 
// Each day, we load the ship with packages on the conveyor belt (in the order given by weights). 
// We may not load more weight than the maximum weight capacity of the ship.

// Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within days days.

// Example 1:
// Input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5
// Output: 15
// Explanation: A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
// 1st day: 1, 2, 3, 4, 5
// 2nd day: 6, 7
// 3rd day: 8
// 4th day: 9
// 5th day: 10
// Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.

// Example 2:
// Input: weights = [3,2,2,4,1,4], days = 3
// Output: 6
// Explanation: A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
// 1st day: 3, 2
// 2nd day: 2, 4
// 3rd day: 1, 4

// Example 3:
// Input: weights = [1,2,3,1,1], days = 4
// Output: 3
// Explanation:
// 1st day: 1
// 2nd day: 2
// 3rd day: 3
// 4th day: 1, 1
 
// Constraints:
//     1 <= days <= weights.length <= 5 * 10^4
//     1 <= weights[i] <= 500

import "fmt"

func shipWithinDays(weights []int, days int) int {
    // Defining a function to check if it's possible to send
    // all the packages withing a given days with some capacity
    fitsCapacity := func(cap int) bool {
        s, d := 0, 0
        for i := 0; i < len(weights) && d < days; i++ {
            w := weights[i]
            if s + w <= cap {
                s += w
            } else {
                d++
                s = w
            }
        }
        return d < days
    }
    // 得到累加值 和 最大值 
    max, sum := 0, 0
    for _, v := range weights {
        if v > max { max = v; }
        sum += v
    }
    // Using binary search from the min cap to the max cap
    // to find the answer
    for max < sum  {
        mid := (max + sum) >> 1
        if !fitsCapacity(mid) { // 判断是否合适
            max = mid + 1
        } else {
            sum = mid
        }
    }
    return max
}

func main() {
    // Explanation: A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
    // 1st day: 1, 2, 3, 4, 5
    // 2nd day: 6, 7
    // 3rd day: 8
    // 4th day: 9
    // 5th day: 10
    // Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.
    fmt.Println(shipWithinDays([]int{1,2,3,4,5,6,7,8,9,10}, 5)) // 15
    // Explanation: A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
    // 1st day: 3, 2
    // 2nd day: 2, 4
    // 3rd day: 1, 4
    fmt.Println(shipWithinDays([]int{3,2,2,4,1,4}, 3)) // 6
    // Explanation:
    // 1st day: 1
    // 2nd day: 2
    // 3rd day: 3
    // 4th day: 1, 1
    fmt.Println(shipWithinDays([]int{1,2,3,1,1}, 4)) // 3
}