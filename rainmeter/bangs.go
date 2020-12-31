package rainmeter

import (
	"log"
	"os/exec"
)

var path string = "C:\\Program Files\\Rainmeter\\Rainmeter.exe"

// LoadLayout loads the Rainmeter layout with the given name.
func LoadLayout(layout string) {
	cmd := exec.Command(path + "!LoadLayout" + layout)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
