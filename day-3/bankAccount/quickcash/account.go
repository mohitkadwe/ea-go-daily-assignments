package quickcash

import (
	"fmt"
)

type SavingsAccount struct {
	balance float64
}

type CreditCardAccount struct {
	balance float64
}

type paytmWallet struct {
	balance float64
}

type GpayWallet struct {
	balance float64
}

type customError struct {
	currentBalance float64
	lackingAmount  float64
	accountType    string
}

func (ce *customError) Error() string {
	return fmt.Sprintf("Insufficient balance in %s account. Current balance: %f, Lacking amount: %f", ce.accountType, ce.currentBalance, ce.lackingAmount)
}

func (fsa *SavingsAccount) WithDraw(amount float64) float64 {
	return fsa.balance - amount
}

func (fsa *SavingsAccount) CanWithDraw(amount float64) (bool, error) {
	if fsa.balance < amount {
		return false, &customError{fsa.balance, amount, "Savings"}
	}
	return true, nil
}

func (fcc *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}

func (fcc *CreditCardAccount) WithDraw(amount float64) float64 {
	return fcc.balance - amount
}

func (fcc *CreditCardAccount) CanWithDraw(amount float64) (bool, error) {
	if fcc.balance < amount {
		return false, &customError{fcc.balance, amount, "Credit Card"}
	}
	return true, nil
}

func (fcc *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}

func (fpw *paytmWallet) WithDraw(amount float64) float64 {
	return fpw.balance - amount
}

func (fpw *paytmWallet) CanWithDraw(amount float64) (bool, error) {
	if fpw.balance < amount {
		return false, &customError{fpw.balance, amount, "Paytm Wallet"}
	}
	return true, nil
}

func (fpw *paytmWallet) GetIdentifier() string {
	return "PAYTM_WALLET"
}

func (fgw *GpayWallet) WithDraw(amount float64) float64 {
	return fgw.balance - amount
}

func (fgw *GpayWallet) CanWithDraw(amount float64) (bool, error) {
	if fgw.balance < amount {
		return false, &customError{fgw.balance, amount, "Gpay Wallet"}
	}
	return true, nil
}

func (fgw *GpayWallet) GetIdentifier() string {
	return "GPAY_WALLET"
}
