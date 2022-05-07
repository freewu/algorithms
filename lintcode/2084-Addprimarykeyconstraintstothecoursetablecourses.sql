-- 2084 · Add primary key constraints to the course table courses
-- # Description
-- Now we need to add a primary key constraint to the courses table courses and 
-- set the id column as the primary key because the database staff forgot to add the primary key constraint 
-- when creating the data table, write the corresponding SQL statement.

-- Table Definition: courses

-- column name	type	comment
-- id	int	not a primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	class start time
-- teacher_id	int	instructor id

-- The primary key column cannot be a NULL value
-- You can try to name and define PRIMARY KEY constraints for multiple columns
-- Example
-- Sample:
-- After running your SQL statement, we will execute the following statement to check if the id is added as a primary key constraint:

-- desc `courses`;
-- The table should return:

-- Field	Type	Null	Key	Default	Extra
-- id	int(10) unsigned	NO	PRI	NULL	
-- name	varchar(64)	NO		NULL	
-- student_count	int(10) unsigned	NO		NULL	
-- created_at	date	NO		NULL	
-- teacher_id	int(10) unsigned	NO		NULL	
-- Returns the table showing that the id has been set as the primary key, then the primary key constraint has been added to the table successfully

-- 先清除原有主键属性
--ALTER Table `courses` DROP PRIMARY KEY;

-- use MODIFY
ALTER TABLE `courses` MODIFY id int(10) unsigned PRIMARY KEY;

-- use ADD CONSTRAINT
ALTER TABLE courses ADD CONSTRAINT PRIMARY KEY(id)