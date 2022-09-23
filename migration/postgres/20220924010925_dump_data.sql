-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO customer
("name", email, "password")
VALUES('customer-01', 'customer-01@mailinator.com', md5('password'));
INSERT INTO customer
( "name", email, "password")
VALUES('customer-02', 'customer-01@mailinator.com', md5('password'));





-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
