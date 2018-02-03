package main

/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3
Example 2:

Input: J = "z", S = "ZZ"
Output: 0
Note:

S and J will consist of letters and have length at most 50.
The characters in J are distinct.
*/
import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	var n = 0
	var jl = len(J)
	for i := 0; i < len(S); i++ {
		for j := 0; j < jl; j++ {
			if S[i] == J[j] {
				n++
				break // 匹配成功就跳出本循环
			}
		}
	}
	return n
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
}
