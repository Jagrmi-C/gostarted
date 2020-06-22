
-- +migrate Up
CREATE TABLE tasks (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(20) NOT NULL UNIQUE,
    group_uuid UUID REFERENCES groups(uuid) on update cascade on delete cascade
);
-- +migrate Down
DROP TABLE tasks;