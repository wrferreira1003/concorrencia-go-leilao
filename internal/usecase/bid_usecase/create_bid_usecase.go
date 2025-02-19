package bidusecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/bid_entity"
)

type BidInputDto struct {
	UserID    string  `json:"user_id"`
	AuctionID string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDto struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	AuctionID string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02T15:04:05Z"`
}

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInputDto *BidInputDto) error
	FindBidByIDAuctionId(ctx context.Context, auctionID string) ([]BidOutputDto, error)
	FindWinnerBidByAuctionId(ctx context.Context, auctionID string) (*BidOutputDto, error)
}

type BidUseCase struct {
	bidRepository       bid_entity.BidRepositoryInterface
	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

func NewBidUseCase(
	bidRepository bid_entity.BidRepositoryInterface,
) BidUseCaseInterface {

	bidUseCase := &BidUseCase{
		bidRepository:       bidRepository,
		maxBatchSize:        getMaxBatchSize(),
		batchInsertInterval: getBatchInsertInterval(),
		timer:               time.NewTimer(getBatchInsertInterval()),
		bidChannel:          make(chan bid_entity.Bid, getMaxBatchSize()),
	}

	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

var bidBatch []bid_entity.Bid

func (u *BidUseCase) triggerCreateRoutine(ctx context.Context) {

	go func() {
		defer close(u.bidChannel)

		for {
			select {
			case bidEntity, ok := <-u.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := u.bidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("error creating bid", err)
						}

					}
					return
				}

				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= u.maxBatchSize {
					if err := u.bidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("error creating bid", err)
					}
					bidBatch = nil
					u.timer.Reset(u.batchInsertInterval)
				}

			case <-u.timer.C:

				if err := u.bidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("error creating bid", err)
				}
				bidBatch = nil
				u.timer.Reset(u.batchInsertInterval)

			}
		}
	}()

}

func (u *BidUseCase) CreateBid(ctx context.Context, bidInputDto *BidInputDto) error {

	bidEntity, err := bid_entity.NewBidEntity(
		bidInputDto.UserID,
		bidInputDto.AuctionID,
		bidInputDto.Amount,
	)
	if err != nil {
		return err
	}

	u.bidChannel <- *bidEntity

	return nil
}

func getBatchInsertInterval() time.Duration {
	batchInsertInterval := os.Getenv("BID_BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}
	return duration

}

func getMaxBatchSize() int {
	maxBatchSize, err := strconv.Atoi(os.Getenv("BID_MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}
	return maxBatchSize
}
