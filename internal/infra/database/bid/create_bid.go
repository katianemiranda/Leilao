package bid

import (
	"context"
	"os"
	"sync"
	"time"

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
	Collection            *mongo.Collection
	AuctionRepository     *auction.AuctionRepository
	auctionInterval       time.Duration
	auctionStatusMap      map[string]auction_entity.AuctionStatus
	auctionEndTimeMap     map[string]time.Time
	auctionStatusMapMutex *sync.Mutex
	auctionEndTimeMutex   *sync.Mutex
}

var _ bid_entity.BidEntityRepository = (*BidRepository)(nil)

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		auctionInterval:       getAuctionInterval(),
		auctionStatusMap:      make(map[string]auction_entity.AuctionStatus),
		auctionEndTimeMap:     make(map[string]time.Time),
		auctionStatusMapMutex: &sync.Mutex{},
		auctionEndTimeMutex:   &sync.Mutex{},
		Collection:            database.Collection("bids"),
		AuctionRepository:     auctionRepository,
	}
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

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}

func (r *BidRepository) FindBidAndAuctionById(
	ctx context.Context,
	auctionId string,
) (*bid_entity.Bid, *internal_error.InternalError) {

	filter := map[string]interface{}{
		"auction_id": auctionId,
	}

	var bidMongo BidEntityMongo

	err := r.Collection.FindOne(ctx, filter).Decode(&bidMongo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		logger.Error("Error finding bid by auction id", err)
		return nil, internal_error.NewInternalServerError("Error finding bid")
	}

	bid := &bid_entity.Bid{
		ID:        bidMongo.ID,
		UserID:    bidMongo.UserID,
		AuctionID: bidMongo.AuctionID,
		Amount:    bidMongo.Amount,
		Timestamp: time.Unix(bidMongo.Timestamp, 0),
	}

	return bid, nil
}

func (r *BidRepository) FindBidByAuctionId(
	ctx context.Context,
	auctionId string,
) ([]bid_entity.Bid, *internal_error.InternalError) {

	filter := map[string]interface{}{
		"auction_id": auctionId,
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error finding bids by auction id", err)
		return nil, internal_error.NewInternalServerError("Error finding bids")
	}
	defer cursor.Close(ctx)

	var bids []bid_entity.Bid

	for cursor.Next(ctx) {
		var bidMongo BidEntityMongo

		if err := cursor.Decode(&bidMongo); err != nil {
			logger.Error("Error decoding bid", err)
			continue
		}

		bid := bid_entity.Bid{
			ID:        bidMongo.ID,
			UserID:    bidMongo.UserID,
			AuctionID: bidMongo.AuctionID,
			Amount:    bidMongo.Amount,
			Timestamp: time.Unix(bidMongo.Timestamp, 0),
		}

		bids = append(bids, bid)
	}

	return bids, nil
}
