package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/katianemiranda/leilao/configuration/database/mongodb"
	auction_controller "github.com/katianemiranda/leilao/internal/infra/api/web/controller/auction_controller"
	bid_controller "github.com/katianemiranda/leilao/internal/infra/api/web/controller/bid_controller"
	user_controller "github.com/katianemiranda/leilao/internal/infra/api/web/controller/user_controller"
	"github.com/katianemiranda/leilao/internal/infra/database/auction"
	"github.com/katianemiranda/leilao/internal/infra/database/bid"
	"github.com/katianemiranda/leilao/internal/infra/database/user"
	"github.com/katianemiranda/leilao/internal/usecase/auction_usecase"
	"github.com/katianemiranda/leilao/internal/usecase/bid_usecase"
	"github.com/katianemiranda/leilao/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")

		return
	}

	databaseClient, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
		return
	}

	fmt.Println("Successfully connected to MongoDB:", databaseClient.Name())

	router := gin.Default()

	userController, auctionController, bidController := initDependencies(databaseClient)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionController.FindAuctionById)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindAuctionById)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")

}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	auctionController *auction_controller.AuctionController,
	bidController *bid_controller.BidController,
) {
	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}
