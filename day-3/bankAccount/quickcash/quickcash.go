package quickcash

import "fmt"

type QuickCash struct {
	AccountTypes []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string, error) {
	for _, account := range qc.AccountTypes {

		canWithDraw, err := account.CanWithDraw(amount)

		if err != nil {
			return 0, account.GetIdentifier(), err
		}
		if canWithDraw {
			result := account.WithDraw(amount)
			return result, account.GetIdentifier(), nil
		}

	}
	return 0, "", fmt.Errorf("no account found")
}
