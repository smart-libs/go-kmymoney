-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS kmmTransactions
(
    id         varchar(32) not null primary key,
    txType     char(1),
    postDate   timestamp,
    memo       mediumtext,
    entryDate  timestamp,
    currencyId char(3),
    bankId     mediumtext
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS kmmTransactions;

-- +goose StatementEnd
