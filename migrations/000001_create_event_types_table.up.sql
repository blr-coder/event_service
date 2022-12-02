CREATE TABLE IF NOT EXISTS event_types
(
    title       VARCHAR(55) PRIMARY KEY,
    created_at  TIMESTAMP   NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at  TIMESTAMP   NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    deleted_at  TIMESTAMP
);