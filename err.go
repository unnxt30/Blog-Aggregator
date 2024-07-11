package main

import "net/http"

func GetErr(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 500, "Internal server Error");
}