package main

// 1825. Finding MK Average
// You are given two integers, m and k, and a stream of integers. 
// You are tasked to implement a data structure that calculates the MKAverage for the stream.

// The MKAverage can be calculated using these steps:
//     1. If the number of the elements in the stream is less than m you should consider the MKAverage to be -1. 
//        Otherwise, copy the last m elements of the stream to a separate container.
//     2. Remove the smallest k elements and the largest k elements from the container.
//     3. Calculate the average value for the rest of the elements rounded down to the nearest integer.

// Implement the MKAverage class:
//     MKAverage(int m, int k) 
//         Initializes the MKAverage object with an empty stream and the two integers m and k.
//     void addElement(int num) 
//         Inserts a new element num into the stream.
//     int calculateMKAverage() 
//         Calculates and returns the MKAverage for the current stream rounded down to the nearest integer.

// Example 1:
// Input
// ["MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"]
// [[3, 1], [3], [1], [], [10], [], [5], [5], [5], []]
// Output
// [null, null, null, -1, null, 3, null, null, null, 5]
// Explanation
// MKAverage obj = new MKAverage(3, 1); 
// obj.addElement(3);        // current elements are [3]
// obj.addElement(1);        // current elements are [3,1]
// obj.calculateMKAverage(); // return -1, because m = 3 and only 2 elements exist.
// obj.addElement(10);       // current elements are [3,1,10]
// obj.calculateMKAverage(); // The last 3 elements are [3,1,10].
//                           // After removing smallest and largest 1 element the container will be [3].
//                           // The average of [3] equals 3/1 = 3, return 3
// obj.addElement(5);        // current elements are [3,1,10,5]
// obj.addElement(5);        // current elements are [3,1,10,5,5]
// obj.addElement(5);        // current elements are [3,1,10,5,5,5]
// obj.calculateMKAverage(); // The last 3 elements are [5,5,5].
//                           // After removing smallest and largest 1 element the container will be [5].
//                           // The average of [5] equals 5/1 = 5, return 5

// Constraints:
//     3 <= m <= 10^5
//     1 <= k*2 < m
//     1 <= num <= 10^5
//     At most 10^5 calls will be made to addElement and calculateMKAverage.

import "fmt"
import "sort"

type MKAverage struct {
    L, M, R   Range
    m, k      int
    q         []int
}

type Range struct {
    s   []int
    sum int
}

func (r *Range) insert(x int) {
    i := sort.SearchInts(r.s, x)
    r.s = append(r.s, 0)
    copy(r.s[i+1:], r.s[i:])
    r.s[i] = x
    r.sum += x
}

func (r *Range) remove(x int) {
    i := sort.SearchInts(r.s, x)
    r.s = append(r.s[:i], r.s[i+1:]...)
    r.sum -= x
}

func Constructor(m int, k int) MKAverage {
    return MKAverage{
        m: m,
        k: k,
    }
}

func (mk *MKAverage) AddElement(num int) {
    mk.q = append(mk.q, num)
    n := len(mk.q)
    if n < mk.m {
        return
    }
    if n == mk.m {
        qq := make([]int, mk.m)
        copy(qq, mk.q)
        sort.Ints(qq)
        for i := 0; i < mk.k; i++ {
            mk.L.insert(qq[i])
        }
        for i := mk.k; i < mk.m-mk.k; i++ {
            mk.M.insert(qq[i])
        }
        for i := mk.m - mk.k; i < mk.m; i++ {
            mk.R.insert(qq[i])
        }
    }
    if n > mk.m {
        mk.M.insert(num)
        x := mk.L.s[len(mk.L.s)-1]
        y := mk.M.s[0]
        if x > y {
            mk.L.remove(x)
            mk.M.remove(y)
            mk.L.insert(y)
            mk.M.insert(x)
        }
        x = mk.M.s[len(mk.M.s)-1]
        y = mk.R.s[0]
        if x > y {
            mk.M.remove(x)
            mk.R.remove(y)
            mk.M.insert(y)
            mk.R.insert(x)
        }
        invalid := mk.q[n-mk.m-1]
        if i := sort.SearchInts(mk.M.s, invalid); i < len(mk.M.s) && mk.M.s[i] == invalid {
            mk.M.remove(invalid)
        } else if i := sort.SearchInts(mk.L.s, invalid); i < len(mk.L.s) && mk.L.s[i] == invalid {
            mk.L.remove(invalid)
            x := mk.M.s[0]
            mk.L.insert(x)
            mk.M.remove(x)
        } else {
            mk.R.remove(invalid)
            x := mk.M.s[len(mk.M.s)-1]
            mk.R.insert(x)
            mk.M.remove(x)
        }
    }
}

func (mk *MKAverage) CalculateMKAverage() int {
    if len(mk.q) < mk.m {
        return -1
    }
    return mk.M.sum / len(mk.M.s)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */

func main() {
    // MKAverage obj = new MKAverage(3, 1); 
    obj := Constructor(3,1)
    fmt.Println(obj)
    // obj.addElement(3);        // current elements are [3]
    obj.AddElement(3)
    fmt.Println(obj)
    // obj.addElement(1);        // current elements are [3,1]
    obj.AddElement(1)
    fmt.Println(obj)
    // obj.calculateMKAverage(); // return -1, because m = 3 and only 2 elements exist.
    fmt.Println(obj.CalculateMKAverage()) // -1
    // obj.addElement(10);       // current elements are [3,1,10]
    obj.AddElement(10)
    fmt.Println(obj)
    // obj.calculateMKAverage(); // The last 3 elements are [3,1,10].
    //                           // After removing smallest and largest 1 element the container will be [3].
    //                           // The average of [3] equals 3/1 = 3, return 3
    fmt.Println(obj.CalculateMKAverage()) // 3
    // obj.addElement(5);        // current elements are [3,1,10,5]
    obj.AddElement(5)
    fmt.Println(obj)
    // obj.addElement(5);        // current elements are [3,1,10,5,5]
    obj.AddElement(5)
    fmt.Println(obj)
    // obj.addElement(5);        // current elements are [3,1,10,5,5,5]
    obj.AddElement(5)
    fmt.Println(obj)
    // obj.calculateMKAverage(); // The last 3 elements are [5,5,5].
    //                           // After removing smallest and largest 1 element the container will be [5].
    //                           // The average of [5] equals 5/1 = 5, return 5
    fmt.Println(obj.CalculateMKAverage()) // 5
}