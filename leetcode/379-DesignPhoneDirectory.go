package main

// 379. Design Phone Directory
// Design a phone directory that initially has maxNumbers empty slots that can store numbers. 
// The directory should store numbers, check if a certain slot is empty or not, and empty a given slot.
// Implement the PhoneDirectory class:
//     PhoneDirectory(int maxNumbers) Initializes the phone directory with the number of available slots maxNumbers.
//     int get() Provides a number that is not assigned to anyone. Returns -1 if no number is available.
//     bool check(int number) Returns true if the slot number is available and false otherwise.
//     void release(int number) Recycles or releases the slot number.

// Example 1:
// Input
// ["PhoneDirectory", "get", "get", "check", "get", "check", "release", "check"]
// [[3], [], [], [2], [], [2], [2], [2]]
// Output
// [null, 0, 1, true, 2, false, null, true]
// Explanation
// PhoneDirectory phoneDirectory = new PhoneDirectory(3);
// phoneDirectory.get();      // It can return any available phone number. Here we assume it returns 0.
// phoneDirectory.get();      // Assume it returns 1.
// phoneDirectory.check(2);   // The number 2 is available, so return true.
// phoneDirectory.get();      // It returns 2, the only number that is left.
// phoneDirectory.check(2);   // The number 2 is no longer available, so return false.
// phoneDirectory.release(2); // Release number 2 back to the pool.
// phoneDirectory.check(2);   // Number 2 is available again, return true.
 
// Constraints:
//     1 <= maxNumbers <= 10^4
//     0 <= number < maxNumbers
//     At most 2 * 10^4 calls will be made to get, check, and release.

import "fmt"

type PhoneDirectory struct {
    used map[int]bool
    queue []int
}

func Constructor(maxNumbers int) PhoneDirectory {
    arr := []int{}
    for i := 0; i < maxNumbers; i++{
        arr = append(arr, i)
    }
    return PhoneDirectory{ used: map[int]bool{}, queue: arr }
}

func (this *PhoneDirectory) Get() int {
    if len(this.queue) <= 0{
        return -1
    }
    pop := this.queue[0] // 取出队列第1个
    this.queue = this.queue[1:]
    this.used[pop] = true
    return pop
}

func (this *PhoneDirectory) Check(number int) bool {
    return !this.used[number]
}

func (this *PhoneDirectory) Release(number int)  {
    if _, ok := this.used[number]; ok{
        delete(this.used, number)
        this.queue = append(this.queue, number) // 重新加回队列中
    }
}


type PhoneDirectory1 struct {
    canuse []bool
}

func Constructor1(maxNumbers int) PhoneDirectory1 {
    canuse := make([]bool, maxNumbers)
    for i := range canuse {
        canuse[i] = true
    }
    return PhoneDirectory1{canuse}
}

func (p *PhoneDirectory1) Get() int {
    for i, b := range p.canuse {
        if b {
            p.canuse[i] = false
            return i
        }
    }
    return -1
}

func (p *PhoneDirectory1) Check(number int) bool {
    return p.canuse[number]
}

func (p *PhoneDirectory1) Release(number int)  {
    p.canuse[number] = true
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */

func main() {
    // PhoneDirectory phoneDirectory = new PhoneDirectory(3);
    obj := Constructor(3)
    // phoneDirectory.get();      // It can return any available phone number. Here we assume it returns 0.
    fmt.Println(obj.Get()) // 0
    // phoneDirectory.get();      // Assume it returns 1.
    fmt.Println(obj.Get()) // 1
    // phoneDirectory.check(2);   // The number 2 is available, so return true.
    fmt.Println(obj.Check(2)) // true
    // phoneDirectory.get();      // It returns 2, the only number that is left.
    fmt.Println(obj.Get()) // 2
    // phoneDirectory.check(2);   // The number 2 is no longer available, so return false.
    fmt.Println(obj.Check(2)) // false
    // phoneDirectory.release(2); // Release number 2 back to the pool.
    obj.Release(2) // true
    fmt.Println(obj) // 
    // phoneDirectory.check(2);   // Number 2 is available again, return true.
    fmt.Println(obj.Check(2)) // true

    // PhoneDirectory phoneDirectory = new PhoneDirectory(3);
    obj1 := Constructor1(3)
    // phoneDirectory.get();      // It can return any available phone number. Here we assume it returns 0.
    fmt.Println(obj1.Get()) // 0
    // phoneDirectory.get();      // Assume it returns 1.
    fmt.Println(obj1.Get()) // 1
    // phoneDirectory.check(2);   // The number 2 is available, so return true.
    fmt.Println(obj1.Check(2)) // true
    // phoneDirectory.get();      // It returns 2, the only number that is left.
    fmt.Println(obj1.Get()) // 2
    // phoneDirectory.check(2);   // The number 2 is no longer available, so return false.
    fmt.Println(obj1.Check(2)) // false
    // phoneDirectory.release(2); // Release number 2 back to the pool.
    obj1.Release(2) // true
    fmt.Println(obj1) // 
    // phoneDirectory.check(2);   // Number 2 is available again, return true.
    fmt.Println(obj1.Check(2)) // true
}