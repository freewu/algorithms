package main

// 710. Random Pick with Blacklist
// You are given an integer n and an array of unique integers blacklist. 
// Design an algorithm to pick a random integer in the range [0, n - 1] that is not in blacklist. 
// Any integer that is in the mentioned range and not in blacklist should be equally likely to be returned.

// Optimize your algorithm such that it minimizes the number of calls to the built-in random function of your language.

// Implement the Solution class:
//     Solution(int n, int[] blacklist) Initializes the object with the integer n and the blacklisted integers blacklist.
//     int pick() Returns a random integer in the range [0, n - 1] and not in blacklist.

// Example 1:
// Input
// ["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
// [[7, [2, 3, 5]], [], [], [], [], [], [], []]
// Output
// [null, 0, 4, 1, 6, 1, 0, 4]
// Explanation
// Solution solution = new Solution(7, [2, 3, 5]);
// solution.pick(); // return 0, any integer from [0,1,4,6] should be ok. Note that for every call of pick,
//                  // 0, 1, 4, and 6 must be equally likely to be returned (i.e., with probability 1/4).
// solution.pick(); // return 4
// solution.pick(); // return 1
// solution.pick(); // return 6
// solution.pick(); // return 1
// solution.pick(); // return 0
// solution.pick(); // return 4

// Constraints:
//     1 <= n <= 10^9
//     0 <= blacklist.length <= min(10^5, n - 1)
//     0 <= blacklist[i] < n
//     All the values of blacklist are unique.
//     At most 2 * 10^4 calls will be made to pick.

import "fmt"
import "math/rand"
import "sort"

type Solution struct {
    n int
    black []int
}

func Constructor(n int, blacklist []int) Solution {
    sort.Ints(blacklist)
    return Solution{n: n, black: blacklist, }
}

func (this *Solution) Pick() int {
    a := rand.Intn(this.n - len(this.black))
    return sort.Search(this.n, func(i int) bool {
        return i >= a + sort.Search(len(this.black), func(j int) bool {
            return this.black[j] > i
        })
    })
}

const (
    ModeRange = iota
    ModeSlice
)

type Solution1 struct {
    s    []int
    mode int
    n    int
    m    map[int]struct{}
}

func Constructor1(n int, blacklist []int) Solution1 {
    if float64(len(blacklist))/float64(n) < 0.01 {
        m := make(map[int]struct{}, len(blacklist))
        for _, i := range blacklist {
            m[i] = struct{}{}
        }
        return Solution1 {
            mode: ModeRange,
            n:    n,
            m:    m,
        }
    }
    l := n - len(blacklist)
    s := make([]int, n)
    for _, i := range blacklist {
        if i >= l {
            s[i] = 1
        }
    }
    n--
    for _, i := range blacklist {
        if i >= l {
            continue
        }
        for s[n] == 1 {
            n--
        }
        s[i] = n
        n--
    }
    return Solution1{s: s[:l], mode: ModeSlice }
}

func (this *Solution1) Pick() int {
    if this.mode == ModeSlice {
        i := rand.Intn(len(this.s))
        v := this.s[i]
        if v == 0 {
            return i
        }
        return v
    }
    for {
        v := rand.Intn(this.n)
        if _, ok := this.m[v]; !ok {
            return v
        }
    }
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */

func main() {
    // Solution solution = new Solution(7, [2, 3, 5]);
    obj := Constructor(7, []int{2,3,5})
    fmt.Println(obj)
    // solution.pick(); // return 0, any integer from [0,1,4,6] should be ok. Note that for every call of pick,
    //                  // 0, 1, 4, and 6 must be equally likely to be returned (i.e., with probability 1/4).
    fmt.Println(obj.Pick())
    // solution.pick(); // return 4
    fmt.Println(obj.Pick())
    // solution.pick(); // return 1
    fmt.Println(obj.Pick())
    // solution.pick(); // return 6
    fmt.Println(obj.Pick())
    // solution.pick(); // return 1
    fmt.Println(obj.Pick())
    // solution.pick(); // return 0
    fmt.Println(obj.Pick())
    // solution.pick(); // return 4
    fmt.Println(obj.Pick())

    obj1 := Constructor1(7, []int{2,3,5})
    fmt.Println(obj1)
    // solution.pick(); // return 0, any integer from [0,1,4,6] should be ok. Note that for every call of pick,
    //                  // 0, 1, 4, and 6 must be equally likely to be returned (i.e., with probability 1/4).
    fmt.Println(obj1.Pick())
    // solution.pick(); // return 4
    fmt.Println(obj1.Pick())
    // solution.pick(); // return 1
    fmt.Println(obj1.Pick())
    // solution.pick(); // return 6
    fmt.Println(obj1.Pick())
    // solution.pick(); // return 1
    fmt.Println(obj1.Pick())
    // solution.pick(); // return 0
    fmt.Println(obj1.Pick())
    // solution.pick(); // return 4
    fmt.Println(obj1.Pick())
}