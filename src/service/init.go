package service

import "github.com/wejick/tego/config"

var (
	dockerSocket string
)

//Initialize initiates service
func Initialize() (err error) {
	dockerSocket = config.Get().Upstream.UnixSocket["docker"].FileDescriptor

	return
}
