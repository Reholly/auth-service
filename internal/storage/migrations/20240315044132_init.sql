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
    id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL PRIMARY KEY ,
    account_id VARCHAR(255) NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    claim_id VARCHAR(255) NOT NULL  REFERENCES claim(id) ON DELETE CASCADE
);


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS claim;
DROP TABLE IF EXISTS account_claim
-- +goose StatementEnd
