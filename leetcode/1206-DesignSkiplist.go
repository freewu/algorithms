package main

// 1206. Design Skiplist
// Design a Skiplist without using any built-in libraries.

// A skiplist is a data structure that takes O(log(n)) time to add, erase and search. 
// Comparing with treap and red-black tree which has the same function and performance, the code length of Skiplist can be comparatively short and the idea behind Skiplists is just simple linked lists.

// For example, we have a Skiplist containing [30,40,50,60,70,90] and we want to add 80 and 45 into it. 
// The Skiplist works this way:
// <img src="https://assets.leetcode.com/uploads/2019/09/27/1506_skiplist.gif" />

// Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons

// You can see there are many layers in the Skiplist. 
// Each layer is a sorted linked list. 
// With the help of the top layers, add, erase and search can be faster than O(n). 
// It can be proven that the average time complexity for each operation is O(log(n)) and space complexity is O(n).

// See more about Skiplist: https://en.wikipedia.org/wiki/Skip_list

// Implement the Skiplist class:
//     Skiplist() 
//         Initializes the object of the skiplist.
//     bool search(int target) 
//         Returns true if the integer target exists in the Skiplist or false otherwise.
//     void add(int num) 
//         Inserts the value num into the SkipList.
//     bool erase(int num) 
//         Removes the value num from the Skiplist and returns true. 
//         If num does not exist in the Skiplist, do nothing and return false. 
//         If there exist multiple num values, removing any one of them is fine.

// Note that duplicates may exist in the Skiplist, your code needs to handle this situation.

// Example 1:
// Input
// ["Skiplist", "add", "add", "add", "search", "add", "search", "erase", "erase", "search"]
// [[], [1], [2], [3], [0], [4], [1], [0], [1], [1]]
// Output
// [null, null, null, null, false, null, true, false, true, false]
// Explanation
// Skiplist skiplist = new Skiplist();
// skiplist.add(1);
// skiplist.add(2);
// skiplist.add(3);
// skiplist.search(0); // return False
// skiplist.add(4);
// skiplist.search(1); // return True
// skiplist.erase(0);  // return False, 0 is not in skiplist.
// skiplist.erase(1);  // return True
// skiplist.search(1); // return False, 1 has already been erased.

// Constraints:
//     0 <= num, target <= 2 * 10^4
//     At most 5 * 10^4 calls will be made to search, add, and erase.

import "fmt"
import "math"
import "math/rand"

const (
    MaxLevel = 32
    P        = 1 / math.E
)

type Skiplist struct {
    head  *Node
    level int
}

func Constructor() Skiplist {
    return Skiplist{
        head:  newNode(0, 32),
        level: 1,
    }
}

type Node struct {
    score int
    next  []*Node
}

func newNode(score int, level int) *Node {
    return &Node{ score, make([]*Node, level), }
}

func randomLevel() int {
    level := 1
    for rand.Float64() < P {
        level++
    }
    if level > MaxLevel {
        return MaxLevel
    }
    return level
}

func (s *Skiplist) Search(target int) bool {
    x := s.head
    for i := s.level - 1; i >= 0; i-- {
        for x.next[i] != nil {
            if x.next[i].score == target {
                return true
            } else if x.next[i].score < target {
                x = x.next[i]
            } else {
                break
            }
        }
    }
    return false
}

func (s *Skiplist) Add(score int) {
    update := make([]*Node, MaxLevel)
    for i, x := s.level-1, s.head; i >= 0; i-- {
        for x.next[i] != nil && x.next[i].score < score {
            x = x.next[i]
        }
        update[i] = x
    }
    lvl := randomLevel()
    n := newNode(score, lvl)
    if lvl > s.level {
        for i := s.level; i < lvl; i++ {
            update[i] = s.head
        }
        s.level = lvl
    }
    for i := 0; i < lvl; i++ {
        n.next[i] = update[i].next[i]
        update[i].next[i] = n
    }
}

func (s *Skiplist) Erase(score int) bool {
    update := make([]*Node, MaxLevel)
    for i, x := s.level-1, s.head; i >= 0; i-- {
        for x.next[i] != nil && x.next[i].score < score {
            x = x.next[i]
        }
        update[i] = x
    }
    node := update[0].next[0]
    if node == nil || node.score != score {
        return false
    }
    for i := 0; i < len(node.next); i++ {
        update[i].next[i] = node.next[i]
    }
    for s.level > 1 && s.head.next[s.level-1] == nil {
        s.level--
    }
    return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

func main() {
    // Skiplist skiplist = new Skiplist();
    obj := Constructor()
    fmt.Println(obj)
    // skiplist.add(1);
    obj.Add(1)
    fmt.Println(obj)
    // skiplist.add(2);
    obj.Add(21)
    fmt.Println(obj)
    // skiplist.add(3);
    obj.Add(3)
    fmt.Println(obj)
    // skiplist.search(0); // return False
    fmt.Println(obj.Search(0)) // false
    // skiplist.add(4);
    obj.Add(4)
    fmt.Println(obj)
    // skiplist.search(1); // return True
    fmt.Println(obj.Search(1)) // true
    // skiplist.erase(0);  // return False, 0 is not in skiplist.
    fmt.Println(obj.Erase(0)) // false
    fmt.Println(obj)
    // skiplist.erase(1);  // return True
    // skiplist.search(1); // return False, 1 has already been erased.
    fmt.Println(obj.Erase(1)) // true
    fmt.Println(obj)
    fmt.Println(obj.Search(1)) // false
}