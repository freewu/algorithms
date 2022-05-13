# 2619 · Track and submit document changes
# Description
# Although using staging areas allows you to carefully prepare the details of your commit, 
# sometimes it can be a bit cumbersome, so this time let's practice skipping the staging area 
# and staging all the files you've already tracked and committing the update together.
# Under the repository directory my-repo, there is a tracked file new_file.cpp, which we have made a change to.

# Step 1 use git status to see the current status of the repository
cd my-repo
git status

# Step 2 use a git commit command to skip using the staging area and commit the changes you just made to the file, with any comments you want to make
git commit -am'any message'
# -a 参数设置修改文件后不需要执行 git add 命令，直接来提交
#   git commit -a
# 相当于:
#   git add
#   git commit 

# Step 3  use git status to see the current status of the repository
git status

# Step 4 use git log to see the git commit log
git log
