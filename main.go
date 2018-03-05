package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/watch", watch)
	log.Fatal(http.ListenAndServe(":9999", nil))
	fmt.Println("Running...")
}

func watch(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}
