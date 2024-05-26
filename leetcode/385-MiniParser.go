package main

// 385. Mini Parser
// Given a string s represents the serialization of a nested list, 
// implement a parser to deserialize it and return the deserialized NestedInteger.

// Each element is either an integer or a list whose elements may also be integers or other lists.

// Example 1:
// Input: s = "324"
// Output: 324
// Explanation: You should return a NestedInteger object which contains a single integer 324.

// Example 2:
// Input: s = "[123,[456,[789]]]"
// Output: [123,[456,[789]]]
// Explanation: Return a NestedInteger object containing a nested list with 2 elements:
// 1. An integer containing value 123.
// 2. A nested list containing two elements:
//     i.  An integer containing value 456.
//     ii. A nested list with one element:
//          a. An integer containing value 789
 
// Constraints:
//     1 <= s.length <= 5 * 10^4
//     s consists of digits, square brackets "[]", negative sign '-', and commas ','.
//     s is the serialization of valid NestedInteger.
//     All the values in the input are in the range [-10^6, 10^6].

import "fmt"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

// 根据题意自己构造NestedInteger的构造函数
// 创建一个空的NestedInteger对象
func NewNestedInteger() NestedInteger {
    return NestedInteger{}
}

// 创建一个带数字的NestedInteger对象
func NewNestedIntegerWithNum(i int) NestedInteger {
    obj := NestedInteger{}
    obj.SetInteger(i)
    return obj
}

type NestedInteger struct {
    num   int
    inner []*NestedInteger
}

func (this *NestedInteger) SetInteger(value int) {
    this.num = value
}
func (this *NestedInteger) Add(elem NestedInteger) {
    this.inner = append(this.inner, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
    return this.inner
}

func deserialize(s string) *NestedInteger {
    stack := []NestedInteger{ NestedInteger{} }
    for i := 0; i < len(s); {
        if s[i] == ',' {
            i++
            continue
        }
        if s[i] == '[' {
            stack = append(stack, NestedInteger{})
            i++
            continue
        } else if s[i] == ']' {
            s := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stack[len(stack)-1].Add(s)
            i++
            continue
        }
        if s[i] == '-' || (s[i] >= '0' && s[i] <= '9')  {
            var temp int
            var minus bool
            if s[i] == '-' {
                minus = true
                i++
            }
            for i < len(s) && s[i] != ',' && s[i] != ']' {
                temp = temp*10 + int(s[i]-'0')
                i++
            }
            if minus {
                temp = -temp
            }
            var new NestedInteger
            new.SetInteger(temp)
            stack[len(stack)-1].Add(new)
        }
    }
    return stack[0].GetList()[0]
}

func deserialize1(s string) *NestedInteger {
    stack:=[]NestedInteger{NestedInteger{}}
    for i:=0;i<len(s);{
        if s[i]==','{
            i++
            continue
        }
        if s[i]=='['{
            stack=append(stack,NestedInteger{})
            i++
            continue
        }else if s[i]==']'{
            top:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            stack[len(stack)-1].Add(top)
            i++
            continue
        }
        if s[i]=='-'|| (s[i]>='0' && s[i]<='9'){
            var temp int
            var minus bool
            if s[i]=='-'{
                minus=true
                i++
            }
            for i<len(s) && s[i]!=',' && s[i]!=']'{
                temp=temp*10+int(s[i]-'0')
                i++
            }
            if minus{
                temp=-temp
            }
            var new NestedInteger
            new.SetInteger(temp)
            stack[len(stack)-1].Add(new)
        }
    }
    return stack[0].GetList()[0]
}

func main() {
    // Example 1:
    // Input: s = "324"
    // Output: 324
    // Explanation: You should return a NestedInteger object which contains a single integer 324.
    fmt.Println(deserialize("324"))
    // Example 2:
    // Input: s = "[123,[456,[789]]]"
    // Output: [123,[456,[789]]]
    // Explanation: Return a NestedInteger object containing a nested list with 2 elements:
    // 1. An integer containing value 123.
    // 2. A nested list containing two elements:
    //     i.  An integer containing value 456.
    //     ii. A nested list with one element:
    //          a. An integer containing value 789
    fmt.Println(deserialize("[123,[456,[789]]]"))

    fmt.Println(deserialize1("324"))
    fmt.Println(deserialize1("[123,[456,[789]]]"))
}