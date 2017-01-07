package daemon

import "fmt"
import "os"

//Manage manages daemon
func Manage() (err error) {
	var output string
	if fInstall {
		output, err = d.Install()
	} else if fStart {
		output, err = d.Start()
	} else if fRemove {
		output, err = d.Remove()
	} else if fStop {
		output, err = d.Stop()
	} else if fStatus {
		output, err = d.Status()
	} else {
		return
	}

	fmt.Println(output)
	os.Exit(0)

	return
}
