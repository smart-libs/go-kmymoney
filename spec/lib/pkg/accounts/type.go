package accounts

import "fmt"

type Type int

const (
	Unknown        Type = iota /**< For error handling */
	Checkings                  /**< Standard checking account */
	Savings                    /**< Typical savings account */
	Cash                       /**< Denotes a shoe-box or pillowcase stuffed with cash */
	CreditCard                 /**< Credit card accounts */
	Loan                       /**< Loan and mortgage accounts (liability) */
	CertificateDep             /**< Certificates of Deposit */
	Investment                 /**< Investment account */
	MoneyMarket                /**< Money Market Account */
	Asset                      /**< Denotes a generic asset account.*/
	Liability                  /**< Denotes a generic liability account.*/
	Currency                   /**< Denotes a currency trading account. */
	Income                     /**< Denotes an income account */
	Expense                    /**< Denotes an expense account */
	AssetLoan                  /**< Denotes a loan (asset of the owner of this object) */
	Stock                      /**< Denotes an security account as sub-account for an investment */
	Equity                     /**< Denotes an equity account e.g. opening/closing balance */
)

func (t Type) String() string {
	switch t {
	case Unknown:
		return "Unknown"
	case Checkings:
		return "Checkings"
	case Savings:
		return "Savings"
	case Cash:
		return "Cash"
	case CreditCard:
		return "CreditCard"
	case Loan:
		return "Loan"
	case CertificateDep:
		return "CertificateDep"
	case Investment:
		return "Investment"
	case MoneyMarket:
		return "MoneyMarket"
	case Asset:
		return "Asset"
	case Liability:
		return "Liability"
	case Currency:
		return "Currency"
	case Income:
		return "Income"
	case Expense:
		return "Expense"
	case AssetLoan:
		return "AssetLoan"
	case Stock:
		return "Stock"
	case Equity:
		return "Equity"
	default:
		panic(fmt.Errorf("unknow type(%d)", t))
	}
}
