# 2128 · Use of list built-in methods
# Description
# Please refine the code in solution.py to implement the following functions.

# get_len: return the length of the input list
# get_max: return the maximum value of the elements of the input list
# get_min: returns the smallest element of the input list
# pop_list: returns the result of the input list after the last element value has been removed
# Each of these functions has only one argument list_in.

# We will import your code from solution.py in main.py and run it. If your code is logically correct and runs successfully, 
# the program will return a list as the result of the operation.he result.


# Example
# The evaluation machine will execute your code by executing python main.py {input_path} 
# and the test data will be placed in the file corresponding to input_path. You can see how the code works in main.py.

# Example 1

# Input.

# [23, 65, 4, 5, 1, 78, 3]
# Output.

# 7
# 78
# 1
# [23, 65, 4, 5, 1, 78]
# Sample 2

# Input.

# [34, 2, 54, 3, 2, 6]
# Output.

# 6
# 54
# 2
# [34, 2, 54, 3, 2]

def get_len(list_in: list) -> int:
    """
    :param list_in: The first input list
    :return: The length of the list_in
    """
    # write your code here
    return len(list_in)

def get_max(list_in: list) -> int:
    """
    :param list_in: The first input list
    :return: The largest value of list_in
    """
    # write your code here
    return max(list_in)

def get_min(list_in: list) -> int:
    """
    :param list_in: The first input list
    :return: The smallest value of list_in
    """
    # write your code here
    return min(list_in)

def pop_list(list_in: list) -> list:
    """
    :param tuple_in: The first input tuple
    :return: The list delete the last element
    """
    # write your code here
    list_in.pop()
    return list_in