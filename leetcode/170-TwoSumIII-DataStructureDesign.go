package main

// 170. Two Sum III - Data structure design
// Design a data structure that accepts a stream of integers and checks if it has a pair of integers that sum up to a particular value.
// Implement the TwoSum class:
//     TwoSum() 
//         Initializes the TwoSum object, with an empty array initially.
//     void add(int number) 
//         Adds number to the data structure.
//     boolean find(int value) 
//         Returns true if there exists any pair of numbers whose sum is equal to value, otherwise, it returns false.

// Example 1:
// Input
// ["TwoSum", "add", "add", "add", "find", "find"]
// [[], [1], [3], [5], [4], [7]]
// Output
// [null, null, null, null, true, false]
// Explanation
// TwoSum twoSum = new TwoSum();
// twoSum.add(1);   // [] --> [1]
// twoSum.add(3);   // [1] --> [1,3]
// twoSum.add(5);   // [1,3] --> [1,3,5]
// twoSum.find(4);  // 1 + 3 = 4, return true
// twoSum.find(7);  // No two integers sum up to 7, return false
 
// Constraints:
//     -10^5 <= number <= 10^5
//     -2^31 <= value <= 2^31 - 1
//     At most 10^4 calls will be made to add and find.

import "fmt"
import "sort"

type TwoSum1 struct {
    arr []int
}

func Constructor1() TwoSum1 {
    return TwoSum1{[]int{}}
}

func (this *TwoSum1) Add(number int)  {
    this.arr = append(this.arr, number)
    sort.Ints(this.arr)
}

func (this *TwoSum1) Find(value int) bool {
    i, j := 0, len(this.arr)-1
    for i < j { // 从外向内缩进
        val := this.arr[i] + this.arr[j]
        if val == value { // 如果刚好匹配则返回
            return true
            // return []int{i + 1, j + 1}
        }
        if val < value { // 如果过小，说明开头需要向里走 ->
            i++
        } else { // 过大, 从尾部向头收缩 <-
            j--
        }
    }
    return false
}


type TwoSum struct {
    arr []int
}

func Constructor() TwoSum {
    return TwoSum{[]int{}}
}

func (this *TwoSum) Add(number int)  {
    this.arr = append(this.arr, number)
    i := len(this.arr) - 1
    for i > 0 && this.arr[i - 1] > this.arr[i] {
        this.arr[i - 1], this.arr[i] = this.arr[i], this.arr[i - 1]
        i--
    }
}

func (this *TwoSum) Find(value int) bool {
    for left, right := 0, len(this.arr) - 1; left < right; {
        sum := this.arr[left] + this.arr[right]
        if sum == value {
            return true
        } else if sum > value {
            right--
        } else {
            left++
        }
    }
    return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */

func main() {
    // TwoSum twoSum = new TwoSum();
    obj := Constructor()
    fmt.Println(obj)
    // twoSum.add(1);   // [] --> [1]
    obj.Add(1)
    fmt.Println(obj)
    // twoSum.add(3);   // [1] --> [1,3]
    obj.Add(3)
    fmt.Println(obj)
    // twoSum.add(5);   // [1,3] --> [1,3,5]
    obj.Add(5)
    fmt.Println(obj)
    // twoSum.find(4);  // 1 + 3 = 4, return true
    fmt.Println(obj.Find(4)) // 4
    // twoSum.find(7);  // No two integers sum up to 7, return false
    fmt.Println(obj.Find(7)) // 7


    // TwoSum twoSum = new TwoSum();
    obj1 := Constructor1()
    fmt.Println(obj1)
    // twoSum.add(1);   // [] --> [1]
    obj1.Add(1)
    fmt.Println(obj1)
    // twoSum.add(3);   // [1] --> [1,3]
    obj1.Add(3)
    fmt.Println(obj1)
    // twoSum.add(5);   // [1,3] --> [1,3,5]
    obj1.Add(5)
    fmt.Println(obj1)
    // twoSum.find(4);  // 1 + 3 = 4, return true
    fmt.Println(obj1.Find(4)) // 4
    // twoSum.find(7);  // No two integers sum up to 7, return false
    fmt.Println(obj1.Find(7)) // 7
}