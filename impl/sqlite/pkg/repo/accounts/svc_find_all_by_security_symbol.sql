select a.id,
       institutionId,
       parentId,
       lastReconciled,
       lastModified,
       openingDate,
       accountNumber,
       accountType,
       accountName,
       description,
       currencyId,
       balance,
       transactionCount
from kmmAccounts a,
     kmmSecurities s
where s.id = a.currencyId
  AND s.symbol = $1
