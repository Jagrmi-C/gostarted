
-- +migrate Up
INSERT INTO groups (title, dt)
VALUES
	('toys', '2020-06-04 19:10:25-07'),
    ('pen', '2020-06-04 22:12:34-07'),
    ('bus', '2020-06-04 19:10:25-07'),
    ('car', '2020-06-04 22:12:34-07');

INSERT INTO tasks (title, group_uuid, dt)
SELECT  'mouse', groups.uuid, '2020-06-05 12:12:34-07'
FROM    groups
WHERE   groups.title = 'toys';

-- +migrate Down
DELETE FROM groups WHERE title = 'toys' OR title = 'pen' OR title = 'car' OR title = 'bus';