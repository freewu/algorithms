# 2126 · Import one of the files in a module folder
# Description
# In this question, we'll practice import modules (files) from a folder in python.
# In addition to importing modules from the current environment directory, 
# Python can also import modules from the current environment's directory. This is shown below:

# from folder_name import module_name
# folder_name refers to the name of the folder in the current environment, and module_name refers to the Python module (file) in that folder.
# In this case, we have a folder named branch under PROJECT in our current project, and within that folder, we can see a file named solution.py.
# Please import the solution module at the location specified in main.py, and we will execute the do method in solution.

# Import keywords should be lowercase
# A space is required between the import keyword and the module name
# Example
# If the import is correct, we will return
#   We are running in branch/branch.solution.py

# write your code here
from branch import solution

# keep the code below
solution.do()