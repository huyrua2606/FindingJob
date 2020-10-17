package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	route.HandleFunc("/job/", getJobs).Methods("GET")

	route.HandleFunc("/getaccount", getAccount).Methods("GET")

	route.HandleFunc("/regis", createAccount).Methods("POST")

	route.HandleFunc("/register", createAccount2).Methods("POST")

	route.HandleFunc("/add", createJob).Methods("POST")

	route.HandleFunc("/login", checkLoginRequest).Methods("GET")

	route.HandleFunc("/upload", uploadFileHandler()).Methods("POST")

	route.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./tmp"))))

	route.PathPrefix("/jobimages/").Handler(http.StripPrefix("/jobimages/", http.FileServer(http.Dir("./jobImage"))))

	route.HandleFunc("/applyjob", applyJob).Methods("PUT")

	route.HandleFunc("/getpostedjob", getPostedjob).Methods("GET")

	route.HandleFunc("/checkjob", checkJobExist).Methods("GET")

	//route.HandleFunc("/getjob", getJob).Methods("GET")

	route.HandleFunc("/removejobapplied", removeJobApplied).Methods("POST")

	fmt.Println("Routes are Loaded.")
}
