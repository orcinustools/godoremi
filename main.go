package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const sock = "/var/run/docker.sock"

func main() {
	tr := &http.Transport{
		Dial: fakeDial,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://localhost/images/json")
	dieIf(err)
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
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
