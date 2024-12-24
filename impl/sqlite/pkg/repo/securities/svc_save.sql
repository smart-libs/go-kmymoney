INSERT INTO kmmSecurities(id,
                          name,
                          symbol,
                          type,
                          typeString,
                          smallestAccountFraction,
                          pricePrecision,
                          tradingMarket,
                          tradingCurrency,
                          roundingMethod)
VALUES ($1,
           $2,
           $3,
           $4,
           $5,
           $6,
           $7,
           $8,
           $9,
           $10)
ON CONFLICT DO UPDATE
    SET name = $2, symbol = $3, type = $4, typestring = $5, smallestAccountFraction = $6, pricePrecision = $7, tradingMarket = $8, tradingCurrency = $9, roundingMethod = $10

