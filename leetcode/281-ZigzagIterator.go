package main

// 281. Zigzag Iterator
// Given two vectors of integers v1 and v2, implement an iterator to return their elements alternately.
// Implement the ZigzagIterator class:
//     ZigzagIterator(List<int> v1, List<int> v2) initializes the object with the two vectors v1 and v2.
//     boolean hasNext() returns true if the iterator still has elements, and false otherwise.
//     int next() returns the current element of the iterator and moves the iterator to the next element.
    
// Example 1:
// Input: v1 = [1,2], v2 = [3,4,5,6]
// Output: [1,3,2,4,5,6]
// Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,3,2,4,5,6].

// Example 2:
// Input: v1 = [1], v2 = []
// Output: [1]

// Example 3:
// Input: v1 = [], v2 = [1]
// Output: [1]
 
// Constraints:
//     0 <= v1.length, v2.length <= 1000
//     1 <= v1.length + v2.length <= 2000
//     -2^31 <= v1[i], v2[i] <= 2^31 - 1
 
// Follow up: What if you are given k vectors? How well can your code be extended to such cases?
// Clarification for the follow-up question:
// The "Zigzag" order is not clearly defined and is ambiguous for k > 2 cases. If "Zigzag" does not look right to you, replace "Zigzag" with "Cyclic".
// Follow-up Example:
//     Input: v1 = [1,2,3], v2 = [4,5,6,7], v3 = [8,9]
//     Output: [1,4,8,2,5,9,3,6,7]

import "fmt"

type ZigzagIterator struct {
    queue []int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
    direction, queue := false, []int{}
    for len(v1) != 0 && len(v2) != 0 { // 交替加入队列
        if direction {
            queue = append(queue, v2[0])
            v2 = v2[1:len(v2)]
            direction = false
            continue
        }
        queue = append(queue, v1[0])
        v1 = v1[1:len(v1)]
        direction = true
    }
    for len(v1) != 0 { // v1 还有余全部追加到队列中
        queue = append(queue, v1[0])
        v1 = v1[1:len(v1)]
    }
    for len(v2) != 0 { // v2 还有余全部追加到队列中
        queue = append(queue, v2[0])
        v2 = v2[1:len(v2)]
    }
    return &ZigzagIterator{
        queue: queue,
    }
}

func (z *ZigzagIterator) next() int { // 每次出队一个
    v := z.queue[0]
    z.queue = z.queue[1:len(z.queue)]
    return v
}

func (z *ZigzagIterator) hasNext() bool {
    return len(z.queue) != 0
}


type Iterator struct {
    arr    []int
    cursor int
    size   int
}

func NewIterator(arr []int) *Iterator {
    return &Iterator{
        arr:    arr,
        cursor: 0,
        size:   len(arr),
    }
}

func (i *Iterator) next() int {
    if !i.hasNext() {
        return 0
    }
    i.cursor++
    return i.arr[i.cursor-1]
}

func (i *Iterator) hasNext() bool {
    return i.cursor != i.size
}

type ZigzagIterator1 struct {
    i1   *Iterator
    i2   *Iterator
    flag bool
}

func Constructor1(v1, v2 []int) *ZigzagIterator1 {
    zi := &ZigzagIterator1{
        i1:   NewIterator(v1),
        i2:   NewIterator(v2),
        flag: true,
    }
    if !zi.i1.hasNext() {
        zi.flag = false
    }
    return zi
}

func (this *ZigzagIterator1) next() int {
    val := 0
    if this.flag {
        val = this.i1.next()
        if this.i2.hasNext() {
            this.flag = false
        }
    } else {
        val = this.i2.next()
        if this.i1.hasNext() {
            this.flag = true
        }
    }
    return val
}

func (this *ZigzagIterator1) hasNext() bool {
    return this.i1.hasNext() || this.i2.hasNext()
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */

func main() {
    // Example 1:
    // Input: v1 = [1,2], v2 = [3,4,5,6]
    // Output: [1,3,2,4,5,6]
    // Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,3,2,4,5,6].
    obj1 := Constructor([]int{1,2},[]int{3,4,5,6})
    for obj1.hasNext() {
        fmt.Printf("%v ", obj1.next())
    }
    fmt.Println() // [1,3,2,4,5,6]
    // Example 2:
    // Input: v1 = [1], v2 = []
    // Output: [1]
    obj2 := Constructor([]int{1},[]int{})
    for obj2.hasNext() {
        fmt.Printf("%v ", obj2.next())
    }
    fmt.Println() // [1]
    // Example 3:
    // Input: v1 = [], v2 = [1]
    // Output: [1]
    obj3 := Constructor([]int{},[]int{1})
    for obj3.hasNext() {
        fmt.Printf("%v ", obj3.next())
    }
    fmt.Println() // [1]
    // Example 4:
    // Input: v1 = [1,2,3], v2 = [4,5,6,7], v3 = [8,9]
    // Output: [1,4,8,2,5,9,3,6,7]
    obj4 := Constructor([]int{1,2,3},[]int{4,5,6,7})
    for obj4.hasNext() {
        fmt.Printf("%v ", obj4.next())
    }
    fmt.Println() // [1,4,2,5,3,6,7]

    obj11 := Constructor1([]int{1,2},[]int{3,4,5,6})
    for obj11.hasNext() {
        fmt.Printf("%v ", obj11.next())
    }
    fmt.Println() // [1,3,2,4,5,6]
    obj12 := Constructor1([]int{1},[]int{})
    for obj12.hasNext() {
        fmt.Printf("%v ", obj12.next())
    }
    fmt.Println() // [1]
    obj13 := Constructor1([]int{},[]int{1})
    for obj13.hasNext() {
        fmt.Printf("%v ", obj13.next())
    }
    fmt.Println() // [1]
    obj14 := Constructor1([]int{1,2,3},[]int{4,5,6,7})
    for obj14.hasNext() {
        fmt.Printf("%v ", obj14.next())
    }
    fmt.Println() // [1,4,2,5,3,6,7]
}