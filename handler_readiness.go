package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // specific func signature
	respondwithJSON(w,200,struct{}{})
}
