package auction_usecase

import (
	"context"
	"time"

	"github.com/katianemiranda/leilao/internal/entity/auction_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02T15:04:05Z07:00"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(auctionInput.ProductName, auctionInput.Category, auctionInput.Description, auction_entity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return err
	}
	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}
	return nil
}

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	return nil, nil
}

func (au *AuctionUseCase) FindAuctions(ctx context.Context) ([]*AuctionOutputDTO, *internal_error.InternalError) {
	return nil, nil
}
