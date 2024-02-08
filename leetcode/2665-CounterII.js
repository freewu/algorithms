// 2665. Counter II
// Write a function createCounter. It should accept an initial integer init. 
// It should return an object with three functions.
// The three functions are:
//         increment() increases the current value by 1 and then returns it.
//         decrement() reduces the current value by 1 and then returns it.
//         reset() sets the current value to init and then returns it.
 
// Example 1:
// Input: init = 5, calls = ["increment","reset","decrement"]
// Output: [6,5,4]
// Explanation:
//         const counter = createCounter(5);
//         counter.increment(); // 6
//         counter.reset(); // 5
//         counter.decrement(); // 4

// Example 2:
// Input: init = 0, calls = ["increment","increment","decrement","reset","reset"]
// Output: [1,2,1,0,0]
// Explanation:
//         const counter = createCounter(0);
//         counter.increment(); // 1
//         counter.increment(); // 2
//         counter.decrement(); // 1
//         counter.reset(); // 0
//         counter.reset(); // 0
 
// Constraints:
//         -1000 <= init <= 1000
//         0 <= calls.length <= 1000
//         calls[i] is one of "increment", "decrement" and "reset"

/**
 * @param {integer} init
 * @return { increment: Function, decrement: Function, reset: Function }
 */
var createCounter = function(init) {
    let value = init;
    let increment = function() {
        return value += 1;
    };
    let reset = function() {
        return value = init;
    };
    let decrement = function() {
        return value -= 1;
    };
    return {
        increment,
        reset,
        decrement,
    }
};

// best solution 
var createCounter1 = function(init) {
    let origin = init
    return{
        increment : (value)=>++origin ,
        decrement : (value)=>--origin ,
        reset : (value)=>origin = init ,
    }
};

/**
 * const counter = createCounter(5)
 * counter.increment(); // 6
 * counter.reset(); // 5
 * counter.decrement(); // 4
 */

const counter = createCounter(5)
console.log(counter.increment()); // 6
console.log(counter.reset()); // 5
console.log(counter.decrement()); // 4

const counter1 = createCounter(0);
console.log(counter1.increment()); // 1
console.log(counter1.increment()); // 2
console.log(counter1.decrement()); // 1
console.log(counter1.reset()); // 0
console.log(counter1.reset()); // 0


const counter11 = createCounter1(5)
console.log(counter11.increment()); // 6
console.log(counter11.reset()); // 5
console.log(counter11.decrement()); // 4

const counter12 = createCounter1(0);
console.log(counter12.increment()); // 1
console.log(counter12.increment()); // 2
console.log(counter12.decrement()); // 1
console.log(counter12.reset()); // 0
console.log(counter12.reset()); // 0