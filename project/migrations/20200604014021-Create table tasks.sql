
-- +migrate Up
CREATE TABLE tasks (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(20) NOT NULL UNIQUE,
    group_uuid UUID REFERENCES groups(uuid),
    dt timestamptz
);
-- +migrate Down
DROP TABLE tasks;