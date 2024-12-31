-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS kmmAccounts
(
    id                varchar(32) not null primary key,
    institutionId     varchar(32),
    parentId          varchar(32),
    lastReconciled    timestamp,
    lastModified      timestamp,
    openingDate       date,
    accountNumber     mediumtext,
    accountType       varchar(16) not null,
    accountTypeString mediumtext,
    isStockAccount    char(1),
    accountName       mediumtext,
    description       mediumtext,
    currencyId        varchar(32),
    balance           mediumtext,
    balanceFormatted  mediumtext,
    transactionCount  bigint unsigned
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS kmmAccounts;

-- +goose StatementEnd
