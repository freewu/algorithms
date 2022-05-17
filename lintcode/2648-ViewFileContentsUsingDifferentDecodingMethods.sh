# 2648 · View file contents using different decoding methods
# Description
# Here, you will complete a 3-step exercise to view the contents of a file.

# Step 1 use the cat command to view the contents of the file file
cat file

# Step 2 use the od command to output the contents of the file file using single-byte octal interpretation
od -c file

# Step 3 use the od command to output the contents of the file file using ASCII code
od -t d1 file
