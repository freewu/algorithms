-- 2801 · Delete Stored Procedure (I)
-- # Description
-- Now you need to delete a stored procedure getTeachers() from the current database, please write SQL to implement it.

-- Example
-- Input：

-- teachers：

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	18	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- Return：

-- Name	Type	Security_type	character_set_client	collation_connection	Database Collation
-- 'getTeachers_1'	'PROCEDURE'	'DEFINER'	'utf8mb4'	'utf8mb4_general_ci'	'utf8mb4_bin'
-- 'getTeachers_2'	'PROCEDURE'	'DEFINER'	'utf8mb4'	'utf8mb4_general_ci'	'utf8mb4_bin'
-- 'getTeachers_3'	'PROCEDURE'	'DEFINER'	'utf8mb4'	'utf8mb4_general_ci'	'utf8mb4_bin'

-- Write your SQL here --
DROP PROCEDURE getTeachers;