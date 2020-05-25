package middleware

import (
	"encoding/json"
	"fmt"
	"go-file/models"
	"log"
	"net/http"
	"os"
	//"path/filepath"
	"io/ioutil"
)

var (
	file *os.File
	err  error
)

type response struct {
	//ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func create(name string) {
	if _, err := os.Stat(name); err == nil {
		fmt.Printf("File exists\n")
		return

	}

	file, err = os.Create(name)

	if err != nil {
		panic(err)
	}
	log.Println("File Created")
	//	log.Println(file)
	//file.Close()
}
func Createfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	create(f.Name)
	res := response{
		Message: "File created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func rename(name, rname string) {
	err = os.Rename(name, rname)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File Renamed")
}
func Renamefile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	rename(f.Name, f.Rname)
	res := response{
		Message: "File renamed successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func delete(name string) {
	err = os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s Deleted", name)
}

func Deletefile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	delete(f.Name)
	res := response{
		Message: "File deleted successfully",
	}
	json.NewEncoder(w).Encode(res)

}
// func get(dir string) error {
// 	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			fmt.Println("NO FILE")
// 		}

// 		// check if it is a regular file (not dir)
// 		if info.Mode().IsRegular() {
// 			fmt.Println("file name:", info.Name())

// 		}
// 		fmt.Println("NO FILE")
// 	})
// }
func get(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
	   return files, err
	}
	for _,file := range fileInfo {
	   files = append(files, file.Name())
	}
	return files, nil
   }

func Getallfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	users, err := get(".")

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(users)
}
