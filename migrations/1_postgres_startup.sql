-- +goose Up
-- +goose StatementBegin
CREATE TABLE blogs (
    id          SERIAL  PRIMARY KEY,
    title       TEXT NOT NULL,
    content     TEXT NOT NULL,
    tags        TEXT[] NOT NULL,
    created_on  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_on  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_tags ON blogs(Tags);
-- +goose StatementEnd

-- +goose StatementBegin
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE blogs;
-- +goose StatementEnd