package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transaction
}

func NewOrder(investor *Investor, asset *Asset, shares int, orderType string) *Order {
	return &Order{
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}
