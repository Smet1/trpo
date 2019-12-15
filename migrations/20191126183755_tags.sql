-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tag
(
    id   SERIAL PRIMARY KEY,
    name citext NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_tags
(
    post_id int REFERENCES posts (id) NOT NULL,
    tag_id  int REFERENCES tag (id)   NOT NULL,
    UNIQUE (post_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table post_tags;
-- +goose StatementEnd

-- +goose StatementBegin
DROP table tag;
-- +goose StatementEnd