package bid

import (
	"context"
	"sync"

	"github.com/katianemiranda/leilao/configuration/logger"
	"github.com/katianemiranda/leilao/internal/entity/auction_entity"
	"github.com/katianemiranda/leilao/internal/entity/bid_entity"
	"github.com/katianemiranda/leilao/internal/infra/database/auction"
	"github.com/katianemiranda/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	ID        string  `bson:"_id"`
	UserID    string  `bson:"user_id"`
	AuctionID string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auctionEntity, err := bd.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionID)
			if err != nil {
				logger.Error("Error finding auction by ID", err)
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				ID:        bidValue.ID,
				UserID:    bidValue.UserID,
				AuctionID: bidValue.AuctionID,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			if _, err := bd.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error inserting bid into MongoDB", err)
				return
			}

		}(bid)
	}

	wg.Wait()
	return nil

}
