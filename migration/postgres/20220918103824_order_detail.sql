-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE order_detail (id SERIAL PRIMARY KEY, order_id BIGINT, product_id BIGINT, amount NUMERIC,quantity INT,
constraint fk_order_id FOREIGN KEY(order_id) REFERENCES orders(id),
constraint fk_product_id FOREIGN KEY(product_id) REFERENCES product(id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
