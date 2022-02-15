package repository

import (
	"context"

	"github.com/aharal/wager/api/models"
	"github.com/aharal/wager/configs"
	"github.com/aharal/wager/database/connection"
)

//WagerConnRepository Struct
type WagerConnRepository struct {
	ConnectionService connection.ConnectionInterface
}

//NewWagerRepository Func
func NewWagerRepository(connectionService connection.ConnectionInterface) *WagerConnRepository {
	return &WagerConnRepository{connectionService}
}

//GetAllWagers Func
func (u WagerConnRepository) GetAllWagers(ctx context.Context) ([]models.Wager, error) {
	album := []models.Wager{}
	err := u.ConnectionService.GetDB().Table("album").Find(&album).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:No Wagers Found In DB: ", err)
		return album, err
	}
	return album, nil
}

//GetBuyWagerByWagerIDAndBuyWagerID Func : Search By Id
func (u WagerConnRepository) GetBuyWagerByWagerIDAndBuyWagerID(ctx context.Context, aid int, pid int) (models.BuyWager, error) {
	photo := models.BuyWager{}
	err := u.ConnectionService.GetDB().Table("photo").Where("album_id=? AND id=?", aid, pid).First(&photo).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Wager Not Found In DB: ", err)
		return photo, err
	}
	return photo, nil
}

//GetWagerByID Func : Search By Id
func (u WagerConnRepository) GetWagerByID(ctx context.Context, id int) (models.Wager, error) {
	albums := models.Wager{}
	err := u.ConnectionService.GetDB().Table("album").Where("id=?", id).First(&albums).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Wager Not Found In DB: ", err)
		return albums, err
	}
	return albums, nil
}
