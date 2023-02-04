package SkyLine

import (
	"log"
	"os/exec"
	"strings"
)

// currently this only supports linux
// SupportOS: linux

// replace this later
func VerErr(x error) {
	if x != nil {
		log.Fatal(x)
	}
}

type ProcessInformation struct {
	ProcessName string // Name of the given process             | csc
	ProcessID   string // Process ID of the selected process    | 7886
	ProcessPath string // Path of the currently running process | /proc/7886/fd/0
}

func (PI *ProcessInformation) PIDbyProgramName(Progname string) {
	cout, x := exec.Command("pidof", Progname).Output()
	VerErr(x)
	o := string(cout)
	oline := strings.Split(o, "\n")
	for _, l := range oline {
		if l != "" {
			PI.ProcessID = l
			PI.ProcessName = Progname
		}
	}
}
