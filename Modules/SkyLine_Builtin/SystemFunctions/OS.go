package SkyLine_SYSTEM

import "runtime"

var GrabOperatingSystemDataBasedOnKey = map[string]func() (string, error){
	"os_name": func() (string, error) {
		funct := runtime.GOOS
		return funct, nil
	},
	"os_arch": func() (string, error) {
		funct := runtime.GOARCH
		return funct, nil
	},
}
