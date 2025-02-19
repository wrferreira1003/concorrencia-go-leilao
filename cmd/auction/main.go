package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/database/mongodb"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/api/web/controller/auction_controller"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/api/web/controller/bid_controller"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/api/web/controller/use_controller"
	auction_repository "github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/database/auction"
	bid_repository "github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/database/bid"
	user_repository "github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/database/user"
	auctionusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/auction_usecase"
	bidusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/bid_usecase"
	userusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// NewMongoDBConnection creates a new connection to the MongoDB database using the provided context.
	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error trying to connect with mongodb", err)
		return
	}

	router := gin.Default()

	auctionController, bidController, userController := initDependencies(databaseConnection)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionController.FindAuctionByID)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionController.FindAuctionsWinningBid)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidAuctionByID)
	router.GET("/user/:userId", userController.FindUserByID)

	router.Run(":8080")

}

func initDependencies(database *mongo.Database) (
	auctionController *auction_controller.AuctionController,
	bidController *bid_controller.BidController,
	userController *use_controller.UseController,
) {

	auctionRepo := auction_repository.NewAuctionRepositoryMongo(database)
	bidRepo := bid_repository.NewBidRepositoryMongo(database, auctionRepo)
	userRepo := user_repository.NewUserRepositoryMongo(database)

	userController = use_controller.NewUseController(userusecase.NewUserUseCase(userRepo))
	auctionController = auction_controller.NewAuctionController(auctionusecase.NewAuctionUseCase(auctionRepo, bidRepo))
	bidController = bid_controller.NewBidController(bidusecase.NewBidUseCase(bidRepo))

	return
}
