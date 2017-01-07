package daemon

import (
	"flag"

	"github.com/takama/daemon"
	"github.com/wejick/tego/config"
)

var (
	//daemon instance
	d daemon.Daemon

	fInstall = flag.Bool("install", false, "Install godoremi as service")
	fRemove  = flag.Bool("remove", false, "Remove godoremi service")
	fStart   = flag.Bool("start", false, "Start godoremi service")
	fStop    = flag.Bool("stop", false, "Stop godoremi service")
	fStatus  = flag.Bool("status", false, "Get godoremi service status")
)

//Initialize initiates service package
func Initialize() (err error) {
	d, err = daemon.New(config.Get().Name, config.Get().Description)

	return
}
