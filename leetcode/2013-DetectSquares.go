package main

// 2013. Detect Squares
// You are given a stream of points on the X-Y plane. 
// Design an algorithm that:
//     1. Adds new points from the stream into a data structure. 
//        Duplicate points are allowed and should be treated as different points.
//     2. Given a query point, counts the number of ways to choose three points from the data structure 
//        such that the three points and the query point form an axis-aligned square with positive area.

// An axis-aligned square is a square whose edges are all the same length 
// and are either parallel or perpendicular to the x-axis and y-axis.

// Implement the DetectSquares class:
//     DetectSquares() 
//         Initializes the object with an empty data structure.
//     void add(int[] point) 
//         Adds a new point point = [x, y] to the data structure.
//     int count(int[] point) 
//         Counts the number of ways to form axis-aligned squares with point point = [x, y] as described above.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/01/image.png" />
// Input
// ["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
// [[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
// Output
// [null, null, null, null, 1, 0, null, 2]
// Explanation
// DetectSquares detectSquares = new DetectSquares();
// detectSquares.add([3, 10]);
// detectSquares.add([11, 2]);
// detectSquares.add([3, 2]);
// detectSquares.count([11, 10]); // return 1. You can choose:
//                                //   - The first, second, and third points
// detectSquares.count([14, 8]);  // return 0. The query point cannot form a square with any points in the data structure.
// detectSquares.add([11, 2]);    // Adding duplicate points is allowed.
// detectSquares.count([11, 10]); // return 2. You can choose:
//                                //   - The first, second, and third points
//                                //   - The first, third, and fourth points

// Constraints:
//     point.length == 2
//     0 <= x, y <= 1000
//     At most 3000 calls in total will be made to add and count.

import "fmt"

type DetectSquares struct {
    PointsCount map[[2]int]int
    Points      [][]int
}

func Constructor() DetectSquares {
    return DetectSquares{ map[[2]int]int{}, [][]int{}, }
}

func (this *DetectSquares) Add(point []int)  {
    key := [2]int{point[0], point[1]}
    this.PointsCount[key] += 1
    this.Points = append(this.Points, point)
}

func (this *DetectSquares) Count(point []int) int {
    res, px, py:= 0, point[0],point[1]
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, p := range this.Points { 
        if abs(px - p[0]) != abs(py-p[1]) || px == p[0] || py == p[1] { continue } // 判读是否正方形(两条边长是否一样) 且不为 0
        res += (this.PointsCount[[2]int{px, p[1]}] * this.PointsCount[[2]int{p[0], py}])
    }
    return res
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/01/image.png" />
    // Input
    // ["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
    // [[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
    // Output
    // [null, null, null, null, 1, 0, null, 2]
    // Explanation
    // DetectSquares detectSquares = new DetectSquares();
    obj := Constructor()
    fmt.Println(obj)
    // detectSquares.add([3, 10]);
    obj.Add([]int{3, 10})
    fmt.Println(obj)
    // detectSquares.add([11, 2]);
    obj.Add([]int{11, 2})
    fmt.Println(obj)
    // detectSquares.add([3, 2]);
    obj.Add([]int{3, 2})
    fmt.Println(obj)
    // detectSquares.count([11, 10]); // return 1. You can choose:
    //                                //   - The first, second, and third points
    fmt.Println(obj.Count([]int{11, 10})) // 1
    // detectSquares.count([14, 8]);  // return 0. The query point cannot form a square with any points in the data structure.
    fmt.Println(obj.Count([]int{14, 8})) // 0
    // detectSquares.add([11, 2]);    // Adding duplicate points is allowed.
    obj.Add([]int{11, 2})
    fmt.Println(obj)
    // detectSquares.count([11, 10]); // return 2. You can choose:
    //                                //   - The first, second, and third points
    //                                //   - The first, third, and fourth points
    fmt.Println(obj.Count([]int{11, 10})) // 2
}