// 2797. Partial Function with Placeholders
// Given a function fn and an array args, return a function partialFn. 
// Placeholders "_" in the args should be replaced with values from restArgs starting from index 0. 
// Any remaining values in the restArgs should be added at the end of the args.

// partialFn should return a result of fn. 
// fn should be called with the elements of the modified args passed as separate arguments.

// Example 1:
// Input: fn = (...args) => args, args = [2,4,6], restArgs = [8,10]
// Output: [2,4,6,8,10]
// Explanation: 
// const partialFn = partial(fn, args)
// const result = partialFn(...restArgs) 
// console.log(result) // [2,4,6,8,10]
// There are no placeholders "_" in args therefore restArgs is just added at the end of args. Then the elements of the args are passed as separate arguments to fn, which returns passed arguments as an array.

// Example 2:
// Input: fn = (...args) => args, args = [1,2,"_",4,"_",6], restArgs = [3,5]
// Output: [1,2,3,4,5,6]
// Explanation: 
// const partialFn = partial(fn, args) 
// const result = partialFn(...restArgs) 
// console.log(result) // [1,2,3,4,5,6] 
// Placeholders "_" are replaced with values from the restArgs. Then the elements of the args are passed as separate arguments to fn, which returns passed arguments as an array.

// Example 3:
// Input: fn = (a, b, c) => b + a - c, args = ["_", 5], restArgs = [5, 20]
// Output: -10
// Explanation: 
// const partialFn = partial(fn, args)
// const result = partialFn(...restArgs)
// console.log(result) // -10
// Placeholder "_" is replaced with 5 and 20 is added at the end of args. Then the elements of the args are passed as separate arguments to fn, which returns -10 (5 + 5 - 20).
 
// Constraints:
//         fn is a function
//         args and restArgs are valid JSON arrays
//         1 <= args.length <= 5 * 10^4
//         1 <= restArgs.length <= 5 * 10^4
//         0 <= number of placeholders <= restArgs.length

/**
 * @param {Function} fn
 * @param {Array} args
 * @return {Function}
 */
var partial = function(fn, args) {
    return function(...restArgs) {
		let newArgs = [...args];
        // 替换 _ 
        args.forEach((item, index) => {
            if (item === '_') {
                newArgs[index] = restArgs.shift();
            }
        });
        // 多出的追加到参数后面
        if (restArgs.length) newArgs = newArgs.concat(restArgs);
        return fn(...newArgs);
    }
};

var partial1 = function(fn, args) {
	return function(...restArgs) {
		let i = 0
        args = args.map(v => v === "_" ? restArgs[i++] : v)
        if (i < restArgs.length) args = Array.of(...args, ...restArgs.slice(i))
        return fn(...args)
    }
};

// Example 1:
// There are no placeholders "_" in args therefore restArgs is just added at the end of args.
let partialFn = partial((...args) => args, [2,4,6])
let result = partialFn(...[8,10]) 
console.log(result) // [2,4,6,8,10]


// Example 2:
// Placeholders "_" are replaced with values from the restArgs. 
// Then the elements of the args are passed as separate arguments to fn, which returns passed arguments as an array.
partialFn = partial((...args) => args, [1,2,"_",4,"_",6]) 
result = partialFn(...[3,5]) 
console.log(result) // [1,2,3,4,5,6] 

// Example 3:
// Placeholder "_" is replaced with 5 and 20 is added at the end of args. 
// Then the elements of the args are passed as separate arguments to fn, which returns -10 (5 + 5 - 20).
partialFn = partial((a, b, c) => b + a - c, ["_", 5])
result = partialFn(...[5, 20])
console.log(result) // -10


partialFn = partial1((...args) => args, [2,4,6])
result = partialFn(...[8,10]) 
console.log(result) // [2,4,6,8,10]


// Example 2:
// Placeholders "_" are replaced with values from the restArgs. 
// Then the elements of the args are passed as separate arguments to fn, which returns passed arguments as an array.
partialFn = partial1((...args) => args, [1,2,"_",4,"_",6]) 
result = partialFn(...[3,5]) 
console.log(result) // [1,2,3,4,5,6] 

// Example 3:
// Placeholder "_" is replaced with 5 and 20 is added at the end of args. 
// Then the elements of the args are passed as separate arguments to fn, which returns -10 (5 + 5 - 20).
partialFn = partial1((a, b, c) => b + a - c, ["_", 5])
result = partialFn(...[5, 20])
console.log(result) // -10
