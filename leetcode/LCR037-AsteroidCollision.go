package main

// LCR 037. 行星碰撞
// 给定一个整数数组 asteroids，表示在同一行的小行星。

// 对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。
// 每一颗小行星以相同的速度移动。

// 找出碰撞后剩下的所有小行星。碰撞规则：
//     两个行星相互碰撞，较小的行星会爆炸。
//     如果两颗行星大小相同，则两颗行星都会爆炸。
//     两颗移动方向相同的行星，永远不会发生碰撞。

// 示例 1：
// 输入：asteroids = [5,10,-5]
// 输出：[5,10]
// 解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。

// 示例 2：
// 输入：asteroids = [8,-8]
// 输出：[]
// 解释：8 和 -8 碰撞后，两者都发生爆炸。

// 示例 3：
// 输入：asteroids = [10,2,-5]
// 输出：[10]
// 解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

// 示例 4：
// 输入：asteroids = [-2,-1,1,2]
// 输出：[-2,-1,1,2]
// 解释：-2 和 -1 向左移动，而 1 和 2 向右移动。 由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞。 
 
// 提示：
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
    res, i, a := make([]int, 0), 0, 0
    for i < len(asteroids) {
        a = asteroids[i]
        if len(res) == 0 || a > 0 || res[len(res)-1] < 0 {
            res = append(res, a)
            i++
            continue
        }
        if res[len(res)-1] > a * -1 {
            i++
            continue
        }
        if res[len(res)-1] == a * -1 {
            res = res[:len(res)-1]
            i++
            continue
        }
        res = res[:len(res)-1]
    }
    return res
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