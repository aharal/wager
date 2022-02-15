package handlers

import (
	"net/http"
	"strconv"

	v1 "github.com/aharal/wager/api/v1"
	"github.com/aharal/wager/configs"

	"github.com/aharal/wager/api"
	"github.com/aharal/wager/api/v1/services"
)

//HTTPBuyWagerHandler Struct
type HTTPBuyWagerHandler struct {
	BuyWagerService services.BuyWagerServices
}

//NewBuyWagerHTTPHandler Function
func NewBuyWagerHTTPHandler(photosService services.BuyWagerServices, router api.Route) {
	handler := &HTTPBuyWagerHandler{BuyWagerService: photosService}
	router.Router.HandleFunc("/photos", handler.GetAllBuyWager).Methods("GET")
}

//GetAllBuyWager Method
func (httpBuyWager HTTPBuyWagerHandler) GetAllBuyWager(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	albumID := params.Get("albumId")
	album, err := strconv.Atoi(albumID)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	photos, err1 := httpBuyWager.BuyWagerService.GetAllBuyWager(r.Context(), album)
	if err1 != nil {
		v1.WriteErrorResponse(w, http.StatusNotFound, "No BuyWager in the Database!")
		return
	}
	v1.WriteOKResponse(w, photos)
}
