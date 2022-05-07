# 2089 · Implement timer decorator
# Description
# Decorator is a special python syntax, You can written on the above line of the function to be decorated with the syntax @decorator_name like this:

# @decorator_name
# def func():
#     # do something
#     pass

# Decorator support some processing with certain universality before and after the function runs. 
# These kind of processings usually need to package multiple functions, so decorator could help us to avoid duplicate codes and improve the readability.

# The task of this problem is to implement a timer decorator, this decorator named timer. 
# We could wrap timer to any functions so that when the function be called, the cost time of the function will be recorded and automatically printed out.

# @timer
# def func():
#     pass

# Your task is to edit the file decorators.py and implement a timer decorator which will print the the function's name and the time consumed.

import time

# 定义 decorator f 参数就是需要执行的方法体
def timer(f):
    def log_time():
        t1 = time.time()
        res = f()
        t2 = time.time()
        print('function %s cost %.1f seconds' % (f.__name__, (t2 - t1)))
        return res
    return log_time

@timer
def func_a():
    time.sleep(0.1)

@timer
def func_b():
    time.sleep(0.2)

# python 2089-ImplementTimerDecorator.py
if __name__=="__main__":
    func_a()
    func_b()