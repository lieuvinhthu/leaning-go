package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-project/model"
	"io/ioutil"
	"net/http"

	"go-project/service"

	"github.com/gorilla/mux"
)

type (
	ReserveController struct {
		reserveService *service.ReserveService
	}

	ReserveControllerImpl interface {
		//CreateNewBooking(w http.ResponseWriter, r *http.Request)
	}
)

func NewController(reserveService *service.ReserveService) *ReserveController {
	return &ReserveController{
		reserveService: reserveService,
	}
}

func (rc *ReserveController) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (rc *ReserveController) CreateNewBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reserve model.Reserve
	json.Unmarshal(reqBody, &reserve)
	ctx := context.Background()
	err := rc.reserveService.CreateNewBooking(ctx, reserve)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode("Success")
}

func (rc *ReserveController) GetAllTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqDate := vars["date"]
	ctx := context.Background()
	filterReserves, err := rc.reserveService.GetAllTable(ctx, reqDate)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(filterReserves)
}

func (rc *ReserveController) CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqPhone := vars["phoneNumber"]
	ctx := context.Background()
	err := rc.reserveService.CancelBooking(ctx, reqPhone)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode("success")
}

func (rc *ReserveController) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reserve model.Reserve
	json.Unmarshal(reqBody, &reserve)
	ctx := context.Background()
	err := rc.reserveService.UpdateBooking(ctx, reserve)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode("Success")
}
