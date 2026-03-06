package bid_entity

import (
	"context"
	"time"

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
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError

	FindBidAndAuctionById(ctx context.Context, id string) (*Bid, *internal_error.InternalError)

	FindBidsByAuctionId(ctx context.Context, auctionId string) ([]*Bid, *internal_error.InternalError)

	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
