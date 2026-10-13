package main

import (
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request) { // specific func signature
	respongwithError(w,400,"something went wrong")
}


