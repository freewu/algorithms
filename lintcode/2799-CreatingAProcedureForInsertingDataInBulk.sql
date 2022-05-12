-- 2799 · Creating a procedure for inserting data in bulk
-- Description
-- We need to insert 30000 test data into the teachers table, 
-- containing teacher name name = 'teacher' + test id, (test id increases from 1 to 30000), 
-- teacher email email = name + '@chapter.com', teacher age age = 26 + (id%20)
-- Please use the stored procedure of SQl to do this, the procedure named addTeachers.

-- Example
-- Input：

-- teachers：

-- id	name	email	age	country
-- Return：

-- id	name	email	age	country
-- 1	'teacher1'	'teacher1@chapter.com'	27	''
-- 2	'teacher2'	'teacher2@chapter.com'	28	''
-- 3	'teacher3'	'teacher3@chapter.com'	29	''
-- 4	'teacher4'	'teacher4@chapter.com'	30	''
-- 5	'teacher5'	'teacher5@chapter.com'	31	''
-- 6	'teacher6'	'teacher6@chapter.com'	32	''
-- 7	'teacher7'	'teacher7@chapter.com'	33	''
-- '...'	'...'	'...'

-- Write your SQL here --
CREATE PROCEDURE addTeachers()
BEGIN
	DECLARE id INT DEFAULT 0;
	SET id = 1;
	WHILE id <= 30000 DO
		INSERT INTO 
			`teachers`(
				`id`,
				`name`,
				`email`,
				`age`,
				`country`
			)
		VALUES (
			id,
			CONCAT('teacher',id),
			CONCAT('teacher',id,'@chapter.com'),
			(26 + (id % 20)),
			''
		);
		SET id = id + 1;
	END WHILE;
END