package main

// 1622. Fancy Sequence
// Write an API that generates fancy sequences using the append, addAll, and multAll operations.
// Implement the Fancy class:
//     Fancy() Initializes the object with an empty sequence.
//     void append(val) Appends an integer val to the end of the sequence.
//     void addAll(inc) Increments all existing values in the sequence by an integer inc.
//     void multAll(m) Multiplies all existing values in the sequence by an integer m.
//     int getIndex(idx) Gets the current value at index idx (0-indexed) of the sequence modulo 10^9 + 7. If the index is greater or equal than the length of the sequence, return -1.

// Example 1:
// Input
// ["Fancy", "append", "addAll", "append", "multAll", "getIndex", "addAll", "append", "multAll", "getIndex", "getIndex", "getIndex"]
// [[], [2], [3], [7], [2], [0], [3], [10], [2], [0], [1], [2]]
// Output
// [null, null, null, null, null, 10, null, null, null, 26, 34, 20]
// Explanation
// Fancy fancy = new Fancy();
// fancy.append(2);   // fancy sequence: [2]
// fancy.addAll(3);   // fancy sequence: [2+3] -> [5]
// fancy.append(7);   // fancy sequence: [5, 7]
// fancy.multAll(2);  // fancy sequence: [5*2, 7*2] -> [10, 14]
// fancy.getIndex(0); // return 10
// fancy.addAll(3);   // fancy sequence: [10+3, 14+3] -> [13, 17]
// fancy.append(10);  // fancy sequence: [13, 17, 10]
// fancy.multAll(2);  // fancy sequence: [13*2, 17*2, 10*2] -> [26, 34, 20]
// fancy.getIndex(0); // return 26
// fancy.getIndex(1); // return 34
// fancy.getIndex(2); // return 20

// Constraints:
//     1 <= val, inc, m <= 100
//     0 <= idx <= 10^5
//     At most 10^5 calls total will be made to append, addAll, multAll, and getIndex.

import "fmt"

type Fancy1 struct {
    data []int
    adds []int
    muls []int
}

func Constructor1() Fancy1 {
    return Fancy1{
        data: make([]int, 0),
        adds: make([]int, 0),
        muls: make([]int, 0),
    }
}

// Append 将整数 val 添加在序列末尾
func (this *Fancy1) Append(val int) {
    this.data = append(this.data, val)
    this.adds = append(this.adds, 0)
    this.muls = append(this.muls, 1)
}

// AddAll 将所有序列中的现有数值都增加 inc
func (this *Fancy1) AddAll(inc int) {
    if len(this.data) == 0 {
        return
    }
    this.adds[len(this.data)-1] = (this.adds[len(this.data)-1] + inc) % 1_000_000_007
}

func (this *Fancy1) MultAll(m int) {
    if len(this.data) == 0 {
        return
    }
    this.muls[len(this.data)-1] = this.muls[len(this.data)-1] * m % 1_000_000_007
    this.adds[len(this.data)-1] = this.adds[len(this.data)-1] * m % 1_000_000_007
}

// GetIndex 得到下标为 idx 处的数值
// 前缀和 计算出来要加多少，前缀积计算出来要乘多少
func (this *Fancy1) GetIndex(idx int) int {
    res := -1
    if idx >= len(this.data) {
        return res
    }
    res = this.data[idx]
    for i := idx; i < len(this.data); i++ {
        add, mul := this.adds[i], this.muls[i]
        // fmt.Println(res,add,mul)
        res = (res*mul + add) % 1_000_000_007
    }
    return res
}

const ArrayUnit = 10
const Mod = 1_000_000_007

type offset struct {
    mult_value int64
    add_value  int64
}

type Fancy struct {
    len            int
    vals           []int
    offsets        []offset
    current_offset offset
    last_value     int
    last_index     int
}

func Constructor() Fancy {
	return Fancy{
        vals: make([]int, ArrayUnit), 
        offsets: make([]offset, ArrayUnit), 
        current_offset: offset{mult_value: 1, add_value: 0}, 
        last_value: -1, 
        last_index: -1,
    }
}

// 快速幂
func (this *Fancy) quickPow(a, n, mod int64) int64 {
    result := int64(1)
    a %= mod
    for n > 0 {
        if n&1 == 1 {
            result = result * a % mod
        }
        a = a * a % mod
        n >>= 1
    }
    return result
}

// 求逆元
func (this *Fancy) modInverse(a, mod int64) int64 {
    return this.quickPow(a, mod-2, mod)
}

func (this *Fancy) Append(val int) {
    this.len++
    if this.len % ArrayUnit == 0 {
        this.vals = append(this.vals, make([]int, ArrayUnit)...)
        this.offsets = append(this.offsets, make([]offset, ArrayUnit)...)
    }
    this.vals[this.len-1] = val
    this.offsets[this.len-1] = this.current_offset
}

