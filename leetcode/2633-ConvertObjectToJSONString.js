// 2633. Convert Object to JSON String
// Given a value, return a valid JSON string of that value. 
// The value can be a string, number, array, object, boolean, or null. 
// The returned string should not include extra spaces. 
// The order of keys should be the same as the order returned by Object.keys().
// Please solve it without using the built-in JSON.stringify method.

// Example 1:
// Input: object = {"y":1,"x":2}
// Output: {"y":1,"x":2}
// Explanation: 
// Return the JSON representation.
// Note that the order of keys should be the same as the order returned by Object.keys().

// Example 2:
// Input: object = {"a":"str","b":-12,"c":true,"d":null}
// Output: {"a":"str","b":-12,"c":true,"d":null}
// Explanation:
// The primitives of JSON are strings, numbers, booleans, and null.

// Example 3:
// Input: object = {"key":{"a":1,"b":[{},null,"Hello"]}}
// Output: {"key":{"a":1,"b":[{},null,"Hello"]}}
// Explanation:
// Objects and arrays can include other objects and arrays.

// Example 4:
// Input: object = true
// Output: true
// Explanation:
// Primitive types are valid inputs.
 
// Constraints:
//         value is a valid JSON value
//         1 <= JSON.stringify(object).length <= 10^5
//         maxNestingLevel <= 1000
//         all strings contain only alphanumeric characters

/**
 * @param {null|boolean|number|string|Array|Object} object
 * @return {string}
 */
var jsonStringify = function(object) {
    // null
    if (object === null) return "null";
    // Array
    if (Array.isArray(object)) return `[${(object.map(jsonStringify)).join(",")}]`
    // object 递归处理每个 value
    if (typeof object === "object") {
        return `{${Object.keys(object).map((key) => `"${key}":${jsonStringify(object[key])}`).join(",")}}`
    }
    // 处理 String
    if (typeof object === "string") return `"${object}"`
    // 处理 true / false / undefined / NaN / 1 / 1.0
    return String(object)
};

console.log(jsonStringify({"y":1,"x":2})) // {"y":1,"x":2}
console.log(jsonStringify({"a":"str","b":-12,"c":true,"d":null})) // {"a":"str","b":-12,"c":true,"d":null}
console.log(jsonStringify({"key":{"a":1,"b":[{},null,"Hello"]}})) // {"key":{"a":1,"b":[{},null,"Hello"]}}
console.log(jsonStringify(true)) // true
console.log(jsonStringify(false)) // false
console.log(jsonStringify(undefined)) // undefined

console.log(jsonStringify(123)) // 123
console.log(jsonStringify(123.456)) // 123.456
console.log(jsonStringify(NaN)) // NaN