package main

// 2594. Minimum Time to Repair Cars
// You are given an integer array ranks representing the ranks of some mechanics. 
// ranksi is the rank of the ith mechanic.
// A mechanic with a rank r can repair n cars in r * n2 minutes.

// You are also given an integer cars representing the total number of cars waiting in the garage to be repaired.

// Return the minimum time taken to repair all the cars.

// Note: All the mechanics can repair the cars simultaneously.

// Example 1:
// Input: ranks = [4,2,3,1], cars = 10
// Output: 16
// Explanation: 
// - The first mechanic will repair two cars. The time required is 4 * 2 * 2 = 16 minutes.
// - The second mechanic will repair two cars. The time required is 2 * 2 * 2 = 8 minutes.
// - The third mechanic will repair two cars. The time required is 3 * 2 * 2 = 12 minutes.
// - The fourth mechanic will repair four cars. The time required is 1 * 4 * 4 = 16 minutes.
// It can be proved that the cars cannot be repaired in less than 16 minutes.​​​​​

// Example 2:
// Input: ranks = [5,1,8], cars = 6
// Output: 16
// Explanation: 
// - The first mechanic will repair one car. The time required is 5 * 1 * 1 = 5 minutes.
// - The second mechanic will repair four cars. The time required is 1 * 4 * 4 = 16 minutes.
// - The third mechanic will repair one car. The time required is 8 * 1 * 1 = 8 minutes.
// It can be proved that the cars cannot be repaired in less than 16 minutes.​​​​​

// Constraints:
//     1 <= ranks.length <= 10^5
//     1 <= ranks[i] <= 100
//     1 <= cars <= 10^6

import "fmt"
import "math"
import "slices"

func repairCars(ranks []int, cars int) int64 {
    left, right := int64(0), int64(math.MaxInt)
    for left < right {
        curr, mid := 0, (left + right) / 2 // mid is the upper bound of time that a mechnic can take
        for _, v := range ranks {
            curr += int(math.Sqrt(float64(mid / int64(v))))
        }
        if curr < cars {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func repairCars1(ranks []int, cars int) int64 {
    mx := slices.Max(ranks)
    left, right := int64(0), int64(mx) * int64(cars) * int64(cars)
    check := func(v int64) bool {
        sum := int64(0)
        for _, rank := range ranks {
            sum += int64(math.Sqrt(float64(v / int64(rank))))
            if sum >= int64(cars) {
                return true
            }
        }
        return false
    }
    for left <= right {
        mid := left + (right - left) / 2
        if check(mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: ranks = [4,2,3,1], cars = 10
    // Output: 16
    // Explanation: 
    // - The first mechanic will repair two cars. The time required is 4 * 2 * 2 = 16 minutes.
    // - The second mechanic will repair two cars. The time required is 2 * 2 * 2 = 8 minutes.
    // - The third mechanic will repair two cars. The time required is 3 * 2 * 2 = 12 minutes.
    // - The fourth mechanic will repair four cars. The time required is 1 * 4 * 4 = 16 minutes.
    // It can be proved that the cars cannot be repaired in less than 16 minutes.​​​​​
    fmt.Println(repairCars([]int{4,2,3,1}, 10)) // 16
    // Example 2:
    // Input: ranks = [5,1,8], cars = 6
    // Output: 16
    // Explanation: 
    // - The first mechanic will repair one car. The time required is 5 * 1 * 1 = 5 minutes.
    // - The second mechanic will repair four cars. The time required is 1 * 4 * 4 = 16 minutes.
    // - The third mechanic will repair one car. The time required is 8 * 1 * 1 = 8 minutes.
    // It can be proved that the cars cannot be repaired in less than 16 minutes.​​​​​
    fmt.Println(repairCars([]int{5,1,8}, 6)) // 16

    fmt.Println(repairCars([]int{1,2,3,4,5,6,7,8,9}, 6)) // 5
    fmt.Println(repairCars([]int{9,8,7,6,5,4,3,2,1}, 6)) // 5

    fmt.Println(repairCars1([]int{4,2,3,1}, 10)) // 16
    fmt.Println(repairCars1([]int{5,1,8}, 6)) // 16
    fmt.Println(repairCars1([]int{1,2,3,4,5,6,7,8,9}, 6)) // 5
    fmt.Println(repairCars1([]int{9,8,7,6,5,4,3,2,1}, 6)) // 5
}