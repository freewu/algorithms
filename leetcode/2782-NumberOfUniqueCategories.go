package main

// 2782. Number of Unique Categories
// You are given an integer n and an object categoryHandler of class CategoryHandler.

// There are n elements, numbered from 0 to n - 1. 
// Each element has a category, and your task is to find the number of unique categories.

// The class CategoryHandler contains the following function, which may help you:
//     boolean haveSameCategory(integer a, integer b): 
//         Returns true if a and b are in the same category and false otherwise. 
//         Also, if either a or b is not a valid number (i.e. it's greater than or equal to nor less than 0), it returns false.

// Return the number of unique categories.

// Example 1:
// Input: n = 6, categoryHandler = [1,1,2,2,3,3]
// Output: 3
// Explanation: There are 6 elements in this example. The first two elements belong to category 1, the second two belong to category 2, and the last two elements belong to category 3. So there are 3 unique categories.

// Example 2:
// Input: n = 5, categoryHandler = [1,2,3,4,5]
// Output: 5
// Explanation: There are 5 elements in this example. Each element belongs to a unique category. So there are 5 unique categories.

// Example 3:
// Input: n = 3, categoryHandler = [1,1,1]
// Output: 1
// Explanation: There are 3 elements in this example. All of them belong to one category. So there is only 1 unique category.

// Constraints:
//     1 <= n <= 100

import "fmt"

type CategoryHandler interface {
    HaveSameCategory(int, int) bool
}

type CategoryService struct {
    data []int
}

func NewCategoryService(data []int) CategoryHandler {
    return CategoryService{ data: data }
}

// 实现接口
func (this CategoryService) HaveSameCategory(a, b int) bool {
    n := len(this.data)
    if a > n - 1 || a < 0 || b > n - 1 || b < 0  { return false }
    return this.data[a] == this.data[b]
}

/**
 * Definition for a category handler.
 * type CategoryHandler interface {
 *  HaveSameCategory(int, int) bool
 * }
 */
// 采用联通分量（QuickUnion）的方式，把满足haveSameCategory(a, b) == true的两个点用同一个id值关联起来
func numberOfCategories(n int, categoryHandler CategoryHandler) int {
    res, parent := 0, make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if categoryHandler.HaveSameCategory(i, j) {
                parent[find(i)] = find(j)
            }
        }
    }
    for i, v := range parent {
        if i == v {
            res++
        }
    }
    return res
}

func numberOfCategories1(n int, categoryHandler CategoryHandler) int {
    res, distinct := n, make([]bool, n)
    for i := 0; i < n - 1; i++ {
        if distinct[i] { continue }
        for j := i + 1; j < n; j++ {
            if categoryHandler.HaveSameCategory(i, j) {
                distinct[j] = true
                res--
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 6, categoryHandler = [1,1,2,2,3,3]
    // Output: 3
    // Explanation: There are 6 elements in this example. The first two elements belong to category 1, the second two belong to category 2, and the last two elements belong to category 3. So there are 3 unique categories.
    fmt.Println(numberOfCategories(6, NewCategoryService([]int{1,1,2,2,3,3}))) // 3
    // Example 2:
    // Input: n = 5, categoryHandler = [1,2,3,4,5]
    // Output: 5
    // Explanation: There are 5 elements in this example. Each element belongs to a unique category. So there are 5 unique categories.
    fmt.Println(numberOfCategories(5, NewCategoryService([]int{1,2,3,4,5}))) // 5
    // Example 3:
    // Input: n = 3, categoryHandler = [1,1,1]
    // Output: 1
    // Explanation: There are 3 elements in this example. All of them belong to one category. So there is only 1 unique category.
    fmt.Println(numberOfCategories(2, NewCategoryService([]int{1,1,1}))) // 1

    fmt.Println(numberOfCategories1(6, NewCategoryService([]int{1,1,2,2,3,3}))) // 3
    fmt.Println(numberOfCategories1(5, NewCategoryService([]int{1,2,3,4,5}))) // 5
    fmt.Println(numberOfCategories1(2, NewCategoryService([]int{1,1,1}))) // 1
}