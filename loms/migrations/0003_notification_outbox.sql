-- +goose Up
-- +goose StatementBegin

CREATE FUNCTION send_status_pending() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 1 :: SMALLINT$$;

CREATE FUNCTION send_status_wait_confirmation() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 2 :: SMALLINT$$;

CREATE FUNCTION send_status_send() RETURNS SMALLINT
    IMMUTABLE
    LANGUAGE sql
AS
$$SELECT 3 :: SMALLINT$$;

CREATE TABLE notification_outbox (
    fk_order_info_id BIGINT NOT NULL REFERENCES order_info,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    send_status SMALLINT NOT NULL CONSTRAINT send_status
    CHECK (send_status = ANY(ARRAY[
        send_status_pending(),
        send_status_wait_confirmation(),
        send_status_send()
        ])),
    notification_status SMALLINT NOT NULL CONSTRAINT check_notification_status
    CHECK (notification_status = ANY(ARRAY[
         status_new(),
         status_failed(),
         status_awaiting_payment(),
         status_payed(),
         status_cancelled()
     ]))
);

CREATE UNIQUE INDEX fk_order_info_id_uniq_idx ON notification_outbox (fk_order_info_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX fk_order_info_id_uniq_idx;
DROP TABLE notification_outbox;

DROP FUNCTION send_status_send;
DROP FUNCTION send_status_wait_confirmation;
DROP FUNCTION send_status_pending;

-- +goose StatementEnd
