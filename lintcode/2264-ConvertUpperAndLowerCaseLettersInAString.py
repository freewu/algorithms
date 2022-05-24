# 2264 · Convert upper and lower case letters in a string
# Description
# Please refine the code in solution.py to implement the str_conversion function.
# The str_conversion function has one argument str_in,
# please convert the upper and lower case letters in the argument str_in. 
# We will import the code you have refined in solution.py into main.py and run it. 
# If your code is logically correct and runs successfully, the program will return a new string as the result of the operation.

# Example
# The evaluation machine will execute your code by executing python main.py {input_path}
# and the test data will be placed in the file corresponding to input_path. 
# You can see how the code works in main.py.

# Example 1
# Input.
#   this is a string example ABCD
# Output.
#   THIS IS A STRING EXAMPLE abcd

# Example 2
# Input.
#   JIU ZHANG da fa HAO
# Output.
#   jiu zhang DA FA hao

def str_conversion(str_in: str) -> str:
    """
    :param str_in: The first input str
    :return: The new str after conversion str case
    """
    return str_in.swapcase()

def str_conversion1(str_in: str) -> str:
    """
    :param str_in: The first input str
    :return: The new str after conversion str case
    """
    chars = list()
    for i in str_in:
        if i == ' ':
            chars.append(' ')
        elif 97 <= ord(i) <= 122: # lower -> upper
            chars.append(i.upper())
        elif 65 <= ord(i) < 91: # upper -> lower
            chars.append(i.lower()) 
    return ''.join(chars)

if __name__ == "__main__":
    print(str_conversion("this is a string example ABCD")) # THIS IS A STRING EXAMPLE abcd
    print(str_conversion("JIU ZHANG da fa HAO")) # jiu zhang da fa hao

    print(str_conversion1("this is a string example ABCD")) # THIS IS A STRING EXAMPLE abcd
    print(str_conversion1("JIU ZHANG da fa HAO")) # jiu zhang da fa hao