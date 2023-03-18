-- +goose Up
-- +goose StatementBegin

CREATE TABLE warehouse(
                          warehouse_id BIGINT NOT NULL,
                          sku INT NOT NULL ,
                          available_to_order INT NOT NULL DEFAULT 0
                              CONSTRAINT not_negative_sku_count_check
                              CHECK (available_to_order >= 0)
);

CREATE UNIQUE INDEX warehouse_uniq_idx ON warehouse (warehouse_id, sku);

CREATE FUNCTION status_new() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 1 :: SMALLINT$$;

CREATE FUNCTION status_failed() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 2 :: SMALLINT$$;

CREATE FUNCTION status_awaiting_payment() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 3 :: SMALLINT$$;

CREATE FUNCTION status_payed() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 4 :: SMALLINT$$;

CREATE FUNCTION status_cancelled() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 5 :: SMALLINT$$;


CREATE TABLE order_info (
                            id BIGSERIAL PRIMARY KEY,
                            user_id BIGINT NOT NULL,
                            created_at TIMESTAMP DEFAULT now() NOT NULL,
                            status SMALLINT NOT NULL CONSTRAINT check_status
                                CHECK (status = ANY(ARRAY[
                                    status_new(),
                                    status_failed(),
                                    status_awaiting_payment(),
                                    status_payed(),
                                    status_cancelled()
                                    ]))
);

CREATE TABLE order_item (
                            sku INT NOT NULL,
                            fk_order_info_id BIGINT NOT NULL REFERENCES order_info,
                            count INT NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX order_item_uniq_idx ON order_item(sku, fk_order_info_id);

CREATE TABLE reserve (
                         sku INT NOT NULL,
                         warehouse_id BIGINT NOT NULL,
                         count INT NOT NULL DEFAULT 0,
                         fk_order_info_id BIGINT NOT NULL REFERENCES order_info

);

CREATE UNIQUE INDEX reserve_uniq_idx ON reserve(sku, warehouse_id, fk_order_info_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX warehouse_uniq_idx;
DROP TABLE warehouse;

DROP FUNCTION status_new;
DROP FUNCTION status_failed;
DROP FUNCTION status_awaiting_payment;
DROP FUNCTION status_payed;
DROP FUNCTION status_cancelled;
DROP TABLE order_info;

DROP INDEX order_item_uniq_idx;
DROP TABLE order_item;

DROP INDEX reserve_uniq_idx;
DROP TABLE reserve;

-- +goose StatementEnd
