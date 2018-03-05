package main

import (
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
)

func main() {
	http.HandleFunc("/watch", watch)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func watch(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(body)
}
