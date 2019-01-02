package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	html := "This is hapa! <br> A simple HAProxy Agent"
	fmt.Fprintf(w, html)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	html := "This is hapa! <br> A simple HAProxy Agent"
	fmt.Fprintf(w, html)
}

func blacklistHandler(w http.ResponseWriter, r *http.Request) {

	restartCMD := exec.Command("/bin/sh", "-c", "sudo systemctl restart haproxy.service")
	statusCMD := exec.Command("/bin/sh", "-c", "sudo systemctl status haproxy.service")

	html := "restarting ...<br>"
	var out bytes.Buffer

	restartCMD.Stdout = &out
	restartErr := restartCMD.Run()
	if restartErr != nil {
		log.Fatal(restartErr)
	}
	html += out.String()

	html += "<br>statusing ...<br>"
	statusCMD.Stdout = &out
	statusErr := statusCMD.Run()
	if statusErr != nil {
		log.Fatal(statusErr)
	}
	html += out.String()

	fmt.Fprintf(w, html)
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/bl", blacklistHandler)
	http.ListenAndServe(":9923", nil)

}
