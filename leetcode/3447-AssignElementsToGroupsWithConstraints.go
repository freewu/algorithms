package main

// 3447. Assign Elements to Groups with Constraints
// You are given an integer array groups, where groups[i] represents the size of the ith group. 
// You are also given an integer array elements.

// Your task is to assign one element to each group based on the following rules:
//     1. An element j can be assigned to a group i if groups[i] is divisible by elements[j].
//     2. If there are multiple elements that can be assigned, assign the element with the smallest index j.
//     3. If no element satisfies the condition for a group, assign -1 to that group.

// Return an integer array assigned, where assigned[i] is the index of the element chosen for group i, 
// or -1 if no suitable element exists.

// Note: An element may be assigned to more than one group.

// Example 1:
// Input: groups = [8,4,3,2,4], elements = [4,2]
// Output: [0,0,-1,1,0]
// Explanation:
// elements[0] = 4 is assigned to groups 0, 1, and 4.
// elements[1] = 2 is assigned to group 3.
// Group 2 cannot be assigned any element.

// Example 2:
// Input: groups = [2,3,5,7], elements = [5,3,3]
// Output: [-1,1,0,-1]
// Explanation:
// elements[1] = 3 is assigned to group 1.
// elements[0] = 5 is assigned to group 2.
// Groups 0 and 3 cannot be assigned any element.

// Example 3:
// Input: groups = [10,21,30,41], elements = [2,1]
// Output: [0,1,0,1]
// Explanation:
// elements[0] = 2 is assigned to the groups with even values, and elements[1] = 1 is assigned to the groups with odd values.

// Constraints:
//     1 <= groups.length <= 10^5
//     1 <= elements.length <= 10^5
//     1 <= groups[i] <= 10^5
//     1 <= elements[i] <= 10^5

import "fmt"
import "math"
import "slices"

func assignElements(groups []int, elements []int) []int {
    mp := make(map[int]int, len(elements))
    res := make([]int, len(groups))
    for i, v := range elements {
        if index, ok := mp[v]; ok {
            mp[v] = min(index, i)
        } else {
            mp[v] = i
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    mn := 0 
    for i, v := range groups {
        res[i], mn = -1, 1 << 31
        for j := 1; j <= int(math.Sqrt(float64(v))); j++ {
            if v % j == 0 {
                if index, ok := mp[j]; ok {
                    mn = min(mn, index)
                }
                if j != v / j {
                    if index, ok := mp[v / j]; ok {
                        mn = min(mn, index)
                    }
                }
            }
        }
        if mn != 1 << 31 {
            res[i] = mn
        }
    }
    return res
}

func assignElements1(groups []int, elements []int) []int {
    mx := slices.Max(groups)
    target := make([]int, mx + 1)
    for i := range target {
        target[i] = -1
    }
    for i, v := range elements {
        if v > mx || target[v] >= 0 { continue }
        for j := v; j <= mx; j += v {
            if target[j] < 0{
                target[j] = i
            }
        }
    }
    res := make([]int, len(groups))
    for i, v := range groups {
        res[i] = target[v]
    }
    return res
}

func main() {
    // Example 1:
    // Input: groups = [8,4,3,2,4], elements = [4,2]
    // Output: [0,0,-1,1,0]
    // Explanation:
    // elements[0] = 4 is assigned to groups 0, 1, and 4.
    // elements[1] = 2 is assigned to group 3.
    // Group 2 cannot be assigned any element.
    fmt.Println(assignElements([]int{8,4,3,2,4}, []int{4,2})) // [0,0,-1,1,0]
    // Example 2:
    // Input: groups = [2,3,5,7], elements = [5,3,3]
    // Output: [-1,1,0,-1]
    // Explanation:
    // elements[1] = 3 is assigned to group 1.
    // elements[0] = 5 is assigned to group 2.
    // Groups 0 and 3 cannot be assigned any element.
    fmt.Println(assignElements([]int{2,3,5,7}, []int{5,3,3})) // [-1,1,0,-1]
    // Example 3:
    // Input: groups = [10,21,30,41], elements = [2,1]
    // Output: [0,1,0,1]
    // Explanation:
    // elements[0] = 2 is assigned to the groups with even values, and elements[1] = 1 is assigned to the groups with odd values.
    fmt.Println(assignElements([]int{10,21,30,41}, []int{2,1})) // [0,1,0,1]

    fmt.Println(assignElements([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [8 7 6 5 4 3 2 1 0]
    fmt.Println(assignElements([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(assignElements([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]

    fmt.Println(assignElements1([]int{8,4,3,2,4}, []int{4,2})) // [0,0,-1,1,0]
    fmt.Println(assignElements1([]int{2,3,5,7}, []int{5,3,3})) // [-1,1,0,-1]
    fmt.Println(assignElements1([]int{10,21,30,41}, []int{2,1})) // [0,1,0,1]
    fmt.Println(assignElements1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [8 7 6 5 4 3 2 1 0]
    fmt.Println(assignElements1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(assignElements1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
}