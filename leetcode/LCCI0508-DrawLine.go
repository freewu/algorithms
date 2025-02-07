package main

// 面试题 05.08. Draw Line LCCI
// A monochrome screen is stored as a single array of int, allowing 32 consecutive pixels to be stored in one int. 
// The screen has width w, where w is divisible by 32 (that is, no byte will be split across rows). 
// The height of the screen, of course, can be derived from the length of the array and the width. 
// Implement a function that draws a horizontal line from (x1, y) to (x2, y).

// Given the length of the array, the width of the array (in bit), 
// start position x1 (in bit) of the line, end position x2 (in bit) of the line and the row number y of the line, 
// return the array after drawing.

// Example1:
// Input: length = 1, w = 32, x1 = 30, x2 = 31, y = 0
// Output: [3]
// Explanation: After drawing a line from (30, 0) to (31, 0), the screen becomes [0b000000000000000000000000000000011].

// Example2:
// Input: length = 3, w = 96, x1 = 0, x2 = 95, y = 0
// Output: [-1, -1, -1]

import "fmt"

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
    res := make([]int, length)
    // 绘制从x1,y 到 x2, y的直线
    // x1 / 32,表示从第几行开始,如: x1 = 0,0
    // x2 / 32,表示到第几行结束,如 x2 = 95,x2 = 2
    for i := x1 / 32; i <= x2 / 32;i++  {
        // end 表示这一行结束位置 start表示这一行开始位置
        end, start := (i + 1) * 32 - 1, i * 32
        if end > x2 { // 是否超出位置
            end = x2 
        }
        if start < x1 { 
            start = x1 
        }
        end, start = end % 32, start % 32 // 取这一行.
        res[y * w / 32 + i] =  int(int32((1 << (end - start + 1) - 1) << (31 - end)))
    }
    return res
}

func drawLine1(length int, w int, x1 int, x2 int, y int) []int {
    if x1 > x2 {
        x1, x2 = x2, x1
    }
    tmp := make([]int32, length)
    for i := x1; i <= x2; i++ {
        tmp[y * (w >> 5) + i >> 5] |= 1 << (31 - i % 32)
    }
    res := make([]int, 0, length)
    for _, v := range tmp {
        res = append(res, int(v))
    }
    return res
}

func main() {
    // Example1:
    // Input: length = 1, w = 32, x1 = 30, x2 = 31, y = 0
    // Output: [3]
    // Explanation: After drawing a line from (30, 0) to (31, 0), the screen becomes [0b000000000000000000000000000000011].
    fmt.Println(drawLine(1, 32, 30, 31, 0)) // [3]
    // Example2:
    // Input: length = 3, w = 96, x1 = 0, x2 = 95, y = 0
    // Output: [-1, -1, -1]
    fmt.Println(drawLine(3, 96, 0, 95, 0)) // [-1, -1, -1]

    fmt.Println(drawLine1(1, 32, 30, 31, 0)) // [3]
    fmt.Println(drawLine1(3, 96, 0, 95, 0)) // [-1, -1, -1]
}