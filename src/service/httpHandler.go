package service

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	tegoHttp "github.com/wejick/tego/http"
)

func dockerDial(network, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", dockerSocket)
}

// Misc
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
func HTTPGetInspectNetwork(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/networks/" + ps.ByName("id"))
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

// Images
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
func HTTPGetInspectImage(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://localhost/images/" + ps.ByName("name") + "/json")
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

	resp, err := client.Get("http://localhost/images/" + ps.ByName("name") + "/history")
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
	resp, err := client.Get("http://localhost/images/search?term=" + r.FormValue("term"))
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

// Containers
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
func HTTPGetInspectContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/" + ps.ByName("name") + "/json")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPPostCreateContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func HTTPPostUpdateContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	var conf = []byte(`{"CpuShares": 512,"CpuPeriod": 100000,"CpuQuota": 50000,"CpusetCpus": "0,1","CpusetMems": "0","Memory": 314572800,"MemorySwap": 514288000}`)
	req, _ := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/update", bytes.NewBuffer(conf))
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
func HTTPPostStartContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func HTTPPostRestartContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func HTTPGetRenameContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func HTTPPostStopContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func HTTPPostKillContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
func HTTPGetTopContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/" + ps.ByName("name") + "/top")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetChangesContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/" + ps.ByName("name") + "/changes")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetStatsContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/" + ps.ByName("name") + "/stats")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetLogsContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	urlReq := "http://localhost/containers/" + ps.ByName("name") + "/logs?" +
		"stderr=" + m["stderr"][0] +
		"&stdout=" + m["stdout"][0] +
		"&timestamps=" + m["timestamps"][0] +
		"&follow=" + m["timestamps"][0] +
		"&tail=" + m["tail"][0]
	resp, err := client.Get(urlReq)
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetExportContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/octet-stream")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/containers/" + ps.ByName("name") + "/export")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPPostResizeContainer(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "text/plain")
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
	req, err := http.NewRequest("POST", "http://localhost/containers/"+ps.ByName("name")+"/resize?h="+m["h"][0]+"&w="+m["w"][0], nil)
	resp, err := client.Do(req)
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}

// Events
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

// Volumes
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
func HTTPGetInspectVolume(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/volumes/" + ps.ByName("name"))
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetServices(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/services")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetTasks(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/tasks")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
func HTTPGetNodes(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	tr := &http.Transport{
		Dial: dockerDial,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("http://localhost/nodes")
	if err != nil {
		tegoHttp.ResponseJSONCode(writer, "Oops!. Something went wrong.", 500)
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	tegoHttp.ResponsePlain(writer, string(body[:]), 400, "")
}
