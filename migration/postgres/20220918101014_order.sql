-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table orders (id serial,customer_id INT, order_date TIMESTAMP, create_at TIMESTAMP, update_at TIMESTAMP, amount numeric,PRIMARY KEY(id),constraint fk_customer_id FOREIGN KEY(customer_id) REFERENCES customer(id));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
