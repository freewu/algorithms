# 2606 · View user ID and group ID
# Description
# Use the id command to view and print your ID and the ID of the group to which you belong, 
# including the actual ID and the valid ID, by completing the 3 steps in order.

# Step 1 View and print your ID and the ID of the group to which you belong, including the actual ID and the valid ID
id

# Step 2 View and print your valid ID
id -un

# Step 3 View and print your real ID
id -ur

# -- 打印用户名、UID 和该用户所属的所有组
#   id -a

# -- 输出所有不同的组ID (有效的，真实的和补充的)
#   id -G

# -- 只输出有效的组ID
#   id -g

# --输出特定用户信息
#   输出特定的用户信息相关的 UID 和 GID。只需要在 id 命令后跟上用户名
#   id root