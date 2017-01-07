package daemon

import "fmt"

//Manage manages daemon
func Manage() (err error) {
	var output string
	if *fInstall {
		output, err = d.Install()
	} else if *fStart {
		output, err = d.Start()
	} else if *fStop {
		output, err = d.Stop()
	} else if *fStatus {
		output, err = d.Status()
	}
	fmt.Println(output)

	return
}
