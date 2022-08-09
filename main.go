package main

import (
	"context"
	"go-project/controller"
	"go-project/repository"
	"go-project/service"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	uri := "mongodb+srv://lieuvinhthu:Abc123456@cluster0.m2rac.mongodb.net/?retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	//quickstartDatabase := client.Database("Cluster0")
	//reservesCollection := quickstartDatabase.Collection("reserves")

	repo := repository.NewReserveRepo(client)
	service := service.NewService(repo)
	controller := controller.NewController(service)

	myRouter.HandleFunc("/", controller.HomePage)
	myRouter.HandleFunc("/booking", controller.CreateNewBooking).Methods("POST")
	myRouter.HandleFunc("/getAllTable", controller.GetAllTable).Queries("date", "{date}")
	myRouter.HandleFunc("/cancelBooking", controller.CancelBooking).Queries("phoneNumber", "{phoneNumber}")
	myRouter.HandleFunc("/updateBooking", controller.UpdateBooking).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	logrus.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
