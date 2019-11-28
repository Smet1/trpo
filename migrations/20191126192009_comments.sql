-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments
(
    id        SERIAL PRIMARY KEY,
    parent_id int REFERENCES comments (id),
    user_id   int  NOT NULL REFERENCES users (id),
    post_id   int  NOT NULL REFERENCES posts (id),
    payload   text NOT NULL,
    show      bool        DEFAULT True,
    created   timestamptz DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table comments;
-- +goose StatementEnd
