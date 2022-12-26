package quickcash

type QuickCash struct {
	AccountTypes []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	for _, account := range qc.AccountTypes {
		if account.CanWithDraw(amount) {
			result := account.WithDraw(amount)

			if result < 0 {
				return 0, "Insufficient balance"
			}

			return result, account.GetIdentifier()
		}
	}
	return 0, "No account found"
}
