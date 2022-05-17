# 2698 · Find a compliant document (ii)
# Description
# Use the grep command to find the line of the file containing the given string.

# Step 1 the cat command to view the contents of the file main.c
cat main.c

# Step 2 use the grep command to find the line containing the string test in all files with the .c suffix
grep test *.c

# Step 3 on the basis of step 2, display the contents of the line after that line, in addition to the column that matches the model style
grep -A1 test *.c

# Step 4 on the basis of step 2, display the contents of the line before that line, in addition to the column that matches the model style
grep -B1 test *.c
