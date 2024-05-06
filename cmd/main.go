package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const addr = ":8000"

func main() {
	podName, _ := os.LookupEnv("HOSTNAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, fmt.Sprintf("Pod #%s is running", podName))
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server running and listening on port %s\n", addr)
}
