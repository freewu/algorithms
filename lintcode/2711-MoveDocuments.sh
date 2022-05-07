# 2711 · Move documents
# Description
# In this topic, let's do a simple exercise with the git command and complete a simple mv.
# Use the mv command to complete the following 3 steps:

# Step 1 go to the my-repo repository directory and use git status to see the current status of the repository
cd /nobody/my-repo
git status

# Step 2 rename the staging area file first.cpp to second.cpp using git mv
# git mv 命令用于移动或重命名一个文件、目录或软连接。
# git mv [file] [newfile]
# 如果新但文件名已经存在，但还是要重命名它，可以使用 -f 参数：
# git mv -f [file] [newfile]
git mv first.cpp second.cpp

# Step 3 use git status to see the current status of the repository
git status
