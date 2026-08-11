package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getLineCnt() int {
	file, err := os.OpenFile("/data/runs.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lineCnt := 0
	for scanner.Scan() {
		lineCnt++
	}

	return lineCnt
}

func bump() {
	file, err := os.OpenFile("/data/runs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = file.WriteString("1\n")
	if err != nil {
		panic(err)
	}
}

func pong(w http.ResponseWriter, req *http.Request) {
	times := getLineCnt()
	bump()

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "pong ("+strconv.Itoa(times)+")")
}

func main() {
	http.HandleFunc("/ping", pong)
	log.Fatal(http.ListenAndServe(":6767", nil))
}
