package auctioncontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	resterr "github.com/katianemiranda/leilao/configuration/rest_err"
	"github.com/katianemiranda/leilao/internal/infra/api/web/validation"
	"github.com/katianemiranda/leilao/internal/usecase/auction_usecase"
)

type auctionController struct {
	auctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *auctionController {
	return &auctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (au *auctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := au.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := resterr.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)

}
