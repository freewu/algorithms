-- 3204. Bitwise User Permissions Analysis
-- Table: user_permissions
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | permissions | int     |
-- +-------------+---------+
-- user_id is the primary key.
-- Each row of this table contains the user ID and their permissions encoded as an integer.
-- Consider that each bit in the permissions integer represents a different access level or feature that a user has.

-- Write a solution to calculate the following:
--     common_perms: The access level granted to all users. This is computed using a bitwise AND operation on the permissions column.
--     any_perms: The access level granted to any user. This is computed using a bitwise OR operation on the permissions column.

-- Return the result table in any order.
-- The result format is shown in the following example.

-- Example:
-- Input:
-- user_permissions table:
-- +---------+-------------+
-- | user_id | permissions |
-- +---------+-------------+
-- | 1       | 5           |
-- | 2       | 12          |
-- | 3       | 7           |
-- | 4       | 3           |
-- +---------+-------------+
-- Output:
-- +-------------+--------------+
-- | common_perms | any_perms   |
-- +--------------+-------------+
-- | 0            | 15          |
-- +--------------+-------------+
    
-- Explanation:
-- common_perms: Represents the bitwise AND result of all permissions:
-- For user 1 (5): 5 (binary 0101)
-- For user 2 (12): 12 (binary 1100)
-- For user 3 (7): 7 (binary 0111)
-- For user 4 (3): 3 (binary 0011)
-- Bitwise AND: 5 & 12 & 7 & 3 = 0 (binary 0000)
-- any_perms: Represents the bitwise OR result of all permissions:
-- Bitwise OR: 5 | 12 | 7 | 3 = 15 (binary 1111)

-- Create table if not exists user_permissions(user_id int, permissions int)
-- Truncate table user_permissions
-- insert into user_permissions (user_id, permissions) values ('1', '5')
-- insert into user_permissions (user_id, permissions) values ('2', '12')
-- insert into user_permissions (user_id, permissions) values ('3', '7')
-- insert into user_permissions (user_id, permissions) values ('4', '3')

-- 位运算符是在二进制数上进行计算的运算符。位运算会先将操作数变成二进制数，进行位运算。
-- 然后再将计算结果从二进制数变回十进制数。

-- 运算符号	作用
-- &	        按位与
-- |	        按位或
-- ^	        按位异或
-- !	        取反
-- <<	        左移
-- >>	        右移

-- # 按位与
-- mysql> select 3&5;
-- +-----+
-- | 3&5 |
-- +-----+
-- |   1 |
-- +-----+

-- # 按位或
-- mysql> select 3|5;
-- +-----+
-- | 3|5 |
-- +-----+
-- |   7 |
-- +-----+

-- # 按位异或
-- mysql> select 3^5;
-- +-----+
-- | 3^5 |
-- +-----+
-- |   6 |
-- +-----+

-- # 按位取反
-- mysql> select ~18446744073709551612;
-- +-----------------------+
-- | ~18446744073709551612 |
-- +-----------------------+
-- |                     3 |
-- +-----------------------+

-- # 按位右移
-- mysql> select 3>>1;
-- +------+
-- | 3>>1 |
-- +------+
-- |    1 |
-- +------+

-- # 按位左移
-- mysql> select 3<<1;
-- +------+
-- | 3<<1 |
-- +------+
-- |    6 |
-- +------+

SELECT
    MIN(@bitand := @bitand & `permissions`) AS common_perms, 
    MAX(@bitor := @bitor | `permissions`) AS any_perms
FROM 
    user_permissions, 
    (SELECT @bitand := ~0, @bitor := 0) AS init