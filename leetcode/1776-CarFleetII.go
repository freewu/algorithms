package main

// 1776. Car Fleet II
// There are n cars traveling at different speeds in the same direction along a one-lane road. 
// You are given an array cars of length n, where cars[i] = [positioni, speedi] represents:
//     1. positioni is the distance between the ith car and the beginning of the road in meters. 
//        It is guaranteed that positioni < positioni+1.
//     2. speedi is the initial speed of the ith car in meters per second.

// For simplicity, cars can be considered as points moving along the number line. 
// Two cars collide when they occupy the same position. 
// Once a car collides with another car, they unite and form a single car fleet. 
// The cars in the formed fleet will have the same position and the same speed, 
// which is the initial speed of the slowest car in the fleet.

// Return an array answer, where answer[i] is the time, 
// in seconds, at which the ith car collides with the next car, or -1 if the car does not collide with the next car. 
// Answers within 10^-5 of the actual answers are accepted.

// Example 1:
// Input: cars = [[1,2],[2,1],[4,3],[7,2]]
// Output: [1.00000,-1.00000,3.00000,-1.00000]
// Explanation: After exactly one second, the first car will collide with the second car, and form a car fleet with speed 1 m/s. After exactly 3 seconds, the third car will collide with the fourth car, and form a car fleet with speed 2 m/s.

// Example 2:
// Input: cars = [[3,4],[5,4],[6,3],[9,1]]
// Output: [2.00000,1.00000,1.50000,-1.00000]

// Constraints:
//     1 <= cars.length <= 10^5
//     1 <= positioni, speedi <= 10^6
//     positioni < positioni+1

import "fmt"

func getCollisionTimes(cars [][]int) []float64 {
    res := make([]float64, len(cars))
    stack := []int{}
    f := func(x int) float64 { return float64(x) }
    for i := len(cars) - 1; i >= 0; i-- {
        res[i] = -1
        curr_car := cars[i]
        for len(stack) > 0 {
            last_car_index := stack[len(stack)-1]
            last_car := cars[last_car_index]
            if curr_car[1] <= last_car[1] {
                stack = stack[:len(stack)-1] // pop
                continue
            }
            estimated_collision_time := (f(last_car[0]) - f(curr_car[0])) / (f(curr_car[1]) - f(last_car[1]))
            if res[last_car_index]  > 0 && estimated_collision_time >= res[last_car_index] {
                stack = stack[:len(stack)-1] // pop
            } else {
                res[i] = estimated_collision_time
                break
            }
        }
        stack = append(stack, i) // push
    }
    return res
}

func getCollisionTimes1(cars [][]int) []float64 {
    n := len(cars)
    res, stack := make([]float64, n), []int{}
    for i := n-1; i >= 0; i-- {
        pos, speed := cars[i][0], cars[i][1]
        size := len(stack)
        for size > 0 {
            index := stack[size - 1]
            top := cars[index]
            if speed <= top[1] {
                size--
                continue
            }
            if res[index] < 0 { break }
            dist := res[index] * float64(speed - top[1])
            if dist > float64(top[0] - pos) { break }
            size--
        }
        if size == 0 {
            res[i] = -1
        } else {
            top := cars[stack[size - 1]]
            res[i] = float64(top[0] - pos) / float64(speed - top[1])
        }
        stack = append(stack[:size], i) // pop & push
    }
    return res
}

func main() {
    // Example 1:
    // Input: cars = [[1,2],[2,1],[4,3],[7,2]]
    // Output: [1.00000,-1.00000,3.00000,-1.00000]
    // Explanation: After exactly one second, the first car will collide with the second car, and form a car fleet with speed 1 m/s. After exactly 3 seconds, the third car will collide with the fourth car, and form a car fleet with speed 2 m/s.
    fmt.Println(getCollisionTimes([][]int{{1,2},{2,1},{4,3},{7,2}})) // [1.00000,-1.00000,3.00000,-1.00000]
    // Example 2:
    // Input: cars = [[3,4],[5,4],[6,3],[9,1]]
    // Output: [2.00000,1.00000,1.50000,-1.00000]
    fmt.Println(getCollisionTimes([][]int{{3,4},{5,4},{6,3},{9,1}})) // [2.00000,1.00000,1.50000,-1.00000]

    fmt.Println(getCollisionTimes1([][]int{{1,2},{2,1},{4,3},{7,2}})) // [1.00000,-1.00000,3.00000,-1.00000]
    fmt.Println(getCollisionTimes1([][]int{{3,4},{5,4},{6,3},{9,1}})) // [2.00000,1.00000,1.50000,-1.00000]
}