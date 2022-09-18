-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table customer (id serial, name varchar,email VARCHAR,password VARCHAR, PRIMARY KEY(id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
