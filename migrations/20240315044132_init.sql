-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS account (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    is_email_confirmed BOOLEAN NOT NULL,
    is_banned BOOLEAN NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS account_claim (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    username VARCHAR(255) NOT NULL,
    claim_title VARCHAR(255) NOT NULL,
    claim_value VARCHAR(255) NOT NULL
);


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS account_claim
-- +goose StatementEnd
