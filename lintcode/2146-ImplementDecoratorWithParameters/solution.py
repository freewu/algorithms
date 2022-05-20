# 2146 · Implement Decorator with Parameters
# Description
# Please implement a decorator named repeat_func, it takes an integer n as parameter which indicate how many times to repeat running the decorated function.

# Besides repeat running the decorated function, you code need to print before function run before the function be called,
# and print after function run after all functions finished running.

# You can check the code in main.py to see how repeat_func be called.

# Example
# The input data contains an integer n, this value will be put into the parameter list when running your code.
# For example, if n = 1, the judge system will execute command python main.py 1 to run the whole project, and then your code should output

# before function run
# function run
# after function run
# If n = 2, your code need to output

# before function run
# function run
# function run
# after function run

def repeat_func(n):
    def wrapper(func):
        def inner():
            print('before function run')
            for i in range(n):
                func()
            print('after function run')
        return inner
    return wrapper