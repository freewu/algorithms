package main

// 2933. High-Access Employees
// You are given a 2D 0-indexed array of strings, access_times, with size n. 
// For each i where 0 <= i <= n - 1, access_times[i][0] represents the name of an employee, 
// and access_times[i][1] represents the access time of that employee. 
// All entries in access_times are within the same day.

// The access time is represented as four digits using a 24-hour time format, 
// for example, "0800" or "2250".

// An employee is said to be high-access if he has accessed the system three or more times within a one-hour period.

// Times with exactly one hour of difference are not considered part of the same one-hour period. 
// For example, "0815" and "0915" are not part of the same one-hour period.

// Access times at the start and end of the day are not counted within the same one-hour period. 
// For example, "0005" and "2350" are not part of the same one-hour period.

// Return a list that contains the names of high-access employees with any order you want.

// Example 1:
// Input: access_times = [["a","0549"],["b","0457"],["a","0532"],["a","0621"],["b","0540"]]
// Output: ["a"]
// Explanation: "a" has three access times in the one-hour period of [05:32, 06:31] which are 05:32, 05:49, and 06:21.
// But "b" does not have more than two access times at all.
// So the answer is ["a"].

// Example 2:
// Input: access_times = [["d","0002"],["c","0808"],["c","0829"],["e","0215"],["d","1508"],["d","1444"],["d","1410"],["c","0809"]]
// Output: ["c","d"]
// Explanation: "c" has three access times in the one-hour period of [08:08, 09:07] which are 08:08, 08:09, and 08:29.
// "d" has also three access times in the one-hour period of [14:10, 15:09] which are 14:10, 14:44, and 15:08.
// However, "e" has just one access time, so it can not be in the answer and the final answer is ["c","d"].

// Example 3:
// Input: access_times = [["cd","1025"],["ab","1025"],["cd","1046"],["cd","1055"],["ab","1124"],["ab","1120"]]
// Output: ["ab","cd"]
// Explanation: "ab" has three access times in the one-hour period of [10:25, 11:24] which are 10:25, 11:20, and 11:24.
// "cd" has also three access times in the one-hour period of [10:25, 11:24] which are 10:25, 10:46, and 10:55.
// So the answer is ["ab","cd"].

// Constraints:
//     1 <= access_times.length <= 100
//     access_times[i].length == 2
//     1 <= access_times[i][0].length <= 10
//     access_times[i][0] consists only of English small letters.
//     access_times[i][1].length == 4
//     access_times[i][1] is in 24-hour time format.
//     access_times[i][1] consists only of '0' to '9'.

import "fmt"
import "sort"
import "strconv"

func findHighAccessEmployees(access_times [][]string) []string {
    sort.Slice(access_times, func(i,j int)bool{
        return access_times[i][1] < access_times[j][1]
    })  
    freq, employees := make(map[string]int), make(map[string]bool)
    i, j, n := 0, 0, len(access_times)
    for i < n && j < n {
        curr, _ := strconv.Atoi(access_times[j][1])
        global, _ := strconv.Atoi(access_times[i][1])
        for curr < global + 100 {
            freq[access_times[j][0]]++ 
            if freq[access_times[j][0]] >= 3 {
                employees[access_times[j][0]] = true
            } 
            j++  
            if j >= n { break }
            curr, _ = strconv.Atoi(access_times[j][1])
        }
        freq[access_times[i][0]]--
        i++
    }
    res := []string{}
    for k := range employees {
        res = append(res, k)
    }
    return res
}

func findHighAccessEmployees1(access_times [][]string) []string {
    res, groups := []string{}, make(map[string][]int)
    for _, v := range access_times {
        name, s := v[0], v[1]
        t := int(s[0] & 15 * 10 + s[1] & 15) * 60 + int(s[2] & 15 * 10 + s[3] & 15)
        groups[name] = append(groups[name], t)
    }
    for name, row := range groups {
        if len(row) < 3 { continue }
        sort.Ints(row)
        for i := 2; i < len(row); i++ {
            if row[i] - row[i - 2] < 60 {
                res = append(res, name)
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: access_times = [["a","0549"],["b","0457"],["a","0532"],["a","0621"],["b","0540"]]
    // Output: ["a"]
    // Explanation: "a" has three access times in the one-hour period of [05:32, 06:31] which are 05:32, 05:49, and 06:21.
    // But "b" does not have more than two access times at all.
    // So the answer is ["a"].
    fmt.Println(findHighAccessEmployees([][]string{{"a","0549"},{"b","0457"},{"a","0532"},{"a","0621"},{"b","0540"}})) // ["a"]
    // Example 2:
    // Input: access_times = [["d","0002"],["c","0808"],["c","0829"],["e","0215"],["d","1508"],["d","1444"],["d","1410"],["c","0809"]]
    // Output: ["c","d"]
    // Explanation: "c" has three access times in the one-hour period of [08:08, 09:07] which are 08:08, 08:09, and 08:29.
    // "d" has also three access times in the one-hour period of [14:10, 15:09] which are 14:10, 14:44, and 15:08.
    // However, "e" has just one access time, so it can not be in the answer and the final answer is ["c","d"].
    fmt.Println(findHighAccessEmployees([][]string{{"d","0002"},{"c","0808"},{"c","0829"},{"e","0215"},{"d","1508"},{"d","1444"},{"d","1410"},{"c","0809"}})) // ["c","d"]
    // Example 3:
    // Input: access_times = [["cd","1025"],["ab","1025"],["cd","1046"],["cd","1055"],["ab","1124"],["ab","1120"]]
    // Output: ["ab","cd"]
    // Explanation: "ab" has three access times in the one-hour period of [10:25, 11:24] which are 10:25, 11:20, and 11:24.
    // "cd" has also three access times in the one-hour period of [10:25, 11:24] which are 10:25, 10:46, and 10:55.
    // So the answer is ["ab","cd"].
    fmt.Println(findHighAccessEmployees([][]string{{"cd","1025"},{"ab","1025"},{"cd","1046"},{"cd","1055"},{"ab","1124"},{"ab","1120"}})) // ["ab","cd"]

    fmt.Println(findHighAccessEmployees1([][]string{{"a","0549"},{"b","0457"},{"a","0532"},{"a","0621"},{"b","0540"}})) // ["a"]
    fmt.Println(findHighAccessEmployees1([][]string{{"d","0002"},{"c","0808"},{"c","0829"},{"e","0215"},{"d","1508"},{"d","1444"},{"d","1410"},{"c","0809"}})) // ["c","d"]
    fmt.Println(findHighAccessEmployees1([][]string{{"cd","1025"},{"ab","1025"},{"cd","1046"},{"cd","1055"},{"ab","1124"},{"ab","1120"}})) // ["ab","cd"]
}