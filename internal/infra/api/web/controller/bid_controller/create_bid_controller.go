package bidcontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	resterr "github.com/katianemiranda/leilao/configuration/rest_err"
	"github.com/katianemiranda/leilao/internal/infra/api/web/validation"
	"github.com/katianemiranda/leilao/internal/usecase/bid_usecase"
)

type BidController struct {
	bidUsecase bid_usecase.BidUseCaseInterface
}

func NewBidController(bidUsecase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		bidUsecase: bidUsecase,
	}
}

func (b *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		resterr := validation.ValidateErr(err)
		c.JSON(resterr.Code, resterr)

		return
	}

	err := b.bidUsecase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		resterr := resterr.ConvertError(err)
		c.JSON(resterr.Code, resterr)

		return
	}

	c.Status(http.StatusCreated)
}
