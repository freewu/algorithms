package main

// 1086. High Five
// Given a list of the scores of different students, items, 
// where items[i] = [IDi, scorei] represents one score from a student with IDi, calculate each student's top five average.

// Return the answer as an array of pairs result, where result[j] = [IDj, topFiveAveragej] 
// represents the student with IDj and their top five average. 
// Sort result by IDj in increasing order.

// A student's top five average is calculated by taking the sum of their top five scores and dividing it by 5 using integer division.

// Example 1:
// Input: items = [[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
// Output: [[1,87],[2,88]]
// Explanation: 
// The student with ID = 1 got scores 91, 92, 60, 65, 87, and 100. Their top five average is (100 + 92 + 91 + 87 + 65) / 5 = 87.
// The student with ID = 2 got scores 93, 97, 77, 100, and 76. Their top five average is (100 + 97 + 93 + 77 + 76) / 5 = 88.6, but with integer division their average converts to 88.

// Example 2:
// Input: items = [[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100]]
// Output: [[1,100],[7,100]]

// Constraints:
//     1 <= items.length <= 1000
//     items[i].length == 2
//     1 <= IDi <= 1000
//     0 <= scorei <= 100
//     For each IDi, there will be at least five scores.

import "fmt"
import "sort"

func highFive(items [][]int) [][]int {
    mp := make(map[int][]int)
    for _, v := range items { // map[学生ID][]int{所有的学科分数}
        if mp[v[0]] == nil {
            mp[v[0]] = make([]int, 0)
        }
        mp[v[0]] = append(mp[v[0]], v[1])
    }
    res := make([][]int, 0)
    for i:= 1; i <= 1000; i++ { // 1 <= IDi <= 1000
        if mp[i] != nil {
            sort.Ints(mp[i])
            sum := 0
            for j := len(mp[i]) - 1; j > len(mp[i]) - 6; j-- {
                sum += mp[i][j]
            }
            res = append(res, []int{i, sum/5})
        }
    }
    // for i, v := range mp {
    //     if v != nil {
    //         sort.Ints(v) // 排序
    //         sum := 0
    //         for j := len(v) - 1; j > len(v) - 6; j-- { // 累加前5科成绩
    //             sum += v[j]
    //         }
    //         res = append(res, []int{i, sum / 5})
    //     }
    // }
    return res
}

func highFive1(items [][]int) [][]int {
    sort.Slice(items, func(i, j int) bool {
        if items[i][0] == items[j][0] {
            return items[i][1] > items[j][1]
        }
        return items[i][0] < items[j][0]
    })
    res, id, score, times := [][]int{}, 0, 0, 0
    for _, v := range items {
        if v[0] == id {
            score += v[1]
            times++
            if times == 5 {
                res = append(res, []int{id, score / 5})
            }
        }
        if v[0] != id {
            id = v[0]
            times = 1
            score = v[1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: items = [[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
    // Output: [[1,87],[2,88]]
    // Explanation: 
    // The student with ID = 1 got scores 91, 92, 60, 65, 87, and 100. Their top five average is (100 + 92 + 91 + 87 + 65) / 5 = 87.
    // The student with ID = 2 got scores 93, 97, 77, 100, and 76. Their top five average is (100 + 97 + 93 + 77 + 76) / 5 = 88.6, but with integer division their average converts to 88.
    fmt.Println(highFive([][]int{{1,91},{1,92},{2,93},{2,97},{1,60},{2,77},{1,65},{1,87},{1,100},{2,100},{2,76}})) // [[1,87],[2,88]]
    // Example 2:
    // Input: items = [[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100]]
    // Output: [[1,100],[7,100]]
    fmt.Println(highFive([][]int{{1,100},{7,100},{1,100},{7,100},{1,100},{7,100},{1,100},{7,100},{1,100},{7,100}})) //  [[1,100],[7,100]]
    
    fmt.Println(highFive1([][]int{{1,91},{1,92},{2,93},{2,97},{1,60},{2,77},{1,65},{1,87},{1,100},{2,100},{2,76}})) // [[1,87],[2,88]]
    fmt.Println(highFive1([][]int{{1,100},{7,100},{1,100},{7,100},{1,100},{7,100},{1,100},{7,100},{1,100},{7,100}})) //  [[1,100],[7,100]]
}