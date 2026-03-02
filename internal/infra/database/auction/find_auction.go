package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/katianemiranda/leilao/configuration/logger"
	"github.com/katianemiranda/leilao/internal/entity/auction_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	fillter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo
	err := ar.Collection.FindOne(ctx, fillter).Decode(&auctionEntityMongo)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by id = %s", id), err)
		return nil, internal_error.NewInternalServerError("Error trying to find auction")
	}

	auctionEntity := &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}

	return auctionEntity, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, categoty, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {

	fillter := bson.M{}

	if status != 0 {
		fillter["status"] = status
	}

	if categoty != "" {
		fillter["category"] = categoty
	}

	if productName != "" {
		fillter["product_name"] = primitive.Regex{Pattern: productName, Options: "i"}
	}

	cursor, err := ar.Collection.Find(ctx, fillter)
	if err != nil {
		logger.Error("Error trying to find auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find auctions")
	}

	defer cursor.Close(ctx)

	var auctionsEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionsEntityMongo); err != nil {
		logger.Error("Error trying to decode auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to decode auctions")
	}

	var auctionsEntity []auction_entity.Auction
	for _, auctionEntityMongo := range auctionsEntityMongo {
		auctionsEntity = append(auctionsEntity, auction_entity.Auction{
			Id:          auctionEntityMongo.Id,
			ProductName: auctionEntityMongo.ProductName,
			Category:    auctionEntityMongo.Category,
			Description: auctionEntityMongo.Description,
			Condition:   auctionEntityMongo.Condition,
			Status:      auctionEntityMongo.Status,
			Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
		})
	}

	return auctionsEntity, nil
}
