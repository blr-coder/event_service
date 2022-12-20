CREATE TABLE IF NOT EXISTS events
(
    id            BIGSERIAL PRIMARY KEY,
    type_title    VARCHAR(55) REFERENCES event_types (title) NOT NULL,
    campaign_id   BIGINT,
    insertion_id  BIGINT,
    user_id       BIGINT,
    cost_amount   BIGINT,
    cost_currency VARCHAR(3),
    created_at    TIMESTAMP                                  NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);