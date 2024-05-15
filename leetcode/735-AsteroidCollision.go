package main

// 735. Asteroid Collision
// We are given an array asteroids of integers representing asteroids in a row.

// For each asteroid, the absolute value represents its size, 
// and the sign represents its direction (positive meaning right, negative meaning left). 
// Each asteroid moves at the same speed.

// Find out the state of the asteroids after all collisions. 
// If two asteroids meet, the smaller one will explode. 
// If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

// Example 1:
// Input: asteroids = [5,10,-5]
// Output: [5,10]
// Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

// Example 2:
// Input: asteroids = [8,-8]
// Output: []
// Explanation: The 8 and -8 collide exploding each other.

// Example 3:
// Input: asteroids = [10,2,-5]
// Output: [10]
// Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

// Constraints:
//     2 <= asteroids.length <= 10^4
//     -1000 <= asteroids[i] <= 1000
//     asteroids[i] != 0

import "fmt"

// stack
func asteroidCollision(asteroids []int) []int {
    res := []int{}
    for _, v := range asteroids {
        for len(res) != 0 && res[len(res)-1] > 0 && res[len(res)-1] < -v {
            res = res[:len(res)-1]
        }
        if len(res) == 0 || v > 0 || res[len(res)-1] < 0 {
            res = append(res, v)
        } else if v < 0 && res[len(res)-1] == -v {
            res = res[:len(res)-1]
        }
    }
    return res
}

func asteroidCollision1(asteroids []int) []int {
    result, i, a := make([]int, 0), 0, 0
    for i < len(asteroids) {
        a = asteroids[i]
        if len(result) == 0 || a > 0 || result[len(result)-1] < 0 {
            result = append(result, a)
            i++
            continue
        }
        if result[len(result)-1] > a * -1 {
            i++
            continue
        }
        if result[len(result)-1] == a * -1 {
            result = result[:len(result)-1]
            i++
            continue
        }
        result = result[:len(result)-1]
    }
    return result
}

func asteroidCollision2(asteroids []int) []int {
    stack := []int{}
    for _, v := range asteroids {
        if len(stack) == 0 || v > 0{
            stack = append(stack, v)
            continue
        }
        for {
            if len(stack) == 0 || stack[len(stack)-1] < 0 {
                stack = append(stack, v)
                break
            }
            if stack[len(stack)-1] == -v {
                stack = stack[:len(stack)-1]
                break
            }
            if stack[len(stack)-1] > -v {
                break
            }
            stack = stack[:len(stack)-1]
        }
    }
    return stack
}

func main() {
    // Example 1:
    // Input: asteroids = [5,10,-5]
    // Output: [5,10]
    // Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
    fmt.Println(asteroidCollision([]int{5,10,-5})) // [5,10]
    // Example 2:
    // Input: asteroids = [8,-8]
    // Output: []
    // Explanation: The 8 and -8 collide exploding each other.
    fmt.Println(asteroidCollision([]int{8,-8})) // []
    // Example 3:
    // Input: asteroids = [10,2,-5]
    // Output: [10]
    // Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
    fmt.Println(asteroidCollision([]int{10,2,-5})) // [10]

    fmt.Println(asteroidCollision1([]int{5,10,-5})) // [5,10]
    fmt.Println(asteroidCollision1([]int{8,-8})) // []
    fmt.Println(asteroidCollision1([]int{10,2,-5})) // [10]

    fmt.Println(asteroidCollision2([]int{5,10,-5})) // [5,10]
    fmt.Println(asteroidCollision2([]int{8,-8})) // []
    fmt.Println(asteroidCollision2([]int{10,2,-5})) // [10]
}