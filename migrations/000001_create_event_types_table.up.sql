CREATE TABLE IF NOT EXISTS event_types
(
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255)                NOT NULL,
    created_at  TIMESTAMP                   NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    deleted_at  TIMESTAMP
);