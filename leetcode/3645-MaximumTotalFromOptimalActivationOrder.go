package main

// 3645. Maximum Total from Optimal Activation Order
// You are given two integer arrays value and limit, both of length n.

// Initially, all elements are inactive. You may activate them in any order.

//     1. To activate an inactive element at index i, 
//        the number of currently active elements must be strictly less than limit[i].
//     2. When you activate the element at index i, 
//        it adds value[i] to the total activation value 
//        (i.e., the sum of value[i] for all elements that have undergone activation operations).
//     3. After each activation, 
//        if the number of currently active elements becomes x, 
//        then all elements j with limit[j] <= x become permanently inactive, 
//        even if they are already active.

// Return the maximum total you can obtain by choosing the activation order optimally.

// Example 1:
// Input: value = [3,5,8], limit = [2,1,3]
// Output: 16
// Explanation:
// One optimal activation order is:
// Step	Activated i	value[i]	Active Before i	Active After i	Becomes Inactive j	Inactive Elements	Total
// 1	1	5	0	1	j = 1 as limit[1] = 1	[1]	5
// 2	0	3	0	1	-	[1]	8
// 3	2	8	1	2	j = 0 as limit[0] = 2	[1, 2]	16
// Thus, the maximum possible total is 16.

// Example 2:
// Input: value = [4,2,6], limit = [1,1,1]
// Output: 6
// Explanation:
// One optimal activation order is:
// Step	Activated i	value[i]	Active Before i	Active After i	Becomes Inactive j	Inactive Elements	Total
// 1	2	6	0	1	j = 0, 1, 2 as limit[j] = 1	[0, 1, 2]	6
// Thus, the maximum possible total is 6.

// Example 3:
// Input: value = [4,1,5,2], limit = [3,3,2,3]
// Output: 12
// Explanation:
// One optimal activation order is:​​​​​​​​​​​​​​
// Step	Activated i	value[i]	Active Before i	Active After i	Becomes Inactive j	Inactive Elements	Total
// 1	2	5	0	1	-	[ ]	5
// 2	0	4	1	2	j = 2 as limit[2] = 2	[2]	9
// 3	1	1	1	2	-	[2]	10
// 4	3	2	2	3	j = 0, 1, 3 as limit[j] = 3	[0, 1, 2, 3]	12
// Thus, the maximum possible total is 12.

// Constraints:
//     1 <= n == value.length == limit.length <= 10^5
//     1 <= value[i] <= 10^5​​​​​​​
//     1 <= limit[i] <= n

import "fmt"
import "sort"
import "slices"

func maxTotal(value []int, limit []int) int64 {
    res, pair, active := 0, make([][]int, 0), make([]int, 0)
    for i := 0; i < len(value); i++ {
        pair = append(pair, []int{value[i], limit[i]})
    }
    sort.Slice(pair, func(i, j int) bool {
        return pair[i][1] < pair[j][1] || (pair[i][1] == pair[j][1] && pair[i][0] > pair[j][0])
    })
    for i := 0; i < len(pair); i++ {
        if pair[i][1] == -1 {continue}
        res += pair[i][0]
        active = append(active, i)
        for j := i; j < len(pair); j++ {
            if len(active) >= pair[j][1] {
                pair[j][1] = -1
            } else {
                break
            }
        }
        for len(active) > 0 {
            if pair[active[0]][1] > len(active) {break}
            active = active[1:]
        }
    }
    return int64(res)
}

func maxTotal1(value []int, limit []int) int64 {
    res, n := 0, len(limit)
    groups := make([][]int, n + 1)
    for i, v := range limit {
        groups[v] = append(groups[v], value[i])
    }
    for v, row := range groups {
        slices.SortFunc(row, func(x, y int) int { return y - x })
        if len(row) > v {
            row = row[:v]
        }
        for _, x := range row {
            res += x
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: value = [3,5,8], limit = [2,1,3]
    // Output: 16
    // Explanation:
    // One optimal activation order is:
    // Step	Activated i	value[i]	Active Before i	Active After i	Becomes Inactive j	Inactive Elements	Total
    // 1	1	5	0	1	j = 1 as limit[1] = 1	[1]	5
    // 2	0	3	0	1	-	[1]	8
    // 3	2	8	1	2	j = 0 as limit[0] = 2	[1, 2]	16
    // Thus, the maximum possible total is 16.
    fmt.Println(maxTotal([]int{3,5,8}, []int{2,1,3})) // 16
    // Example 2:
    // Input: value = [4,2,6], limit = [1,1,1]
    // Output: 6
    // Explanation:
    // One optimal activation order is:
    // Step	Activated i	value[i]	Active Before i	Active After i	Becomes Inactive j	Inactive Elements	Total
    // 1	2	6	0	1	j = 0, 1, 2 as limit[j] = 1	[0, 1, 2]	6
    // Thus, the maximum possible total is 6.
    fmt.Println(maxTotal([]int{4,2,6}, []int{1,1,1})) // 6
    // Example 3:
    // Input: value = [4,1,5,2], limit = [3,3,2,3]
    // Output: 12
    // Explanation:
    // One optimal activation order is:​​​​​​​​​​​​​​
    // Step	Activated i	value[i]	Active Before i	Active After i	Becomes Inactive j	Inactive Elements	Total
    // 1	2	5	0	1	-	[ ]	5
    // 2	0	4	1	2	j = 2 as limit[2] = 2	[2]	9
    // 3	1	1	1	2	-	[2]	10
    // 4	3	2	2	3	j = 0, 1, 3 as limit[j] = 3	[0, 1, 2, 3]	12
    // Thus, the maximum possible total is 12.
    fmt.Println(maxTotal([]int{4,1,5,2}, []int{3,3,2,3})) // 12

    fmt.Println(maxTotal([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxTotal([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 45
    fmt.Println(maxTotal([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxTotal([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxTotal1([]int{3,5,8}, []int{2,1,3})) // 16
    fmt.Println(maxTotal1([]int{4,2,6}, []int{1,1,1})) // 6
    fmt.Println(maxTotal1([]int{4,1,5,2}, []int{3,3,2,3})) // 12
    fmt.Println(maxTotal1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxTotal1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 45
    fmt.Println(maxTotal1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxTotal1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 45
}