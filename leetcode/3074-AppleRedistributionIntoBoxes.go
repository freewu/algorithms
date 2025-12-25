package main

// 3074. Apple Redistribution into Boxes
// You are given an array apple of size n and an array capacity of size m.
// There are n packs where the ith pack contains apple[i] apples. 
// There are m boxes as well, and the ith box has a capacity of capacity[i] apples.
// Return the minimum number of boxes you need to select to redistribute these n packs of apples into boxes.
// Note that, apples from the same pack can be distributed into different boxes.

// Example 1:
// Input: apple = [1,3,2], capacity = [4,3,1,5,2]
// Output: 2
// Explanation: We will use boxes with capacities 4 and 5.
// It is possible to distribute the apples as the total capacity is greater than or equal to the total number of apples.

// Example 2:
// Input: apple = [5,5,5], capacity = [2,4,2,7]
// Output: 4
// Explanation: We will need to use all the boxes.
 
// Constraints:
//     1 <= n == apple.length <= 50
//     1 <= m == capacity.length <= 50
//     1 <= apple[i], capacity[i] <= 50
//     The input is generated such that it's possible to redistribute packs of apples into boxes.

import "fmt"
import "sort"

func minimumBoxes(apple []int, capacity []int) int {
    sum := 0
    for _, v := range apple { // 累加总和
        sum += v
    }
    // 容量由大到小排序
    sort.Slice(capacity, func(i, j int) bool { return capacity[i] > capacity[j] })
    for i, c := range capacity {
        sum -= c
        if sum <= 0 { // 所有苹果都装入了箱子
            return i + 1 // 0 到 i 有 i + 1 个箱子
        }
    }
    return -1
}

func minimumBoxes1(apple []int, capacity []int) int {
    apples := 0
    for i := range apple { // 累加苹果数量
        apples += apple[i]
    }
    // 箱子容量由大到小排序
    sort.Slice(capacity, func(i, j int) bool { return capacity[i] > capacity[j] })
    sum, res := 0, 0
    for i := range capacity {
        sum += capacity[i]
        res++ // 累加箱子个数
        if sum >= apples { // 装满了
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: apple = [1,3,2], capacity = [4,3,1,5,2]
    // Output: 2
    // Explanation: We will use boxes with capacities 4 and 5.
    // It is possible to distribute the apples as the total capacity is greater than or equal to the total number of apples.
    fmt.Println(minimumBoxes([]int{1,3,2},[]int{4,3,1,5,2})) // 2
    // Example 2:
    // Input: apple = [5,5,5], capacity = [2,4,2,7]
    // Output: 4
    // Explanation: We will need to use all the boxes.
    fmt.Println(minimumBoxes([]int{5,5,5},[]int{2,4,2,7})) // 4

    fmt.Println(minimumBoxes([]int{1,2,3,4,5,6,7,8,9},[]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minimumBoxes([]int{1,2,3,4,5,6,7,8,9},[]int{9,8,7,6,5,4,3,2,1})) // 9
    fmt.Println(minimumBoxes([]int{9,8,7,6,5,4,3,2,1},[]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minimumBoxes([]int{9,8,7,6,5,4,3,2,1},[]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(minimumBoxes1([]int{1,3,2},[]int{4,3,1,5,2})) // 2
    fmt.Println(minimumBoxes1([]int{5,5,5},[]int{2,4,2,7})) // 4
    fmt.Println(minimumBoxes1([]int{1,2,3,4,5,6,7,8,9},[]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minimumBoxes1([]int{1,2,3,4,5,6,7,8,9},[]int{9,8,7,6,5,4,3,2,1})) // 9
    fmt.Println(minimumBoxes1([]int{9,8,7,6,5,4,3,2,1},[]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minimumBoxes1([]int{9,8,7,6,5,4,3,2,1},[]int{9,8,7,6,5,4,3,2,1})) // 9
}