func (this *Fancy) AddAll(inc int) {
    this.last_value = -1
    this.last_index = -1
    if this.len == 0 {
        return
    }
    this.current_offset.add_value = this.current_offset.add_value + int64(inc)
    if this.current_offset.add_value >= Mod {
        this.current_offset.add_value -= Mod
    }
}

func (this *Fancy) MultAll(m int) {
    this.last_value = -1
    this.last_index = -1
    if this.len == 0 {
        return
    }
    this.current_offset.mult_value = this.current_offset.mult_value * int64(m)
    if this.current_offset.mult_value >= Mod {
        this.current_offset.mult_value = this.getMod(this.current_offset.mult_value)
    }
    this.current_offset.add_value = this.current_offset.add_value * int64(m)
    if this.current_offset.add_value >= Mod {
        this.current_offset.add_value = this.getMod(this.current_offset.add_value)
    }
}

func (this *Fancy) GetIndex(idx int) int {
    if idx >= this.len {
        return -1
    }
    if idx == this.last_index {
        return this.last_value
    }
    val := int64(this.vals[idx])
    mult := this.modInverse(this.offsets[idx].mult_value, Mod) * this.current_offset.mult_value
    if mult >= Mod {
        mult = this.getMod(mult)
    }
    add := this.current_offset.add_value - this.offsets[idx].add_value*mult
    if add >= Mod {
        add = this.getMod(add)
    }
    val2 := this.getMod(val*mult + add)
    return int(val2)
}

func (this *Fancy) getMod(value int64) int64 {
    val := value % Mod
    if val < 0 {
        val += Mod
    }
    return int64(val)
}

/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */

func main() {
    // Fancy fancy = new Fancy();
    obj := Constructor()
    fmt.Println(obj)
    // fancy.append(2);   // fancy sequence: [2]
    obj.Append(2)
    fmt.Println(obj) // [2]
    // fancy.addAll(3);   // fancy sequence: [2+3] -> [5]
    obj.AddAll(3)
    fmt.Println(obj) // [5]
    // fancy.append(7);   // fancy sequence: [5, 7]
    obj.Append(7)
    fmt.Println(obj) // [5, 7]
    // fancy.multAll(2);  // fancy sequence: [5*2, 7*2] -> [10, 14]
    obj.MultAll(2)
    fmt.Println(obj) // [10, 14]
    // fancy.getIndex(0); // return 10
    fmt.Println(obj.GetIndex(0)) // 10
    // fancy.addAll(3);   // fancy sequence: [10+3, 14+3] -> [13, 17]
    obj.AddAll(3)
    fmt.Println(obj) // [13, 17]
    // fancy.append(10);  // fancy sequence: [13, 17, 10]
    obj.Append(10)
    fmt.Println(obj) // [13, 17, 10]
    // fancy.multAll(2);  // fancy sequence: [13*2, 17*2, 10*2] -> [26, 34, 20]
    obj.MultAll(2)
    fmt.Println(obj) //  [26, 34, 20]
    // fancy.getIndex(0); // return 26
    fmt.Println(obj.GetIndex(0)) // 26
    // fancy.getIndex(1); // return 34
    fmt.Println(obj.GetIndex(1)) // 34
    // fancy.getIndex(2); // return 20
    fmt.Println(obj.GetIndex(2)) // 20

    // Fancy fancy = new Fancy();
    obj1 := Constructor1()
    fmt.Println(obj1)
    // fancy.append(2);   // fancy sequence: [2]
    obj1.Append(2)
    fmt.Println(obj1) // [2]
    // fancy.addAll(3);   // fancy sequence: [2+3] -> [5]
    obj1.AddAll(3)
    fmt.Println(obj1) // [5]
    // fancy.append(7);   // fancy sequence: [5, 7]
    obj1.Append(7)
    fmt.Println(obj1) // [5, 7]
    // fancy.multAll(2);  // fancy sequence: [5*2, 7*2] -> [10, 14]
    obj1.MultAll(2)
    fmt.Println(obj1) // [10, 14]
    // fancy.getIndex(0); // return 10
    fmt.Println(obj1.GetIndex(0)) // 10
    // fancy.addAll(3);   // fancy sequence: [10+3, 14+3] -> [13, 17]
    obj1.AddAll(3)
    fmt.Println(obj1) // [13, 17]
    // fancy.append(10);  // fancy sequence: [13, 17, 10]
    obj1.Append(10)
    fmt.Println(obj1) // [13, 17, 10]
    // fancy.multAll(2);  // fancy sequence: [13*2, 17*2, 10*2] -> [26, 34, 20]
    obj1.MultAll(2)
    fmt.Println(obj1) //  [26, 34, 20]
    // fancy.getIndex(0); // return 26
    fmt.Println(obj1.GetIndex(0)) // 26
    // fancy.getIndex(1); // return 34
    fmt.Println(obj1.GetIndex(1)) // 34
    // fancy.getIndex(2); // return 20
    fmt.Println(obj1.GetIndex(2)) // 20
}