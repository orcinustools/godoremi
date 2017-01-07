package service

import (
	"io/ioutil"
	"net"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
	tegoHttp "github.com/wejick/tego/http"
)

//HTTPGetInfo get info
func HTTPGetInfo(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/info")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "sorry bro", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

//HTTPGetImages get images
func HTTPGetImages(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/images")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "sorry bro", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

//HTTPGetContainers get Containers
func HTTPGetContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/containers")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "sorry bro", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func dockerDial(network, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", dockerSocket)
}
