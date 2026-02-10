package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Started")
	err := http.ListenAndServe(":" + os.Getenv("HTTP_SIMPLE_SERVER_PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
