package bid_entity

import "time"

type Bid struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}
