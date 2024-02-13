// 2628. JSON Deep Equal
// Given two values o1 and o2, return a boolean value indicating whether two values, o1 and o2, are deeply equal.
// For two values to be deeply equal, the following conditions must be met:
//         If both values are primitive types, they are deeply equal if they pass the === equality check.
//         If both values are arrays, they are deeply equal if they have the same elements in the same order, and each element is also deeply equal according to these conditions.
//         If both values are objects, they are deeply equal if they have the same keys, and the associated values for each key are also deeply equal according to these conditions.

// You may assume both values are the output of JSON.parse. In other words, they are valid JSON.
// Please solve it without using lodash's _.isEqual() function

 
// Example 1:
// Input: o1 = {"x":1,"y":2}, o2 = {"x":1,"y":2}
// Output: true
// Explanation: The keys and values match exactly.

// Example 2:
// Input: o1 = {"y":2,"x":1}, o2 = {"x":1,"y":2}
// Output: true
// Explanation: Although the keys are in a different order, they still match exactly.

// Example 3:
// Input: o1 = {"x":null,"L":[1,2,3]}, o2 = {"x":null,"L":["1","2","3"]}
// Output: false
// Explanation: The array of numbers is different from the array of strings.

// Example 4:
// Input: o1 = true, o2 = false
// Output: false
// Explanation: true !== false
 
// Constraints:
//         1 <= JSON.stringify(o1).length <= 10^5
//         1 <= JSON.stringify(o2).length <= 10^5
//         maxNestingDepth <= 1000

/**
 * @param {null|boolean|number|string|Array|Object} o1
 * @param {null|boolean|number|string|Array|Object} o2
 * @return {boolean}
 */
var areDeeplyEqual = function(o1, o2) {
    // 类型不同 直接返回 false
    if (Object.prototype.toString.call(o1) !== Object.prototype.toString.call(o2)) {
        return false;
    }
    // 数组和对象判断
    if (o1 !== null && typeof o1 === 'object') {
        // 判断 key 数量相同 && key 值都相同
        return Object.keys(o1).length === Object.keys(o2).length &&
            Object.entries(o1).every(
                // 如果是对象递归判断
                ([key, value]) => areDeeplyEqual(value, o2[key])
            )
    }
    // 基本类型判断
    return o1 === o2;
};

// Input: o1 = {"x":1,"y":2}, o2 = {"x":1,"y":2}
// Output: true
// Explanation: The keys and values match exactly.

// Example 2:
// Input: o1 = {"y":2,"x":1}, o2 = {"x":1,"y":2}
// Output: true
// Explanation: Although the keys are in a different order, they still match exactly.

// Example 3:
// Input: o1 = {"x":null,"L":[1,2,3]}, o2 = {"x":null,"L":["1","2","3"]}
// Output: false
// Explanation: The array of numbers is different from the array of strings.

// Example 4:
// Input: o1 = true, o2 = false
// Output: false

console.log(areDeeplyEqual({"x":1,"y":2},{"x":1,"y":2})); // true
console.log(areDeeplyEqual({"x":1,"y":2},{"y":2,"x":1})); // true
console.log(areDeeplyEqual({"x":null,"L":[1,2,3]},{"x":null,"L":["1","2","3"]})); // false
console.log(areDeeplyEqual(true,false)); // false