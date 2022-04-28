-- 1943 · Looking for Students of a Certain Height
-- # Description
-- The student_heights table records the height of the students, including many repeated heights
-- Please write a SQL statement to find the tallest height in a height that has only appeared once

-- Table definition: student_heights

-- columns_name	type	explaination
-- height	int	student's height
-- Find the tallest height in a height that has only appeared once, if not found, please return null

-- Example
-- Example 1:

-- Table content: student_heights

-- height
-- 178
-- 178
-- 173
-- 173
-- 171
-- 174
-- 175
-- 176
-- After running your SQL statement, the table should return:

-- height
-- 176
-- Example 2:

-- Table content: student_heights

-- height
-- 178
-- 178
-- 173
-- 173
-- After running your SQL statement, the table should return:

-- height
-- null

SELECT
	MAX(a.height) AS height
FROM
	(
		SELECT 
			height,
			COUNT(*) AS num
		FROM
			student_heights
		GROUP BY height
		HAVING num = 1 -- 只取出现在了一次的数据
	) AS a