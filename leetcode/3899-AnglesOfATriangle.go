package main

// 3899. Angles of a Triangle
// You are given a positive integer array sides of length 3.

// Determine if there exists a triangle with positive area whose three side lengths are given by the elements of sides.

// If such a triangle exists, return an array of three floating-point numbers representing its internal angles (in degrees),sorted in non-decreasing order. 
// Otherwise, return an empty array.

// Answers within 10-5 of the actual answer will be accepted.

// Example 1:
// Input: sides = [3,4,5]
// Output: [36.86990,53.13010,90.00000]
// Explanation:
// You can form a right-angled triangle with side lengths 3, 4, and 5. The internal angles of this triangle are approximately 36.869897646, 53.130102354, and 90 degrees respectively.

// Example 2:
// Input: sides = [2,4,2]
// Output: []
// Explanation:
// You cannot form a triangle with positive area using side lengths 2, 4, and 2.

// Constraints:
//     sides.length == 3
//     1 <= sides[i] <= 1000

import "fmt"
import "math"
import "slices"

// 余弦定理
func internalAngles(sides []int) []float64 {
    slices.Sort(sides)
    a, b, c := sides[0], sides[1], sides[2]
    if a + b <= c {
        return nil
    }
    const rad = 180 / math.Pi
    v1 := math.Acos(float64(b * b + c * c - a * a) / float64(b * c * 2)) * rad
    v2 := math.Acos(float64(a * a + c * c - b * b) / float64(a * c * 2)) * rad
    return []float64{v1, v2, 180 - v1 - v2} // 小边对小角
}

func main() {
    // Example 1:
    // Input: sides = [3,4,5]
    // Output: [36.86990,53.13010,90.00000]
    // Explanation:
    // You can form a right-angled triangle with side lengths 3, 4, and 5. The internal angles of this triangle are approximately 36.869897646, 53.130102354, and 90 degrees respectively.
    fmt.Println(internalAngles([]int{3,4,5})) // [36.86990,53.13010,90.00000] 
    // Example 2:
    // Input: sides = [2,4,2]
    // Output: []
    // Explanation:
    // You cannot form a triangle with positive area using side lengths 2, 4, and 2.
    fmt.Println(internalAngles([]int{2,4,2})) // []

    fmt.Println(internalAngles([]int{1,2,3})) // []
    fmt.Println(internalAngles([]int{4,5,6})) // [41.409622109270856 55.77113367218742 82.81924421854173]
    fmt.Println(internalAngles([]int{7,8,9})) // [48.189685104221404 58.41186449479883 73.39845040097975]
}