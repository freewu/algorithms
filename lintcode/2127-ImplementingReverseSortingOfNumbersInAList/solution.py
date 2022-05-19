# 2127 · Implementing reverse sorting of numbers in a list
# Description
# Please refine the code in solution.py to implement the list_sort function. 
# The list_sort function has one argument list_in,
# please sort the numbers in the argument list_in in reverse order and return it.
# We will import the code you refined in solution.py in main.py and run it. 
# If your code is logically correct and runs successfully, the program will return a list as the result of the calculation.

# The elements in the passed parameter list_in are all numbers

# Example
# The evaluation machine will execute your code by executing python main.py {input_path} 
# and the test data will be placed in the file corresponding to input_path. You can see how the code works in main.py.

# Example 1
# Input.
#   list_in = [23,65,4,5,1,78,3]
# Output.
#    [78, 65, 23, 5, 4, 3, 1]

# Example 2
# Input.
#   list_in = [34,2,54,3,2,6]
# Output.
#   [54, 34, 6, 3, 2, 2]
def list_sort(list_in: list) -> list:
    """
    :param list_in: The first input list
    :return: A list sorted from largest to smallest
    """
    # write your code here
    list_in.sort(reverse = True)
    return list_in

if __name__ == '__main__':
    print(list_sort([23,65,4,5,1,78,3])) # [78, 65, 23, 5, 4, 3, 1]
    print(list_sort([34,2,54,3,2,6])) # [54, 34, 6, 3, 2, 2]