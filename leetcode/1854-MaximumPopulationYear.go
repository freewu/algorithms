package main

// 1854. Maximum Population Year
// You are given a 2D integer array logs where each logs[i] = [birthi, deathi] indicates the birth and death years of the ith person.

// The population of some year x is the number of people alive during that year. 
// The ith person is counted in year x's population if x is in the inclusive range [birthi, deathi - 1]. 
// Note that the person is not counted in the year that they die.

// Return the earliest year with the maximum population.

// Example 1:
// Input: logs = [[1993,1999],[2000,2010]]
// Output: 1993
// Explanation: The maximum population is 1, and 1993 is the earliest year with this population.

// Example 2:
// Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
// Output: 1960
// Explanation: 
// The maximum population is 2, and it had happened in years 1960 and 1970.
// The earlier year between them is 1960.

// Constraints:
//     1 <= logs.length <= 100
//     1950 <= birthi < deathi <= 2050

import "fmt"

func maximumPopulation(logs [][]int) int {
    tmp, res := [101]int{}, 0
    for _, l := range logs {
        tmp[l[0] - 1950]++
        tmp[l[1] - 1950]--
    }
    for i := 1; i < 101; i++ {
        tmp[i] += tmp[i - 1]
        if tmp[i] > tmp[res] { // 找到最大的年份
            res = i 
        }
    }
    return res + 1950
}

func main() {
    // Example 1:
    // Input: logs = [[1993,1999],[2000,2010]]
    // Output: 1993
    // Explanation: The maximum population is 1, and 1993 is the earliest year with this population.
    fmt.Println(maximumPopulation([][]int{{1993,1999},{2000,2010}})) // 1993
    // Example 2:
    // Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
    // Output: 1960
    // Explanation: 
    // The maximum population is 2, and it had happened in years 1960 and 1970.
    // The earlier year between them is 1960.
    fmt.Println(maximumPopulation([][]int{{1950,1961},{1960,1971},{1970,1981}})) // 1960
}