package main

// 386. Lexicographical Numbers
// Given an integer n, return all the numbers in the range [1, n] sorted in lexicographical order.
// You must write an algorithm that runs in O(n) time and uses O(1) extra space. 

// Example 1:
// Input: n = 13
// Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]

// Example 2:
// Input: n = 2
// Output: [1,2]
 
// Constraints:
//     1 <= n <= 5 * 10^4

import "fmt"
import "container/list"

// stack
func lexicalOrder(n int) []int {
    res := []int{}
    for i := 1; len(res) < n && i <= 9; i++ {
        stack := list.New()
        stack.PushBack(i)
        for stack.Len() > 0 {
            current := stack.Remove(stack.Back()).(int)
            if current <= n {
                res = append(res, current)
                for j := 9; j >= 0; j-- {
                    stack.PushBack(current * 10 + j)
                }
            }
        }
    }
    return res
}

// dfs
func lexicalOrder1(n int) []int {
	res := make([]int, 0, n)
    var dfs func(x, n int, res *[]int) 
    dfs = func(x, n int, res *[]int) {
        limit := (x + 10) / 10 * 10
        for x <= n && x < limit {
            *res = append(*res, x)
            if x*10 <= n {
                dfs(x*10, n, res)
            }
            x++
        }
    }
    dfs(1, n, &res)
    return res
}

func lexicalOrder2(n int) []int {
    res := []int{}
    var dfs func(cur, limit int)
    dfs = func(cur, limit int) {
        if cur > limit { return; }
        res = append(res, cur)
        for i := 0; i <= 9; i++ {
            if cur * 10 + i > limit {
                return
            }
            dfs(cur*10 + i, limit)
        }
    }
    for i := 1; i <= 9; i++ {
        dfs(i, n)
    }
    return res
}

func lexicalOrder3(n int) []int {
    res, num  := make([]int, n), 1
    for i := 0; i < len(res); i++ {
        res[i] = num
        if num * 10 <= n {
            num = num * 10
        } else {
            for num % 10 == 9 || num + 1 > n {
                num = num / 10
            }
            num++
        }
    }
    return res
}

func main() {
    fmt.Println(lexicalOrder(13)) // [1,10,11,12,13,2,3,4,5,6,7,8,9]
    fmt.Println(lexicalOrder(1)) // [1,2]

    fmt.Println(lexicalOrder1(13)) // [1,10,11,12,13,2,3,4,5,6,7,8,9]
    fmt.Println(lexicalOrder1(1)) // [1,2]

    fmt.Println(lexicalOrder2(13)) // [1,10,11,12,13,2,3,4,5,6,7,8,9]
    fmt.Println(lexicalOrder2(1)) // [1,2]

    fmt.Println(lexicalOrder3(13)) // [1,10,11,12,13,2,3,4,5,6,7,8,9]
    fmt.Println(lexicalOrder3(1)) // [1,2]
}