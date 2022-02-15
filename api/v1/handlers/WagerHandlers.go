package handlers

import (
	"net/http"
	"strconv"

	v1 "github.com/aharal/wager/api/v1"
	"github.com/aharal/wager/configs"

	"github.com/aharal/wager/api"
	"github.com/aharal/wager/api/v1/services"
)

//HTTPWagerHandler struct
type HTTPWagerHandler struct {
	WagerService services.WagerServices
}

//NewWagerHTTPHandler Function
func NewWagerHTTPHandler(albumsService services.WagerServices, router api.Route) {
	handler := &HTTPWagerHandler{WagerService: albumsService}
	router.Router.HandleFunc("/wagers", handler.PlaceWager).Methods("POST")
	router.Router.HandleFunc("/wagers", handler.GetWagerList).Methods("GET")
}

// PlaceWager : Create a New Wager
func (httpWager HTTPWagerHandler) PlaceWager(w http.ResponseWriter, r *http.Request) {
	albums, err := httpWager.WagerService.GetAllWager(r.Context())
	if err != nil {
		v1.WriteErrorResponse(w, http.StatusNotFound, "No Wagers in the Database!")
		return
	}
	v1.WriteOKResponse(w, albums)
}

func (httpWager HTTPWagerHandler) GetWagerList(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	searchType := params.Get("type")

	if "album" == searchType {
		albumID := params.Get("id")
		aid, err := strconv.Atoi(albumID)
		if err != nil {
			configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
			v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input For Getting Wagers!")
			return
		}
		response, err := httpWager.WagerService.Search(r.Context(), aid)
		if err != nil {
			v1.WriteErrorResponse(w, http.StatusNotFound, "No Wagers in the Database!")
			return
		}
		v1.WriteOKResponse(w, response)
	} else if "photo" == searchType {
		photoID := params.Get("id")
		pid, err := strconv.Atoi(photoID)
		if err != nil {
			configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
			v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input For Getting Wagers!")
			return
		}
		albumID := params.Get("album")
		aid, err := strconv.Atoi(albumID)
		if err != nil {
			configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
			v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input For Getting Wagers!")
			return
		}

		response, err := httpWager.WagerService.BuyWagerSearch(r.Context(), aid, pid)
		if err != nil {
			v1.WriteErrorResponse(w, http.StatusNotFound, "No Wagers in the Database!")
			return
		}
		v1.WriteOKResponse(w, response)
	}
	/*else {
		pID := params.Get("album")
		id, err := strconv.Atoi(pID)
		response := nil
	}*/

}
