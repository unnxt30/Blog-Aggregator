package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/unnxt30/Blog-Aggregator/internal/database"

	_ "github.com/lib/pq"
)


type apiConfig struct {
	DB *database.Queries;
}


func main(){
	godotenv.Load();
	portNumber := os.Getenv("PORT")
	connectionString := os.Getenv("CONN_STRING");

	db, err := sql.Open("postgres", connectionString);
	if err != nil{
		log.Fatal(err);
	}

	dbQueries := database.New(db);

	var myConfig apiConfig;
	myConfig.DB = dbQueries	;

	mux := http.NewServeMux()
	var server http.Server

	mux.HandleFunc("GET /v1/healthz", GetHealthz);
	mux.HandleFunc("GET /v1/err", GetErr);
	mux.HandleFunc("POST /v1/users", myConfig.HandleUserFunc);

	server.Addr = fmt.Sprintf(":%v", portNumber)
	server.Handler = mux;

	server.ListenAndServe();
}
