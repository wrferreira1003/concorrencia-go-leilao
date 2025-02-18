package auction_entity

import "time"

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	// Assume that the auction status is a number between 1 and 2
	Active    AuctionStatus = iota // 0
	Completed                      // 1
)

const (
	// Assume that the product condition is a number between 1 and 3
	New         ProductCondition = iota // 0
	Used                                // 1
	Refurbished                         // 2
)
