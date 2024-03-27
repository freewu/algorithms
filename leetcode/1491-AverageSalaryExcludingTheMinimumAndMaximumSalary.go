package main

// 1491. Average Salary Excluding the Minimum and Maximum Salary
// You are given an array of unique integers salary where salary[i] is the salary of the ith employee.
// Return the average salary of employees excluding the minimum and maximum salary. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// Input: salary = [4000,3000,1000,2000]
// Output: 2500.00000
// Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
// Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500

// Example 2:
// Input: salary = [1000,2000,3000]
// Output: 2000.00000
// Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
// Average salary excluding minimum and maximum salary is (2000) / 1 = 2000

// Constraints:
//     3 <= salary.length <= 100
//     1000 <= salary[i] <= 10^6
//     All the integers of salary are unique.

import "fmt"
import "sort"

func average(salary []int) float64 {
    sort.Ints(salary)
    sum := 0
    for i := 1; i < len(salary) - 1; i++ {
        sum += salary[i]
    }
    return float64(sum) / float64(len(salary) - 2)
}

func average1(salary []int) float64 {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    minimum, maximum, sum := salary[0],salary[0],salary[0]
    for i := 1; i < len(salary); i++ {
        sum += salary[i]
        minimum = min(minimum,salary[i])
        maximum = max(maximum,salary[i])
    }
    return float64(sum - minimum - maximum) / float64(len(salary) - 2)
}

func main() {
    // Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
    // Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500
    fmt.Println(average([]int{4000,3000,1000,2000})) // 2500.00000

    // Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
    // Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
    fmt.Println(average([]int{1000,2000,3000})) // 2000.00000

    fmt.Println(average1([]int{4000,3000,1000,2000})) // 2500.00000
    fmt.Println(average1([]int{1000,2000,3000})) // 2000.00000
}