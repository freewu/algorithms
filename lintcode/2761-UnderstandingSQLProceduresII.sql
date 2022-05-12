-- 2761 · Understanding SQL Procedures (II)
-- # Description
-- We have given you a stored procedure getTeachers, please write SQL statement to see the details of this procedure.
-- Example
-- Input：

-- teachers：

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	18	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- Return：

-- Procedure	sql_mode	Create Procedure	character_set_client	collation_connection	Database Collation
-- 'getTeachers'	'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'	'CREATE DEFINER=lintcode@% PROCEDURE getTeachers()\nBEGIN\r\n\tSELECT * FROM teachers;\r\nEND'	'utf8mb4'	'utf8mb4_general_ci'	'utf8mb4_bin'

-- Write your SQL here --
SHOW PROCEDURE STATUS;