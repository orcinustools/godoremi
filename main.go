package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
	"github.com/orcinustools/godoremi/src/daemon"
	"github.com/orcinustools/godoremi/src/service"
	"github.com/wejick/tego/config"
)

var (
	signalChan = make(chan os.Signal, 1)
)

//Initialize initiates godoremi
func Initialize() (err error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err = config.LoadConfigFromFile("/etc/godoremi/config.json")
	if err != nil {
		err = config.LoadConfigFromFile("./config.json")
	}

	signal.Notify(signalChan,
		syscall.SIGTERM)

	//handle signal here
	go func() {
		for {
			s := <-signalChan
			switch s {
			case syscall.SIGTERM:
				os.Exit(0)
			}
		}
	}()

	return
}

func main() {
	err := Initialize()
	if err != nil {
		log.Fatalln("Coudln't initialize godoremi", err)
	}

	err = daemon.Initialize()
	if err != nil {
		log.Fatalln("Coudln't initialize daemon", err)
	}

	flag.Parse()

	daemon.Manage()

	err = service.Initialize()
	if err != nil {
		log.Fatalln("Coudln't initialize service", err)
	}

	router := httprouter.New()

	router.GET("/", service.HTTPGetInfo)
	router.GET("/version", service.HTTPGetVersion)
	router.GET("/networks", service.HTTPGetNetworks)
	router.GET("/networks/:id", service.HTTPGetInspectNetwork)
	router.GET("/ping", service.HTTPGetPing)

	router.GET("/images", service.HTTPGetImages)
	router.GET("/images/search", service.HTTPGetSearchImage)
	router.GET("/images/inspect/:name", service.HTTPGetInspectImage)
	router.GET("/images/history/:name", service.HTTPGetHistoryImage)

	router.GET("/containers", service.HTTPGetContainers)
	router.POST("/containers/create", service.HTTPPostCreateContainer)
	router.POST("/containers/start/:name", service.HTTPPostStartContainer)
	router.POST("/containers/restart/:name", service.HTTPPostRestartContainer)
	router.GET("/containers/rename/:name", service.HTTPGetRenameContainer)
	router.POST("/containers/stop/:name", service.HTTPPostStopContainer)
	router.POST("/containers/kill/:name", service.HTTPPostKillContainer)
	router.GET("/containers/inspect/:name", service.HTTPGetInspectContainer)
	router.GET("/containers/top/:name", service.HTTPGetTopContainer)
	router.GET("/containers/changes/:name", service.HTTPGetChangesContainer)
	router.GET("/containers/stats/:name", service.HTTPGetStatsContainer)
	router.GET("/containers/logs/:name", service.HTTPGetLogsContainer)
	router.GET("/containers/export/:name", service.HTTPGetExportContainer)
	router.POST("/containers/resize/:name", service.HTTPPostResizeContainer)
	router.POST("/containers/update/:name", service.HTTPPostUpdateContainer)

	router.GET("/events", service.HTTPGetEvents)

	router.GET("/volumes", service.HTTPGetVolumes)
	router.GET("/volumes/:name", service.HTTPGetInspectVolume)

	router.GET("/services", service.HTTPGetServices)

	listenTo := config.Get().HTTP.Listen + ":" + config.Get().HTTP.Port
	log.Fatal(http.ListenAndServe(listenTo, router))
}
