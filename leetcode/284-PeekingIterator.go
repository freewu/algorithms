package main

// 284. Peeking Iterator
// Design an iterator that supports the peek operation on an existing iterator in addition to the hasNext and the next operations.

// Implement the PeekingIterator class:
//     PeekingIterator(Iterator<int> nums) Initializes the object with the given integer iterator iterator.
//     int next() Returns the next element in the array and moves the pointer to the next element.
//     boolean hasNext() Returns true if there are still elements in the array.
//     int peek() Returns the next element in the array without moving the pointer.

// Note: Each language may have a different implementation of the constructor and Iterator,
// but they all support the int next() and boolean hasNext() functions.

// Example 1:
// Input
// ["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
// [[[1, 2, 3]], [], [], [], [], []]
// Output
// [null, 1, 2, 2, 3, false]
// Explanation
// PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
// peekingIterator.next();    // return 1, the pointer moves to the next element [1,2,3].
// peekingIterator.peek();    // return 2, the pointer does not move [1,2,3].
// peekingIterator.next();    // return 2, the pointer moves to the next element [1,2,3]
// peekingIterator.next();    // return 3, the pointer moves to the next element [1,2,3]
// peekingIterator.hasNext(); // return False

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 1000
//     All the calls to next and peek are valid.
//     At most 1000 calls will be made to next, hasNext, and peek.

// Below is the interface for Iterator, which is already defined for you.

import "fmt"

type Iterator struct {
    data  []int  // 该迭代器中存放的元素集合的指针
    index int  // 该迭代器当前指向的元素下标，-1 即不存在元素
}

func NewIterator(arr []int) *Iterator {
    return &Iterator{arr, 0}
}

func (this *Iterator) hasNext() bool {
    if this == nil { // 迭代器为nil时不能后移
        return false
    }
    if len(this.data) == 0 { // 元素集合为空时不能后移
        return false
    }
    return this.index < len(this.data) // 下标到达元素集合上限时不能后移,否则可以后移
}

func (this *Iterator) next() int {
    if this.index < len(this.data) {
        this.index++
        return this.data[this.index - 1]
    }
    return -1
}

type PeekingIterator struct {
    iter    *Iterator
    val     int
    has     bool
    cached  bool
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter, 0, false, false}
}

func (this *PeekingIterator) hasNext() bool {
    if this.cached {
        return this.has
    }
    return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
    if this.cached {
        this.cached = false
        return this.val
    }
    return this.iter.next()
}

func (this *PeekingIterator) peek() int {
    if !this.cached {
        this.has = this.iter.hasNext()
        this.val = this.iter.next()
        this.cached = true
    }
    return this.val
}

func main() {
    // PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
    obj := Constructor(NewIterator([]int{1,2,3}))
    fmt.Println(obj)
    // peekingIterator.next();    // return 1, the pointer moves to the next element [1,2,3].
    fmt.Println(obj.next()) // 1
    fmt.Println(obj)
    // peekingIterator.peek();    // return 2, the pointer does not move [1,2,3].
    fmt.Println(obj.peek()) // 2
    fmt.Println(obj)
    // peekingIterator.next();    // return 2, the pointer moves to the next element [1,2,3]
    fmt.Println(obj.next()) // 2
    fmt.Println(obj)
    fmt.Println(obj.hasNext()) // true
    // peekingIterator.next();    // return 3, the pointer moves to the next element [1,2,3]
    fmt.Println(obj.next()) // 3
    fmt.Println(obj)
    // peekingIterator.hasNext(); // return False
    fmt.Println(obj.hasNext()) // false
    fmt.Println(obj)
}