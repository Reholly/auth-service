-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS account (
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS claim (
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    title VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS account_claim (
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL,
    account_id VARCHAR(255) NOT NULL,
    claim_id VARCHAR(255) NOT NULL
);


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS claim;
DROP TABLE IF EXISTS user_claim
-- +goose StatementEnd
