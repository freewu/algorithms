# 2668 · Create and switch branches
# Description
# In this topic, let's do a simple practice of git branch instruction. 
# Please use the git command to complete the operation of renaming the branch. Go through the following 4 steps:

# Step 1 Query the status of the current branch
git status
git branch -a

# Step 2 Switch to the branch1 branch
git checkout branch1

# Step 3 Query the status of the current branch
git status
git branch -a

# Step 4 Use the -b parameter to create a new and switch to the branch2 branch
# git checkout -b 本地分支名 origin/远程分支名 使用该方式会在本地新建分支，并自动切换到该本地分支。
git checkout -b branch2

# Step 5 Query the status of the current branch
git branch -a