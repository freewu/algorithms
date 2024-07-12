package main

// 338. Counting Bits
// Given an integer n, return an array ans of length n + 1 such 
// that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

// Example 1:
// Input: n = 2
// Output: [0,1,1]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10

// Example 2:
// Input: n = 5
// Output: [0,1,1,2,1,2]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101

// Constraints:
//         0 <= n <= 10^5

// Follow up:
//         It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
//         Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?

import "fmt"

//   X & 1 ==1 or ==0，可以用 X & 1 判断奇偶性，X & 1 > 0 即奇数  X & 1 == 0 即偶数
//   X = X & (X-1) 清零最低位的1
//   X & -X => 得到最低位的1 
//   X &~X=>0
func countBits(n int) []int {
    bits := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        //fmt.Printf("i: %v i & (i-1): %b, %v\n", i,i & (i-1),i & (i-1))
        // X & (X-1) 清零最低位的 1
        bits[i] += bits[ i & (i-1) ] + 1
    }
    return bits
}

func countBits1(n int) []int {
    res := []int{0}
    k := 1
    for i := 1; i <= n; i++ {
        if k*2 == i {
            k = i
        }
        res = append(res, res[i-k]+1)
    }
    return res
}

func countBits2(n int) []int {
    onesCount := func(v int) int {
        res := 0
        for v != 0 {
            res++
            v = v & (v - 1)
        }
        return res
    }
    res := []int{}
    for i := 0; i <= n; i++ {
        res = append(res, onesCount(i))
    }
    return res
}

func countBits3(n int) []int {
    countOne := func (x int) int {
        res := 0
        for ; x > 0; x &= x - 1 {
            res++
        }
        return res
    }
    bits := make([]int, n + 1)
    for i := range bits {
        bits[i] = countOne(i)
    }
    return bits
}


func main() {
    fmt.Println(countBits(2)) // [0,1,1]
    fmt.Println(countBits(5)) // [0,1,1,2,1,2]
    fmt.Println(countBits(10)) // [0 1 1 2 1 2 2 3 1 2 2]
    fmt.Println(countBits(16)) // [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1]
    //fmt.Println(countBits(1024)) // 

    fmt.Println(countBits1(2)) // [0,1,1]
    fmt.Println(countBits1(5)) // [0,1,1,2,1,2]
    fmt.Println(countBits1(10)) // [0 1 1 2 1 2 2 3 1 2 2]
    fmt.Println(countBits1(16)) // [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1]
    //fmt.Println(countBits1(1024)) // 

    fmt.Println(countBits2(2)) // [0,1,1]
    fmt.Println(countBits2(5)) // [0,1,1,2,1,2]
    fmt.Println(countBits2(10)) // [0 1 1 2 1 2 2 3 1 2 2]
    fmt.Println(countBits2(16)) // [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1]
    //fmt.Println(countBits2(1024)) // 

    
    fmt.Println(countBits3(2)) // [0,1,1]
    fmt.Println(countBits3(5)) // [0,1,1,2,1,2]
    fmt.Println(countBits3(10)) // [0 1 1 2 1 2 2 3 1 2 2]
    fmt.Println(countBits3(16)) // [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1]
    //fmt.Println(countBits3(1024)) // 
}