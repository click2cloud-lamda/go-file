package main

import (
	"fmt"

	//"encoding/json"
	//"go-file/models"
"go-file/router"
	"log"
	"net/http"
	//"os"
	//"path/filepath"

	//"github.com/gorilla/mux"
)


func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	// var host string
	var x=[]int{1,2,3,4,5,}
	for _,z:=range x{
		fmt.Println(z)
	} 
	fmt.Print("for test purpose007")
	fmt.Println("localhost server starting check")
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
