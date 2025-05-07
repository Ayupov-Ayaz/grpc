package actions

type Bet struct {
	id     string
	userID string
	amount uint64
}

func NewBetRequest(id, userID string, amount uint64) Bet {
	return Bet{
		id:     id,
		userID: userID,
		amount: amount,
	}
}

func (b Bet) GetID() string {
	return b.id
}

func (b Bet) GetUserID() string {
	return b.userID
}

func (b Bet) GetAmount() uint64 {
	return b.amount
}
