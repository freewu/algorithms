# 2240 · Update string
# Description
# Write Python code that implements a function called str_update. 
# This function will extract the first 7 bytes of the given string, 
# add `` welcomes!`'' to the string, and finally return the updated string.

# Please complete the str_update function in solution.py, 
# and we will run your code in main.py by importing it and checking that the end result is correct.t.

# Space between letters after change
# Letters need to be strictly case sensitive
# Example
# The quizzer runs main.py by executing python main.py,
# which outputs the updated string after running, and you can see how the code is running in main.py.

# Example 1
# If the string is:
#   Hello, Mr.Green
# then the updated string will be:
#   Mr.Green welcomes!

# Example 2
# If the string is:
#   Hello, John
# then the updated string will be:
#   John welcomes!

def str_update(txt: str) -> str:
    """
    :param txt: a input string
    :return: a changed string
    """
    # write your code here
    # # 使用 , 分割 
    # arr = txt.split(", ")
    # return  arr[1] + " welcomes!"
    return txt[7:] + " welcomes!"

if __name__ == "__main__":
    print(str_update("Hello, Mr.Green")) # Mr.Green welcomes!
    print(str_update("Hello, John")) # John welcomes!