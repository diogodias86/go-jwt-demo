package main

import (
	"jwt_demo/user"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", user.Signin)
	http.HandleFunc("/welcome", user.Welcome)
	// http.HandleFunc("refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
