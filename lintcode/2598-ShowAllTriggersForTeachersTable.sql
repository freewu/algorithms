-- 2598 · Show all triggers for teachers table
-- # Description
-- We want to see all the triggers in the teachers table, please write SQL statements to achieve this

-- table definition : teachers (teachers table)
-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality

-- Example
-- After running your SQL statement, the table should return :

-- Trigger	Event	Table	Statement	Timing	sql_mode	Definer	character_set_client	collation_connection	Database Collation
-- 'before_teachers_insert'	'INSERT'	'teachers'	"SET new.country = 'RS'"	'BEFORE'	'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'	'lintcode@%'	'utf8mb4'	'utf8mb4_general_ci'	'utf8mb4_bin'
-- 'before_teachers_update'	'UPDATE'	'teachers'	"SET new.country = 'CN'"	'BEFORE'	'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'	'lintcode@%'	'utf8mb4'	'utf8mb4_general_ci'	'utf8mb4_bin'

-- Write your SQL here --
-- SELECT 
--     * 
-- FROM
--     information_schema.triggers
-- WHERE
--     `Table` = 'teachers'

-- mysql 查看触发器：
-- 1 可以使用“show triggers;”查看触发器
--      show triggers;
-- 2 由于MySQL创建的触发器保存在“information_schema库中的triggers表中，所以还可以通过查询此表查看触发器：
-- 通过information_schema.triggers表查看触发器：
--      select * from information_schema.triggers;

-- SELECT 
--     * 
-- FROM
--     information_schema.triggers
-- WHERE
--     `table` = 'teachers'

-- information_schema.triggers
-- select * from information_schema.triggers;

show triggers;