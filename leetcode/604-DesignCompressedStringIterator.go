package main

// 604. Design Compressed String Iterator
// Design and implement a data structure for a compressed string iterator. 
// The given compressed string will be in the form of each letter followed by a positive integer representing the number of this letter existing in the original uncompressed string.

// Implement the StringIterator class:
//     next() 
//         Returns the next character if the original string still has uncompressed characters, 
//         otherwise returns a white space.
//     hasNext() 
//         Returns true if there is any letter needs to be uncompressed in the original string, 
//         otherwise returns false.
 
// Example 1:
// Input
// ["StringIterator", "next", "next", "next", "next", "next", "next", "hasNext", "next", "hasNext"]
// [["L1e2t1C1o1d1e1"], [], [], [], [], [], [], [], [], []]
// Output
// [null, "L", "e", "e", "t", "C", "o", true, "d", true]
// Explanation
// StringIterator stringIterator = new StringIterator("L1e2t1C1o1d1e1");
// stringIterator.next(); // return "L"
// stringIterator.next(); // return "e"
// stringIterator.next(); // return "e"
// stringIterator.next(); // return "t"
// stringIterator.next(); // return "C"
// stringIterator.next(); // return "o"
// stringIterator.hasNext(); // return True
// stringIterator.next(); // return "d"
// stringIterator.hasNext(); // return True

// Constraints:
//     1 <= compressedString.length <= 1000
//     compressedString consists of lower-case an upper-case English letters and digits.
//     The number of a single character repetitions in compressedString is in the range [1, 10^9]
//     At most 100 calls will be made to next and hasNext.

import "fmt"
import "strconv"

type StringIterator struct {
    stack [][]int
}

func Constructor(s string) StringIterator {
    stack :=[][]int{}
    for i := 1; i < len(s); i++{
        j := i + 1
        for ; j < len(s) && isDigital(s[j]); j++ {
        }
        d := int(s[i-1])
        c,_ := strconv.Atoi(s[i:j])
        stack = append(stack, []int{d, c})
        i = j
    }
    return StringIterator{stack: stack}
}

func isDigital(x byte) bool{
    return x >= '0' && x <= '9'
}

func (this *StringIterator) Next() byte {
    if len(this.stack) == 0 {
        return ' '
    }
    d := this.stack[0][0]
    this.stack[0][1]--
    if this.stack[0][1]==0{
        this.stack = this.stack[1:]
    }
    return byte(d)
}

func (this *StringIterator) HasNext() bool {
    return len(this.stack) > 0
}

type StringIterator1 struct {
    s []byte
    c []int
}

func Constructor1(compressedString string) StringIterator1 {
    s, c := make([]byte, 0), make([]int, 0)
    for i := 0; i < len(compressedString); {
        if compressedString[i] < '0' || compressedString[i] > '9' {
            s = append(s, compressedString[i])
            i++
            continue
        }
        n := 0
        for ; i < len(compressedString); i++ {
            if compressedString[i] < '0' || compressedString[i] > '9' {
                break
            }
            n = n*10 + int(compressedString[i]-'0')
        }
        c = append(c, n)
    }
    return StringIterator1{ s: s, c: c }
}

func (this *StringIterator1) Next() byte {
    if len(this.s) == 0 {
        return ' '
    }
    res := this.s[0]
    if this.c[0] == 1 {
        this.s, this.c = this.s[1:], this.c[1:]
    } else {
        this.c[0]--
    }
    return res
}

func (this *StringIterator1) HasNext() bool {
    return len(this.s) != 0
}


/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
    // StringIterator stringIterator = new StringIterator("L1e2t1C1o1d1e1");
    obj := Constructor("L1e2t1C1o1d1e1")
    // stringIterator.next(); // return "L"
    fmt.Printf("%c\n", obj.Next()) // L
    // stringIterator.next(); // return "e"
    fmt.Printf("%c\n", obj.Next()) // e
    // stringIterator.next(); // return "e"
    fmt.Printf("%c\n", obj.Next()) // e
    // stringIterator.next(); // return "t"
    fmt.Printf("%c\n", obj.Next()) // t
    // stringIterator.next(); // return "C"
    fmt.Printf("%c\n", obj.Next()) // C
    // stringIterator.next(); // return "o"
    fmt.Printf("%c\n", obj.Next()) // o
    // stringIterator.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // stringIterator.next(); // return "d"
    fmt.Printf("%c\n", obj.Next()) // d
    // stringIterator.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true

    // StringIterator stringIterator = new StringIterator("L1e2t1C1o1d1e1");
    obj1 := Constructor1("L1e2t1C1o1d1e1")
    // stringIterator.next(); // return "L"
    fmt.Printf("%c\n", obj1.Next()) // L
    // stringIterator.next(); // return "e"
    fmt.Printf("%c\n", obj1.Next()) // e
    // stringIterator.next(); // return "e"
    fmt.Printf("%c\n", obj1.Next()) // e
    // stringIterator.next(); // return "t"
    fmt.Printf("%c\n", obj1.Next()) // t
    // stringIterator.next(); // return "C"
    fmt.Printf("%c\n", obj1.Next()) // C
    // stringIterator.next(); // return "o"
    fmt.Printf("%c\n", obj1.Next()) // o
    // stringIterator.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
    // stringIterator.next(); // return "d"
    fmt.Printf("%c\n", obj1.Next()) // d
    // stringIterator.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
}