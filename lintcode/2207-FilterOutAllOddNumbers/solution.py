# 2207 · Filter out all odd numbers
# Description
# Write Python code that implements a function named odd. The function will take a list of integers,
# return all the odd numbers by using a generator, and output them one by one in order as a list.
# Write the generator that creates the odd numbers in solution.py, 
# and we will run your code in main.py by importing it to check that your code does the above.

# The number to be generated is an odd number

# Example 1
# When the input data is 15,10,54,30,74,12,11,24,9,43,52,8,14,38,47,77,99,89,26, the
# The output data is:
#   [15, 11, 9, 43, 47, 77, 99, 89]

# Example 2
# When the input data is 100,25,14,15,80,61,23,30,7,31,95,42,100,21,96,64,31,59,66,47,27,82,79,77, the
# The output data is:
#   [25, 15, 61, 23, 7, 31, 95, 21, 31, 59, 47, 27, 79, 77]

def odd(list_in):
    # write your code here
    # 取余
    # for x in list_in:
    #     if x % 2 == 1:
    #         yield x
    
    # 位运算
    for x in list_in:
        if x & 1:
            yield x