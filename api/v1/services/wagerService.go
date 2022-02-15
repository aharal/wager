package services

import (
	"context"

	"github.com/aharal/wager/api/models"
	"github.com/aharal/wager/api/v1/repository"
)

type AllWagerResponse struct {
	Wagers []models.Wager
}

//WagerService to manage album persistence
type WagerService struct {
	albumsRepository repository.WagerRepository
}

//NewWagerService func
func NewWagerService(albumsRepo repository.WagerRepository) *WagerService {
	return &WagerService{albumsRepository: albumsRepo}
}

//GetAllWagers Func
func (u WagerService) GetAllWager(ctx context.Context) (AllWagerResponse, error) {
	var allWagerResponse AllWagerResponse
	albumsData, err := u.albumsRepository.GetAllWagers(ctx)
	if err != nil {
		return allWagerResponse, err
	}
	albums := models.Wager{}
	for _, value := range albumsData {
		albums.ID = value.ID
		albums.UserId = value.UserId
		albums.Title = value.Title
		allWagerResponse.Wagers = append(allWagerResponse.Wagers, albums)
	}
	return allWagerResponse, nil
}

//Search Func
func (u WagerService) Search(ctx context.Context, id int) (models.Wager, error) {
	albums := models.Wager{}
	albumsData, err := u.albumsRepository.GetWagerByID(ctx, id)
	if err != nil {
		return albums, err
	}
	albums.ID = albumsData.ID
	albums.UserId = albumsData.UserId
	albums.Title = albumsData.Title
	return albums, nil
}

//Search Func
func (u WagerService) BuyWagerSearch(ctx context.Context, aid int, id int) (models.BuyWager, error) {
	photosData := models.BuyWager{}
	photos, err := u.albumsRepository.GetBuyWagerByWagerIDAndBuyWagerID(ctx, aid, id)
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
