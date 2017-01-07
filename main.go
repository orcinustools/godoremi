package main

import (
	"net/http"
	"os"

	"log"
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
		err = config.LoadConfigFromFile("./files/etc/godoremi/config.json")
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

	err = service.Initialize()
	if err != nil {
		log.Fatalln("Coudln't initialize service", err)
	}

	router := httprouter.New()

	router.GET("/", service.HTTPGetInfo)
	router.GET("/images", service.HTTPGetImages)
	router.GET("/containers", service.HTTPGetContainers)

	listenTo := config.Get().HTTP.Listen + ":" + config.Get().HTTP.Port
	log.Fatal(http.ListenAndServe(listenTo, router))
}
