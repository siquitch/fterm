package cmd

import "os/exec"

func FlutterRun(fvm bool) *exec.Cmd {
	if fvm {
		return exec.Command("fvm", "flutter")
	} else {
		return exec.Command("flutter")
	}
}
