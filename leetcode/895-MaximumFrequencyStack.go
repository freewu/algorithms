package main

// 895. Maximum Frequency Stack
// Design a stack-like data structure to push elements to the stack and pop the most frequent element from the stack.
// Implement the FreqStack class:
//     FreqStack() constructs an empty frequency stack.
//     void push(int val) pushes an integer val onto the top of the stack.
//     int pop() removes and returns the most frequent element in the stack.

// If there is a tie for the most frequent element, the element closest to the stack's top is removed and returned.

// Example 1:
// Input
// ["FreqStack", "push", "push", "push", "push", "push", "push", "pop", "pop", "pop", "pop"]
// [[], [5], [7], [5], [7], [4], [5], [], [], [], []]
// Output
// [null, null, null, null, null, null, null, 5, 7, 5, 4]
// Explanation
// FreqStack freqStack = new FreqStack();
// freqStack.push(5); // The stack is [5]
// freqStack.push(7); // The stack is [5,7]
// freqStack.push(5); // The stack is [5,7,5]
// freqStack.push(7); // The stack is [5,7,5,7]
// freqStack.push(4); // The stack is [5,7,5,7,4]
// freqStack.push(5); // The stack is [5,7,5,7,4,5]
// freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
// freqStack.pop();   // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
// freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
// freqStack.pop();   // return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the top. The stack becomes [5,7].

// Constraints:
//     0 <= val <= 10^9
//     At most 2 * 10^4 calls will be made to push and pop.
//     It is guaranteed that there will be at least one element in the stack before calling pop

import "fmt"

type FreqStack struct {
    stack map[int][]int
    freq map[int]int
    mx int
}

func Constructor() FreqStack {
    return FreqStack{ make(map[int][]int), make(map[int]int), 0 }
}

func (this *FreqStack) Push(val int)  {
    this.freq[val]++
    max := func (x, y int) int { if x > y { return x; }; return y; }
    this.mx = max(this.mx, this.freq[val])
    this.stack[this.freq[val]] = append(this.stack[this.freq[val]], val)
}

func (this *FreqStack) Pop() int {
    l := len(this.stack[this.mx])
    n := this.stack[this.mx][l-1]
    this.stack[this.mx] = this.stack[this.mx][:l-1]
    this.freq[n]--
    if len(this.stack[this.mx]) == 0 {
        this.mx--
    }
    return n
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

func main() {
    // FreqStack freqStack = new FreqStack();
    obj := Constructor()
    fmt.Println(obj)
    // freqStack.push(5); // The stack is [5]
    obj.Push(5)
    fmt.Println(obj)
    // freqStack.push(7); // The stack is [5,7]
    obj.Push(7)
    fmt.Println(obj)
    // freqStack.push(5); // The stack is [5,7,5]
    obj.Push(5)
    fmt.Println(obj)
    // freqStack.push(7); // The stack is [5,7,5,7]
    obj.Push(7)
    fmt.Println(obj)
    // freqStack.push(4); // The stack is [5,7,5,7,4]
    obj.Push(4)
    fmt.Println(obj)
    // freqStack.push(5); // The stack is [5,7,5,7,4,5]
    obj.Push(5)
    fmt.Println(obj)
    // freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
    fmt.Println(obj.Pop()) // 5
    fmt.Println(obj)
    // freqStack.pop();   // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
    fmt.Println(obj.Pop()) // 7
    fmt.Println(obj)
    // freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
    fmt.Println(obj.Pop()) // 5
    fmt.Println(obj)
    // freqStack.pop();   // return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the top. The stack becomes [5,7].
    fmt.Println(obj.Pop()) // 4
    fmt.Println(obj)
}