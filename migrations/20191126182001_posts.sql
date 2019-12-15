-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts
(
    id          SERIAL PRIMARY KEY,
    header      CITEXT NOT NULL UNIQUE,
    short_topic citext NOT NULL UNIQUE,
    main_topic  CITEXT NOT NULL UNIQUE,
    user_id     int    NOT NULL REFERENCES users (id),
    show        bool        DEFAULT False,
    created     timestamptz DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table posts;
-- +goose StatementEnd
