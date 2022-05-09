-- 2091 · Adding Foreign Key Constraints to Course Tables
-- # Description
-- Write an SQL statement to add a foreign key constraint to teacher_id in the courses table courses 
-- so that it can be associated with id in the teachers table teachers.

-- Table Definition 1: courses (course table)

-- column name	type	comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Table Definition 2 : teachers (teachers table)

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- Add FOREIGN KEY to the teacher_id column of the created table courses.
-- The FOREIGN KEY of teacher_id is the primary key id in the teachers table that does not have a NULL value.
-- Example
-- Sample:
-- After running your SQL statement, we will execute the following statement to check if a foreign key constraint has been added:

-- desc `courses`;
-- The table should return:

-- Field	Type	Null	Key	Default	Extra
-- id	int(10) unsigned	NO	PRI	NULL	auto_increment
-- name	varchar(64)	NO		NULL	
-- student_count	int(10) unsigned	NO		NULL	
-- created_at	date	YES		NULL	
-- teacher_id	int(10) unsigned	NO	MUL	NULL	
-- If teacher_id is shown as a foreign key in the returned data, it is added successfully.

-- ALTER table_name
-- ADD CONSTRAINT constraint_name
-- FOREIGN KEY foreign_key_name(columns)
-- REFERENCES parent_table(columns)
-- ON DELETE action
-- ON UPDATE action;

-- The FOREIGN KEY clause specifies the columns in the child table that refers to primary key columns in the parent table. You can put a foreign key name after FOREIGN KEY clause or leave it to let MySQL create a name for you. Notice that MySQL automatically creates an index with the foreign_key_name name.
-- The REFERENCES clause specifies the parent table and its columns to which the columns in the child table refer. The number of columns in the child table and parent table specified in the FOREIGN KEYand REFERENCES must be the same.
-- The ON DELETE clause allows you to define what happens to the records in the child table when the records in the parent table are deleted. If you omit the ON DELETE clause and delete a record in the parent table that has records in the child table refer to, MySQL will reject the deletion. In addition, MySQL also provides you with actions so that you can have other options such as ON DELETE CASCADE that ask MySQL to delete records in the child table that refers to a record in the parent table when the record in the parent table is deleted. If you don’t want the related records in the child table to be deleted, you use the ON DELETE SET NULL action instead. MySQL will set the foreign key column values in the child table to NULL when the record in the parent table is deleted, with a condition that the foreign key column in the child table must accept NULL values. Notice that if you use ON DELETE NO ACTION or ON DELETE RESTRICT action, MySQL will reject the deletion.
-- The ON UPDATE clause enables you to specify what happens to the rows in the child table when rows in the parent table are updated. You can omit the ON UPDATE clause to let MySQL reject any updates to the rows in the child table when the rows in the parent table are updated. The ON UPDATE CASCADE action allows you to perform a cross-table update, and the ON UPDATE SET NULLaction resets the values in the rows in the child table to NULL values when the rows in the parent table are updated. The ON UPDATE NO ACTION or UPDATE RESTRICT actions reject any updates.

ALTER TABLE 
	`courses` 
ADD CONSTRAINT FOREIGN KEY ( `teacher_id` )
REFERENCES 
	`teachers` ( `id` );