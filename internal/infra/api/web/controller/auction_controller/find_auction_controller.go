package auctioncontroller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	rest_err "github.com/katianemiranda/leilao/configuration/rest_err"
	"github.com/katianemiranda/leilao/internal/usecase/auction_usecase"
)

func (au *AuctionController) FindAuctionById(c *gin.Context) {
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

	auctionData, err := au.auctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)

}

func (au *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := au.auctionUseCase.FindAuctions(context.Background(), auction_usecase.AuctionStatus(statusNumber), category, productName)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)

}

func (au *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
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

	auctionData, err := au.auctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)

}
