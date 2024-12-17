// Flutter commands go here
package utils

import "os/exec"

func FlutterRun(args []string) *exec.Cmd {
	args = append([]string{"run"}, args...)
	return exec.Command("flutter", args...)
}

func FlutterDevices() *exec.Cmd {
	return exec.Command("flutter", "devices", "--machine")
}

func FlutterEmulators(args []string) *exec.Cmd {
	args = append([]string{"emulators"}, args...)
	return exec.Command("flutter", args...)
}
