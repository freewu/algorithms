package main

// 251. Flatten 2D Vector
// Design an iterator to flatten a 2D vector. It should support the next and hasNext operations.
// Implement the Vector2D class:
//     Vector2D(int[][] vec) initializes the object with the 2D vector vec.
//     next() returns the next element from the 2D vector and moves the pointer one step forward. You may assume that all the calls to next are valid.
//     hasNext() returns true if there are still some elements in the vector, and false otherwise.

// Example 1:
// Input
// ["Vector2D", "next", "next", "next", "hasNext", "hasNext", "next", "hasNext"]
// [[[[1, 2], [3], [4]]], [], [], [], [], [], [], []]
// Output
// [null, 1, 2, 3, true, true, 4, false]
// Explanation
// Vector2D vector2D = new Vector2D([[1, 2], [3], [4]]);
// vector2D.next();    // return 1
// vector2D.next();    // return 2
// vector2D.next();    // return 3
// vector2D.hasNext(); // return True
// vector2D.hasNext(); // return True
// vector2D.next();    // return 4
// vector2D.hasNext(); // return False
 
// Constraints:
//     0 <= vec.length <= 200
//     0 <= vec[i].length <= 500
//     -500 <= vec[i][j] <= 500
//     At most 10^5 calls will be made to next and hasNext.
 
// Follow up: As an added challenge, try to code it using only iterators in C++ or iterators in Java.

import "fmt"

type Vector2D struct {
    data []int // 存放展开的数组
    index int // 游标
}

func Constructor(vec [][]int) Vector2D {
    arr := []int{}
    for _, row := range vec {
        for _, v := range row {
            arr = append(arr,v)
        }
    }
    return Vector2D{ arr, 0 }
}

func (this *Vector2D) Next() int {
    if !this.HasNext() {
        return -1
    }
    val := this.data[this.index]
    this.index++
    return val
}

func (this *Vector2D) HasNext() bool {
    return this.index < len(this.data)
}


// best solution
type Vector2D1 struct {
    i, j int
    m int
    vec [][]int
}


func Constructor1(vec [][]int) Vector2D1 {
    ret := Vector2D1{m: len(vec), vec: vec}
    for ;ret.i < ret.m && ret.j == len(ret.vec[ret.i]); {
        ret.j = 0
        ret.i += 1
    }
    return ret
}


func (this *Vector2D1) Next() int {
    ret := this.vec[this.i][this.j]
    this.j += 1
    for ;this.i < this.m && this.j == len(this.vec[this.i]); {
        this.j = 0
        this.i += 1
    }

    return ret
}

func (this *Vector2D1) HasNext() bool {
    return this.i < this.m
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
    // Vector2D vector2D = new Vector2D([[1, 2], [3], [4]]);
    obj := Constructor([][]int{{1, 2},{3},{4}})
    fmt.Println(obj) // 
    // vector2D.next();    // return 1
    fmt.Println(obj.Next()) // 1
    // vector2D.next();    // return 2
    fmt.Println(obj.Next()) // 2
    // vector2D.next();    // return 3
    fmt.Println(obj.Next()) // 3
    // vector2D.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // vector2D.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // vector2D.next();    // return 4
    fmt.Println(obj.Next()) // 4
    // vector2D.hasNext(); // return False
    fmt.Println(obj.HasNext()) // false

    // Vector2D vector2D = new Vector2D([[1, 2], [3], [4]]);
    obj1 := Constructor1([][]int{{1, 2},{3},{4}})
    fmt.Println(obj1) // 
    // vector2D.next();    // return 1
    fmt.Println(obj1.Next()) // 1
    // vector2D.next();    // return 2
    fmt.Println(obj1.Next()) // 2
    // vector2D.next();    // return 3
    fmt.Println(obj1.Next()) // 3
    // vector2D.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
    // vector2D.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
    // vector2D.next();    // return 4
    fmt.Println(obj1.Next()) // 4
    // vector2D.hasNext(); // return False
    fmt.Println(obj1.HasNext()) // false
}