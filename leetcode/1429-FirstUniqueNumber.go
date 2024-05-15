package main

// 1429. First Unique Number
// You have a queue of integers, you need to retrieve the first unique integer in the queue.
// Implement the FirstUnique class:
//     FirstUnique(int[] nums) Initializes the object with the numbers in the queue.
//     int showFirstUnique() returns the value of the first unique integer of the queue, and returns -1 if there is no such integer.
//     void add(int value) insert value to the queue.
 
// Example 1:
// Input: 
// ["FirstUnique","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","showFirstUnique"]
// [[[2,3,5]],[],[5],[],[2],[],[3],[]]
// Output: 
// [null,2,null,2,null,3,null,-1]
// Explanation: 
// FirstUnique firstUnique = new FirstUnique([2,3,5]);
// firstUnique.showFirstUnique(); // return 2
// firstUnique.add(5);            // the queue is now [2,3,5,5]
// firstUnique.showFirstUnique(); // return 2
// firstUnique.add(2);            // the queue is now [2,3,5,5,2]
// firstUnique.showFirstUnique(); // return 3
// firstUnique.add(3);            // the queue is now [2,3,5,5,2,3]
// firstUnique.showFirstUnique(); // return -1

// Example 2:
// Input: 
// ["FirstUnique","showFirstUnique","add","add","add","add","add","showFirstUnique"]
// [[[7,7,7,7,7,7]],[],[7],[3],[3],[7],[17],[]]
// Output: 
// [null,-1,null,null,null,null,null,17]
// Explanation: 
// FirstUnique firstUnique = new FirstUnique([7,7,7,7,7,7]);
// firstUnique.showFirstUnique(); // return -1
// firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7]
// firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3]
// firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3,3]
// firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7,3,3,7]
// firstUnique.add(17);           // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
// firstUnique.showFirstUnique(); // return 17

// Example 3:
// Input: 
// ["FirstUnique","showFirstUnique","add","showFirstUnique"]
// [[[809]],[],[809],[]]
// Output: 
// [null,809,null,-1]
// Explanation: 
// FirstUnique firstUnique = new FirstUnique([809]);
// firstUnique.showFirstUnique(); // return 809
// firstUnique.add(809);          // the queue is now [809,809]
// firstUnique.showFirstUnique(); // return -1
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^8
//     1 <= value <= 10^8
//     At most 50000 calls will be made to showFirstUnique and add.

import "fmt"

type FirstUnique struct {
    m map[int]int // 通过一个哈希表来记录每个元素出现的次数
    data []int // 维护一个队列来保证头部元素永远是不重复的元素
}

func Constructor(nums []int) FirstUnique {
    m := make(map[int]int)
    for _, n := range nums { // 统计每个元素出现次数
        m[n]++
    }
    for len(nums) > 0 && m[nums[0]] != 1 { // 保证队列第一个值是唯一的即可,不是唯一 m[nums[0]] != 1 则从队列移出  nums[1:]
        nums = nums[1:]
    }
    return FirstUnique{ m: m, data: nums }
}

func (this *FirstUnique) ShowFirstUnique() int {
    if len(this.data) == 0 {
        return -1
    }
    return this.data[0]
}

func (this *FirstUnique) Add(value int)  {
    this.m[value]++
    this.data = append(this.data, value)
    for len(this.data) > 0 && this.m[this.data[0]] != 1 { // 保证队列第一个值是唯一的即可,不是唯一 this.m[this.data[0]] != 1 则从队列移出 this.data[1:]
        this.data = this.data[1:]
    }
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */

func main() {
    // Example 1:
    // FirstUnique firstUnique = new FirstUnique([2,3,5]);
    obj1 := Constructor([]int{2,3,5})
    fmt.Println(obj1)
    // firstUnique.showFirstUnique(); // return 2
    fmt.Println(obj1.ShowFirstUnique()) // 2
    // firstUnique.add(5);            // the queue is now [2,3,5,5]
    obj1.Add(5)
    fmt.Println(obj1)
    // firstUnique.showFirstUnique(); // return 2
    fmt.Println(obj1.ShowFirstUnique()) // 2
    // firstUnique.add(2);            // the queue is now [2,3,5,5,2]
    obj1.Add(2)
    fmt.Println(obj1)
    // firstUnique.showFirstUnique(); // return 3
    fmt.Println(obj1.ShowFirstUnique()) // 3
    // firstUnique.add(3);            // the queue is now [2,3,5,5,2,3]
    obj1.Add(3)
    fmt.Println(obj1)
    // firstUnique.showFirstUnique(); // return -1
    fmt.Println(obj1.ShowFirstUnique()) // -1

    // Example 2:
    // FirstUnique firstUnique = new FirstUnique([7,7,7,7,7,7]);
    obj2 := Constructor([]int{7,7,7,7,7,7})
    fmt.Println(obj2)
    // firstUnique.showFirstUnique(); // return -1
    fmt.Println(obj2.ShowFirstUnique()) // -1
    // firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7]
    obj2.Add(7)
    fmt.Println(obj2)
    // firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3]
    obj2.Add(3)
    fmt.Println(obj2)
    fmt.Println(obj2.ShowFirstUnique()) // 3
    // firstUnique.add(3);            // the queue is now [7,7,7,7,7,7,7,3,3]
    obj2.Add(3)
    fmt.Println(obj2)
    fmt.Println(obj2.ShowFirstUnique()) // -1
    // firstUnique.add(7);            // the queue is now [7,7,7,7,7,7,7,3,3,7]
    obj2.Add(7)
    fmt.Println(obj2)
    // firstUnique.add(17);           // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
    obj2.Add(17)
    fmt.Println(obj2)
    // firstUnique.showFirstUnique(); // return 17
    fmt.Println(obj2.ShowFirstUnique()) // 17

    // Example 3:
    // FirstUnique firstUnique = new FirstUnique([809]);
    obj3 := Constructor([]int{809})
    fmt.Println(obj3)
    // firstUnique.showFirstUnique(); // return 809
    fmt.Println(obj3.ShowFirstUnique()) // 809
    // firstUnique.add(809);          // the queue is now [809,809]
    obj3.Add(809)
    fmt.Println(obj3)
    // firstUnique.showFirstUnique(); // return -1
    fmt.Println(obj3.ShowFirstUnique()) // -1
}