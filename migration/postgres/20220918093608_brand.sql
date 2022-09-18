-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table brand (id serial, name varchar, PRIMARY key(id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
