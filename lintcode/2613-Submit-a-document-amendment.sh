# 2613 · Submit a document amendment
# Description
# In this topic, let's do a simple git command exercise to complete a simple commit.

# Use the git command to complete the following 6 steps: 
# Step 1 go to the my-repo repository directory and use git status to see the current status of the repository
cd my-repo
git status

# Step 2 use git add to keep track of existing new files new_file.cpp
touch new_file.cpp
git add new_file.cpp

# Step 3 use git status to see the current status of the repository
git status

# Step 4 commit the update using git commit, you can comment anything
git commit -m'comment'

# Step 5 use git status to see the current status of the repository
git status

# Step 6 use git log to view the git commit log
git log
