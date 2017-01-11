package service

import (
	"io/ioutil"
	"net"
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
	tegoHttp "github.com/wejick/tego/http"
)

func HTTPGetPing(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/_ping")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

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
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetVersion(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/version")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetImages(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/images/json")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/containers/json")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetNetworks(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/networks")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetInspectNetworks(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/networks/"+ps.ByName("id"))
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetInspectImage(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/images/"+ps.ByName("name")+"/json")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetHistoryImage(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/images/"+ps.ByName("name")+"/history")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetSearchImage(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/images/search?term="+r.FormValue("term"))
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetEvents(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/events")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetVolumes(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/volumes")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetInspectVolumes(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/volumes/"+ps.ByName("name"))
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func dockerDial(network, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", dockerSocket)
}
