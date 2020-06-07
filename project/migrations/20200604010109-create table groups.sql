
-- +migrate Up
CREATE TABLE groups (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(20) NOT NULL UNIQUE,
    dt timestamptz
);

-- +migrate Down
DROP TABLE groups;
