package main

// 1476. Subrectangle Queries
// Implement the class SubrectangleQueries which receives a rows x cols rectangle as a matrix of integers in the constructor and supports two methods:
//     1. updateSubrectangle(int row1, int col1, int row2, int col2, int newValue)

// Updates all values with newValue in the subrectangle whose upper left coordinate is (row1,col1) and bottom right coordinate is (row2,col2).
//     2. getValue(int row, int col)

// Returns the current value of the coordinate (row,col) from the rectangle.

// Example 1:
// Input
// ["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue","getValue"]
// [[[[1,2,1],[4,3,4],[3,2,1],[1,1,1]]],[0,2],[0,0,3,2,5],[0,2],[3,1],[3,0,3,2,10],[3,1],[0,2]]
// Output
// [null,1,null,5,5,null,10,5]
// Explanation
// SubrectangleQueries subrectangleQueries = new SubrectangleQueries([[1,2,1],[4,3,4],[3,2,1],[1,1,1]]);  
// // The initial rectangle (4x3) looks like:
// // 1 2 1
// // 4 3 4
// // 3 2 1
// // 1 1 1
// subrectangleQueries.getValue(0, 2); // return 1
// subrectangleQueries.updateSubrectangle(0, 0, 3, 2, 5);
// // After this update the rectangle looks like:
// // 5 5 5
// // 5 5 5
// // 5 5 5
// // 5 5 5 
// subrectangleQueries.getValue(0, 2); // return 5
// subrectangleQueries.getValue(3, 1); // return 5
// subrectangleQueries.updateSubrectangle(3, 0, 3, 2, 10);
// // After this update the rectangle looks like:
// // 5   5   5
// // 5   5   5
// // 5   5   5
// // 10  10  10 
// subrectangleQueries.getValue(3, 1); // return 10
// subrectangleQueries.getValue(0, 2); // return 5

// Example 2:
// Input
// ["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue"]
// [[[[1,1,1],[2,2,2],[3,3,3]]],[0,0],[0,0,2,2,100],[0,0],[2,2],[1,1,2,2,20],[2,2]]
// Output
// [null,1,null,100,100,null,20]
// Explanation
// SubrectangleQueries subrectangleQueries = new SubrectangleQueries([[1,1,1],[2,2,2],[3,3,3]]);
// subrectangleQueries.getValue(0, 0); // return 1
// subrectangleQueries.updateSubrectangle(0, 0, 2, 2, 100);
// subrectangleQueries.getValue(0, 0); // return 100
// subrectangleQueries.getValue(2, 2); // return 100
// subrectangleQueries.updateSubrectangle(1, 1, 2, 2, 20);
// subrectangleQueries.getValue(2, 2); // return 20

// Constraints:
//     There will be at most 500 operations considering both methods: updateSubrectangle and getValue.
//     1 <= rows, cols <= 100
//     rows == rectangle.length
//     cols == rectangle[i].length
//     0 <= row1 <= row2 < rows
//     0 <= col1 <= col2 < cols
//     1 <= newValue, rectangle[i][j] <= 10^9
//     0 <= row < rows
//     0 <= col < cols

import "fmt"

type SubrectangleQueries struct {
    rect []int
    rows int
    cols int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
    rows, cols := len(rectangle), len(rectangle[0])
    arr := make([]int, 0, rows * cols)
    for _, v := range rectangle {
        arr = append(arr, v...)
    }
    return SubrectangleQueries{ arr, rows, cols}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for i := row1; i <= row2; i++ {
        for j := col1; j <= col2; j++ {
            index := (i * this.cols) + j
            this.rect[index] = newValue
        }
    }
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
    index := (this.cols * row) + col
    return this.rect[index]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */


type SubrectangleQueries1 struct {
    rectangle [][]int
    history [][5]int
}

func Constructor1(rectangle [][]int) SubrectangleQueries1 {
    return SubrectangleQueries1{ rectangle: rectangle }
}

func (s *SubrectangleQueries1) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    s.history = append(s.history, [5]int{ row1, col1, row2, col2, newValue })
}

func (s *SubrectangleQueries1) GetValue(row int, col int) int {
    for i := len(s.history)-1; i >= 0; i-- {
        h := s.history[i]
        if row >= h[0] && row <= h[2] && col >= h[1] && col <= h[3] {
            return h[4]
        }
    }
    return s.rectangle[row][col]
}

