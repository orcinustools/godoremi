package daemon

import (
	"flag"

	"github.com/takama/daemon"
	"github.com/wejick/tego/config"
)

var (
	//daemon instance
	d daemon.Daemon

	fInstall bool
	fRemove  bool
	fStart   bool
	fStop    bool
	fStatus  bool
)

//Initialize initiates service package
func Initialize() (err error) {
	flag.BoolVar(&fInstall, "install", false, "Install godoremi as service")
	flag.BoolVar(&fRemove, "remove", false, "Remove godoremi service")
	flag.BoolVar(&fStart, "start", false, "Start godoremi service")
	flag.BoolVar(&fStop, "stop", false, "Stop godoremi service")
	flag.BoolVar(&fStop, "status", false, "Get godoremi service status")

	d, err = daemon.New(config.Get().Name, config.Get().Description)

	return
}
