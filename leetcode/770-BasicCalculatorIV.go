package main

// 770. Basic Calculator IV
// Given an expression such as expression = "e + 8 - a + 5" 
// and an evaluation map such as {"e": 1} (given in terms of evalvars = ["e"] 
// and evalints = [1]), return a list of tokens representing the simplified expression, such as ["-1*a","14"]
//     1. An expression alternates chunks and symbols, with a space separating each chunk and symbol.
//     2. A chunk is either an expression in parentheses, a variable, or a non-negative integer.
//     3. A variable is a string of lowercase letters (not including digits.) Note that variables can be multiple letters, and note that variables never have a leading coefficient or unary operator like "2x" or "-x".
    
// Expressions are evaluated in the usual order: brackets first, then multiplication, then addition and subtraction.
//     For example, expression = "1 + 2 * 3" has an answer of ["7"].

// The format of the output is as follows:
//     1. For each term of free variables with a non-zero coefficient, 
//        we write the free variables within a term in sorted order lexicographically.
//             For example, we would never write a term like "b*a*c", only "a*b*c".
//     2. Terms have degrees equal to the number of free variables being multiplied, counting multiplicity. 
//        We write the largest degree terms of our answer first, breaking ties by lexicographic order ignoring the leading coefficient of the term.
//             For example, "a*a*b*c" has degree 4.
//     3. The leading coefficient of the term is placed directly to the left with an asterisk separating it from the variables (if they exist.) A leading coefficient of 1 is still printed.
//     4. An example of a well-formatted answer is ["-2*a*a*a", "3*a*a*b", "3*b*b", "4*a", "5*c", "-6"].
//     5. Terms (including constant terms) with coefficient 0 are not included.
//         For example, an expression of "0" has an output of [].

// Note: 
//     You may assume that the given expression is always valid. 
//     All intermediate results will be in the range of [-2^31, 2^31 - 1].

// Example 1:
// Input: expression = "e + 8 - a + 5", evalvars = ["e"], evalints = [1]
// Output: ["-1*a","14"]

// Example 2:
// Input: expression = "e - 8 + temperature - pressure", evalvars = ["e", "temperature"], evalints = [1, 12]
// Output: ["-1*pressure","5"]

// Example 3:
// Input: expression = "(e + 8) * (e - 8)", evalvars = [], evalints = []
// Output: ["1*e*e","-64"]

// Constraints:
//     1 <= expression.length <= 250
//     expression consists of lowercase English letters, digits, '+', '-', '*', '(', ')', ' '.
//     expression does not contain any leading or trailing spaces.
//     All the tokens in expression are separated by a single space.
//     0 <= evalvars.length <= 100
//     1 <= evalvars[i].length <= 20
//     evalvars[i] consists of lowercase English letters.
//     evalints.length == evalvars.length
//     -100 <= evalints[i] <= 100

import "fmt"
import "strconv"
import "sort"
import "strings"

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
    evalMap := map[string]int{}
    for i := range evalvars {
        evalMap[evalvars[i]] = evalints[i]
    }
    return Parse(expression).Evaluate(evalMap).ToList()
}

type Polynomial map[string]int

func MakePolynomial(expr string) Polynomial {
    h := map[string]int{}
    if isNumber(expr[0]) {
        num, _ := strconv.Atoi(expr)
        h[""] += num
    } else {
        h[expr]++
    }
    return h
}

func Combine(this, that Polynomial, op byte) Polynomial {
    switch op {
    case '+':
        return this.Add(that)
    case '-':
        return this.Sub(that)
    case '*':
        return this.Mul(that)
    }
    return nil
}

func Parse(expr string) Polynomial {
    n := len(expr)
    polys := []Polynomial{}
    ops := []byte{}
    for i := 0; i < n; i++ {
        c := expr[i]
        if c == '(' {
            pa, j := 0, i
            for ; j < n; j++ {
                if expr[j] == '(' {
                    pa++
                } else if expr[j] == ')' {
                    pa--
                }
                if pa == 0 {
                    break
                }
            }
            polys = append(polys, Parse(expr[i+1:j]))
            i = j
        } else if isLetterOrNumber(c) {
            j := i
            ok := true
            for ; j < n; j++ {
                if expr[j] == ' ' {
                    polys = append(polys, MakePolynomial(expr[i:j]))
                    ok = false
                    break
                }
            }
            if ok {
                polys = append(polys, MakePolynomial(expr[i:]))
            }
            i = j
        } else if c != ' ' {
            ops = append(ops, c)
        }
    }

    if len(polys) == 0 {
        return nil
    }

    for j := len(ops) - 1; j >= 0; j-- {
        if ops[j] == '*' {
            polys[j] = Combine(polys[j], polys[j+1], ops[j])
            copy(polys[j+1:], polys[j+2:])
            polys = polys[:len(polys)-1]
            copy(ops[j:], ops[j+1:])
            ops = ops[:len(ops)-1]
        }
    }
    res := polys[0]
    for j := range ops {
        res = Combine(res, polys[j+1], ops[j])
    }
    return res
}

