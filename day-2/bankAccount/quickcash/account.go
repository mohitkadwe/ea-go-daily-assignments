package quickcash

type SavingsAccount struct {
	balance float64
}

func (fsa *SavingsAccount) WithDraw(amount float64) float64 {
	return fsa.balance - amount
}

func (fsa *SavingsAccount) CanWithDraw(amount float64) bool {
	if fsa.balance < amount {
		return false
	}
	return true
}

func (fsa *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}

type CreditCardAccount struct {
	balance float64
}

func (fsa *CreditCardAccount) WithDraw(amount float64) float64 {
	return fsa.balance - amount
}

func (fsa *CreditCardAccount) CanWithDraw(amount float64) bool {
	if fsa.balance < amount {
		return false
	}
	return true
}

func (fsa *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}

type paytmWallet struct {
	balance float64
}

func (fsa *paytmWallet) WithDraw(amount float64) float64 {
	return fsa.balance - amount
}

func (fsa *paytmWallet) CanWithDraw(amount float64) bool {
	if fsa.balance < amount {
		return false
	}
	return true
}

func (fsa *paytmWallet) GetIdentifier() string {
	return "PAYTM_WALLET"
}

type GpayWallet struct {
	balance float64
}

func (fsa *GpayWallet) WithDraw(amount float64) float64 {
	return fsa.balance - amount
}

func (fsa *GpayWallet) CanWithDraw(amount float64) bool {
	if fsa.balance < amount {
		return false
	}
	return true
}

func (fsa *GpayWallet) GetIdentifier() string {
	return "GPAY_WALLET"
}

type creditCardAccountWithZeroBalance struct {
	balance float64
}

func (fsa *creditCardAccountWithZeroBalance) WithDraw(amount float64) float64 {
	return 0
}

func (fsa *creditCardAccountWithZeroBalance) CanWithDraw(amount float64) bool {
	return false
}

func (fsa *creditCardAccountWithZeroBalance) GetIdentifier() string {
	return "ZERO_CREDIT_CARD_ACCOUNT"
}

type savingsAccountWithZeroBalance struct {
	balance float64
}

func (fsa *savingsAccountWithZeroBalance) WithDraw(amount float64) float64 {
	return 0
}

func (fsa *savingsAccountWithZeroBalance) CanWithDraw(amount float64) bool {
	return false
}

func (fsa *savingsAccountWithZeroBalance) GetIdentifier() string {
	return "ZERO_SAVINGS_ACCOUNT"
}

type paytmWalletWithZeroBalance struct {
	balance float64
}

func (fsa *paytmWalletWithZeroBalance) WithDraw(amount float64) float64 {
	return 0
}

func (fsa *paytmWalletWithZeroBalance) CanWithDraw(amount float64) bool {
	return false
}

func (fsa *paytmWalletWithZeroBalance) GetIdentifier() string {
	return "ZERO_PAYTM_WALLET"
}

type GpayWalletWithZeroBalance struct {
	balance float64
}

func (fsa *GpayWalletWithZeroBalance) WithDraw(amount float64) float64 {
	return 0
}

func (fsa *GpayWalletWithZeroBalance) CanWithDraw(amount float64) bool {
	return false
}

func (fsa *GpayWalletWithZeroBalance) GetIdentifier() string {
	return "ZERO_GPAY_WALLET"
}
