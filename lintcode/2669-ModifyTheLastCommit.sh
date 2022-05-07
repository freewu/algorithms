# 2669 · Modify the last commit
# Description
# We've already committed the commit message once, but the notes are wrong, 
# so please change them to the correct notes by completing the following 3 steps:

# Step 1 use git log to view the git commit log
cd my-repo
git log

# Step 2 Using the git commit --amend command, change the last commit comment to 'This is a correct commit'
git commit --amend -m'This is a correct commit.'

# Step 3 Use git log to view the git commit log
git log