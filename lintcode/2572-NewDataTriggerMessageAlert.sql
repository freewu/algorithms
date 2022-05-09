-- 2572 · New data trigger message alert
-- # Description
-- A INSERT trigger is created for a table named members 
-- to implement a reminder message 'Hi (member name), please update your date of birth.' 
-- into the reminders table if the member's date of birth is NULL.

-- Table Definition 1:members (members table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	member name
-- email	varchar	member email
-- birthDate	date	date of birth
-- Table Definition 2: reminders(reminders table)

-- column name	type	comment
-- id	int unsigned	primary key
-- message	varchar	message content
-- memberId	int	member id

-- Example
-- After executing your SQL statement, the following code will be executed in the background:

-- INSERT INTO `members` (`name`, `email`, `birthDate`) 
--     VALUES ('eastern heretic', 'eastern.heretic@gmail.com', '2001-02-05');
-- INSERT INTO `members` (`name`) VALUES ('Northern Beggar');
-- INSERT INTO `members` (`name`, `email`, `birthDate`) 
--     VALUES ('Western Venom', 'western.venom@163.com', '2000-02-28');
-- INSERT INTO `members` (`name`, `email`) 
--     VALUES ('LingHu', 'LingHu@1234.com');
-- INSERT INTO `members` (`name`, `birthDate`) 
--     VALUES ('Southern Emperor', '2014-12-01');
-- SELECT * FROM `reminders`;
-- Return results:

-- id	memberId	message
-- 1	2	'Hi Northern Beggar, please update your date of birth.'
-- 2	4	'Hi LingHu, please update your date of birth.'

-- Write your SQL here --
DROP TRIGGER IF EXISTS `after_members_insert`;
CREATE TRIGGER `after_members_insert`
AFTER INSERT ON `members`
FOR EACH ROW
BEGIN
	-- 如果发现生日为 null
	IF new.birthDate IS NULL THEN
		-- 插入一条 reminders 数据
		INSERT INTO
			`reminders`(
				`memberId`,
				`message`
			) VALUES (
				-- 直接取自增ID
				new.id,
				CONCAT('Hi ', new.name , ', please update your date of birth.')
			);
	END IF;
END