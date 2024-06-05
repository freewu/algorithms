package main

// 772. Basic Calculator III
// Implement a basic calculator to evaluate a simple expression string.
// The expression string contains only non-negative integers, '+', '-', '*', '/' operators, and open '(' and closing parentheses ')'. The integer division should truncate toward zero.

// You may assume that the given expression is always valid. 
// All intermediate results will be in the range of [-2^31, 2^31 - 1].

// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

// Example 1:
// Input: s = "1+1"
// Output: 2

// Example 2:
// Input: s = "6-4/2"
// Output: 4

// Example 3:
// Input: s = "2*(5+5*2)/3+(6/2+8)"
// Output: 21

// Constraints:
//     1 <= s <= 10^4
//     s consists of digits, '+', '-', '*', '/', '(', and ')'.
//     s is a valid expression.

import "fmt"

// stack
func calculate(s string) int {
    // 也可以直接用栈处理，但是因为数字和运算符类型不一样，通常使用两个栈分别存储数字和运算符。
    // 不过为了体现一个栈的思路，这里用一个栈处理
    getNum := func(i int) (int, int) {
        p, sum := i, 0
        for ;p<len(s) && s[p]>='0' && s[p]<='9'; p++ {
            sum = sum*10 + int(s[p]-'0')
        }
        if p > i {
            return sum, p-1
        } 
        return -1, i
    }
    // 二元加减乘除计算
    oper := func (a,b int, oper byte ) int {
        switch(oper){
            case '+': return a+b
            case '-': return a-b
            case '*': return a*b
            case '/': return a/b
        }
        return 0
    }
    // 往栈中push一个num，视情形弹栈计算和压栈
    pushNum := func (stack *[]interface{}, num int) {
        st := *stack
        for len(st)>0 {
            //获取前一个符号（肯定不是数字）
            sign := st[len(st)-1].(byte)
            if sign == '*' || sign == '/' {
                a := st[len(st)-2].(int)
                st = st[:len(st)-2]
                num = oper(a, num, sign)
            } else { // 前一个是+-号或者左括号(，当前数字都直接压栈
                break
            }
        }
        *stack = append(st, num)
    }
    // 计算只有加减号的中缀表达式
    // 整个字符串解析完后，stack中就只剩下包含加减号的表达式
    // 注意：这里需要从前往后计算，因为有加减法，正反顺序计算结果是不一样的
    calcSubAdd := func (stack []interface{}, p int) int {
        num := stack[p].(int)
        for i := p + 1; i < len(stack); i += 2 {
            sign := stack[i].(byte)
            b := stack[i+1].(int)
            num = oper(num,b, sign)
        }
        return num
    }
    stack := []interface{}{}
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c >= '0' && c <= '9' {
            num, ed := getNum(i)
            i = ed
            pushNum(&stack, num)
        } else if c == ')' {//弹栈计算直到遇到'('
            // 查找 '('
            k := len(stack)-2 //倒数第一个肯定是数字，倒数第二才是符号
            for {
                //不用考虑越界，因为题目说了表达式有效，前面肯定有左括号对应
                
                sign := stack[k].(byte)
                if sign == '(' {
                    break
                } else {
                    //a := stack[k-1].(int)
                    //num = oper(a, num, sign)
                    k -= 2
                }
            }
            num := calcSubAdd(stack, k+1)
            stack = stack[:k]
            // 处理完括号后的当前num处理方法与直接读取到的一个num无异
            pushNum(&stack, num)
        } else {// 运算符和左括号都直接压栈
            stack = append(stack, c)
        }
    }
    // fmt.Println(stack)
    return calcSubAdd(stack, 0)
}

func calculate1(s string) int {
    var cal func(*[]byte) int
    cal = func(ss *[]byte) int {
        res, st, num, sign := 0, []int{}, 0, byte('+')
        for len(*ss) > 0 {
            c := (*ss)[0]
            (*ss) = (*ss)[1:] // pop left
            if c >= '0' && c <= '9' {
                num = num * 10 + int(c-'0')
            }
            if c == '(' {
                num = cal(ss)
            }
            if (!(c >= '0' && c <= '9') && c != ' ') || len(*ss) == 0 {
                switch sign {
                case '+':
                    st = append(st, num)
                case '-':
                    st = append(st, -num)
                case '*':
                    st[len(st)-1] *= num
                case '/':
                    st[len(st)-1] /= num
                }
                num, sign = 0, c
            }
            if c == ')' {
                break
            }
        }
        for _, x := range st {
            res += x
        }
        return res
    }
    ss := []byte(s)
    return cal(&ss)
}

func main() {
    // Example 1:
    // Input: s = "1+1"
    // Output: 2
    fmt.Println(calculate("1+1")) // 2
    // Example 2:
    // Input: s = "6-4/2"
    // Output: 4
    fmt.Println(calculate("6-4/2")) // 4
    // Example 3:
    // Input: s = "2*(5+5*2)/3+(6/2+8)"
    // Output: 21
    fmt.Println(calculate("2*(5+5*2)/3+(6/2+8)")) // 21

    fmt.Println(calculate1("1+1")) // 2
    fmt.Println(calculate1("6-4/2")) // 4
    fmt.Println(calculate1("2*(5+5*2)/3+(6/2+8)")) // 21
}