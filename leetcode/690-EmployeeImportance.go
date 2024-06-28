package main

// 690. Employee Importance
// You have a data structure of employee information, including the employee's unique ID, importance value, and direct subordinates' IDs.
// You are given an array of employees employees where:
//     employees[i].id is the ID of the ith employee.
//     employees[i].importance is the importance value of the ith employee.
//     employees[i].subordinates is a list of the IDs of the direct subordinates of the ith employee.

// Given an integer id that represents an employee's ID, 
// return the total importance value of this employee and all their direct and indirect subordinates.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/31/emp1-tree.jpg" />
// Input: employees = [[1,5,[2,3]],[2,3,[]],[3,3,[]]], id = 1
// Output: 11
// Explanation: Employee 1 has an importance value of 5 and has two direct subordinates: employee 2 and employee 3.
// They both have an importance value of 3.
// Thus, the total importance value of employee 1 is 5 + 3 + 3 = 11.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/31/emp2-tree.jpg" />
// Input: employees = [[1,2,[5]],[5,-3,[]]], id = 5
// Output: -3
// Explanation: Employee 5 has an importance value of -3 and has no direct subordinates.
// Thus, the total importance value of employee 5 is -3.

// Constraints:
//     1 <= employees.length <= 2000
//     1 <= employees[i].id <= 2000
//     All employees[i].id are unique.
//     -100 <= employees[i].importance <= 100
//     One employee has at most one direct leader and may have several subordinates.
//     The IDs in employees[i].subordinates are valid IDs.

import "fmt"

type Employee struct {
    Id int
    Importance int
    Subordinates []int
}

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */
// dfs
func getImportance(employees []*Employee, id int) int {
    mp := make(map[int]*Employee)  // 使用一个数组存放员工ID和信息的映射
    for _, v := range employees { 
        mp[v.Id] = v 
    }
    var dfs func(id int) int
    dfs = func(id int) int {
        e := mp[id]
        res := e.Importance
        for _, v := range e.Subordinates { // 累加直接下属的重要值
            res += dfs(v) 
        }
        return res
    }
    return dfs(id)
}

// bfs
func getImportance1(employees []*Employee, id int) int {
    mp, res, queue := make(map[int]*Employee), 0, []*Employee{} // 使用一个数组存放员工ID和信息的映射
    for _, v := range employees { 
        mp[v.Id] = v 
    }
    queue = append(queue, mp[id])
    for len(queue) > 0 {
        e := queue[0]
        queue = queue[1:]
        res += e.Importance
        for _, v := range e.Subordinates { 
            queue = append(queue, mp[v])
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/31/emp1-tree.jpg" />
    // Input: employees = [[1,5,[2,3]],[2,3,[]],[3,3,[]]], id = 1
    // Output: 11
    // Explanation: Employee 1 has an importance value of 5 and has two direct subordinates: employee 2 and employee 3.
    // They both have an importance value of 3.
    // Thus, the total importance value of employee 1 is 5 + 3 + 3 = 11.
    employees1 := []*Employee{
        &Employee{1,5,[]int{2,3}},
        &Employee{2,3,[]int{}},
        &Employee{3,3,[]int{}},
    }
    fmt.Println(getImportance(employees1, 1)) // 11
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/31/emp2-tree.jpg" />
    // Input: employees = [[1,2,[5]],[5,-3,[]]], id = 5
    // Output: -3
    // Explanation: Employee 5 has an importance value of -3 and has no direct subordinates.
    // Thus, the total importance value of employee 5 is -3.
    employees2 := []*Employee{
        &Employee{1,2,[]int{5}},
        &Employee{5,-3,[]int{}},
    }
    fmt.Println(getImportance(employees2, 5)) // -3

    fmt.Println(getImportance1(employees1, 1)) // 11
    fmt.Println(getImportance1(employees2, 5)) // -3
}