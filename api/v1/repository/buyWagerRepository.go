package repository

import (
	"context"

	"github.com/aharal/wager/api/models"
	"github.com/aharal/wager/configs"
	"github.com/aharal/wager/database/connection"
)

//BuyWagerConnRepository Struct
type BuyWagerConnRepository struct {
	ConnectionService connection.ConnectionInterface
}

//NewBuyWagerRepository Func
func NewBuyWagerRepository(connectionService connection.ConnectionInterface) *BuyWagerConnRepository {
	return &BuyWagerConnRepository{connectionService}
}

//GetAllBuyWager Func
func (c BuyWagerConnRepository) GetAllBuyWager(ctx context.Context, albumID int) ([]models.BuyWager, error) {
	photos := []models.BuyWager{}
	err := c.ConnectionService.GetDB().Table("photo").Where("album_id=?", albumID).Find(&photos).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:No BuyWager Found In DB: ", err)
		return photos, err
	}
	return photos, nil
}

//GetBuyWagerByID Func
func (c BuyWagerConnRepository) GetBuyWagerByWagerIDAndID(ctx context.Context, albumID int, id int) (models.BuyWager, error) {
	photos := models.BuyWager{}
	err := c.ConnectionService.GetDB().Table("photo").Where("album_id=? AND id=?", albumID, id).First(&photos).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:BuyWager Not Found In DB: ", err)
		return photos, err
	}
	return photos, nil
}
