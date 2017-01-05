package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

const sock = "/var/run/docker.sock"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tr := &http.Transport{
			Dial: fakeDial,
		}
		client := &http.Client{Transport: tr}
		var req string
		switch r.URL.Path {
		case "/":
			req = "info"
		case "/images":
			req = "images/json"
		case "/containers":
			req = "containers/json"
		}
		resp, err := client.Get("http://localhost/" + req)
		dieIf(err)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		w.Write(body)
		resp.Body.Close()
	})
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func dieIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

func fakeDial(proto, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", sock)
}