func isLetterOrNumber(c byte) bool {
    return isLetter(c) || isNumber(c)
}

func isLetter(c byte) bool {
    return 'a' <= c && c <= 'z'
}

func isNumber(c byte) bool {
    return '0' <= c && c <= '9'
}

func (this Polynomial) Add(that Polynomial) Polynomial {
    h := map[string]int{}
    for k, v := range this {
        h[k] += v
    }
    for k, v := range that {
        h[k] += v
    }
    return h
}

func (this Polynomial) Sub(that Polynomial) Polynomial {
    h := map[string]int{}
    for k, v := range this {
        h[k] += v
    }
    for k, v := range that {
        h[k] -= v
    }
    return h
}

func (this Polynomial) Mul(that Polynomial) Polynomial {
    h := map[string]int{}
    for k1, v1 := range this {
        for k2, v2 := range that {
            ks1 := strings.Split(k1, "$")
            ks2 := strings.Split(k2, "$")
            ks := make([]string, 0, len(ks1)+len(ks2))
            for _, k := range ks1 {
                if k == "" {
                    continue
                }
                ks = append(ks, k)
            }
            for _, k := range ks2 {
                if k == "" {
                    continue
                }
                ks = append(ks, k)
            }
            sort.Strings(ks)
            h[strings.Join(ks, "$")] += v1 * v2
        }
    }
    return h
}

func (this Polynomial) Evaluate(evalMap map[string]int) Polynomial {
    h := map[string]int{}
    for k, v := range this {
        ks := []string{}
        for _, x := range strings.Split(k, "$") {
            if val, ok := evalMap[x]; ok {
                v *= val
            } else {
                ks = append(ks, x)
            }
        }
        h[strings.Join(ks, "$")] += v
    }
    return h
}

func (this Polynomial) ToList() []string {
    res, ks := []string{}, make([]string, 0, len(this))
    for k := range this {
        ks = append(ks, k)
    }
    sort.Slice(ks, func(i, j int) bool {
        if ks[i] == "" {
            return false
        } else if ks[j] == "" {
            return true
        }
        ss1 := strings.Split(ks[i], "$")
        ss2 := strings.Split(ks[j], "$")
        if len(ss1) != len(ss2) {
            return len(ss1) > len(ss2)
        }
        for i := range ss1 {
            if ss1[i] != ss2[i] {
                return ss1[i] < ss2[i]
            }
        }
        return true
    })
    for _, k := range ks {
        v := this[k]
        if v == 0 {
            continue
        }
        s := strconv.Itoa(v)
        for _, x := range strings.Split(k, "$") {
            if x == "" {
                continue
            }
            s += "*" + x
        }
        res = append(res, s)
    }
    return res
}

func main() {
    // Example 1:
    // Input: expression = "e + 8 - a + 5", evalvars = ["e"], evalints = [1]
    // Output: ["-1*a","14"]
    fmt.Println(basicCalculatorIV("e + 8 - a + 5", []string{"e"}, []int{1})) // ["-1*a","14"]
    // Example 2:
    // Input: expression = "e - 8 + temperature - pressure", evalvars = ["e", "temperature"], evalints = [1, 12]
    // Output: ["-1*pressure","5"]
    fmt.Println(basicCalculatorIV("e - 8 + temperature - pressure", []string{"e", "temperature"}, []int{1, 12})) // ["-1*pressure","5"]
    // Example 3:
    // Input: expression = "(e + 8) * (e - 8)", evalvars = [], evalints = []
    // Output: ["1*e*e","-64"]
    fmt.Println(basicCalculatorIV("(e + 8) * (e - 8)", []string{}, []int{})) // ["1*e*e","-64"]
}