package main

// 3391. Design a 3D Binary Matrix with Efficient Layer Tracking
// You are given a n x n x n binary 3D array matrix.

// Implement the Matrix3D class:
//     Matrix3D(int n) 
//         Initializes the object with the 3D binary array matrix, where all elements are initially set to 0.
//     void setCell(int x, int y, int z) 
//         Sets the value at matrix[x][y][z] to 1.
//     void unsetCell(int x, int y, int z) 
//         Sets the value at matrix[x][y][z] to 0.
//     int largestMatrix() 
//         Returns the index x where matrix[x] contains the most number of 1's. 
//         If there are multiple such indices, return the largest x.

// Example 1:
// Input:
// ["Matrix3D", "setCell", "largestMatrix", "setCell", "largestMatrix", "setCell", "largestMatrix"]
// [[3], [0, 0, 0], [], [1, 1, 2], [], [0, 0, 1], []]
// Output:
// [null, null, 0, null, 1, null, 0]
// Explanation
// Matrix3D matrix3D = new Matrix3D(3); // Initializes a 3 x 3 x 3 3D array matrix, filled with all 0's.
// matrix3D.setCell(0, 0, 0); // Sets matrix[0][0][0] to 1.
// matrix3D.largestMatrix(); // Returns 0. matrix[0] has the most number of 1's.
// matrix3D.setCell(1, 1, 2); // Sets matrix[1][1][2] to 1.
// matrix3D.largestMatrix(); // Returns 1. matrix[0] and matrix[1] tie with the most number of 1's, but index 1 is bigger.
// matrix3D.setCell(0, 0, 1); // Sets matrix[0][0][1] to 1.
// matrix3D.largestMatrix(); // Returns 0. matrix[0] has the most number of 1's.

// Example 2:
// Input:
// ["Matrix3D", "setCell", "largestMatrix", "unsetCell", "largestMatrix"]
// [[4], [2, 1, 1], [], [2, 1, 1], []]
// Output:
// [null, null, 2, null, 3]
// Explanation
// Matrix3D matrix3D = new Matrix3D(4); // Initializes a 4 x 4 x 4 3D array matrix, filled with all 0's.
// matrix3D.setCell(2, 1, 1); // Sets matrix[2][1][1] to 1.
// matrix3D.largestMatrix(); // Returns 2. matrix[2] has the most number of 1's.
// matrix3D.unsetCell(2, 1, 1); // Sets matrix[2][1][1] to 0.
// matrix3D.largestMatrix(); // Returns 3. All indices from 0 to 3 tie with the same number of 1's, but index 3 is the biggest.

// Constraints:
//     1 <= n <= 100
//     0 <= x, y, z < n
//     At most 10^5 calls are made in total to setCell and unsetCell.
//     At most 10^4 calls are made to largestMatrix.

import "fmt"

type Matrix3D struct {
    n int
    matrix [][][]int
    sums []int
}

func Constructor(n int) Matrix3D {
    matrix := make([][][]int, n)
    for i := range matrix {
        matrix[i] = make([][]int, n)
        for j := range matrix[i] {
            matrix[i][j] = make([]int, n)
        }
    }
    return Matrix3D{ n: n, matrix: matrix, sums: make([]int, n) }
}

func (this *Matrix3D) SetCell(x int, y int, z int)  {
    this.update(x, y, z, 1)
}

func (this *Matrix3D) UnsetCell(x int, y int, z int)  {
    this.update(x, y, z, 0)
}

func (this *Matrix3D) LargestMatrix() int {
    res := this.n - 1
    for i := this.n - 2; i >= 0; i-- {
        if this.sums[i] > this.sums[res] {
            res = i
        }
    }
    return res
}

func (this *Matrix3D) update(x int, y int, z int, val int) {
    if val != this.matrix[x][y][z] {
        this.sums[x] -= this.matrix[x][y][z]
        this.matrix[x][y][z] = val
        this.sums[x] += this.matrix[x][y][z]
    }
}


type Matrix3D1 struct {
    count []int
    matrix [][][]int
}

func Constructor1(n int) Matrix3D1 {
    matrix := make([][][]int, n)
    for i := 0; i < n; i++ {
        matrix[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            matrix[i][j] = make([]int, n)
        }
    }
    return Matrix3D1{ count: make([]int, n), matrix: matrix }
}

func (this *Matrix3D1) SetCell(x int, y int, z int) {
    if this.matrix[x][y][z] == 0 {
        this.count[x]++
        this.matrix[x][y][z] = 1
    }
}

func (this *Matrix3D1) UnsetCell(x int, y int, z int) {
    if this.matrix[x][y][z] == 1 {
        this.count[x]--
        this.matrix[x][y][z] = 0
    }
}

func (this *Matrix3D1) LargestMatrix() int {
    res, c := -1,  0
    for i := 0; i < len(this.count); i++ {
        if c <= this.count[i] {
            c, res = this.count[i], i
        }
    }
    return res
}

