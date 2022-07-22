package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Table struct {
	Id          string `json:"Id"`
	MaxPosition int    `json:"maxPosition"`
}

type Reserve struct {
	Id          string `json:"Id"`
	DateTime    string `json:"maxPosition"`
	TotalPeople int    `json:"totalPeople"`
	PhoneNumber string `json:"phoneNumber"`
	TableId     string `json:"tableId"`
}

var Reserves = []*Reserve{
	{Id: "1", DateTime: "21/10/2022", TotalPeople: 2, PhoneNumber: "0799004177", TableId: "1"},
	{Id: "2", DateTime: "22/10/2022", TotalPeople: 5, PhoneNumber: "0913616957", TableId: "2"},
}

// Tables := []Table{
// 	Table{Id: "1", maxPosition: 4},
// 	Table{Id: "2", maxPosition: 8},
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// http.HandleFunc("/", homePage)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/getAllTable", getAllTable).Queries("date", "{date}")
	myRouter.HandleFunc("/cancelBooking", cancelBooking).Queries("phoneNumber", "{phoneNumber}")
	myRouter.HandleFunc("/updateBooking", updateBooking).Methods("POST")
	// myRouter.HandleFunc("/all", returnAllArticles)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reserve Reserve
	json.Unmarshal(reqBody, &reserve)
	Reserves = append(Reserves, &reserve)

	json.NewEncoder(w).Encode(Reserves)
}

func getAllTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqDate := vars["date"]
	var filterReserves []*Reserve
	for _, reserve := range Reserves {
		if reserve.DateTime == reqDate {
			filterReserves = append(filterReserves, reserve)
		}
	}
	json.NewEncoder(w).Encode(filterReserves)
}

func cancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqPhoneNumber := vars["phoneNumber"]
	var filterReserves []*Reserve
	for _, reserve := range Reserves {
		if reserve.PhoneNumber != reqPhoneNumber {
			filterReserves = append(filterReserves, reserve)
		}
	}
	json.NewEncoder(w).Encode(filterReserves)
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reqReserve Reserve
	json.Unmarshal(reqBody, &reqReserve)
	for _, reserve := range Reserves {
		if reserve.PhoneNumber == reqReserve.PhoneNumber {
			reserve.Id = reqReserve.Id
			reserve.DateTime = reqReserve.DateTime
			reserve.PhoneNumber = reqReserve.PhoneNumber
			reserve.TableId = reqReserve.TableId
			reserve.TotalPeople = reqReserve.TotalPeople
		}
	}
	json.NewEncoder(w).Encode(Reserves)
}

func main() {
	handleRequests()
}
