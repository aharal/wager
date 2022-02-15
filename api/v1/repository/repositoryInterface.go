package repository

import (
	"context"

	"github.com/aharal/wager/api/models"
)

//WagerRepository Interface
type WagerRepository interface {
	GetAllWagers(context.Context) ([]models.Wager, error)
	GetWagerByID(context.Context, int) (models.Wager, error)
	GetBuyWagerByWagerIDAndBuyWagerID(context.Context, int, int) (models.BuyWager, error)
}

//BuyWagerRepository Inteface
type BuyWagerRepository interface {
	GetAllBuyWager(context.Context, int) ([]models.BuyWager, error)
	GetBuyWagerByWagerIDAndID(context.Context, int, int) (models.BuyWager, error)
}
