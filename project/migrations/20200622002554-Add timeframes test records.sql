
-- +migrate Up
INSERT INTO timeframes (task_uuid, dtfrom, dtto)
SELECT  tasks.uuid, '2020-05-19T20:34:09Z', '2020-05-19T20:34:09Z'
FROM    tasks
WHERE   tasks.title = 'mouse';

INSERT INTO timeframes (task_uuid, dtfrom, dtto)
SELECT  tasks.uuid, '2020-05-19T21:34:09Z', '2020-05-19T21:34:09Z'
FROM    tasks
WHERE   tasks.title = 'mouse';
-- +migrate Down
TRUNCATE timeframes;