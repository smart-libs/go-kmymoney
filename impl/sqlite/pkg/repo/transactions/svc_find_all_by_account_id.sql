SELECT
    t.id,
    t.txType,
    t.postDate,
    t.memo,
    t.entryDate,
    t.currencyId,
    t.bankId,
    s.splitID,
    s.payeeID,
    s.reconcileDate,
    s.action,
    s.reconcileFlag,
    s.value,
    s.valueFormatted,
    s.shares,
    s.sharesFormatted,
    s.price,
    s.priceFormatted,
    s.costCenterID,
    s.checkNumber,
    s.postDate,
    s.bankID,
    s.memo
FROM
    kmmTransactions t,
    kmmSplits s
WHERE
    t.ID = s.transactionID AND
    s.accountID = $1