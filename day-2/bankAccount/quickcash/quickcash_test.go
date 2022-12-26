package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	accountTypes := []Withdrawable{
		&SavingsAccount{100},
		&CreditCardAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(100)
	assert.Equal(t, float64(0), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {

	accountTypes := []Withdrawable{
		&CreditCardAccount{100},
		&SavingsAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(50)
	assert.Equal(t, float64(50), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromPaytmWallet(t *testing.T) {

	accountTypes := []Withdrawable{
		&paytmWallet{100},
		&CreditCardAccount{100},
		&SavingsAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(20)
	assert.Equal(t, float64(80), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromGpayWallet(t *testing.T) {

	accountTypes := []Withdrawable{
		&GpayWallet{200},
		&paytmWallet{100},
		&CreditCardAccount{100},
		&SavingsAccount{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(30)
	assert.Equal(t, float64(170), amt)
	assert.Equal(t, accountTypes[0].GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccountWithZeroBalance(t *testing.T) {

	accountTypes := []Withdrawable{
		&creditCardAccountWithZeroBalance{0},
		&SavingsAccount{100},
		&GpayWallet{100},
		&paytmWallet{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(50)
	assert.Equal(t, float64(50), amt)
	assert.Equal(t, accountTypes[1].GetIdentifier(), accType)
}

func TestGetCashFromSavingAccountWithZeroBalance(t *testing.T) {

	accountTypes := []Withdrawable{
		&savingsAccountWithZeroBalance{0},
		&CreditCardAccount{1000},
		&GpayWallet{100},
		&paytmWallet{100},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, accountTypes[1].GetIdentifier(), accType)
}

func TestGetCashFromPaytmWalletWithZeroBalance(t *testing.T) {

	accountTypes := []Withdrawable{
		&paytmWalletWithZeroBalance{0},
		&CreditCardAccount{100},
		&GpayWallet{100},
		&SavingsAccount{200},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(40)
	assert.Equal(t, float64(60), amt)
	assert.Equal(t, accountTypes[1].GetIdentifier(), accType)
}

func TestGetCashFromGpayWalletWithZeroBalance(t *testing.T) {

	accountTypes := []Withdrawable{
		&GpayWalletWithZeroBalance{0},
		&paytmWallet{200},
		&CreditCardAccount{130},
		&SavingsAccount{356},
	}

	qc := QuickCash{
		accountTypes,
	}

	amt, accType := qc.getCash(100)
	assert.Equal(t, float64(100), amt)
	assert.Equal(t, accountTypes[1].GetIdentifier(), accType)
}
