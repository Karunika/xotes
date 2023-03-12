-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE DATABASE xotesDB;
CREATE TABLE user (
    username varchar(32),

);

CREATE TABLE notes (
    uuid char(),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
