package main

// 858. Mirror Reflection
// There is a special square room with mirrors on each of the four walls. 
// Except for the southwest corner, there are receptors on each of the remaining corners, numbered 0, 1, and 2.

// The square room has walls of length p and a laser ray from the southwest corner first meets the east wall at a distance q from the 0th receptor.
// Given the two integers p and q, return the number of the receptor that the ray meets first.
// The test cases are guaranteed so that the ray will meet a receptor eventually.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/18/reflection.png" />
// Input: p = 2, q = 1
// Output: 2
// Explanation: The ray meets receptor 2 the first time it gets reflected back to the left wall.

// Example 2:
// Input: p = 3, q = 1
// Output: 1

// Constraints:
//     1 <= q <= p <= 1000

import "fmt"

func mirrorReflection(p int, q int) int {
    n := 0
    for true {
        n += q
        if n % p == 0 {
            return n / p % 2
        }
        n += q
        if n % p == 0 && n/p % 2 == 1 {
            return 2
        }
    }
    return 0
}

func mirrorReflection1(p int, q int) int {
    // 设横线移动的距离为w,竖向为h, 要到达四个角落,需要数值移动的距离为 a*w(正方形),并且这个a*w还要是h的倍数, 则a*w=lcm
    // 既然数值方向移动的距离是lcm,那么水平方向移动的距离为 lcm/h * w = w*w / gcd, 移动的次数是  w/gcd
    // 水平方向 w/gcd如果是奇数,则移动到了右侧,偶数则在左侧
    // 数值方向 移动的距离lcm,因为正方形,移动的边角是 lcm/w = h/gcd, 如果是奇数则是1,如果是偶数则是0(镜子镜像映射上去)
    // 因为gcd的奇偶性是综合了w,h两者的奇偶性
    gcd := func(a, b int) int {
        for b != 0 {
            a, b = b, a % b
        }
        return a
    }
    g := gcd(p, q)
    if (p/g) % 2 == 0 { // 水平移动,偶数在左侧,奇数在右侧
        return 2
    }
    return (q / g) % 2 // 竖直移动,如果是奇数,就在1位置,偶数在镜像后的0位置 正好结果如果是奇数就是1,偶数就是0
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/18/reflection.png" />
    // Input: p = 2, q = 1
    // Output: 2
    // Explanation: The ray meets receptor 2 the first time it gets reflected back to the left wall.
    fmt.Println(mirrorReflection(2,1)) // 2
    // Example 2:
    // Input: p = 3, q = 1
    // Output: 1
    fmt.Println(mirrorReflection(3,1)) // 1

    fmt.Println(mirrorReflection1(2,1)) // 2
    fmt.Println(mirrorReflection1(3,1)) // 1
}