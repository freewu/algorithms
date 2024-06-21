package main

// 1282. Group the People Given the Group Size They Belong To
// There are n people that are split into some unknown number of groups. 
// Each person is labeled with a unique ID from 0 to n - 1.

// You are given an integer array groupSizes, where groupSizes[i] is the size of the group that person i is in. 
// For example, if groupSizes[1] = 3, then person 1 must be in a group of size 3.

// Return a list of groups such that each person i is in a group of size groupSizes[i].

// Each person should appear in exactly one group, and every person must be in a group. 
// If there are multiple answers, return any of them. 
// It is guaranteed that there will be at least one valid solution for the given input.

// Example 1:
// Input: groupSizes = [3,3,3,3,3,1,3]
// Output: [[5],[0,1,2],[3,4,6]]
// Explanation: 
// The first group is [5]. The size is 1, and groupSizes[5] = 1.
// The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
// The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
// Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].

// Example 2:
// Input: groupSizes = [2,1,3,3,3,2]
// Output: [[1],[0,5],[2,3,4]]
 
// Constraints:
//     groupSizes.length == n
//     1 <= n <= 500
//     1 <= groupSizes[i] <= n

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
    res, m := [][]int{}, make(map[int][]int)
    for i, v := range groupSizes {
        m[v] = append(m[v], i)
    }
    for i, v := range m {
        if len(v) == i { // 其中 groupSizes[i] 是第 i 个人所在的组的大小
            res = append(res, v)
        } else if len(v) % i == 0 {
            for j := 0; j < len(v); j += i {
                res = append(res, v[j:j+i])
            }
        }
    }
    return res
}

func groupThePeople1(groupSizes []int) [][]int {
    res, m := [][]int{}, make(map[int][]int)
    for i, v := range groupSizes {
        m[v] = append(m[v], i)
    }
    for i, v := range m {
        for j := 0; j < len(v); j += i {
            res = append(res, v[j:j+i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: groupSizes = [3,3,3,3,3,1,3]
    // Output: [[5],[0,1,2],[3,4,6]]
    // Explanation: 
    // The first group is [5]. The size is 1, and groupSizes[5] = 1.
    // The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
    // The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
    // Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
    fmt.Println(groupThePeople([]int{3,3,3,3,3,1,3})) // [[5],[0,1,2],[3,4,6]]
    // Example 2:
    // Input: groupSizes = [2,1,3,3,3,2]
    // Output: [[1],[0,5],[2,3,4]]
    fmt.Println(groupThePeople([]int{2,1,3,3,3,2})) // [[1],[0,5],[2,3,4]]

    fmt.Println(groupThePeople1([]int{3,3,3,3,3,1,3})) // [[5],[0,1,2],[3,4,6]]
    fmt.Println(groupThePeople1([]int{2,1,3,3,3,2})) // [[1],[0,5],[2,3,4]]
}