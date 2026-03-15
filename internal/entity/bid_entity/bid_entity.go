package bid_entity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/katianemiranda/leilao/internal/internal_error"
)

type Bid struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepository interface {
	CreateBid(
		ctx context.Context,
		bidEntities []Bid) *internal_error.InternalError

	FindBidByAuctionId(
		ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}

func CreateBid(userId string, auctionId string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		ID:        uuid.New().String(),
		UserID:    userId,
		AuctionID: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	if err := uuid.Validate(b.UserID); err != nil {
		return internal_error.NewBadRequestError("invalid user_id format")
	}
	if err := uuid.Validate(b.AuctionID); err != nil {
		return internal_error.NewBadRequestError("invalid auction_id format")
	}
	if b.UserID == "" {
		return internal_error.NewBadRequestError("user_id is required")
	}

	if b.AuctionID == "" {
		return internal_error.NewBadRequestError("auction_id is required")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("amount must be greater than zero")
	}

	return nil
}
