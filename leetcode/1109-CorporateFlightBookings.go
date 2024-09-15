package main

// 1109. Corporate Flight Bookings
// There are n flights that are labeled from 1 to n.

// You are given an array of flight bookings bookings, 
// where bookings[i] = [firsti, lasti, seatsi] represents a booking for flights firsti through lasti (inclusive) with seatsi seats reserved for each flight in the range.

// Return an array answer of length n, where answer[i] is the total number of seats reserved for flight i.

// Example 1:
// Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
// Output: [10,55,45,25,25]
// Explanation:
// Flight labels:        1   2   3   4   5
// Booking 1 reserved:  10  10
// Booking 2 reserved:      20  20
// Booking 3 reserved:      25  25  25  25
// Total seats:         10  55  45  25  25
// Hence, answer = [10,55,45,25,25]

// Example 2:
// Input: bookings = [[1,2,10],[2,2,15]], n = 2
// Output: [10,25]
// Explanation:
// Flight labels:        1   2
// Booking 1 reserved:  10  10
// Booking 2 reserved:      15
// Total seats:         10  25
// Hence, answer = [10,25]

// Constraints:
//     1 <= n <= 2 * 10^4
//     1 <= bookings.length <= 2 * 10^4
//     bookings[i].length == 3
//     1 <= firsti <= lasti <= n
//     1 <= seatsi <= 10^4

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
    res := make([]int, n + 1)
    for _, record := range bookings {
        start, end, value := record[0], record[1], record[2]
        for start <= end {
            res[start] += value
            start++ 
        }
    }
    return res[1:]
}

func corpFlightBookings1(bookings [][]int, n int) []int {
    diffSeat := make([]int, n + 1)
    for _, record := range bookings {
        start, end, value := record[0], record[1], record[2]
        diffSeat[start] += value
        if end != n {
            diffSeat[end+1] -= value
        }
    }
    res, diff := make([]int, n), 0
    for i := 0; i < n; i++ {
        res[i] = diff + diffSeat[i+1]
        diff = res[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
    // Output: [10,55,45,25,25]
    // Explanation:
    // Flight labels:        1   2   3   4   5
    // Booking 1 reserved:  10  10
    // Booking 2 reserved:      20  20
    // Booking 3 reserved:      25  25  25  25
    // Total seats:         10  55  45  25  25
    // Hence, answer = [10,55,45,25,25]
    fmt.Println(corpFlightBookings([][]int{{1,2,10},{2,3,20},{2,5,25}}, 5)) // [10,55,45,25,25]
    // Example 2:
    // Input: bookings = [[1,2,10],[2,2,15]], n = 2
    // Output: [10,25]
    // Explanation:
    // Flight labels:        1   2
    // Booking 1 reserved:  10  10
    // Booking 2 reserved:      15
    // Total seats:         10  25
    // Hence, answer = [10,25]
    fmt.Println(corpFlightBookings([][]int{{1,2,10},{2,2,15}}, 2)) // [10,25]

    fmt.Println(corpFlightBookings1([][]int{{1,2,10},{2,3,20},{2,5,25}}, 5)) // [10,55,45,25,25]
    fmt.Println(corpFlightBookings1([][]int{{1,2,10},{2,2,15}}, 2)) // [10,25]
}