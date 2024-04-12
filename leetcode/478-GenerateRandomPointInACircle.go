package main

// 478. Generate Random Point in a Circle
// Given the radius and the position of the center of a circle, implement the function randPoint which generates a uniform random point inside the circle.
// Implement the Solution class:
//     Solution(double radius, double x_center, double y_center) 
//         initializes the object with the radius of the circle radius and the position of the center (x_center, y_center).
//     randPoint() 
//         returns a random point inside the circle. 
//         A point on the circumference of the circle is considered to be in the circle. 
//         The answer is returned as an array [x, y].
    
// Example 1:
// Input
// ["Solution", "randPoint", "randPoint", "randPoint"]
// [[1.0, 0.0, 0.0], [], [], []]
// Output
// [null, [-0.02493, -0.38077], [0.82314, 0.38945], [0.36572, 0.17248]]
// Explanation
// Solution solution = new Solution(1.0, 0.0, 0.0);
// solution.randPoint(); // return [-0.02493, -0.38077]
// solution.randPoint(); // return [0.82314, 0.38945]
// solution.randPoint(); // return [0.36572, 0.17248]
 
// Constraints:
//     0 < radius <= 10^8
//     -10^7 <= x_center, y_center <= 10^7
//     At most 3 * 10^4 calls will be made to randPoint.

import "fmt"
import "math/rand"

// type Solution struct {
//     r float64
//     x float64
//     y float64
// }

// func Constructor(radius float64, x_center float64, y_center float64) Solution {
//     return Solution {
//         r: radius,
//         x: x_center,
//         y: y_center,
//     }
// }

// func (this *Solution) RandPoint() []float64 {
//     return []float64 {
//         this.x + rand.Float64() * this.r,
//         this.y + rand.Float64() * this.r,
//     }
// }

type Solution struct {
    r float64
    x float64
    y float64
}

func Constructor(radius, xCenter, yCenter float64) Solution {
    return Solution{
        r: radius,
        x: xCenter,
        y: yCenter,
    }
}

func (this *Solution) RandPoint() []float64 {
    // [-1,1) 的随机数
    xFactor, yFactor := 1., 1.
    for xFactor * xFactor + yFactor * yFactor > 1 {
        xFactor = 2 * rand.Float64() - 1
        yFactor = 2 * rand.Float64() - 1
    }
    return []float64{
        this.x + this.r * xFactor,
        this.y + this.r * yFactor,
    }
}
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

func main() {
    // Solution solution = new Solution(1.0, 0.0, 0.0);
    obj := Constructor(1.0, 0.0, 0.0)
    // solution.randPoint(); // return [-0.02493, -0.38077]
    fmt.Println(obj.RandPoint())
    // solution.randPoint(); // return [0.82314, 0.38945]
    fmt.Println(obj.RandPoint())
    // solution.randPoint(); // return [0.36572, 0.17248]
    fmt.Println(obj.RandPoint())
}