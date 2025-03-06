package main

// 面试题 16.10. Living People LCCI
// Given a list of people with their birth and death years, implement a method to compute the year with the most number of people alive. 
// You may assume that all people were born between 1900 and 2000 (inclusive). 
// If a person was alive during any portion of that year, they should be included in that year's count. 
// For example, Person (birth= 1908, death= 1909) is included in the counts for both 1908 and 1909.

// If there are more than one years that have the most number of people alive, return the smallest one.

// Example:
// Input: 
// birth = [1900, 1901, 1950]
// death = [1948, 1951, 2000]
// Output:  1901

// Note:
//     0 < birth.length == death.length <= 10000
//     birth[i] <= death[i]

import "fmt"
import "sort"

func maxAliveYear(birth []int, death []int) int {
    sort.Ints(birth)
    sort.Ints(death)
    res, count, end, mx := 0, 0, 0, 0
    for _, v := range birth {
        count++
        if v > death[end] {
            count--
            end++
        }
        if count > mx {
            res, mx = v, count
        }
    }
    return res
}

func maxAliveYear1(birth []int, death []int) int {
    diff := make([]int, 200)
    for i := range birth {
        diff[birth[i] - 1900]++
        diff[death[i] - 1900 + 1]--
    }
    for i := 1; i < 200; i++ {
        diff[i] = diff[i] + diff[i - 1]
    }
    year, mx := 0, -1 << 31
    for i := 0; i < 200; i++ {
        if diff[i] > mx {
            mx, year = diff[i], i
        }
    }
    return year + 1900
}

func main() {
    // Example:
    // Input: 
    // birth = [1900, 1901, 1950]
    // death = [1948, 1951, 2000]
    // Output:  1901
    fmt.Println(maxAliveYear([]int{1900, 1901, 1950}, []int{1948, 1951, 2000})) // 1901

    fmt.Println(maxAliveYear1([]int{1900, 1901, 1950}, []int{1948, 1951, 2000})) // 1901
}