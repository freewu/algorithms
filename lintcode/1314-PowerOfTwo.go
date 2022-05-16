package main

/**
1314 · Power of Two

Description
Given an integer, write a function to determine if it is a power of two.

Example1

	Input: n = 3
	Output: false

Example2

	Input: n = 3
	Output: true

考察位操作。如果一个数n是2的幂，例如 4（ 100）
4-1=3（011） ，减去一后会将原来高位1 后面的位全部补为1，4（100）&3（011）=0 ，进行与操作，结果为0.

10 2
01 1 
10 & 01 = 00

100 4
011 3
100 & 011 = 000

1000 8
0111 7 
1000 & 0111 = 0000
*/

/**
 * @param n: an integer
 * @return: if n is a power of two
 */
func IsPowerOfTwo(n int) bool {
    // Write your code here
	if n <= 0 {
		return false;
	}
	return n && (n - 1) == 0
}

func main() {
	fmt.Printf("IsPowerOfTwo(3) = %v\n",IsPowerOfTwo(3)) // false
	fmt.Printf("IsPowerOfTwo(4) = %v\n",IsPowerOfTwo(4)) // true
	fmt.Printf("IsPowerOfTwo(6) = %v\n",IsPowerOfTwo(6)) // false
	fmt.Printf("IsPowerOfTwo(8) = %v\n",IsPowerOfTwo(8)) // true
}

