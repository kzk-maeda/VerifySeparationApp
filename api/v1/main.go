package main

import (
	"fmt"
	"net/http"
	"time"

	"main/controller"
	"main/util"
)

func main() {
	fmt.Println("Start App v1", util.Version(), "started at", util.Config.Address)

	mux := http.NewServeMux()

	// Routing
	mux.HandleFunc("/index", Index)

	mux.HandleFunc("/auth/login", controller.Login)
	mux.HandleFunc("/auth/logout", controller.Logout)
	mux.HandleFunc("/auth/signup", controller.Signup)

	// Setting Server
	server := &http.Server{
		Addr:           util.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(util.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(util.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}

// Index is healthcheck method
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /index")
}
