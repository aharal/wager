package services

import (
	"context"

	"github.com/aharal/wager/api/models"
	"github.com/aharal/wager/api/v1/repository"
)

//AllBuyWagerResponse Struct
type AllBuyWagerResponse struct {
	BuyWager []models.BuyWager
}

//BuyWagerService Service
type BuyWagerService struct {
	photosRepository repository.BuyWagerRepository
}

//NewBuyWagerService func
func NewBuyWagerService(photosRepo repository.BuyWagerRepository) *BuyWagerService {
	return &BuyWagerService{photosRepository: photosRepo}
}

//GetAllBuyWager Func
func (c BuyWagerService) GetAllBuyWager(ctx context.Context, albumID int) (AllBuyWagerResponse, error) {
	var allBuyWagerResponse AllBuyWagerResponse
	photosData, err := c.photosRepository.GetAllBuyWager(ctx, albumID)
	if err != nil {
		return allBuyWagerResponse, err
	}
	photos := models.BuyWager{}
	for _, value := range photosData {
		photos.ID = value.ID
		photos.WagerId = value.WagerId
		photos.BuyWagerId = value.BuyWagerId
		photos.Title = value.Title
		photos.Url = value.Url
		photos.ThumbnailUrl = value.ThumbnailUrl
		allBuyWagerResponse.BuyWager = append(allBuyWagerResponse.BuyWager, photos)
	}
	return allBuyWagerResponse, nil
}

//Search Func
func (c BuyWagerService) Search(ctx context.Context, albumId int, id int) (models.BuyWager, error) {
	photosData := models.BuyWager{}
	photos, err := c.photosRepository.GetBuyWagerByWagerIDAndID(ctx, albumId, id)
	if err != nil {
		return photos, err
	}
	photosData.ID = photos.ID
	photosData.WagerId = photos.WagerId
	photosData.BuyWagerId = photos.BuyWagerId
	photosData.Url = photos.Url
	photosData.ThumbnailUrl = photos.ThumbnailUrl
	return photos, nil
}
