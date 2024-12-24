-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS kmmCurrencies
(
    ISOcode                 char(3) not null primary key,
    name                    text    not null,
    type                    smallint unsigned,
    typeString              mediumtext,
    symbol1                 smallint unsigned,
    symbol2                 smallint unsigned,
    symbol3                 smallint unsigned,
    symbolString            varchar(255),
    smallestCashFraction    varchar(24),
    smallestAccountFraction varchar(24),
    pricePrecision          smallint unsigned default 4 not null
);

INSERT INTO kmmCurrencies (ISOcode, name, type, typeString, symbol1, symbol2, symbol3, symbolString, smallestCashFraction, smallestAccountFraction, pricePrecision) VALUES ('BRL', 'Brazilian Real', 3, 'Currency', 82, 36, 32, 'R$   ', '100', '100', 4);

-- CREATE UNIQUE INDEX sqlite_autoindex_kmmCurrencies_1
--     ON kmmCurrencies (ISOcode);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS kmmCurrencies;

-- +goose StatementEnd
