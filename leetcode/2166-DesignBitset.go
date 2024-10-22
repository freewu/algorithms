package main

// 2166. Design Bitset
// A Bitset is a data structure that compactly stores bits.

// Implement the Bitset class:
//     Bitset(int size) 
//         Initializes the Bitset with size bits, all of which are 0.
//     void fix(int idx) 
//         Updates the value of the bit at the index idx to 1. 
//         If the value was already 1, no change occurs.
//     void unfix(int idx) 
//         Updates the value of the bit at the index idx to 0. 
//         If the value was already 0, no change occurs.
//     void flip() 
//         Flips the values of each bit in the Bitset. 
//         In other words, all bits with value 0 will now have value 1 and vice versa.
//     boolean all() 
//         Checks if the value of each bit in the Bitset is 1. 
//         Returns true if it satisfies the condition, false otherwise.
//     boolean one() 
//         Checks if there is at least one bit in the Bitset with value 1. 
//         Returns true if it satisfies the condition, false otherwise.
//     int count() 
//         Returns the total number of bits in the Bitset which have value 1.
//     String toString() 
//         Returns the current composition of the Bitset. 
//         Note that in the resultant string, the character at the ith index should coincide with the value at the ith bit of the Bitset.

// Example 1:
// Input
// ["Bitset", "fix", "fix", "flip", "all", "unfix", "flip", "one", "unfix", "count", "toString"]
// [[5], [3], [1], [], [], [0], [], [], [0], [], []]
// Output
// [null, null, null, null, false, null, null, true, null, 2, "01010"]
// Explanation
// Bitset bs = new Bitset(5); // bitset = "00000".
// bs.fix(3);     // the value at idx = 3 is updated to 1, so bitset = "00010".
// bs.fix(1);     // the value at idx = 1 is updated to 1, so bitset = "01010". 
// bs.flip();     // the value of each bit is flipped, so bitset = "10101". 
// bs.all();      // return False, as not all values of the bitset are 1.
// bs.unfix(0);   // the value at idx = 0 is updated to 0, so bitset = "00101".
// bs.flip();     // the value of each bit is flipped, so bitset = "11010". 
// bs.one();      // return True, as there is at least 1 index with value 1.
// bs.unfix(0);   // the value at idx = 0 is updated to 0, so bitset = "01010".
// bs.count();    // return 2, as there are 2 bits with value 1.
// bs.toString(); // return "01010", which is the composition of bitset.

// Constraints:
//     1 <= size <= 10……5
//     0 <= idx <= size - 1
//     At most 10^5 calls will be made in total to fix, unfix, flip, all, one, count, and toString.
//     At least one call will be made to all, one, count, or toString.
//     At most 5 calls will be made to toString.

import "fmt"

type Bitset struct {
    OneBits map[int]struct{} 
    ZeroBits map[int]struct{}
    Size int
}

func Constructor(size int) Bitset {
    obj := Bitset{
        OneBits:    map[int]struct{}{}, // with empty struct will us save some memory
        ZeroBits:   map[int]struct{}{}, // with empty struct will us save some memory
        Size:       size,
    }
    for i := 0; i < size; i++ { // fill zero map
        obj.ZeroBits[i] = struct{}{}
    }
    return obj
}

func (this *Bitset) Fix(idx int)  {
    // sets it to 1
    if _, ok := this.OneBits[idx]; ok{ return }
    this.OneBits[idx] = struct{}{}
    delete(this.ZeroBits, idx)
}

func (this *Bitset) Unfix(idx int)  {
    // sets to 0
    if _, ok := this.ZeroBits[idx]; ok { return }
    this.ZeroBits[idx] = struct{}{}
    delete(this.OneBits, idx)
}

func (this *Bitset) Flip()  {
    // 1 -> 0 vica versa
    this.OneBits, this.ZeroBits = this.ZeroBits, this.OneBits  
}


func (this *Bitset) All() bool {
    // true if all is 1 
    return len(this.OneBits) == this.Size
}

func (this *Bitset) One() bool {
    // true, at least 1 bit == 1
    return len(this.OneBits) > 0
   
}

func (this *Bitset) Count() int {
    // sum of 1s
    return len(this.OneBits)
}

func (this *Bitset) ToString() string {
    res := []byte{}
    for i := 0; i < this.Size; i++ {
        if _, ok := this.ZeroBits[i]; ok {
            res = append(res, '0')
        } else {
            res = append(res, '1')
        }
    }
    return string(res)
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */

func main() {
    // Bitset bs = new Bitset(5); // bitset = "00000".
    obj := Constructor(5)
    fmt.Println(obj) // 00000
    // bs.fix(3);     // the value at idx = 3 is updated to 1, so bitset = "00010".
    obj.Fix(3)
    fmt.Println(obj) // 00010
    // bs.fix(1);     // the value at idx = 1 is updated to 1, so bitset = "01010". 
    obj.Fix(1)
    fmt.Println(obj) // 01010
    // bs.flip();     // the value of each bit is flipped, so bitset = "10101". 
    obj.Flip()
    fmt.Println(obj) // 10101
    // bs.all();      // return False, as not all values of the bitset are 1.
    fmt.Println(obj.All()) // false
    // bs.unfix(0);   // the value at idx = 0 is updated to 0, so bitset = "00101".
    obj.Unfix(0)
    fmt.Println(obj) // 00101
    // bs.flip();     // the value of each bit is flipped, so bitset = "11010". 
    obj.Flip()
    fmt.Println(obj) // 11010
    // bs.one();      // return True, as there is at least 1 index with value 1.
    fmt.Println(obj.One()) // true
    // bs.unfix(0);   // the value at idx = 0 is updated to 0, so bitset = "01010".
    obj.Unfix(0)
    fmt.Println(obj) // 01010
    // bs.count();    // return 2, as there are 2 bits with value 1.
    fmt.Println(obj.Count()) // 2
    // bs.toString(); // return "01010", which is the composition of bitset.
    fmt.Println(obj.ToString()) // 01010
}