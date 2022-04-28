-- 1941 · Find Right Triangle
-- # Description
-- Li Hua’s job is to determine whether three line segments can form a right triangle
-- Assuming that the table line_segments saves all groups consisting of three line segments with lengths a, b, c, please help Li Hua write a SQL statement to determine whether each group of line segments can form a right triangle

-- Table definition: line_segments

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- a	int	The length of a line segment
-- b	int	The length of b line segment
-- c	int	The length of c line segment
-- Example
-- Example 1:

-- Table content: line_segments

-- id	a	b	c
-- 1	3	4	5
-- 2	10	20	15
-- 3	1	2	10
-- After running your SQL statement, the table should return:

-- id	a	b	c	right_triangle
-- 1	3	4	5	Yes
-- 2	10	20	15	No
-- 3	1	2	10	No
-- Example 2:

-- Table content: line_segments

-- id	a	b	c
-- 1	6	6	6
-- 2	5	12	13
-- After running your SQL statement, the table should return:

-- id	a	b	c	right_triangle
-- 1	6	6	6	No
-- 2	5	12	13	Yes

-- 3 * 3 + 4 * 4 = 5 * 5
SELECT
	id,
	a,
	b,
	c,
	IF((a * a + b * b = c * c),'Yes','No') AS right_triangle
FROM
	line_segments