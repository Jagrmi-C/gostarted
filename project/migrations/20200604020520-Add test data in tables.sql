
-- +migrate Up
INSERT INTO groups (title, dt)
VALUES
	('toys', '2020-06-04 19:10:25-07'),
    ('pen', '2020-06-04 22:12:34-07'),
    ('bus', '2020-06-04 19:10:25-07'),
    ('car', '2020-06-04 22:12:34-07');

INSERT INTO tasks (title, group_uuid)
SELECT  'mouse', groups.uuid
FROM    groups
WHERE   groups.title = 'toys';

INSERT INTO tasks (title, group_uuid)
SELECT  'cat', groups.uuid
FROM    groups
WHERE   groups.title = 'toys';

INSERT INTO tasks (title, group_uuid)
SELECT  'dog', groups.uuid
FROM    groups
WHERE   groups.title = 'toys';

-- +migrate Down
TRUNCATE tasks;
DELETE FROM groups WHERE title = 'toys' OR title = 'pen' OR title = 'car' OR title = 'bus';
