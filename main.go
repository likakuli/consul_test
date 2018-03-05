package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/watch", watch)
	fmt.Println("Start running...")
	log.Fatal(http.ListenAndServe(":9999", nil))

	select {}
}

func watch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(body))
	default:
		fmt.Println("not support")
	}
}
