package main

/**
8 · Rotate Character Array

Description
Given a character array and an offset, rotate the array by offset in place. (from left to right).

offset >= 0
the length of s >= 0
In place means you should change array in the function. You don't return anything.

Example 1:

	Input:
		s = "abcdefg"
		offset = 3
	Output: 
		"efgabcd"
	Explanation:
		Note that it is rotated in place, that is, after s is rotated, it becomes "efgabcd".

Example 2:

	Input:
		s = "abcdefg"
		offset = 0
	Output:
		"abcdefg"
	Explanation:
		Note that it is rotated in place, that is, after s is rotated, it becomes "abcdefg".

Example 3:

	Input:
		s = "abcdefg"
		offset = 1
	Output:
		"gabcdef"
	Explanation:
		Note that it is rotated in place, that is, after s is rotated, it becomes "gabcdef".

Example 4:

	Input:
		s = "abcdefg"
		offset = 2
	Output:
		"fgabcde"
	Explanation:
		Note that it is rotated in place, that is, after s is rotated, it becomes "fgabcde".

Example 5:

	Input:
		s = "abcdefg"
		offset = 10
	Output:
		"efgabcd"
	Explanation:
		Note that it is rotated in place, that is, after s is rotated, it becomes "efgabcd".
*/

import (
	"fmt"
	"unsafe"
	"reflect"
)

 // 要尽量避免[]byte和string的转换，因为转换过程会存在内存拷贝，影响性能。
 // 此外在fasthttp中还提出了一个解决方案，用于[]byte和string的高性能转换。直接看下源码：
// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func b2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// s2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func s2b(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

/**
 * @param s: An array of char
 * @param offset: An integer
 * @return: nothing
 */
 func RotateString(s []byte, offset int) []byte {
    // write your code here
	if offset == 0 {
		return s
	}
	l := len(s)
	s2 := make([]byte,l)
	for i := 0; i < l; i++ {
		fmt.Printf("(i + offset) = %v\n",(i + offset))
		fmt.Printf("(i + offset) %% l = %v\n",((i + offset) % l) ) 
		s2[(i + offset) % l] = s[i] // 当前 index  + 偏移 再取数组长度的 模 就能得到要转成的
	}
	for i := 0; i < l; i++ {
		s[i] = s2[i]
	}
	return s
}

func main() {
	fmt.Printf("RotateString([]byte{'a','b','c','d','e','f','g'},3) = %#v\n",b2s(RotateString([]byte{'a','b','c','d','e','f','g'},3))) // efgabcd
	fmt.Printf("RotateString([]byte{'a','b','c','d','e','f','g'},0) = %#v\n",b2s(RotateString([]byte{'a','b','c','d','e','f','g'},0))) // abcdefg
	fmt.Printf("RotateString([]byte{'a','b','c','d','e','f','g'},1) = %#v\n",b2s(RotateString([]byte{'a','b','c','d','e','f','g'},1))) // gabcdef
	fmt.Printf("RotateString([]byte{'a','b','c','d','e','f','g'},2) = %#v\n",b2s(RotateString([]byte{'a','b','c','d','e','f','g'},2))) // fgabcde
	fmt.Printf("RotateString([]byte{'a','b','c','d','e','f','g'},10) = %#v\n",b2s(RotateString([]byte{'a','b','c','d','e','f','g'},10))) // efgabcd
}