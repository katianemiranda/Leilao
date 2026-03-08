package bid_usecase

import (
	"context"

	"github.com/katianemiranda/leilao/internal/internal_error"
)

func (bu *BidUseCase) FindBidByAuctionId(
	ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bidList, err := bu.BidRepository.FindBidsByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var bidOutputList []BidOutputDTO
	for _, bid := range bidList {
		bidOutputList = append(bidOutputList, BidOutputDTO{
			Id:        bid.ID,
			UserId:    bid.UserID,
			AuctionId: bid.AuctionID,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidOutputList, nil
}

func (bu *BidUseCase) FindWinningBidByAuctionId(
	ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError) {
	bidEntity, err := bu.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	bidOutput := &BidOutputDTO{
		Id:        bidEntity.ID,
		UserId:    bidEntity.UserID,
		AuctionId: bidEntity.AuctionID,
		Amount:    bidEntity.Amount,
		Timestamp: bidEntity.Timestamp,
	}

	return bidOutput, nil
}
