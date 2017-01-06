package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/takama/daemon"
	"log"
	"os/signal"
	"syscall"
)

const (
	name        = "godoremi"
	description = "Golang daemon for docker remote API"
	// port for daemon lister
	port = ":9977"
	sock = "/var/run/docker.sock"
)

var stdlog, errlog *log.Logger

type Service struct {
	daemon.Daemon
}

func init() {
	stdlog = log.New(os.Stdout, "", 0)
	errlog = log.New(os.Stderr, "", 0)
}

func main() {
	srv, err := daemon.New(name, description, "")
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}

/**
 * This method for managing daemon and collecting it into a channel
 * @param service Struct
 * @return usage, error
 */
func (service *Service) Manage() (string, error) {
	usage := "Usage: godoremi install | remove | start | stop | status"
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}
	// handle the incoming request and set response to JSON
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
	stdlog.Println("remote api starting at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return "Possibly was a problem with the port binding", err
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			stdlog.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
	return usage, nil
}

/**
 * Accept a client connection and collect it in a channel
 * @param listener, listen channel
 * @return no return
 */
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}
/**
 * Handler incoming client connection
 * @param client
 * @return no return
 */
func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}

func dieIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

/**
 * Setup connection for docker socket
 * @param proto, address
 * @return connection, error
 */
func fakeDial(proto, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", sock)
}
