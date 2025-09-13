package routes

import (
	"customer-api/repository"
	"customer-api/service"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql" // this is important import otherwise the driver will not be registered
	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()

	// Wiring the components together....
	dbClient := NewDBClient()
	repo := repository.NewCustomerRepositoryDB(dbClient)
	cs := service.NewCustomerService(repo)
	ch := NewCustomerHandler(cs)

	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{id}", ch.getCustomer).Methods(http.MethodGet)

	log.Println("Starting server on 8080 port....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func WriteResponse(w http.ResponseWriter, response any, httpStatusCode int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}

// this function needs to be called only once and that too at the time of server start
func NewDBClient() *sql.DB {
	db, err := sql.Open("mysql", "root:thecodecamp@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
