-- 2085 · Remove the primary key constraint from the course table `courses`
-- # Description
-- Write an SQL statement to delete the primary key constraint in the course table courses.

-- Table Definition : courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	total number of students
-- created_at	date	course creation time
-- teacher_id	int	instructor id

-- If there are self-growing attributes, you need to delete the self-growing attributes first when you delete the primary key constraint.
-- Example
-- Sample:
-- After running your SQL statement, we will execute the following statement to check if the primary key constraint is removed from the course table course :

-- desc `courses`;
-- The table should return:

-- Field	Type	Null	Key	Default	Extra
-- id	int(10) unsigned	NO		NULL	
-- name	varchar(64)	NO		NULL	
-- student_count	int(10) unsigned	NO		NULL	
-- created_at	date	NO		NULL	
-- teacher_id	int(10) unsigned	NO		NULL	
-- The return data shows that the original primary key id is no longer the primary key, indicating a successful deletion.

-- 直接删除 
ALTER TABLE `courses` DROP PRIMARY KEY;

-- 如果主键是自增 运行前面会出现
-- ERROR 1075 (42000): Incorrect table definition; there can be only one auto column and it must be defined as a key
-- 先修改掉 AUTO_INCREMENT 属性
ALTER TABLE `courses` MODIFY id int;

-- 再清除主键
ALTER TABLE `courses` DROP PRIMARY KEY;