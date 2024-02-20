// 2803. Factorial Generator
// Write a generator function that takes an integer n as an argument and returns a generator object which yields the factorial sequence.
// The factorial sequence is defined by the relation n! = n * (n-1) * (n-2) * ... * 2 * 1​​​.
// The factorial of 0 is defined as 1.

// Example 1:
// Input: n = 5
// Output: [1,2,6,24,120]
// Explanation: 
// const gen = factorial(5)
// gen.next().value // 1
// gen.next().value // 2
// gen.next().value // 6
// gen.next().value // 24
// gen.next().value // 120

// Example 2:
// Input: n = 2
// Output: [1,2]
// Explanation: 
// const gen = factorial(2) 
// gen.next().value // 1 
// gen.next().value // 2 

// Example 3:
// Input: n = 0
// Output: [1]
// Explanation: 
// const gen = factorial(0) 
// gen.next().value // 1 
 
// Constraints:
//     0 <= n <= 18

/**
 * @param {number} n
 * @yields {number}
 */
function* factorial(n) {
    // 处理输入为 0 的
    if (n <= 1 ) return yield 1;
    function fib(n) {
        if(n < 1) return 1
        return n * fib(n-1)
    }
    // 生成一个 
    let arr = [];
    for ( let i = 1 ; i <= n; i++ ) {
        arr.push(fib(i))
    }
    // 返回一个 Generator
    for(const item of arr) {
        yield item;
    }
};

function* factorial1(n) {
    if( n === 0 ) yield 1
    let res = 1;
    for( let i = 1; i <= n; i++ ) yield res *= i
};

// Example 1:
let gen = factorial(5)
console.log(gen.next().value) // 1
console.log(gen.next().value) // 2
console.log(gen.next().value) // 6
console.log(gen.next().value) // 24
console.log(gen.next().value) // 120

// Example 2:
gen = factorial(2) 
console.log(gen.next().value) // 1 
console.log(gen.next().value) // 2 

// Example 3:
gen = factorial(0) 
console.log(gen.next().value) // 1 

// Example 4:
gen = factorial(1) 
console.log(gen.next().value) // 1 


// Example 1:
gen = factorial1(5)
console.log(gen.next().value) // 1
console.log(gen.next().value) // 2
console.log(gen.next().value) // 6
console.log(gen.next().value) // 24
console.log(gen.next().value) // 120

// Example 2:
gen = factorial1(2) 
console.log(gen.next().value) // 1 
console.log(gen.next().value) // 2 

// Example 3:
gen = factorial1(0) 
console.log(gen.next().value) // 1 

// Example 4:
gen = factorial1(1) 
console.log(gen.next().value) // 1 