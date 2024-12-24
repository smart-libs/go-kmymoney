-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS kmmSecurities
(
    id                      varchar(32) not null primary key,
    name                    text        not null,
    symbol                  mediumtext,
    type                    smallint unsigned           not null,
    typeString              mediumtext,
    smallestAccountFraction varchar(24),
    pricePrecision          smallint unsigned default 4 not null,
    tradingMarket           mediumtext,
    tradingCurrency         char(3),
    roundingMethod          smallint unsigned default 7 not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS kmmSecurities;

-- +goose StatementEnd
