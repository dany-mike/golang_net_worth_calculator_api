package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"golang_net_worth_calculator_api/auth"
	"golang_net_worth_calculator_api/models"
	"golang_net_worth_calculator_api/responses"
	"golang_net_worth_calculator_api/utils/formatError"
)

func (server *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	item := models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	item.Prepare()
	err = item.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != item.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	itemCreated, err := item.SaveItem(server.DB)
	if err != nil {
		formattedError := formatError.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, itemCreated.ID))
	responses.JSON(w, http.StatusCreated, itemCreated)
}

func (server *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	uid, err := auth.ExtractTokenID(r)

	if err == nil {
		items, err := item.FindItemsByUserId(server.DB, uint64(uid))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		//
		responses.JSON(w, http.StatusOK, items)
	}
}

func (server *Server) GetTotalNetWorth(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	uid, err := auth.ExtractTokenID(r)

	if err == nil {
		items, err := item.FindItemsByUserId(server.DB, uint64(uid))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, getTotalNetWorth(*items))
	}
}

func (server *Server) GetItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid item id given to us?
	itemId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("bad request"))
		return
	}

	//Check if the auth token is valid and get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the item exist
	item := models.Item{}
	err = server.DB.Debug().Model(models.Item{}).Where("id = ?", itemId).Take(&item).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("item not found"))
		return
	}

	// If a user attempt to update a item not belonging to him
	if uid != item.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Forbidden"))
		return
	}

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	itemReceived, err := item.FindItemByID(server.DB, itemId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, itemReceived)
}

func (server *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the item id is valid
	itemId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//Check if the auth token is valid and get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the item exist
	item := models.Item{}
	err = server.DB.Debug().Model(models.Item{}).Where("id = ?", itemId).Take(&item).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("item not found"))
		return
	}

	// If a user attempt to update a item not belonging to him
	if uid != item.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Forbidden"))
		return
	}
	// Read the data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	itemUpdate := models.Item{}
	err = json.Unmarshal(body, &itemUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != itemUpdate.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	itemUpdate.Prepare()
	err = itemUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemUpdate.ID = item.ID

	itemUpdated, err := itemUpdate.UpdateItem(server.DB)

	if err != nil {
		formattedError := formatError.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, itemUpdated)
}

func (server *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid item id given to us?
	itemId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the item exist
	item := models.Item{}
	err = server.DB.Debug().Model(models.Item{}).Where("id = ?", itemId).Take(&item).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this item?
	if uid != item.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Forbidden"))
		return
	}
	_, err = item.DeleteItem(server.DB, itemId, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", itemId))
	responses.JSON(w, http.StatusNoContent, "")
}

func getTotalNetWorth(items []models.Item) float64 {

	var (
		total float64 = 0
	)

	for i := 0; i < len(items); i++ {
		total += items[i].Price
	}

	return float64(total)
}
