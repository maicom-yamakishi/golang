package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"peixe-urbano/src/database"
	"peixe-urbano/src/models"
	"peixe-urbano/src/repository"
	"peixe-urbano/src/responses"
)

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var offer models.Offer
	if err = json.Unmarshal(bodyRequest, &offer); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = offer.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewOfferRepository(db)
	offer.ID, err = repository.Create(offer)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, offer)
}

func GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Geting all Offers!"))
}

func GetOffer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Offer!"))
}

func UpdateOffer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update Offer!"))
}

func DeleteOffer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting Offer!"))
}
