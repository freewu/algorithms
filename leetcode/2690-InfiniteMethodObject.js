// 2690. Infinite Method Object
// Write a function that returns an infinite-method object.
// An infinite-method object is defined as an object that allows you to call any method and it will always return the name of the method.
// For example, if you execute obj.abc123(), it will return "abc123".

// Example 1:
// Input: method = "abc123"
// Output: "abc123"
// Explanation:
// const obj = createInfiniteObject();
// obj['abc123'](); // "abc123"
// The returned string should always match the method name.

// Example 2:
// Input: method = ".-qw73n|^2It"
// Output: ".-qw73n|^2It"
// Explanation: The returned string should always match the method name.
 
// Constraints:
//         0 <= method.length <= 1000

/**
 * @return {Object}
 */
var createInfiniteObject = function() {
    return new Proxy(
        {},
        {
            get: (_, prop) => () => prop.toString(),
        },
    );
};

/**
 * const obj = createInfiniteObject();
 * obj['abc123'](); // "abc123"
 */

const obj = createInfiniteObject();
obj['abc123'](); // "abc123"
obj['.-qw73n|^2It'](); // ".-qw73n|^2It"

/*
# Proxy

    Proxy 可以理解成，在目标对象之前架设一层“拦截”，外界对该对象的访问，都必须先通过这层拦截，因此提供了一种机制，可以对外界的访问进行过滤和改写。
    Proxy 这个词的原意是代理，用在这里表示由它来“代理”某些操作，可以译为“代理器”。
    ES6 原生提供 Proxy 构造函数，用来生成 Proxy 实例。

        var proxy = new Proxy(target, handler); //target为目标对象，handler参数也是一个对象，用来定制拦截行为。

    定义一个拦截读取属性行为的例子

        var proxy = new Proxy(
            {}, 
            {
                // 拦截了所有属性 只返回 
                get: function(target, propKey) {
                    return 'bluefrog';
                }
            }
        );

        proxy.time  // bluefrog
        proxy.name  // bluefrog
        proxy.title // bluefrog

# Proxy的方法
    
    (target : 目标对象；propKey : 属性名 ；receiver : proxy实例本身；value : 属性值；)

    get(target, propKey, receiver)：拦截对象属性的读取。
    set(target, propKey, value, receiver)：拦截对象属性的设置(返回布尔值)。
    has(target, propKey)：拦截propKey in proxy的操作(返回布尔值)。
    deleteProperty(target, propKey)：拦截delete proxy[propKey]的操作(返回布尔值)。
    ownKeys(target)：拦截Object.getOwnPropertyNames(proxy)、Object.getOwnPropertySymbols(proxy)、Object.keys(proxy)、for...in循环，返回一个数组。该方法返回目标对象所有自身的属性的属性名，而Object.keys()的返回结果仅包括目标对象自身的可遍历属性。
    getOwnPropertyDescriptor(target, propKey)：拦截Object.getOwnPropertyDescriptor(proxy, propKey)，返回属性的描述对象。
    defineProperty(target, propKey, propDesc)：拦截Object.defineProperty(proxy, propKey, propDesc）、Object.defineProperties(proxy, propDescs)，返回一个布尔值。
    preventExtensions(target)：拦截Object.preventExtensions(proxy)，返回一个布尔值。
    getPrototypeOf(target)：拦截Object.getPrototypeOf(proxy)，返回一个对象。
    isExtensible(target)：拦截Object.isExtensible(proxy)，返回一个布尔值。
    setPrototypeOf(target, proto)：拦截Object.setPrototypeOf(proxy, proto)，返回一个布尔值。如果目标对象是函数，那么还有两种额外操作可以拦截。
    apply(target, object, args)：拦截 Proxy 实例作为函数调用的操作，比如proxy(...args)、proxy.call(object, ...args)、proxy.apply(...)。
    construct(target, args)：拦截 Proxy 实例作为构造函数调用的操作，比如new proxy(...args)。
*/