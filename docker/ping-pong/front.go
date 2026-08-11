package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ping(w http.ResponseWriter, req *http.Request) {
	res, err := http.Get("http://south:6767/ping")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
		_, _ = fmt.Fprint(w, "bad status")
		return
	}

	body, _ := ioutil.ReadAll(res.Body)
	bodystr := strings.Split(string(body), " ")
	if bodystr[0] != "pong" {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "bad response")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, bodystr[1])
}

func main() {
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":7676", nil))
}
