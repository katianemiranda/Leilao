package bid

import (
	"context"
	"time"

	"github.com/katianemiranda/leilao/configuration/logger"
	"github.com/katianemiranda/leilao/internal/entity/bid_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidById(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}
	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error finding bids by auction ID", err)
		return nil, internal_error.NewInternalServerError("Error finding bids by auction ID")
	}

	var bidEntitiesMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntitiesMongo); err != nil {
		logger.Error("Error decoding bids from MongoDB", err)
		return nil, internal_error.NewInternalServerError("Error decoding bids from MongoDB")
	}

	var bidEntities []bid_entity.Bid
	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			ID:        bidEntityMongo.ID,
			UserID:    bidEntityMongo.UserID,
			AuctionID: bidEntityMongo.AuctionID,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error("Error finding winning bid by auction ID", err)
		return nil, internal_error.NewInternalServerError("Error finding winning bid by auction ID")
	}

	return &bid_entity.Bid{
		ID:        bidEntityMongo.ID,
		UserID:    bidEntityMongo.UserID,
		AuctionID: bidEntityMongo.AuctionID,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
