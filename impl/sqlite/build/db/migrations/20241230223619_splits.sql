-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS kmmSplits
(
    transactionId   varchar(32)       not null,
    txType          char(1),
    splitId         smallint unsigned not null,
    payeeId         varchar(32),
    reconcileDate   timestamp,
    action          varchar(16),
    reconcileFlag   char(1),
    value           text              not null,
    valueFormatted  text,
    shares          text              not null,
    sharesFormatted mediumtext,
    price           text,
    priceFormatted  mediumtext,
    memo            mediumtext,
    accountId       varchar(32)       not null,
    costCenterId    varchar(32),
    checkNumber     varchar(32),
    postDate        timestamp,
    bankId          mediumtext,
    primary key (transactionId, splitId)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS kmmSplits;

-- +goose StatementEnd
