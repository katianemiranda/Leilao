package auction_usecase

import (
	"context"

	"github.com/katianemiranda/leilao/internal/entity/auction_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
	"github.com/katianemiranda/leilao/internal/usecase/bid_usecase"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := au.auctionRepositoryInterface.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil

}

func (au *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntities, err := au.auctionRepositoryInterface.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}

	var auctionDTOs []*AuctionOutputDTO
	for _, auction := range auctionEntities {
		auctionDTOs = append(auctionDTOs, &AuctionOutputDTO{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductCondition(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			Timestamp:   auction.Timestamp,
		})
	}

	return auctionDTOs, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError) {
	auction, err := au.auctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionOutputDTO := AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	bidWinning, err := au.bidRepositoryInterface.FindWinningBidByAuctionId(ctx, auction.Id)
	if err != nil {
		return &WinningInfoOutputDTO{
			Auction: auctionOutputDTO,
		}, nil
	}

	bidOutputDTO := &bid_usecase.BidOutputDTO{
		Id:        bidWinning.ID,
		UserId:    bidWinning.UserID,
		AuctionId: bidWinning.AuctionID,
		Amount:    bidWinning.Amount,
		Timestamp: bidWinning.Timestamp,
	}

	return &WinningInfoOutputDTO{
		Auction: auctionOutputDTO,
		Bid:     bidOutputDTO,
	}, nil
}
