package main

// 412. Fizz Buzz
// Given an integer n, return a string array answer (1-indexed) where:
//     answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
//     answer[i] == "Fizz" if i is divisible by 3.
//     answer[i] == "Buzz" if i is divisible by 5.
//     answer[i] == i (as a string) if none of the above conditions are true.
 
// Example 1:
// Input: n = 3
// Output: ["1","2","Fizz"]

// Example 2:
// Input: n = 5
// Output: ["1","2","Fizz","4","Buzz"]

// Example 3:
// Input: n = 15
// Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
 
// Constraints:
//     1 <= n <= 10^4

import "fmt"

func fizzBuzz(n int) []string {
    res := []string{}
    for i := 1; i <= n; i++ {
        if i % 15 == 0 { // 能被 15 整除 3 * 5
            res = append(res,"FizzBuzz")
            continue
        }
        if i % 5 == 0 { // 能被 5 整除
            res = append(res,"Buzz")
            continue
        }
        if i % 3 == 0 { // 能被 3 整除
            res = append(res,"Fizz")
            continue
        }
        res = append(res,fmt.Sprintf("%d",i))
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: ["1","2","Fizz"]
    fmt.Println(fizzBuzz(3)) // ["1","2","Fizz"]
    // Example 2:
    // Input: n = 5
    // Output: ["1","2","Fizz","4","Buzz"]
    fmt.Println(fizzBuzz(5)) // ["1","2","Fizz","4","Buzz"]
    // Example 3:
    // Input: n = 15
    // Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
    fmt.Println(fizzBuzz(15)) // ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
}