func main() {
    // Example 1:
    // Input
    // ["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue","getValue"]
    // [[[[1,2,1],[4,3,4],[3,2,1],[1,1,1]]],[0,2],[0,0,3,2,5],[0,2],[3,1],[3,0,3,2,10],[3,1],[0,2]]
    // Output
    // [null,1,null,5,5,null,10,5]
    // Explanation
    // SubrectangleQueries subrectangleQueries = new SubrectangleQueries([[1,2,1],[4,3,4],[3,2,1],[1,1,1]]);  
    // // The initial rectangle (4x3) looks like:
    // // 1 2 1
    // // 4 3 4
    // // 3 2 1
    // // 1 1 1
    obj1 := Constructor([][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}})
    fmt.Println(obj1)
    // subrectangleQueries.getValue(0, 2); // return 1
    fmt.Println(obj1.GetValue(0,2)) // 1
    // subrectangleQueries.updateSubrectangle(0, 0, 3, 2, 5);
    obj1.UpdateSubrectangle(0, 0, 3, 2, 5)
    fmt.Println(obj1)
    // // After this update the rectangle looks like:
    // // 5 5 5
    // // 5 5 5
    // // 5 5 5
    // // 5 5 5 
    // subrectangleQueries.getValue(0, 2); // return 5
    fmt.Println(obj1.GetValue(0,2)) // 5
    // subrectangleQueries.getValue(3, 1); // return 5
    fmt.Println(obj1.GetValue(3,1)) // 5
    // subrectangleQueries.updateSubrectangle(3, 0, 3, 2, 10);
    // // After this update the rectangle looks like:
    // // 5   5   5
    // // 5   5   5
    // // 5   5   5
    // // 10  10  10 
    obj1.UpdateSubrectangle(3, 0, 3, 2, 10)
    fmt.Println(obj1)
    // subrectangleQueries.getValue(3, 1); // return 10
    fmt.Println(obj1.GetValue(3,1)) // 10
    // subrectangleQueries.getValue(0, 2); // return 5
    fmt.Println(obj1.GetValue(0,2)) // 5

    // Example 2:
    // Input
    // ["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue"]
    // [[[[1,1,1],[2,2,2],[3,3,3]]],[0,0],[0,0,2,2,100],[0,0],[2,2],[1,1,2,2,20],[2,2]]
    // Output
    // [null,1,null,100,100,null,20]
    // Explanation
    // SubrectangleQueries subrectangleQueries = new SubrectangleQueries([[1,1,1],[2,2,2],[3,3,3]]);
    obj2 := Constructor([][]int{{1,1,1},{2,2,2},{3,3,3}})
    fmt.Println(obj2)
    // subrectangleQueries.getValue(0, 0); // return 1
    fmt.Println(obj2.GetValue(0,0)) // 1
    // subrectangleQueries.updateSubrectangle(0, 0, 2, 2, 100);
    obj2.UpdateSubrectangle(0, 0, 2, 2, 100)
    fmt.Println(obj2)
    // subrectangleQueries.getValue(0, 0); // return 100
    fmt.Println(obj2.GetValue(0,0)) // 100
    // subrectangleQueries.getValue(2, 2); // return 100
    fmt.Println(obj2.GetValue(2,2)) // 100
    // subrectangleQueries.updateSubrectangle(1, 1, 2, 2, 20);
    obj2.UpdateSubrectangle(1, 1, 2, 2, 20)
    fmt.Println(obj2)
    // subrectangleQueries.getValue(2, 2); // return 20
    fmt.Println(obj2.GetValue(2,2)) // 20


    obj11 := Constructor1([][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}})
    fmt.Println(obj11)
    // subrectangleQueries.getValue(0, 2); // return 1
    fmt.Println(obj11.GetValue(0,2)) // 1
    // subrectangleQueries.updateSubrectangle(0, 0, 3, 2, 5);
    obj11.UpdateSubrectangle(0, 0, 3, 2, 5)
    fmt.Println(obj11)
    // // After this update the rectangle looks like:
    // // 5 5 5
    // // 5 5 5
    // // 5 5 5
    // // 5 5 5 
    // subrectangleQueries.getValue(0, 2); // return 5
    fmt.Println(obj11.GetValue(0,2)) // 5
    // subrectangleQueries.getValue(3, 1); // return 5
    fmt.Println(obj11.GetValue(3,1)) // 5
    // subrectangleQueries.updateSubrectangle(3, 0, 3, 2, 10);
    // // After this update the rectangle looks like:
    // // 5   5   5
    // // 5   5   5
    // // 5   5   5
    // // 10  10  10 
    obj11.UpdateSubrectangle(3, 0, 3, 2, 10)
    fmt.Println(obj11)
    // subrectangleQueries.getValue(3, 1); // return 10
    fmt.Println(obj11.GetValue(3,1)) // 10
    // subrectangleQueries.getValue(0, 2); // return 5
    fmt.Println(obj11.GetValue(0,2)) // 5

    // Example 2:
    // Input
    // ["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue"]
    // [[[[1,1,1],[2,2,2],[3,3,3]]],[0,0],[0,0,2,2,100],[0,0],[2,2],[1,1,2,2,20],[2,2]]
    // Output
    // [null,1,null,100,100,null,20]
    // Explanation
    // SubrectangleQueries subrectangleQueries = new SubrectangleQueries([[1,1,1],[2,2,2],[3,3,3]]);
    obj12 := Constructor1([][]int{{1,1,1},{2,2,2},{3,3,3}})
    fmt.Println(obj12)
    // subrectangleQueries.getValue(0, 0); // return 1
    fmt.Println(obj12.GetValue(0,0)) // 1
    // subrectangleQueries.updateSubrectangle(0, 0, 2, 2, 100);
    obj12.UpdateSubrectangle(0, 0, 2, 2, 100)
    fmt.Println(obj12)
    // subrectangleQueries.getValue(0, 0); // return 100
    fmt.Println(obj12.GetValue(0,0)) // 100
    // subrectangleQueries.getValue(2, 2); // return 100
    fmt.Println(obj12.GetValue(2,2)) // 100
    // subrectangleQueries.updateSubrectangle(1, 1, 2, 2, 20);
    obj12.UpdateSubrectangle(1, 1, 2, 2, 20)
    fmt.Println(obj12)
    // subrectangleQueries.getValue(2, 2); // return 20
    fmt.Println(obj12.GetValue(2,2)) // 20
}