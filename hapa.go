package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	html := "<html><head>hapa</head><body><h2>This is hapa! </h2><br> A simple HAProxy Agent<br> <br> <a href=\"bl\">restart</a> <a href=\"logs/raw\">show logs</a> </body></html>"
	fmt.Fprintf(w, html)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	html := "This is hapa! \n A simple HAProxy Agent"
	fmt.Fprintf(w, html)
}

func blacklistHandler(w http.ResponseWriter, r *http.Request) {

	restartCMD := exec.Command("/bin/sh", "-c", "sudo systemctl restart haproxy.service")
	statusCMD := exec.Command("/bin/sh", "-c", "sudo systemctl status haproxy.service")

	html := "restarting ...\n"
	var out bytes.Buffer

	restartCMD.Stdout = &out
	restartErr := restartCMD.Run()
	if restartErr != nil {
		log.Fatal(restartErr)
	}
	html += out.String()

	html += "\nstatusing ...\n"
	statusCMD.Stdout = &out
	statusErr := statusCMD.Run()
	if statusErr != nil {
		log.Fatal(statusErr)
	}
	html += out.String()

	fmt.Fprintf(w, html)
}

func logsRawHandler(w http.ResponseWriter, r *http.Request) {

	logsRawCMD := exec.Command("/bin/sh", "-c", "sudo cat /var/log/haproxy-ldap.log")

	var out bytes.Buffer

	logsRawCMD.Stdout = &out
	logsRawErr := logsRawCMD.Run()
	if logsRawErr != nil {
		log.Fatal(logsRawErr)
	}
	html := out.String()

	fmt.Fprintf(w, html)
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/bl", blacklistHandler)
	http.HandleFunc("/logs/raw", logsRawHandler)
	http.ListenAndServe(":9923", nil)

}
