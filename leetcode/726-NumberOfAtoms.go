package main

// 726. Number of Atoms
// Given a string formula representing a chemical formula, return the count of each atom.
// The atomic element always starts with an uppercase character, then zero or more lowercase letters, representing the name.
// One or more digits representing that element's count may follow if the count is greater than 1. 
// If the count is 1, no digits will follow.
//     For example, "H2O" and "H2O2" are possible, but "H1O2" is impossible.

// Two formulas are concatenated together to produce another formula.
//     For example, "H2O2He3Mg4" is also a formula.

// A formula placed in parentheses, and a count (optionally added) is also a formula.
//     For example, "(H2O2)" and "(H2O2)3" are formulas.

// Return the count of all elements as a string in the following form: 
//     the first name (in sorted order), 
//     followed by its count (if that count is more than 1), 
//     followed by the second name (in sorted order), 
//     followed by its count (if that count is more than 1), and so on.

// The test cases are generated so that all the values in the output fit in a 32-bit integer.

// Example 1:
// Input: formula = "H2O"
// Output: "H2O"
// Explanation: The count of elements are {'H': 2, 'O': 1}.

// Example 2:
// Input: formula = "Mg(OH)2"
// Output: "H2MgO2"
// Explanation: The count of elements are {'H': 2, 'Mg': 1, 'O': 2}.

// Example 3:
// Input: formula = "K4(ON(SO3)2)2"
// Output: "K4N2O14S4"
// Explanation: The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.

// Constraints:
//     1 <= formula.length <= 1000
//     formula consists of English letters, digits, '(', and ')'.
//     formula is always valid.

import "fmt"
import "strconv"
import "sort"

// // 解答错误 29 / 33 
// func countOfAtoms(formula string) string {
//     hash, atoms, count, parenthesis := map[string]int{}, []string{}, []int{},[]int{}
//     for i := 0; i < len(formula); {
//         if formula[i] == '(' {
//             parenthesis = append(parenthesis, len(atoms))
//             i++
//         } else if formula[i] == ')' {
//             i++
//             numInBytes := []byte{}
//             for i < len(formula) {
//                 if formula[i] >= '0' && formula[i] <= '9' {
//                     numInBytes = append(numInBytes, formula[i])
//                     i++
//                 } else {
//                     break
//                 }
//             }
//             times, _ := strconv.Atoi(string(numInBytes))
//             for i := parenthesis[len(parenthesis) - 1]; i < len(atoms); i++ {
//                 count[i] *= times
//             }
//             parenthesis = parenthesis[:len(parenthesis) - 1]
//         } else {
//             atomInBytes := []byte{formula[i]}
//             i++
//             for i < len(formula) {
//                 if formula[i] >= 'a' && formula[i] <= 'z' {
//                     atomInBytes = append(atomInBytes, formula[i])
//                     i++
//                 } else {
//                     break
//                 }
//             }
//             atoms = append(atoms, string(atomInBytes))
//             numInBytes := []byte{}
//             for i < len(formula) {
//                 if formula[i] >= '0' && formula[i] <= '9' {
//                     numInBytes = append(numInBytes, formula[i])
//                     i++
//                 } else {
//                     break
//                 }
//             }
//             if len(numInBytes) == 0 {
//                 count = append(count, 1)
//             } else {
//                 times, _ := strconv.Atoi(string(numInBytes))
//                 count = append(count, times)
//             }
//         }
//     }
//     for i := range atoms {
//         hash[atoms[i]] += count[i]
//     }
//     atomsDistinct := []string{}
//     for k := range hash {
//         atomsDistinct = append(atomsDistinct, k)
//     }
//     sort.Strings(atomsDistinct)
//     resultInBytes := []byte{}
//     for _, atom := range atomsDistinct {
//         resultInBytes = append(resultInBytes, []byte(atom)...)
//         if hash[atom] != 1 {
//             resultInBytes = append(resultInBytes, []byte(strconv.Itoa(hash[atom]))...)
//         }
//     }
//     return string(resultInBytes)
// }

func countOfAtoms(formula string) string {
    maps, multiple := make(map[string]int), []int{} // 输出maps, 倍数列表
    str, count := "", "" // 缓存字符，缓存数字
    for i := len(formula) - 1; i >= 0; i-- {
        char := formula[i]
        if char >= '0' && char <= '9' {
            count = string(char) + count
        } else {
            if char == ')' {
                atoi, _ := strconv.Atoi(count)
                if len(multiple) > 0 {
                    atoi = atoi * multiple[len(multiple)-1]
                }
                multiple = append(multiple, atoi)
                count = ""
            } else if char == '(' {
                multiple = multiple[:len(multiple)-1]
            } else if char >= 'a' && char <= 'z' { // 小写
                str = string(char) + str
            } else if char >= 'A' && char <= 'Z' { // 大写
                str = string(char) + str
                nums := 1
                if count == "" {
                    count = "1"
                }
                atoi, _ := strconv.Atoi(count)
                // fmt.Println(multiple)
                // fmt.Println(atoi)
                if len(multiple) > 0 {
                    multi := multiple[len(multiple)-1]
                    if multi == 0 { // 处理 只有()的情况 如: "Mg(H2O)N"
                        nums = atoi
                    } else {
                        nums = multi * atoi
                    }
                } else {
                    nums = atoi
                }
                maps[str] += nums
                str = ""
                count = ""
            }
        }
    }
    //fmt.Println(maps)
    // 按字母排序并输出
    res := ""
    var keys []string
    for k := range maps {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        v := maps[k]
        res += k
        if v != 1 {
            res += strconv.Itoa(v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: formula = "H2O"
    // Output: "H2O"
    // Explanation: The count of elements are {'H': 2, 'O': 1}.
    fmt.Println(countOfAtoms("H2O")) // "H2O"
    // Example 2:
    // Input: formula = "Mg(OH)2"
    // Output: "H2MgO2"
    // Explanation: The count of elements are {'H': 2, 'Mg': 1, 'O': 2}.
    fmt.Println(countOfAtoms("Mg(OH)2")) // "H2MgO2"
    // Example 3:
    // Input: formula = "K4(ON(SO3)2)2"
    // Output: "K4N2O14S4"
    // Explanation: The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.
    fmt.Println(countOfAtoms("K4(ON(SO3)2)2")) // "K4N2O14S4"

    fmt.Println(countOfAtoms("Mg(H2O)N")) // "H2MgNO"
}