-- 2705 · Show view in the current database
-- # Description
-- Now we need to see all the view in the current database, write the SQL statement to do so

-- Example
-- The result should be returned after executing your SQL statement:

-- Tables_in_judge	Table_type
-- 'v_courses'	'VIEW'
-- 'v_courses_teachers'	'VIEW'
-- 'v_teachers'	'VIEW'

-- Write your SQL here --

-- show table 
SHOW FULL TABLES WHERE table_type = 'VIEW';

-- select information_schema.VIEWS
SELECT 
	TABLE_NAME AS Tables_in_judge,
	'VIEW' AS Table_type
FROM 
	information_schema.VIEWS