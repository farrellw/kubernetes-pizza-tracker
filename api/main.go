package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var col *mgo.Collection

func main() {
	// DialInfo holds options for establishing a session.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("databaseURL")},
		Timeout:  15 * time.Second,
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	// to Cosmos database (using Azure Cosmos DB's API for MongoDB).
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()

	col = session.DB("pizza").C("Pizza")

	r := mux.NewRouter()
	r.HandleFunc("/", healthCheck).Methods("GET")
	r.HandleFunc("/time", getTime).Methods("GET")
	r.HandleFunc("/create", createTime).Methods("GET")
	r.HandleFunc("/time/all", allTime).Methods("POST")
	srv := &http.Server{
		Handler:      handlers.CORS(handlers.AllowedOrigins([]string{"*"}), handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "HEAD"}), handlers.AllowedHeaders([]string{"Access-Control-Request-Method", "Access-Control-Request-Headers"}))(r),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Alive")
}

func getTime(w http.ResponseWriter, r *http.Request) {
	var allMyData []timeResponse
	err := col.Find(bson.M{}).All(&allMyData)

	if err != nil {
		fmt.Println(err)
	}

	var lowest timeResponse
	for i, e := range allMyData {
		if i == 0 || e.PizzaTime.After(lowest.PizzaTime) {
			lowest = e
		}

	}

	json.NewEncoder(w).Encode(lowest)
}

type timeResponse struct {
	PizzaTime time.Time     `bson:"pizzaTime" json:"pizzaTime"`
	ID        bson.ObjectId `bson:"_id,omitempty" json:"ID"`
}

func createTime(w http.ResponseWriter, r *http.Request) {
	pizza := timeResponse{
		PizzaTime: time.Now(),
	}

	err := col.Insert(pizza)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(pizza)
}

func allTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}
