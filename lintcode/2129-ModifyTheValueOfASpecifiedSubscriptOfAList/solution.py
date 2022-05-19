# 2129 · Modify the value of a specified subscript of a list
# # Description
# Please refine the code in solution.py to implement the update function, 
# which changes the value of the list list_in subscripted by idx to value. 
# The update function takes three arguments: list_in, idx, and value. 
# We will import the code you refined in solution.py into main.py and run it. 
# If your code is logically correct and runs successfully, the program will return a list as the result of the operation.

# The elements inside the passed parameter list_in are all of string type, and the value is also of string type.

# Example
# The evaluation machine will execute your code by executing python main.py {input_path} 
# and the test data will be placed in the file corresponding to input_path. 
# You can see how the code works in main.py.

# Example 1
# Input.
#   ['angelica', 'uh', 'outer', 'skiff', 'deanna', 'recusant']
#   1
#   'outpouring'
# Output.
#   ['angelica', 'outpouring', 'outer', 'skiff', 'deanna', 'recusant']

# Example 2
# Input.
#   ['annihilator', 'endosperm', 'spicebush']
#   1
#   'ciliated'
# Output.
#   ['annihilator', 'ciliated', 'spicebush']

def update(list_in: list, idx: int, value: str) -> list:
    """
    :param list_in: The first input list
    :param idx: The index of update element
    :param value: The new element after update old element
    :return: The new list after update the index element
    """
    # write your code here
    list_in[idx] = value
    return list_in