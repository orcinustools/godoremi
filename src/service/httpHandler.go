package service

import (
	"io/ioutil"
	"net"
	"net/http"
	"log"
	"bytes"

	"github.com/julienschmidt/httprouter"
	tegoHttp "github.com/wejick/tego/http"
	"net/url"
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

func HTTPGetInspectContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/"+ps.ByName("name")+"/json")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPPostCreateContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	var conf = []byte(`{"Image": "ubuntu", "Labels": {"com.example.vendor": "aksaramaya","com.example.license": "MIT","com.example.version": "1.0"}}`)
	req, _ := http.NewRequest("POST", "http://localhost/containers/create", bytes.NewBuffer(conf))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPPostStartContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/start", nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPPostRestartContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		panic(err)
	}
	tty := "1"
	m, _ := url.ParseQuery(u.RawQuery)
	if m != nil {
		tty = m["t"][0]
	}
	req, _ := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/restart?t="+tty, nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPPostRenameContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	req, _ := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/rename?name="+m["name"][0], nil)
	resp, err := client.Do(req)
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPPostStopContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/stop", nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPPostKillContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/kill", nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetTopContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/"+ps.ByName("name")+"/top")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetChangesContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/"+ps.ByName("name")+"/changes")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

func HTTPGetStatsContainers(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/"+ps.ByName("name")+"/stats")
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
