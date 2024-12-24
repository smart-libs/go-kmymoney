SELECT id,
       name,
       type,
       smallestAccountFraction,
       pricePrecision,
       tradingMarket,
       tradingCurrency,
       roundingMethod
FROM kmmSecurities
WHERE symbol = $1