package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now().Format(time.ANSIC)
		fmt.Fprintf(w, "<b>Hello!!!</b><hr>Current Time is: %q", timeNow)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
