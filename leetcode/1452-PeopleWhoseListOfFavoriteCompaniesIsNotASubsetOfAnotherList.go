package main

// 1452. People Whose List of Favorite Companies Is Not a Subset of Another List
// Given the array favoriteCompanies where favoriteCompanies[i] is the list of favorites companies for the ith person (indexed from 0).

// Return the indices of people whose list of favorite companies is not a subset of any other list of favorites companies. 
// You must return the indices in increasing order.

// Example 1:
// Input: favoriteCompanies = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
// Output: [0,1,4] 
// Explanation: 
// Person with index=2 has favoriteCompanies[2]=["google","facebook"] which is a subset of favoriteCompanies[0]=["leetcode","google","facebook"] corresponding to the person with index 0. 
// Person with index=3 has favoriteCompanies[3]=["google"] which is a subset of favoriteCompanies[0]=["leetcode","google","facebook"] and favoriteCompanies[1]=["google","microsoft"]. 
// Other lists of favorite companies are not a subset of another list, therefore, the answer is [0,1,4].

// Example 2:
// Input: favoriteCompanies = [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]
// Output: [0,1] 
// Explanation: In this case favoriteCompanies[2]=["facebook","google"] is a subset of favoriteCompanies[0]=["leetcode","google","facebook"], therefore, the answer is [0,1].

// Example 3:
// Input: favoriteCompanies = [["leetcode"],["google"],["facebook"],["amazon"]]
// Output: [0,1,2,3]

// Constraints:
//     1 <= favoriteCompanies.length <= 100
//     1 <= favoriteCompanies[i].length <= 500
//     1 <= favoriteCompanies[i][j].length <= 20
//     All strings in favoriteCompanies[i] are distinct.
//     All lists of favorite companies are distinct, that is, If we sort alphabetically each list then favoriteCompanies[i] != favoriteCompanies[j].
//     All strings consist of lowercase English letters only.

import "fmt"
import "sort"

func peopleIndexes(favoriteCompanies [][]string) []int {
    res, dist := make([]int, 0), make(map[int][]map[string]int)
    for _, arr := range favoriteCompanies { // v - []string
        mp := make(map[string]int)
        for _, v := range arr {
            mp[v]++ // v - company name str
        }
        dist[len(arr)] = append(dist[len(arr)], mp)
    }
    subset := func (d1 []string, d2 map[int][]map[string]int) bool {
        res := 0
        for k, v := range d2 {
            if k >= len(d1) {
                for _, m := range v {
                    sub := true
                    for _, s := range d1 {
                        if _, ok := m[s]; !ok {
                            sub = false
                            break
                        }
                    }
                    if sub {
                        res++
                    }
                }
            }
        }
        if res >= 2 {
            return false
        }
        return true
    }
    for k, v := range favoriteCompanies {
        if subset(v, dist) {
            res = append(res, k)
        }
    }
    return res
}

func peopleIndexes1(favoriteCompanies [][]string) []int {
    res, mp := []int{}, map[string]int{}
    total, n := 0, len(favoriteCompanies)
    arr := make([][]int, n)
    for i, s := range favoriteCompanies {
        for _, t := range s {
            if _, ok := mp[t]; !ok {
                mp[t] = total
                total++
            }
            arr[i] = append(arr[i], mp[t])
        }
        sort.Ints(arr[i])
    }
    for i := range favoriteCompanies {
        flag := true
        for j := 0; j < n; j++ {
            if len(arr[j]) < len(arr[i]) || i == j { continue }
            p, q := 0, 0
            for p < len(arr[i]) && q < len(arr[j]) {
                if arr[i][p] == arr[j][q] {
                    p++
                } else {
                    q++
                }
            }
            if p == len(arr[i]) {
                flag = false
                break
            }
        }
        if flag {
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: favoriteCompanies = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
    // Output: [0,1,4] 
    // Explanation: 
    // Person with index=2 has favoriteCompanies[2]=["google","facebook"] which is a subset of favoriteCompanies[0]=["leetcode","google","facebook"] corresponding to the person with index 0. 
    // Person with index=3 has favoriteCompanies[3]=["google"] which is a subset of favoriteCompanies[0]=["leetcode","google","facebook"] and favoriteCompanies[1]=["google","microsoft"]. 
    // Other lists of favorite companies are not a subset of another list, therefore, the answer is [0,1,4].
    fmt.Println(peopleIndexes([][]string{{"leetcode","google","facebook"},{"google","microsoft"},{"google","facebook"},{"google"},{"amazon"}})) // [0,1,4] 
    // Example 2:
    // Input: favoriteCompanies = [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]
    // Output: [0,1] 
    // Explanation: In this case favoriteCompanies[2]=["facebook","google"] is a subset of favoriteCompanies[0]=["leetcode","google","facebook"], therefore, the answer is [0,1].
    fmt.Println(peopleIndexes([][]string{{"leetcode","google","facebook"},{"leetcode","amazon"},{"facebook","google"}})) // [0,1] 
    // Example 3:
    // Input: favoriteCompanies = [["leetcode"],["google"],["facebook"],["amazon"]]
    // Output: [0,1,2,3]
    fmt.Println(peopleIndexes([][]string{{"leetcode"},{"google"},{"facebook"},{"amazon"}})) // [0,1,2,3]


    fmt.Println(peopleIndexes1([][]string{{"leetcode","google","facebook"},{"google","microsoft"},{"google","facebook"},{"google"},{"amazon"}})) // [0,1,4] 
    fmt.Println(peopleIndexes1([][]string{{"leetcode","google","facebook"},{"leetcode","amazon"},{"facebook","google"}})) // [0,1] 
    fmt.Println(peopleIndexes1([][]string{{"leetcode"},{"google"},{"facebook"},{"amazon"}})) // [0,1,2,3]
}