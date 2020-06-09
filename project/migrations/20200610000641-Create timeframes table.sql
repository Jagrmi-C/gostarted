
-- +migrate Up
CREATE TABLE timeframes (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_uuid UUID REFERENCES tasks(uuid),
    dtfrom timestamptz,
    dtto timestamptz
);
-- +migrate Down
DROP TABLE timeframes;