# 2276 · Modify and delete elements in the dictionary
# Description
# We've written the modify_and_delete_elements_in_the_dictionary function in solution.py for you,
# where dict_1 represents the initial dictionary and the function eventually returns a new dictionary. 
# where you need to replace Team with Los_Angeles_Lakers from Cleveland_Cavaliers,
# delete the key-value pair Age and finally return dict_1.

# Example
# The evaluation machine will execute your code by executing python main.py {input_path},
# and the test data will be placed in the file corresponding to input_path. 
# You can see how the code works in main.py.

# Example 1
# When the input dictionary is:
#   {'Name':'Lebron','Age': 35,'Team':'Cleveland_Cavaliers'}
# The output result is:
#   Name: Lebron
#   Team: Los_Angeles_Lakers

# Example 2
# When the input dictionary is:
#   {'Name':'JR_Smith','Age': 36,'Team':'Cleveland_Cavaliers'}
# The output result is:
#   Name: JR_Smith
#   Team: Los_Angeles_Lakers

def modify_and_delete_elements_in_the_dictionary(dict_1: dict) -> dict:
    '''
    :param dict_1: Input dictionary
    :return: Newly generated dictionary after modification and deletion
    '''
    # -- write your code here --
    dict_1["Team"] = "Los_Angeles_Lakers"
    del dict_1["Age"]
    return dict_1