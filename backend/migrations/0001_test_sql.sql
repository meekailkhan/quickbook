-- +goose Up

CREATE TABLE test (
    id BIGSERIAL PRIMARY KEY,
    test_name TEXT NOT NULL
);

-- +goose Down

DROP TABLE test;