/**
 * Your Matrix3D object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.SetCell(x,y,z);
 * obj.UnsetCell(x,y,z);
 * param_3 := obj.LargestMatrix();
 */

func main() {
    // Example 1:
    // Input:
    // ["Matrix3D", "setCell", "largestMatrix", "setCell", "largestMatrix", "setCell", "largestMatrix"]
    // [[3], [0, 0, 0], [], [1, 1, 2], [], [0, 0, 1], []]
    // Output:
    // [null, null, 0, null, 1, null, 0]
    // Explanation
    // Matrix3D matrix3D = new Matrix3D(3); // Initializes a 3 x 3 x 3 3D array matrix, filled with all 0's.
    obj1 := Constructor(3)
    fmt.Println(obj1)
    // matrix3D.setCell(0, 0, 0); // Sets matrix[0][0][0] to 1.
    obj1.SetCell(0, 0, 0)
    fmt.Println(obj1)
    // matrix3D.largestMatrix(); // Returns 0. matrix[0] has the most number of 1's.
    fmt.Println(obj1.LargestMatrix()) // 0
    // matrix3D.setCell(1, 1, 2); // Sets matrix[1][1][2] to 1.
    obj1.SetCell(1, 1, 2)
    fmt.Println(obj1)
    // matrix3D.largestMatrix(); // Returns 1. matrix[0] and matrix[1] tie with the most number of 1's, but index 1 is bigger.
    fmt.Println(obj1.LargestMatrix()) // 1
    // matrix3D.setCell(0, 0, 1); // Sets matrix[0][0][1] to 1.
    obj1.SetCell(0, 0, 1)
    fmt.Println(obj1)
    // matrix3D.largestMatrix(); // Returns 0. matrix[0] has the most number of 1's.
    fmt.Println(obj1.LargestMatrix()) // 0

    // Example 2:
    // Input:
    // ["Matrix3D", "setCell", "largestMatrix", "unsetCell", "largestMatrix"]
    // [[4], [2, 1, 1], [], [2, 1, 1], []]
    // Output:
    // [null, null, 2, null, 3]
    // Explanation
    // Matrix3D matrix3D = new Matrix3D(4); // Initializes a 4 x 4 x 4 3D array matrix, filled with all 0's.
    obj2 := Constructor(4)
    fmt.Println(obj2)
    // matrix3D.setCell(2, 1, 1); // Sets matrix[2][1][1] to 1.
    obj2.SetCell(2, 1, 1)
    fmt.Println(obj2)
    // matrix3D.largestMatrix(); // Returns 2. matrix[2] has the most number of 1's.
    fmt.Println(obj2.LargestMatrix()) // 2
    // matrix3D.unsetCell(2, 1, 1); // Sets matrix[2][1][1] to 0.
    obj2.UnsetCell(2, 1, 1)
    fmt.Println(obj2)
    // matrix3D.largestMatrix(); // Returns 3. All indices from 0 to 3 tie with the same number of 1's, but index 3 is the biggest.
    fmt.Println(obj2.LargestMatrix()) // 3

    obj11 := Constructor1(3)
    fmt.Println(obj11)
    // matrix3D.setCell(0, 0, 0); // Sets matrix[0][0][0] to 1.
    obj11.SetCell(0, 0, 0)
    fmt.Println(obj11)
    // matrix3D.largestMatrix(); // Returns 0. matrix[0] has the most number of 1's.
    fmt.Println(obj11.LargestMatrix()) // 0
    // matrix3D.setCell(1, 1, 2); // Sets matrix[1][1][2] to 1.
    obj11.SetCell(1, 1, 2)
    fmt.Println(obj11)
    // matrix3D.largestMatrix(); // Returns 1. matrix[0] and matrix[1] tie with the most number of 1's, but index 1 is bigger.
    fmt.Println(obj11.LargestMatrix()) // 1
    // matrix3D.setCell(0, 0, 1); // Sets matrix[0][0][1] to 1.
    obj11.SetCell(0, 0, 1)
    fmt.Println(obj11)
    // matrix3D.largestMatrix(); // Returns 0. matrix[0] has the most number of 1's.
    fmt.Println(obj11.LargestMatrix()) // 0

    obj12 := Constructor1(4)
    fmt.Println(obj12)
    // matrix3D.setCell(2, 1, 1); // Sets matrix[2][1][1] to 1.
    obj12.SetCell(2, 1, 1)
    fmt.Println(obj12)
    // matrix3D.largestMatrix(); // Returns 2. matrix[2] has the most number of 1's.
    fmt.Println(obj12.LargestMatrix()) // 2
    // matrix3D.unsetCell(2, 1, 1); // Sets matrix[2][1][1] to 0.
    obj12.UnsetCell(2, 1, 1)
    fmt.Println(obj12)
    // matrix3D.largestMatrix(); // Returns 3. All indices from 0 to 3 tie with the same number of 1's, but index 3 is the biggest.
    fmt.Println(obj12.LargestMatrix()) // 3
}