pavackage auction_usecase

import 
	"context"

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

func (au *AuctionUseCase) FindAuctions(ctx context.Context) ([]*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntities, err := au.auctionRepositoryInterface.FindAuctions(ctx)
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
