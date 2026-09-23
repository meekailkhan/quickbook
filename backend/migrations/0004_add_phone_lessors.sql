-- +goose Up
ALTER TABLE lessors ADD COLUMN phone VARCHAR(10) DEFAULT '0000000000';
-- +goose Down
ALTER TABLE lessors DROP COLUMN phone;