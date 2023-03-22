-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SCHEMA IF NOT EXISTS xotes_schema;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    uuid UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
    email VARCHAR(320) NOT NULL,
    username VARCHAR(32) NOT NULL,
    pwd_hash CHAR(72) NOT NULL,
    country VARCHAR(64) NOT NULL,
    bio VARCHAR(255),
    pfp_blob BYTEA,
    pfp_mime_type VARCHAR(255),
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (username),
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS groups (
    uuid UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
    user_uuid UUID NOT NULL,
    prev_item UUID,
    belongs_to UUID NOT NULL,
    group_name VARCHAR(32) NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (uuid),
    FOREIGN KEY (user_uuid) REFERENCES users,
    FOREIGN KEY (belongs_to) REFERENCES groups,
    FOREIGN KEY (prev_item) REFERENCES groups
);

-- CREATE TABLE notes (
--     uuid UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
--     title VARCHAR(32),
--     ndescription VARCHAR(256) NOT NULL DEFAULT(''),
--     keywords TEXT [],
--     prev_item UUID,
--     belongs_to UUID NOT NULL,
--     folder UUID NOT NULL,
--     blob BLOB(10G) NOT NULL DEFAULT (''),
--     date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     last_modified TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (uuid),
--     FOREIGN KEY (belongs_to) REFERENCES folders
-- );

-- CREATE TABLE followers (
--     follower_uuid UUID NOT NULL,
--     following_uuid UUID NOT NULL,
--     PRIMARY KEY (follower_uuid, following_uuid),
--     FOREIGN KEY (follower_uuid) REFERENCES users,
--     FOREIGN KEY (following_uuid) REFERENCES users
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP SCHEMA IF EXISTS xotes_schema CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS groups CASCADE;
-- +goose StatementEnd
