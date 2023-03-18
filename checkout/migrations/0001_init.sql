-- +goose Up
-- +goose StatementBegin

CREATE TABLE cart (
    user_id BIGINT NOT NULL,
    sku INT NOT NULL ,
    count INT NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX cart_uniq_idx ON cart (user_id, sku);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX cart_uniq_idx;

DROP TABLE cart;

-- +goose StatementEnd
