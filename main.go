package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func read(r http.ResponseWriter, r2 *http.Request) {
	f, err := exec.Command("journalctl").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(r, string(f))
	log.Print("Returned log file to browser!")
}
func main() {
	http.HandleFunc("/", read)
	http.ListenAndServe(":8080", nil)
}
