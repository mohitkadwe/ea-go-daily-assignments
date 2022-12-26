package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	accountTypes := []Withdrawable{
		&SavingsAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType, err := qc.getCash(50)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, float64(50), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {

	accountTypes := []Withdrawable{
		&CreditCardAccount{1000},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType, err := qc.getCash(500)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromPaytmWallet(t *testing.T) {

	accountTypes := []Withdrawable{
		&paytmWallet{1000},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType, err := qc.getCash(500)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromGpayWallet(t *testing.T) {

	accountTypes := []Withdrawable{
		&GpayWallet{1000},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType, err := qc.getCash(500)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestInsufficientBalanceInSavingAccount(t *testing.T) {

	accountTypes := []Withdrawable{
		&SavingsAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	_, _, err := qc.getCash(500)
	if err != nil {
		assert.Equal(t, "Insufficient balance in Savings account. Current balance: 100.000000, Lacking amount: 500.000000", err.Error())
	}
}

func TestInsufficientBalanceInCreditCardAccount(t *testing.T) {

	accountTypes := []Withdrawable{
		&CreditCardAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	_, _, err := qc.getCash(500)
	if err != nil {
		assert.Equal(t, "Insufficient balance in Credit Card account. Current balance: 100.000000, Lacking amount: 500.000000", err.Error())
	}
}

func TestInsufficientBalanceInPaytmWallet(t *testing.T) {

	accountTypes := []Withdrawable{
		&paytmWallet{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	_, _, err := qc.getCash(500)
	if err != nil {
		assert.Equal(t, "Insufficient balance in Paytm Wallet account. Current balance: 100.000000, Lacking amount: 500.000000", err.Error())
	}
}

func TestInsufficientBalanceInGpayWallet(t *testing.T) {

	accountTypes := []Withdrawable{
		&GpayWallet{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	_, _, err := qc.getCash(500)
	if err != nil {
		assert.Equal(t, "Insufficient balance in Gpay Wallet account. Current balance: 100.000000, Lacking amount: 500.000000", err.Error())
	}
}
