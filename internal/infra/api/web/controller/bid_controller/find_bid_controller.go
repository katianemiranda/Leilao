package bidcontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	rest_err "github.com/katianemiranda/leilao/configuration/rest_err"
)

func (u *BidController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auctionId")

	//localhost:8080/user?userId=123456789

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "auctionId must be a valid UUID",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	bidOutputList, err := u.bidUsecase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bidOutputList)

}
