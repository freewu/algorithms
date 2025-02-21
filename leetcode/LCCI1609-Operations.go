package main

// 面试题 16.09. Operations LCCI
// Write methods to implement the multiply, subtract, and divide operations for integers. 
// The results of all of these are integers. Use only the add operator.

// You should implement following methods:
//     Operations()  constructor
//     minus(a, b)  Subtraction, returns a - b
//     multiply(a, b)  Multiplication, returns a * b
//     divide(a, b)  Division, returns a / b

// Example:
// Operations operations = new Operations();
// operations.minus(1, 2); //returns -1
// operations.multiply(3, 4); //returns 12
// operations.divide(5, -2); //returns -2

// Note:
//     You can assume inputs are always valid, that is, e.g., denominator will not be 0 in division.
//     The number of calls will not exceed 1000.

import "fmt"

type Operations struct { }

func Constructor() Operations {
    return Operations{}
}

func (this *Operations) Minus(a int, b int) int {
    return a + (^b + 1)
}

func (this *Operations) Multiply(a int, b int) int {
    isNagtive := false
    if a < 0 {
        a, isNagtive = ^(a - 1), !isNagtive
    }
    if b < 0 {
        b, isNagtive = ^(b - 1), !isNagtive
    }
    res, x := 0, 1
    for x <= a {
        // 确定2的多少倍，指数级增长
        if x & a != 0 {
            res += b
        }
        b += b
        x += x
    }
    if isNagtive { return ^res + 1 }
    return res
}

func (this *Operations) Divide(a int, b int) int {
     res, isNegative := 0, false
    if a < 0 {
        a, isNegative = ^a + 1, !isNegative
    }
    if b < 0 {
        b, isNegative = ^b + 1, !isNegative
    }
    for a > 0 {
        fx := 0 // 记录上一个倍数
        cx := 1 // 记录当前倍数
        fb := b // 记录上一个b
        cb := b // 记录当前倍数的b
        for a >= cb {
            fb = cb
            cb += cb
            fx = cx
            cx += cx
        }
        res += fx
        a -= fb
    }
    if isNegative { return ^res + 1 }
    return res
}

func main() {
    // Operations operations = new Operations();
    obj := Constructor()
    // operations.minus(1, 2); //returns -1
    fmt.Println(obj.Minus(1, 2)) // -1
    fmt.Println(obj.Minus(1, -2)) // 3
    // operations.multiply(3, 4); //returns 12
    fmt.Println(obj.Multiply(3, 4)) // 12
    fmt.Println(obj.Multiply(3, -4)) // -12
    // operations.divide(5, -2); //returns -2
    fmt.Println(obj.Divide(5, -2)) // -2
    fmt.Println(obj.Divide(-5, -2)) // 2
}