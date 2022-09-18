-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE product (id serial PRIMARY KEY, brand_id int, name VARCHAR,price numeric, quantity INT,
constraint fk_brand_id FOREIGN KEY(brand_id) REFERENCES brand (id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
