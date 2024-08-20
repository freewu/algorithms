package main

// 2376. Count Special Integers
// We call a positive integer special if all of its digits are distinct.
// Given a positive integer n, return the number of special integers that belong to the interval [1, n].

// Example 1:
// Input: n = 20
// Output: 19
// Explanation: All the integers from 1 to 20, except 11, are special. Thus, there are 19 special integers.

// Example 2:
// Input: n = 5
// Output: 5
// Explanation: All the integers from 1 to 5 are special.

// Example 3:
// Input: n = 135
// Output: 110
// Explanation: There are 110 integers from 1 to 135 that are special.
// Some of the integers that are not special are: 22, 114, and 131.

// Constraints:
//     1 <= n <= 2 * 10^9

import "fmt"
import "strconv"

func countSpecialNumbers(n int) int {
    if n <= 10 {
        return n
    }
    res, set := 0, make([]bool, 10)
    var dfs func(n int, cur int, set []bool) int 
    dfs = func(n int, cur int, set []bool) int {
        res := 0
        if cur <= n {
            res++
        } else {
            return 0
        }
        for i := 0; i <= 9; i++ {
            if !set[i] {
                set[i] = true
                res += dfs(n, cur * 10 + i, set)
                set[i] = false
            }
            
        }
        return res
    }
    for i := 1; i <= 9; i++ {
        set[i] = true
        res += dfs(n, i, set)
        set[i] = false
    }
    return res
}

// 数位dp + 状态压缩
// 时间复杂度 O(10 *logN), 因为是和n的位数有关, 而位数=log(10,n)即logN
func countSpecialNumbers1(n int) int {
    s := strconv.Itoa(n) // 转为字符串的目的是,可以方便取到f[i]位置的值(如果是数字,需要计算offset,麻烦)
    m := len(s)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, 1<<10)
        for j := range memo[i] { // 重大bug!! 应该是 range memo[i] 写成range momo了, 造成只有 memo[i][0->i]被成功赋值了
            memo[i][j] = -1
        }
    }
    // isLimit:表示之前选择的数是否达是上限,如果是,则当前选择的不能超过限制, 否则,当前可以随意选
    // isNum:之前是否已经做出过选择, 在忽略前导0的情况下,前面可以不选,(比如n有5位,做选择时
    var dfs func(i int, mask int, isLimit, isNum bool) int // i:当前来到第i位, mask:状压,为1表示已经选过这个数字
    dfs = func(i int, mask int, isLimit, isNum bool) int {
        res := 0
        if i == m {
            if isNum { // 重大bug!! 全局不做决定,是非法的,必须形成数字才行(题目的范围[1,n],不允许0)
                return 1
            }
            return res
        }
        // trick!! 虽然是四个参数,但memo只需2个维度即可.(要记住一个要点, memo只是为减少重复计算存在的)
        // 首先: mask的值确定了isNum是否为true. (只要选过数字,isNum为true, mask>0)
        // 而对于isLimit,并不会有重复计算.因为isLimit为false时常态. 对于 isLimit==true的状态,每一位只会遇到一次.不会有重复计算,自然无需做缓存
        if !isLimit && isNum { // 记忆:只有可能重复计算的部分(isLimit是false, isNum是true是常态)
            if memo[i][mask] != -1 {
                v := memo[i][mask]
                memo[i][mask] = res
                return v
            }
        }
        // 之前没有做出选择,当前也可以不做出
        if !isNum {
            res += dfs(i+1, mask, false, false) // 主要!! 当前不选,肯定是上限小了(不能继承isLimit,因为开始时isLimit=true,如果是继承就错了)
        }
        // 当前位置要做出选择
        low := 0
        if !isNum {
            low = 1 // 如果之前没选过数,如果要选择,只能从1开始
        }
        high := 9
        if isLimit { // 如果之间和限制选的数相同,当前选择的上限不能超过限制
            high = int(s[i] - '0')
        }
        for d := low; d <= high; d++ { // 枚举选哪个
            if mask>>d&1 == 0 {
                res += dfs(i+1, mask|(1<<d), isLimit && d == high, true)
            }
        }
        return res
    }
    return dfs(0, 0, true, false) // 注意:开始阶段,isLimit=true,用来限制第一步做决定时,不能超过限制
}

func countSpecialNumbers2(n int) int {
    res, visited, arr := 0, make([]bool, 10), []int{}
    perm := func (n, k int) int {
        res := 1
        for ; k > 0; n, k = n-1, k-1 {
            res *= n
        }
        return res
    }
    for n = n + 1; n > 0; n /= 10 {
        arr = append([]int{n % 10}, arr...)
    }
    for i := 0; i < len(arr)-1; i++ {
        res += 9 * perm(9, i)
    }
    for i, upper := range arr {
        for digit := 0; digit < upper; digit++ {
            if (i > 0 || digit > 0) && !visited[digit] {
                res += perm(9-i, len(arr)-1-i)
            }
        }
        if visited[upper] { break }
        visited[upper] = true
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 20
    // Output: 19
    // Explanation: All the integers from 1 to 20, except 11, are special. Thus, there are 19 special integers.
    fmt.Println(countSpecialNumbers(20)) // 19
    // Example 2:
    // Input: n = 5
    // Output: 5
    // Explanation: All the integers from 1 to 5 are special.
    fmt.Println(countSpecialNumbers(5)) // 5
    // Example 3:
    // Input: n = 135
    // Output: 110
    // Explanation: There are 110 integers from 1 to 135 that are special.
    // Some of the integers that are not special are: 22, 114, and 131.
    fmt.Println(countSpecialNumbers(135)) // 110

    fmt.Println(countSpecialNumbers1(20)) // 19
    fmt.Println(countSpecialNumbers1(5)) // 5
    fmt.Println(countSpecialNumbers1(135)) // 110

    fmt.Println(countSpecialNumbers2(20)) // 19
    fmt.Println(countSpecialNumbers2(5)) // 5
    fmt.Println(countSpecialNumbers2(135)) // 110
}