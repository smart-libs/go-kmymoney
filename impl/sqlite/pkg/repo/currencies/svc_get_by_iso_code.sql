SELECT
    name,
    type
FROM
    kmmCurrencies
WHERE
    ISOCode = $1
