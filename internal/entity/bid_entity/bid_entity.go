package bid_entity

import "time"

type Bid struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	AuctionID string    `json:"auction_id" bson:"auction_id"`
	Amount    float64   `json:"amount" bson:"amount"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
