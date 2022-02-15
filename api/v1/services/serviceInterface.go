package services

import (
	"context"

	"github.com/aharal/wager/api/models"
)

//WagerServices Inteface
type WagerServices interface {
	GetAllWager(context.Context) (AllWagerResponse, error)
	Search(context.Context, int) (models.Wager, error)
	BuyWagerSearch(context.Context, int, int) (models.BuyWager, error)
}

//BuyWagerServices Intefaces
type BuyWagerServices interface {
	GetAllBuyWager(context.Context, int) (AllBuyWagerResponse, error)
	Search(context.Context, int, int) (models.BuyWager, error)
}
