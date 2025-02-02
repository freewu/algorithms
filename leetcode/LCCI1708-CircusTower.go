package main

// 面试题 17.08. Circus Tower LCCI
// A circus is designing a tower routine consisting of people standing atop one anoth­er's shoulders. 
// For practical and aesthetic reasons, each person must be both shorter and lighter than the person below him or her. 
// Given the heights and weights of each person in the circus, write a method to compute the largest possible number of people in such a tower.

// Example:
// Input: height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
// Output: 6
// Explanation: The longest tower is length 6 and includes from top to bottom: (56,90), (60,95), (65,100), (68,110), (70,150), (75,190)

// Note:
//     height.length == weight.length <= 10000

import "fmt"
import "sort"

func bestSeqAtIndex(height []int, weight []int) int {
    n := len(height)
    if n == 0 { return 0 }
    type Person struct { height, weight int }
    persons := make([]Person, n)
    for i := range persons {
        persons[i] = Person{ height[i],weight[i]}
    }
    sort.Slice(persons, func(i, j int) bool { // 身高高的在前边，身高相等则体重轻的在前边
        if persons[i].height == persons[j].height { return persons[i].weight < persons[j].weight }
        return persons[i].height > persons[j].height
    })
    res := []Person{}
    for _, p := range persons {
        j := sort.Search(len(res), func(i int) bool { // 在结果中找到第一个 p 不能叠在上面的人, 二分法
            return res[i].height <= p.height || res[i].weight <= p.weight
        })
        if j == len(res) {
            res = append(res, p)
        } else {
            res[j] = p // 将第 j 个人替换成 p
        }
    }
    return len(res)
}

func main() {
    // Example:
    // Input: height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
    // Output: 6
    // Explanation: The longest tower is length 6 and includes from top to bottom: (56,90), (60,95), (65,100), (68,110), (70,150), (75,190)
    fmt.Println(bestSeqAtIndex([]int{65,70,56,75,60,68}, []int{100,150,90,190,95,110})) // 6

    fmt.Println(bestSeqAtIndex([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 9
}