-- 2808 · MySQL Stored Procedure OUT Parameters I
-- Description
-- Write an SQL statement to create a stored procedure that will return the number of teachers by nationality. 
-- The procedure has two parameters.
--      teacherCountry: is the IN parameter that specifies the nationality of the teachers to be returned
--      total: is the OUT parameter that stores the number of teachers of that nationality
-- Call this procedure to find the number of teachers of Chinese nationality

-- Table definition: teachers (teachers table)

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- Example
-- Input：

-- teachers:
-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN

-- Return：
-- @total
-- 2

-- Write your SQL Query here --
CREATE PROCEDURE GetTeachersByCountry (
	IN teacherCountry VARCHAR(10),
    OUT total INT
)
BEGIN
    -- 
    SET total =  (
        SELECT
            COUNT(*)
        FROM
            `teachers`
        WHERE
            country = teacherCountry
    );
    -- 
	SELECT
		*
	FROM
		`teachers`
	WHERE
		country = teacherCountry;
END;

CALL GetTeachersByCountry('CN',@total);

SELECT @total;