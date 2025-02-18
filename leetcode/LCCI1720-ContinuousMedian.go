package main

// 面试题 17.20. Continuous Median LCCI
// Numbers are randomly generated and passed to a method. 
// Write a program to find and maintain the median value as new values are generated.

// Median is the middle value in an ordered integer list. 
// If the size of the list is even, there is no middle value. 
// So the median is the mean of the two middle value.

// For example,
//     [2,3,4], the median is 3
//     [2,3], the median is (2 + 3) / 2 = 2.5

// Design a data structure that supports the following two operations:
//     void addNum(int num) - Add a integer number from the data stream to the data structure.
//     double findMedian() - Return the median of all elements so far.

// Example:
//     addNum(1)
//     addNum(2)
//     findMedian() -> 1.5
//     addNum(3) 
//     findMedian() -> 2

import "fmt"
import "sort"
import "container/heap"

// type MedianFinder struct {
//     nums        *redblacktree.Tree
//     total       int
//     left, right iterator
// }

// func Constructor() MedianFinder {
//     return MedianFinder{nums: redblacktree.NewWithIntComparator()}
// }

// func (mf *MedianFinder) AddNum(num int) {
//     if count, has := mf.nums.Get(num); has {
//         mf.nums.Put(num, count.(int)+1)
//     } else {
//         mf.nums.Put(num, 1)
//     }
//     if mf.total == 0 {
//         it := mf.nums.Iterator()
//         it.Next()
//         mf.left = iterator{it, 1}
//         mf.right = mf.left
//     } else if mf.total%2 == 1 {
//         if num < mf.left.Key().(int) {
//             mf.left.prev()
//         } else {
//             mf.right.next()
//         }
//     } else {
//         if mf.left.Key().(int) < num && num < mf.right.Key().(int) {
//             mf.left.next()
//             mf.right.prev()
//         } else if num >= mf.right.Key().(int) {
//             mf.left.next()
//         } else {
//             mf.right.prev()
//             mf.left = mf.right
//         }
//     }
//     mf.total++
// }

// func (mf *MedianFinder) FindMedian() float64 {
//     return float64(mf.left.Key().(int)+mf.right.Key().(int)) / 2
// }

// type iterator struct {
//     redblacktree.Iterator
//     count int
// }

// func (it *iterator) prev() {
//     if it.count > 1 {
//         it.count--
//     } else {
//         it.Prev()
//         it.count = it.Value().(int)
//     }
// }

// func (it *iterator) next() {
//     if it.count < it.Value().(int) {
//         it.count++
//     } else {
//         it.Next()
//         it.count = 1
//     }
// }

type MedianFinder struct {
    mn, mx Heap
}

func Constructor() MedianFinder {
    return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
    mn, mx := &mf.mn, &mf.mx
    if mn.Len() == 0 || num <= -mn.IntSlice[0] {
        heap.Push(mn, -num)
        if mx.Len()+1 < mn.Len() {
            heap.Push(mx, -heap.Pop(mn).(int))
        }
    } else {
        heap.Push(mx, num)
        if mx.Len() > mn.Len() {
            heap.Push(mn, -heap.Pop(mx).(int))
        }
    }
}

func (mf *MedianFinder) FindMedian() float64 {
    mn, mx := mf.mn, mf.mx
    if mn.Len() > mx.Len() {
        return float64(-mn.IntSlice[0])
    }
    return float64(mx.IntSlice[0] - mn.IntSlice[0]) / 2
}

type Heap struct{ sort.IntSlice }
func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Heap) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() {
    obj := Constructor()
    fmt.Println(obj)
    //     addNum(1)
    obj.AddNum(1)
    fmt.Println(obj)
    //     addNum(2)
    obj.AddNum(2)
    fmt.Println(obj)
    //     findMedian() -> 1.5
    fmt.Println(obj.FindMedian()) // 1.5
    //     addNum(3) 
    obj.AddNum(3)
    fmt.Println(obj)
    //     findMedian() -> 2
    fmt.Println(obj.FindMedian()) // 2
}